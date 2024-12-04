{ config, pkgs, ... }:

let
  files = "${config.home.homeDirectory}/proj/dotfiles/files";
in
  {
    programs.home-manager.enable = true;
    home.stateVersion = "24.11";
    home.file.".tmux.conf".source = "${files}/.tmux.conf";
    home.file.".zshrc".source = "${files}/.zshrc";
    home.file.".config/zed/settings.json".source = "${files}/zed/settings.json";
    home.file.".config/zed/keymap.json".source = "${files}/zed/keymap.json";
    programs.git = {
      enable = true;
      userName  = "yannic rieger";
      userEmail = "ybr@76k.io";
      lfs.enable = true;
      extraConfig = {
        commit.gpgsign = true;
        user.signingkey = "DFEC74EA9AB2B963";
      };
    };
  }
