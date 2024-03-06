<h1 align=center>Rojberr Dotfiles | <a href="https://rojberr.github.io/dotfiles/" rel="nofollow">Gh-pages 📚</a></h1>

> Here you will find files I use to configure my Linux/MacOS workstations.

The idea behind this repo is that it can be used on any system that's running a Bash shell. Just clone the repo, run and
install everything with the script. Use with Linux/MacOS/WSL.

**Take a look at the [Documentation page](https://rojberr.github.io/dotfiles/)
or [dotfiles Wiki](https://github.com/rojberr/dotfiles/wiki)**

[![Deploy dotfiles Docu to Pages](https://github.com/rojberr/dotfiles/actions/workflows/gh-pages.yml/badge.svg)](https://github.com/rojberr/dotfiles/actions/workflows/gh-pages.yml)

---

<p align="center">
<kbd><img src="images/perfect-setup.jpg" alt="Perfect setup" title="Perfect setup"/></kbd>
</p>

---

## Contents

What's in here?

- all my `brew` dependencies including: applications, fonts, etc.
  See [`Brewfile`](https://github.com/rojberr/dotfiles/blob/master/Brewfile)
- all my `scripts` that I use daily. See [`scripts/`](https://github.com/rojberr/dotfiles/blob/master/scripts/)
- all my `shell` settings. I'm using bash as default.
  See [`shell/`](https://github.com/rojberr/dotfiles/tree/master/shell)
- all my `configuration` files. See [`config/`](https://github.com/rojberr/dotfiles/tree/master/config) for nvim and
  other.

The configuration files are used by symlinks to config location.

## Installation 👨‍💻

*Disclaimer: This is my personal solution! You may need to customize it a bit to make it work with your machine.

Warning: If you want to give these dotfiles a try, you should review the code, and adjust things to your own
need or liking. Don't use it blindly. Use at your own risk!

To update your desktop run:

```bash
git clone https://github.com/rojberr/dotfiles.git && cd dotfiles && set -- -f; source ./setup.sh
```

## Why Dotfiles? 💡

This repository will quickly install all required dependencies, program,s files and configs that I need to work with
code. It sets up the toolchain, that lets change between env versions. It also makes my
env on different working stations identical, which speeds up my work.

It allows me to:

- collect ideas for tooling, config and scripts,
- share it with the community,
- let others benefit from my experience,
- allow me to download it wherever I have access to internet.

## What are Dotfiles?

It's a set of scripts, Ansible roles and other allowing me to provision a bunch of tools on a new machine.

The idea is to build one binary, which upon execution will do the provisioning for me.

This repository contains my configuration files for the tools that I use.

Confused? Read on.

### Brewfile

I use simple Brewfile to manage my apps. The versions aren't pinned, but it's enought to work with my homelab.

To dump current apps to Brewfile I use:

```bash
brew bumdle dump
```

To import from Brewfile:

```bash
brew bundle
```

### TODO

shortcuts generalize
what is not possible to automate should be noted in repo
autocomplete verbessern
create symlinks in loop

## Support 🫶

- Star 🌟 this repository.
- Help spread the word about PaperMod by sharing it on social media and recommending it to your friends. 🗣️
- You can also sponsor 🏅
  on [Github Sponsors](https://github.com/sponsors/adityatelange) / [Ko-Fi](https://ko-fi.com/adityatelange).

### References

https://dotfiles.github.io/
https://missing.csail.mit.edu/2020/potpourri/
