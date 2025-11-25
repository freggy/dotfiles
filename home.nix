{ config, pkgs, ... }:

let
  files = "${config.home.homeDirectory}/proj/dotfiles/files";
in
  {
    programs.home-manager.enable = true;
    home.stateVersion = "24.11";
    home.file.".tmux.conf".source = "${files}/.tmux.conf";
    home.file.".config/zed/settings.json".source = "${files}/zed/settings.json";
    home.file.".config/zed/keymap.json".source = "${files}/zed/keymap.json";
    programs.git = {
      enable = true;
      lfs.enable = true;
      settings = {
        user.name  = "yannic rieger";
        user.email = "ybr@76k.io";
        commit.gpgsign = true;
        # yubikey 29-124-674
        user.signingkey = "A4F2D129FCD88C5EFB6D64767FF0A945E1587BF4";
        push.autoSetupRemote = true;
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
    programs.neovim = {
      enable = true;
      # set clipboard+=unnamedplus makes system clipboard available to neovim
      extraConfig = ''
        set clipboard+=unnamedplus
      '';
    };
    programs.direnv = {
      enable = true;
      enableZshIntegration = true;
      };
    programs.zsh = {
      enable = true;
      oh-my-zsh = {
        enable = true;
        plugins = ["git"];
        theme = "afowler";
      };
      shellAliases = {
        tf = "tofu";
        k = "kubectl";
        kx = "kubectx";
        cd = "z"; # zoxide
        ll = "ls -la";
        vim = "nvim";
        gs = "git status";
        gd = "git diff";
        gdc = "git diff --cached";
        gce = "git commit -v --edit";
        xc = "limactl shell xcomp";
      };
      initContent = ''
        export GOROOT=$(go env GOROOT)
        export GOPATH=$(go env GOPATH)
        export GOBIN=$GOPATH/bin
        export GIT_EDITOR=nvim
        export PATH=$PATH:$GOBIN
        export LANG=en_US

        EDITOR=nvim
        source <(fzf --zsh)
        eval "$(zoxide init zsh)"
        eval "$(direnv hook zsh)"

        bindkey -e
        bindkey '^[[1;3C' forward-word
        bindkey '^[[1;3D' backward-word

        if [ "$TMUX" = "" ]; then tmux; fi
      '';
    };
  }
