 <p align="center">
  <a href="https://alfonsofortunato.com">
    <picture>
      <source media="(prefers-color-scheme: dark)" srcset="https://alfonsofortunato.com/img/logo.png">
      <img src="https://alfonsofortunato.com/img/logo.png" height="90">
    </picture>
    <h1 align="center">Note cli</h1>
  </a>
</p>
<p align="center">
  <a href="https://github.com/MovieMaker93/note-cli/commit">
    <img alt="LastCommit" src="https://img.shields.io/github/last-commit/MovieMaker93/note-cli/main?style=for-the-badge&logo=github&color=%237dcfff">
  </a>
  <a href="https://github.com/MovieMaker93/devpod-dotfiles-chezmoi/blob/main/LICENSE">
    <img alt="License" src="https://img.shields.io/github/license/MovieMaker93/devpod-dotfiles-chezmoi?style=for-the-badge&logo=github">
  </a>
  <!-- <a href="https://github.com/MovieMaker93/devpod-dotfiles-chezmoi/stars"> -->
  <!--   <img alt="stars" src="https://img.shields.io/github/stars/MovieMaker93/devpod-dotfiles-chezmoi?style=for-the-badge&logo=github&color=%23f7768e"> -->
  <!-- </a> -->
</p>
<div style="text-align: center;">
  <h1>
    Introducing Note Cli 
  </h1>
</div>

Note CLI is a tool that allows users to create or open notes without ever leaving the terminal. If you're someone like me who lives in the terminal and is obsessed with note-taking, this might be the perfect tool for you.

Whether you’re using Nvim or Obsidian, this utility helps you create notes on the fly without losing focus losing focus.
![Refine](./public/refine.gif) 

## Table of Contents

- [Install](#install)
- [Environment Configuration](#environment-configuration)
- [Commands and Usage Example](#commands)
- [How to Customize note templates](#template)

<a id="install"></a>
<h2>
  <picture>
    <img src="./public/download.png" width="60px" style="margin-right: 1px;">
  </picture>
   Install
</h2>

### Requirements

- **Golang** version > 1.21
- **Nvim** installed

### Binary Installation
```bash
go install github.com/MovieMaker93/note-cli@latest
```

This installs a go binary that will automatically bind to your $GOPATH

### Building and Installing from Source

Clone the Note-Cli repository from GitHub:
```bash
git-clone https://github.com/MovieMaker93/note-cli
```
Build the Note-Cli binary:
```bash
go build
```
Install in your `PATH` to make it accessible system-wide:
```bash
go install
```
Verify the installation by running:
```bash
note-cli version
```

<a id="environment-configuration"></a>
<h2>
  <picture>
    <img src="./public/conf.png" width="60px" style="margin-right: 1px;">
  </picture>
  Environment Configuration
</h2>

You can set up your Vault path, Daily Note directory, New Note default directory, and date format via environment variables. A full example can be found in the `.env` file.

**Brief description:** 
- `DAILY_PATH`: Defines the directory where you store your daily notes in Obsidian
- `NEW_NOTE_PATH`: Defines the directory for your new notes in Obsidian (consume or refine)
- `VAULT`: Defines the absolute path to your Obsidian Vault
- `DATE_FORMAT`: Set your preferred date format. For `gg/mm/dd` format, use `02-01-2006`; for the US format use `2006-01-02`.

<a id="commands"></a>
<h2>
  <picture>
    <img src="./public/line.png" width="60px" style="margin-right: 1px;">
  </picture>
  Commands
</h2>

You can either create a new note (of type `refine` or `consume`) or open/create today’s note. 

When creating a new note, you have two options:
- Interact with the UI 
- Use an imperative command

Both methods allow you to create a note with a title, content, and type. The available note types are `consume` and `refine`.

### Example
![Example refine note](./public/refine.gif)

> [!NOTE]
> For consume notes, a link to the new note will be automatically created in your Today note.


For the second option, you can use the provided flags to create a refine note without interacting with the UI:
```bash
note-cli new -t example-consume -c example-content --type consume
```
See `note-cli new -h` for all the options and shorthands.


Another useful command is `today`, which allows you to open the existing today’s note or create one if it’s missing:
```bash
note-cli today
```


<a id="template"></a>
<h2>
  <picture>
    <img src="./public/note.png" width="60px" style="margin-right: 1px;">
  </picture>
  How to Customize Note templates
</h2>

The creation of both the Today and New notes is based on custom templates. You can find them inside: `cmd/template/note/files`

These templates are built using Go's `tmpl` package. If you wish to customize them, simply modify the following files:
- `cmd/template/note/file/consume.tmpl` --> for **consume** note
- `cmd/template/note/file/refine.tmpl` --> for **refine** note
- `cmd/template/note/file/today.tmpl` --> for **today** note
