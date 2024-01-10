#!/bin/bash

set -ex

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

# Install Brew
/bin/bash -c "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/HEAD/install.sh)"

# # Install Brewfile apps
brew bundle

# Install NvChar
git clone https://github.com/NvChad/NvChad ~/.config/nvim --depth 1 && nvim

# Load .bash_aliases
rm -f ~/.bash_aliases
ln -s `pwd`/config/.bash_aliases ~/.bash_aliases

