{ config, pkgs, ... }:

let
  files = "${config.home.homeDirectory}/proj/dotfiles/files";
in
  {
    programs.home-manager.enable = true;
    home.stateVersion = "24.11";
    home.file.".tmux.conf".source = "${files}/.tmux.conf";
    home.file.".zshrc".source = "${files}/.zshrc";
    home.file.".zsh_aliases".source = "${files}/.zsh_aliases";
    home.file.".config/zed/settings.json".source = "${files}/zed/settings.json";
    home.file.".config/zed/keymap.json".source = "${files}/zed/keymap.json";
    programs.git = {
      enable = true;
      userName  = "yannic rieger";
      userEmail = "ybr@76k.io";
      lfs.enable = true;
      extraConfig = {
        commit.gpgsign = true;
        # yubikey 29-124-674
        user.signingkey = "A4F2D129FCD88C5EFB6D64767FF0A945E1587BF4";
      };
    };
    programs.alacritty = {
      enable = true;
      settings = {
        font = {
          size = 13;
        };
      };
    };
  }
