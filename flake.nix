{
  description = "darwin config";

  inputs = {
    nixpkgs.url = "github:NixOS/nixpkgs/nixpkgs-unstable";
    nix-darwin.url = "github:LnL7/nix-darwin";
    nix-darwin.inputs.nixpkgs.follows = "nixpkgs";
    nix-homebrew.url = "github:zhaofengli-wip/nix-homebrew";
    homebrew-core = {
      url = "github:homebrew/homebrew-core";
      flake = false;
    };
    homebrew-cask = {
      url = "github:homebrew/homebrew-cask";
      flake = false;
    };
    homebrew-bundle = {
      url = "github:homebrew/homebrew-bundle";
      flake = false;
    };
    home-manager.url = "github:nix-community/home-manager";
    home-manager.inputs.nixpkgs.follows = "nixpkgs";
  };

  outputs = inputs@{ self, nix-darwin, nixpkgs, nix-homebrew, homebrew-core, homebrew-cask, homebrew-bundle, home-manager }:
  let
    configuration = { pkgs, config, ... }: {
      nixpkgs.config.allowUnfree = true;
      environment.systemPackages =
        [
          # tooling
          pkgs.kubectl
          pkgs.kubectx
          pkgs.fluxcd
          pkgs.kind
          pkgs.sops
          pkgs.tmux
          pkgs.opentofu
          pkgs.hcloud
          pkgs.pre-commit
          pkgs.tree
          pkgs.jq
          pkgs.bore-cli
          pkgs.git
          pkgs.neovim
          pkgs.buf
          pkgs.direnv
          pkgs.age
          pkgs.mtr
          pkgs.gnupg
          pkgs.yubikey-manager
          pkgs.yubico-pam
          pkgs.yubikey-personalization

          # langs
          pkgs.go

          # gui apps
          pkgs.jetbrains.goland
          pkgs.jetbrains.idea-community
          pkgs.alacritty

          # misc
          pkgs.mkalias
          pkgs.oh-my-zsh
        ];

      homebrew = {
        enable = true;
        # remove all packages that have
        # not been installed using nix
        onActivation.cleanup = "zap";
        brews = [
          "lima"
          "wget"
          "curl"
          "mas"
          "fzf"
          "zoxide"
        ];
        casks = [
          "google-chrome"
          "bitwarden"
          "discord"
          "zed"
        ];
        masApps = {
          "Magnet" = 441258766;
          "Tailscale" = 1475387142;
        };
      };

      system.defaults.NSGlobalDomain."com.apple.swipescrolldirection" = false;

      users.users.yannic = {
          name = "yannic";
          home = "/Users/yannic";
      };

      system.activationScripts.applications.text = let
        env = pkgs.buildEnv {
          name = "system-applications";
          paths = config.environment.systemPackages;
          pathsToLink = "/Applications";
        };
      in
        pkgs.lib.mkForce ''
          # Set up applications.
          echo "setting up /Applications..." >&2
          rm -rf /Applications/Nix\ Apps
          mkdir -p /Applications/Nix\ Apps
          find ${env}/Applications -maxdepth 1 -type l -exec readlink '{}' + |
          while read -r src; do
            app_name=$(basename "$src")
            echo "copying $src" >&2
            ${pkgs.mkalias}/bin/mkalias "$src" "/Applications/Nix Apps/$app_name"
          done
        '';

      # Necessary for using flakes on this system.
      nix.settings.experimental-features = "nix-command flakes";

      programs.zsh.enable = true;

      # Set Git commit hash for darwin-version.
      system.configurationRevision = self.rev or self.dirtyRev or null;

      # Used for backwards compatibility, please read the changelog before changing.
      # $ darwin-rebuild changelog
      system.stateVersion = 5;

      # The platform the configuration will be used on.
      nixpkgs.hostPlatform = "aarch64-darwin";
    };
  in
  {
    darwinConfigurations."personal" = nix-darwin.lib.darwinSystem {
      modules = [
        configuration
        home-manager.darwinModules.home-manager
        {
          home-manager.useGlobalPkgs = true;
          home-manager.useUserPackages = true;
          home-manager.users.yannic = import ./home.nix;
        }
        nix-homebrew.darwinModules.nix-homebrew
        {
          nix-homebrew = {
            # Install Homebrew under the default prefix
            enable = true;

            # Apple Silicon Only: Also install Homebrew under the default Intel prefix for Rosetta 2
            enableRosetta = true;

            # User owning the Homebrew prefix
            user = "yannic";

            # Optional: Declarative tap management
            taps = {
              "homebrew/homebrew-core" = homebrew-core;
              "homebrew/homebrew-cask" = homebrew-cask;
              "homebrew/homebrew-bundle" = homebrew-bundle;
            };

            # Optional: Enable fully-declarative tap management
            #
            # With mutableTaps disabled, taps can no longer be added imperatively with `brew tap`.
            mutableTaps = false;
          };
        }
      ];
    };
  };
}
