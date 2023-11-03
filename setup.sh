#!/bin/bash

# Install pip for ansible installation
if command -v python3 >/dev/null 2>&1 && echo "Python 3 is installed"; then
  curl https://bootstrap.pypa.io/get-pip.py -o get-pip.py
  python3 get-pip.py --user
else
  echo "ERROR: Install python3"
  exit 1
fi

# Install ansible, dependencies, roles
python3 -m pip install --user ansible
sudo ansible-galaxy install --force -r roles_to_install.yaml

# Install apps
ansible-playbook install-apps.yaml -e ALLOW_ERRORS=no --ask-become-pass