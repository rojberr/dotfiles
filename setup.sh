#!/bin/bash

# Install pip for ansible installation
if python3 -m pip -V; then
  curl https://bootstrap.pypa.io/get-pip.py -o get-pip.py
  python3 get-pip.py --user
else
  echo "Install python3"
  exit 1
fi

# Install ansible, dependencies, roles
python3 -m pip install --user ansible
sudo ansible-galaxy install --force -r roles_to_install.yaml

# Install apps
ansible-playbook install-apps.yaml -e ALLOW_ERRORS=no --ask-become-pass