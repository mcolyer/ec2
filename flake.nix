# https://blog.mplanchard.com/posts/installing-a-specific-version-of-a-package-with-nix.html
{
  description = "some project";

  inputs = {
   # Your preferred primary nix relesae
   nixpkgs.url = "nixpkgs/release-23.11";
   # Provides some nice helpers for multiple system compatibility
   flake-utils.url = "github:numtide/flake-utils";

	 # Some particular revision for installing fd
   nixpkgs-node.url = "github:NixOS/nixpkgs/407f8825b321617a38b86a4d9be11fd76d513da2";
  };

  outputs = { self, nixpkgs, nixpkgs-node, flake-utils }:
    # Calls the provided function for each "default system", which
    # is the standard set.
    flake-utils.lib.eachDefaultSystem
      (system:
        # instantiate the package set for the supported system, with our
        # rust overlay
        let 
					pkgs = import nixpkgs { inherit system; };
          pkgs-node = import nixpkgs-node { inherit system; };
        in
        # "unpack" the pkgs attrset into the parent namespace
        with pkgs;
        {
          devShell = mkShell {
            buildInputs = [
              pkgs-node.nodejs-16_x
            ];
          };
        });
}
