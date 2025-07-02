{
  description = "Hafas go implementation, specifically for unternehmen.vbb.de/digitale-services/api/";
  inputs = {
    nixpkgs.url = "github:nixos/nixpkgs?ref=nixos-unstable";
  };
  outputs = {
    self,
    nixpkgs,
  }: let
    system = "x86_64-linux";
    pkgs = nixpkgs.legacyPackages.${system};

    openapiSpec = pkgs.fetchurl {
      url = "https://vbb-demo.demo2.hafas.cloud/api/fahrinfo/latest/api-doc";
      sha256 = "sha256-yW4BVIEKU7SEybBkNYUdmYgsQbwxJdGWKkxA6j4rYmk=";
    };

    oapi-codegen = pkgs.oapi-codegen;
  in {
    devShells.${system}.default = pkgs.mkShell {
      buildInputs = [pkgs.go oapi-codegen];

      shellHook = ''
        source ./.env
        mkdir -p vbbraw
        oapi-codegen \
          -generate "client,types" \
          -package vbbraw \
          -o vbbraw/vbb.go \
          ${openapiSpec}
        go mod tidy
      '';
    };
  };
}
