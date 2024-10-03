 <p align="center">
  <a href="https://alfonsofortunato.com">
    <picture>
      <source media="(prefers-color-scheme: dark)" srcset="https://alfonsofortunato.com/img/logo.png">
      <img src="https://alfonsofortunato.com/img/logo.png" height="90">
    </picture>
    <h1 align="center">Note cli</h1>
  </a>
</p>
<div style="text-align: center;">
  <h1>
    Introducing Note Cli 
  </h1>
</div>

Note CLI is a tool that allows users to create or open notes without ever leaving the terminal. If you're someone like me who lives in the terminal and is obsessed with note-taking, this might be the perfect tool for you.

Whether you’re using Nvim or Obsidian, this utility helps you create notes on the fly without losing focus losing focus.

## Table of Contents

- [Install](#install)
- [Environment Configuration](#environment-configuration)
- [Commands](#commands)
- [Usage Example](#usage-example)

<a id="install"></a>
<h2>
  <picture>
    <img src="./public/install.gif?raw=true" width="60px" style="margin-right: 1px;">
  </picture>
  Install
</h2>

## Binary Installation
```bash
go install github.com/MovieMaker93/note-cli@latest
```

This installs a go binary that will automatically bind to your $GOPATH

## Building and Installing from Source

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
    <img src="./public/install.gif?raw=true" width="60px" style="margin-right: 1px;">
  </picture>
  Environment Configuration
</h2>

You can set up your own Vault path, Daily note directory, New note defualt directory and date format via `Environment variables`.
You can find a full example inside `.env` file. 

A brief description:
- `DAILY_PATH`: Define your directory where you put your daily note in Obsidian
- `NEW_NOTE_PATH`: Define your directory where you put your new note in Obsidian (consume or refine)
- `VAULT`: Define your absolute Obsidian Vault path
- `DATE_FORMAT`: Define your date format. `02-01-2006` for the standard UTC Europe, for US `2006-01-02`.

<a id="commands"></a>
<h2>
  <picture>
    <img src="./public/install.gif?raw=true" width="60px" style="margin-right: 1px;">
  </picture>
  Commands
</h2>




You can also use the provided flags to set up a project without interacting with the UI.

```bash
go-blueprint create --name my-project --framework gin --driver postgres --git commit
```

See `go-blueprint create -h` for all the options and shorthands.



