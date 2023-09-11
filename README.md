# dotfiles

Here you will find files I use to configure my Linux workstations. 

![Perfect setup ;)](images/perfect-setup.jpg)

## Installation 👨‍💻

Warning: If you want to give these dotfiles a try, you should review the code, and adjust things to your own
need or liking. Don't use it blindly. Use at your own risk!

To update your desktop run:

```bash
git clone https://github.com/rojberr/dotfiles.git && cd dotfiles && set -- -f; source ./setup.sh
```

Git-free installation
To install without using git use:

```bash
cd; \
curl -L https://github.com/rojberr/dotfiles/archive/0.0.1.tar.gz -o dotfiles-0.0.1.tar.gz && \
tar -zxvf dotfiles-0.0.1.tar.gz && \
cd dotfiles-0.0.1 && set -- -f; source ./setup.sh
```

