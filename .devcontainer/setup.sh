#!/usr/bin/env bash
set -eux

# Install dependencies (for GO)
sudo apt-get update
sudo apt-get install -y bsdmainutils bison git curl zip unzip

# Install NVM + Node + Claude
curl -o- https://raw.githubusercontent.com/nvm-sh/nvm/v0.40.3/install.sh | bash
. "$HOME/.nvm/nvm.sh"
nvm install 22
node -v
npm -v
npm install -g @anthropic-ai/claude-code

# Install GVM (Go Version Manager)
bash < <(curl -sSL https://raw.githubusercontent.com/moovweb/gvm/master/binscripts/gvm-installer)
source "$HOME/.gvm/scripts/gvm"
gvm install go1.17.13 --prefer-binary
gvm use go1.17.13 --default

# Install GopherJS
go install github.com/gopherjs/gopherjs@v1.17.2

# # Build lute.min.js
# cd /workspaces/lute/javascript/
# export PATH=$PATH:/usr/local/go/bin:$HOME/go/bin
#   gopherjs build --tags javascript -o lute.min.js -m
