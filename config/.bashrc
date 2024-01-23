# Starship
eval "$(starship init bash)" || true

[ -f ~/.fzf.bash ] && source ~/.fzf.bash

# Get aliases
source ~/.bash_aliases

# Homebrew to PATH
(echo; echo 'eval "$(/home/linuxbrew/.linuxbrew/bin/brew shellenv)"')
eval "$(/home/linuxbrew/.linuxbrew/bin/brew shellenv)"

