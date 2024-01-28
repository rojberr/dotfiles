# dotfiles

Here you will find files I use to configure my Linux/MacOS workstations. 

The idea behind this repo is that it can be used on any system that's running a Bash shell. Just clone the repo, run and install everything with the script. I mostly use Linux/MacOS/WSL.

![Perfect setup ;)](images/perfect-setup.jpg)

## Contents 

What's in here?
- all my `brew` dependencies including: applications, fonts, etc. See [`Brewfile`](https://github.com/rojberr/dotfiles/blob/master/Brewfile)
- all my `scripts` that I use daily. See [`scripts/`](https://github.com/rojberr/dotfiles/blob/master/scripts/)
- all my `shell` settings. See [`shell/`](https://github.com/rojberr/dotfiles/tree/master/shell)
- all my `configuration` files. See [`config/`](https://github.com/rojberr/dotfiles/tree/master/config)

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

This repository install quickly all required dependencies to work with code.
It sets up the toolchain, that lets change between env versions. It also makes my
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


### tbc

For now I am using Brew for binary installation. Another idea
is to use ready-to-use binary.

1. Install Ansible (link:https://docs.ansible.com/ansible/latest/installation_guide/intro_installation.html[Ansible page])
```bash
curl https://bootstrap.pypa.io/get-pip.py -o get-pip.py
python3 get-pip.py --user
python3 -m pip install --user ansible-core==2.12.3
python3 -m pip install --user ansible
```

2. Run the playbook
```bash
ansible-playbook playbooks/install_role.yml
```

3. Enjoy my setup on your PC 👨‍💻

4. To upgrade ansible you may use:
```bash
python3 -m pip install --upgrade --user ansible-core==2.12.3
ansible --version # Checks the ansible-core package version
python3 -m pip show ansible # Checks the ansible package version
```

Add Ansible command shell completion
python3 -m pip install --user argcomplete

### tbc - Git-free installation

To install without using git use:

```bash
cd; \
curl -L https://github.com/rojberr/dotfiles/archive/0.0.1.tar.gz -o dotfiles-0.0.1.tar.gz && \
tar -zxvf dotfiles-0.0.1.tar.gz && \
cd dotfiles-0.0.1 && set -- -f; source ./setup.sh
```

