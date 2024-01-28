# Kubernetes aliases
alias k="kubectl"
alias kn="kubectl config set-context --current --namespace"
alias kl="kubectl logs"
alias kd="kubectl describe"
alias kg="kubectl get"
alias kga="kubectl get all"
alias kgn="kubectl get nodes"
alias kei="kubectl exec -it"
alias kpf="kubectl port-forward"
alias kgnm="kubectl get nodes -o=custom-columns=NAME:.metadata.name,STATUS:.status.conditions[?(@.type==\"Ready\")].status"

# Use neovim
alias vim=nvim

# Docker aliases
alias d="docker"

# Git aliases
alias gst='git status'
alias ga='git add'
alias gaa='git add .'
alias gl='git pull'
alias gp='git push'
alias gd='git diff | mate'
alias gau='git add --update'
alias gc='git commit -v'
alias gca='git commit -v -a'
alias gcm='git commit -m'
alias gb='git branch'
alias gba='git branch -a'
alias gco='git checkout'
alias gcob='git checkout -b'
alias gcot='git checkout -t'
alias gcotb='git checkout --track -b'
alias glog='git log'
alias glogp='git log --pretty=format:"%h %s" --graph'
## This command generates the entire log of the git repo in one line
alias glogo='git log --decorate --graph --oneline --all'
## to view entire log with commit messages in detail
alias gloga='git log --decorate --graph --all'
