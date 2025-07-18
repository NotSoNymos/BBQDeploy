<div align="center">
<h1>🚀BBQDeploy: Online Installer for GitHub Projects🚀</h1>
</div>

[РУССКИЙ](https://github.com/bazelik-null/BBQDeploy/blob/main/README_RU.md)

>[!CAUTION]
>
>🛠️ Work In Progress

## 📚 Table of Contents
- [Description](#description)
- [IMPORTANT](#important)
- [Installation](#installation)
- [Plugins](#plugins)
- [License](#license)

## ✨ Description
The **BBQDeploy** installer is an easy-to-use cross-platform installer based on the [Fyne](https://fyne.io/) framework, designed to work with GitHub projects.

It comes as a single executable file and installs all the necessary resources from the latest release on GitHub.

BBQDeploy offers a high degree of customization through its [plugins](#plugins) system and full configuration in `config.toml`, allowing you to tailor it to your needs perfectly! 🎨

Example usage: [Example repo](https://github.com/bazelik-null/example)

## ⚠️ IMPORTANT
For a detailed manual on installation, configuration, creating your own plugins, and releasing, visit the [documentation](https://github.com/bazelik-null/BBQDeploy/wiki) 📖

## 🛠️ Installation
1. Clone the installer into your project using the submodule system.
```bash
git submodule add https://github.com/bazelik-null/BBQDeploy
```
2. Initialize the submodule.
```bash
git submodule init
git submodule update
```
3. Copy the necessary files from the submodule and paste them into the root of your repository.
```bash
cp -r ./BBQDeploy/resources ./resources
cp -r ./BBQDeploy/MAKEFILE ./MAKEFILE
```
4. Move the necessary files into it.
5. Open `resources/config/meta.json` and insert the path to the source file in the `src` field, and the path to the destination in the `dst` field. If there are multiple files, copy the structure with `src` and `dst` and paste it as many times as needed.
6. Open `resources/config/config.toml` and configure the installer.
7. Open `MAKEFILE` and fill in the variables `OWNER` (repository owner) and `NAME` (repository name).

## ⚙️ Plugins
The BBQDeploy installer features a unique plugin system that allows you to add **your own logic** without modifying the source code.

All plugins are written in GOlang but are executed using the [Yaegi](https://github.com/traefik/yaegi) interpreter, which allows them to be run without compilation.

For complete documentation on plugin development, read the [documentation](https://github.com/bazelik-null/BBQDeploy/wiki).

## 📝 License
This project is licensed under the open-source [MIT](https://mit-license.org/) license. You are free to use, modify, and distribute this installer in accordance with the terms of the license. 🌟