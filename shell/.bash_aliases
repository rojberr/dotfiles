##############################################################################
# Bash Aliases                                                                #
##############################################################################

# General
alias ll="ls -al"

# Kubernetes aliases
alias k="kubectl"
complete -o default -F __start_kubectl k
alias knc="kubectl config set-context --current --namespace"
alias kl="kubectl logs"
alias kd="kubectl describe"
alias kg="kubectl get"
alias kga="kubectl get all"
alias kgn="kubectl get nodes"
alias kei="kubectl exec -it"
alias kpf="kubectl port-forward"
alias kgnm="kubectl get nodes -o=custom-columns=NAME:.metadata.name,STATUS:.status.conditions[?(@.type==\"Ready\")].status"

export drc="--dry-run=client -o yaml" # To run add $drc at the end
export drs="--dry-run=server -o yaml"

# Use neovim
alias v=nvim

# Docker aliases
alias d="docker"

# Git aliases
alias gs='git status'
alias ga='git add'
alias gaa='git add .'
alias gpl='git pull'
alias gp='git push'
#alias gd='git diff | mate'
#alias gau='git add --update'
alias gc='git commit --signoff -v'          # Commit with sign, verbose
alias gca='git commit --signoff -v -a'      # Commit ALL sign, verbose
alias gcm='git commit --signoff -v -m'      # Commit with MESSAGE sign, verbose
alias gb='git branch'
alias gba='git branch -a'
alias gco='git checkout'
alias gcob='git checkout -b'
#alias gcot='git checkout -t'
#alias gcotb='git checkout --track -b'
alias gl='git log --decorate --graph --oneline'
alias glp='git log --pretty=format:"%h %s" --graph'
## This command generates the entire log of the git repo in one line
alias gla='git log --decorate --graph --oneline --all'
## to view entire log with commit messages in detail
alias gld='git log --decorate --graph --all'

# Go up n directories
alias ..="cd .."
alias ...="cd ../.."
alias ....="cd ../../.."


# Recursively remove .DS_Store files
alias cleanupds="find . -type f -name '*.DS_Store' -ls -delete"