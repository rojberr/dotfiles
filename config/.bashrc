# Bashrc
[ -f ~/.fzf.bash ] && source ~/.fzf.bash
[ -f /usr/local/etc/bash_completion ] && . /usr/local/etc/bash_completion
brew_etc="$(brew --prefix)/etc" && [[ -r "${brew_etc}/profile.d/bash_completion.sh" ]] && . "${brew_etc}/profile.d/bash_completion.sh"
source <(kubectl completion bash)
[[ -r "$HOMEBREW_PREFIX/etc/profile.d/bash_completion.sh" ]] && . "$HOMEBREW_PREFIX/etc/profile.d/bash_completion.sh"

# Get aliases
source ~/.bash_aliases

# Homebrew to PATH
(echo; echo 'eval "$(/home/linuxbrew/.linuxbrew/bin/brew shellenv)"')
eval "$(/home/linuxbrew/.linuxbrew/bin/brew shellenv)"

# Starship
eval "$(starship init bash)" || true
