#!/usr/bin/env bash
set -e
mkdir -p ~/.config/
cd ~/.config
git clone https://github.com/Freggy/dotfiles.git
curl -fsSL https://raw.githubusercontent.com/Homebrew/install/HEAD/install.sh | bash
brew install go
go build main.go -o dof
chmod +x dof
mkdir ~/bin
mv dof ~/bin/dof
~/bin/dof apply
