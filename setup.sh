#!/bin/bash

# Install pip for ansible installation
if python3 -m pip -V; then
  curl https://bootstrap.pypa.io/get-pip.py -o get-pip.py
  python3 get-pip.py --user
else
  echo "Install python3"
fi

# Install ansible, dependencies, roles
python3 -m pip install --user ansible
ansible-galaxy install --force -r roles_to_install.yaml

# Install apps
ansible-playbook install-apps.yaml -e ALLOW_ERRORS=no