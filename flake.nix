{
  description = "Go development environment and tools";

  inputs = {
    nixpkgs.url = "github:NixOS/nixpkgs/nixos-unstable";
    flake-utils.url = "github:numtide/flake-utils";
  };

  outputs = { self, nixpkgs, flake-utils }:
    flake-utils.lib.eachDefaultSystem (system:
      let
        pkgs = import nixpkgs {
          inherit system;
        };
      in
      {
        devShells.default = pkgs.mkShell {
          packages = with pkgs; [
            # Go
            go

            # Language Server / editor support
            gopls

            # Debugger
            delve

            # Go tools
            gotools

            # Static analysis / linting
            golangci-lint
          ];

          shellHook = ''
            echo "Go development shell"
            echo "Go version: $(go version)"
            echo "gopls:       $(gopls version 2>/dev/null || echo 'available')"
          '';
        };
      }
    );
}
