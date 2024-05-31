#!/usr/bin/env bash

set -e

cd "$(dirname "$0")"

echo "Installing Oh-My-Posh"

curl -s https://ohmyposh.dev/install.sh | sudo bash -s

curl -s https://raw.githubusercontent.com/JanDeDobbeleer/oh-my-posh/main/themes/catppuccin.omp.json > ~/catppuccin.omp.json
grep -qF "oh-my-posh" ~/.bashrc || echo 'eval "$(oh-my-posh init bash --config ~/catppuccin.omp.json)"' >> ~/.bashrc