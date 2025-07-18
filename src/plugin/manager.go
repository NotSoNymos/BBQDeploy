package plugin

import (
	"log"
	"os"
	"path/filepath"
	"reflect"

	"github.com/traefik/yaegi/interp"
	"github.com/traefik/yaegi/stdlib"
)

// Manager loads plugins, reads their Hooks maps, and dispatches entries.
type Manager struct {
	interpreter *interp.Interpreter
	hooks       map[string][]func()
}

// Global is the shared plugin manager instance
var Global = NewManager()

// NewManager creates a Yaegi interpreter, preloads the standard library, and initializes hook storage.
func NewManager() *Manager {
	m := &Manager{
		interpreter: interp.New(interp.Options{}),
		hooks:       make(map[string][]func()),
	}
	// Import GO stdlib
	_ = m.interpreter.Use(stdlib.Symbols)
	return m
}

// LoadPlugins scans dir for *.go plugins, Eval's them, then looks for a main.Hooks variable of type map[string]func().
func (m *Manager) LoadPlugins(dir string) {
	log.Printf("[INFO]: scanning %s for plugins", dir)

	entries, err := os.ReadDir(dir)
	if err != nil {
		log.Printf("[ERROR]: reading plugins dir: %v", err)
		return
	}

	for _, e := range entries {
		if e.IsDir() || filepath.Ext(e.Name()) != ".go" {
			continue
		}
		path := filepath.Join(dir, e.Name())
		src, err := os.ReadFile(path)
		if err != nil {
			log.Printf("[ERROR]: reading %s: %v", path, err)
			continue
		}

		// Evaluate the plugin source as package main
		_, err = m.interpreter.Eval(string(src))
		if err != nil {
			log.Printf("[ERROR]: evaluating %s: %v", path, err)
			continue
		}

		// Try to fetch main.Hooks
		sym, err := m.interpreter.Eval("main.Hooks")
		if err != nil {
			log.Printf("[INFO]: no Hooks in %s (skipping)", e.Name())
			continue
		}

		hooksVal := sym.Interface()
		rv := reflect.ValueOf(hooksVal)
		if rv.Kind() != reflect.Map {
			log.Printf("[ERROR]: Hooks in %s is not a map, got %T", e.Name(), hooksVal)
			continue
		}

		for _, key := range rv.MapKeys() {
			name := key.String()
			fnVal := rv.MapIndex(key)
			if !fnVal.IsValid() || fnVal.Kind() != reflect.Func {
				log.Printf("[ERROR]: Hooks[%q] in %s is not a func", name, e.Name())
				continue
			}
			// Convert to Go func()
			hookFn := fnVal.Interface().(func())
			log.Printf("[INFO]: registering hook %q from %s", name, e.Name())
			m.hooks[name] = append(m.hooks[name], hookFn)
		}
	}
}

// Entry fires all handlers registered under the given hook name.
func (m *Manager) Entry(name string) {
	log.Printf("[INFO]: reached entry point %q", name)
	for _, h := range m.hooks[name] {
		h()
	}
}
