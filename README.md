<h1 align=center>Rojberr Dotfiles ~/.* | <a href="https://rojberr.github.io/dotfiles/" rel="nofollow">Gh-pages 📚</a></h1>

> ~/.* Here you will find files I use to configure my Linux/MacOS workstations.

The idea behind this repo is that it can be used on any system that's running a Bash shell.
Just clone the repo, run and install everything with the script. Use with Linux/MacOS/WSL.
Supports both Apple Silicon (M1) and Intel chips.

**Take a look at the [Documentation page](https://rojberr.github.io/dotfiles/)
or [dotfiles Wiki](https://github.com/rojberr/dotfiles/wiki)**

[![Deploy .files Docu to Pages](https://github.com/rojberr/dotfiles/actions/workflows/gh-pages.yml/badge.svg)](https://github.com/rojberr/dotfiles/actions/workflows/gh-pages.yml)

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
  other
- `$EDITOR` is [neovim](https://neovim.io/).

The configuration files are used by symlinks to config location.

## Installation 👨‍💻

> The only pre-requisite is to install Git.

```bash
git config --global user.name "your name"
git config --global user.email "your@email.com"
git config --global github.user "your-github-username"
```

**Disclaimer: This is my personal solution! You may need to customize it a bit to make it work with your machine.**
**You can use different branches for different computers, you can replicate you configuration easily on new
installation.**

To install `dotfiles` and update your desktop run:

```bash
git clone https://github.com/rojberr/dotfiles.git ~/dotfiles && cd ~/dotfiles && set -- -f; source ./install.sh
```

**For more information take a look at the [Documentation page](https://rojberr.github.io/dotfiles/)
or [dotfiles Wiki](https://github.com/rojberr/dotfiles/wiki)** 🙂

### TODO

- [ ] shortcuts generalize
- [ ] what is not possible to automate should be noted in repo
- [ ] autocomplete verbessern
- [ ] create symlinks in loop
- [ ] more content to gh pages

## Support 🫶

- Star 🌟 this repository.
- Contact me if you have any questions or suggestions. 📣

![img.png](images/dotsymbol.png)