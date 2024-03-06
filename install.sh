#!/bin/bash
set -ex

## Default to BASH
chsh -s /bin/bash

# Install Brew
/bin/bash -c "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/HEAD/install.sh)"

# Backup, create and load .bash_aliases
mv ~/.bash_aliases ~/.bash_aliases.bak
rm -f ~/.bash_aliases
ln -s `pwd`/shell/.bash_aliases ~/.bash_aliases

# Backup, create and load bashrc
mv  ~/.bash_rc ~/.bash_rc.bak
rm -f ~/.bashrc
ln -s `pwd`/shell/.bashrc ~/.bashrc
source ~/.bashrc

# # Install Brewfile apps
brew bundle

# Install starship
sh -c "$(curl -fsSL https://starship.rs/install.sh)"

# Install NvChar
git clone https://github.com/NvChad/NvChad ~/.config/nvim --depth 1 && nvim

# Install bash completions
kubectl completion bash >/usr/local/etc/bash_completion.d/kubectl

# Install Helm plugins
helm plugin install https://github.com/databus23/helm-diff
helm plugin install https://github.com/jkroepke/helm-secrets
helm plugin install https://github.com/aslafy-z/helm-git #--version 0.15.1

bash

# Install pip for ansible installation
# if command -v python3 >/dev/null 2>&1 && echo "Python 3 is installed"; then
#   curl https://bootstrap.pypa.io/get-pip.py -o get-pip.py
#   python3 get-pip.py --user --break-system-packages
# else
#   echo "ERROR: Install python3"
#   exit 1
# fi

# Install ansible, dependencies, roles
# python3 -m pip install --user ansible --break-system-packages
# sudo ansible-galaxy install --force -r roles_to_install.yaml

# Install apps
# ansible-playbook install-apps.yaml -e ALLOW_ERRORS=no --ask-become-pass