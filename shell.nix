let
    pkgs = import <nixpkgs> {};
in
pkgs.mkShell {
    buildInputs = [ 
        pkgs.skaffold 
        pkgs.minikube 
        pkgs.kubectl 
        pkgs.go 
    ];
    shellHook = ''
        export SHELL=/usr/bin/zsh
        exec zsh
    '';
}
