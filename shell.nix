{ pkgs ? import <nixpkgs> }: pkgs.mkShell {
    nativeBuildInputs = with pkgs; [
        go  
    ];

    buildInputs = with pkgs; [
        curl
    ];
}