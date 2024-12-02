{ config, pkgs, ... }:

let
  files = "${config.home.homeDirectory}/proj/dotfiles/files";
in
  {
    programs.home-manager.enable = true;
    home.stateVersion = "24.11";
    home.file.".tmux.conf".source = "${files}/.tmux.conf";
    home.file.".zshrc".source = "${files}/.zshrc";
  }
