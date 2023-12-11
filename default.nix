{ pkgs ? import <nixpkgs> {} }:

pkgs.mkShell {
  buildInputs = [
    pkgs.go
  ];
  shellHook='''
    export CGO_ENABLED=0
    export GOOS=linux
    export GOARCH=386
  ''';
}
