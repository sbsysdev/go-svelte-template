FROM mcr.microsoft.com/devcontainers/go:1-1.24-bookworm AS dev

RUN git config --global core.autocrlf input
RUN git config --global --add safe.directory /workspaces/go-svelte-template
RUN apt update && apt upgrade -y

RUN sudo chown -R vscode $GOPATH
RUN go install github.com/air-verse/air@latest
