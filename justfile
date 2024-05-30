# alias b := build
alias r := run

default:
  @just --list

install-deps:
  go mod tidy

install-cobra-cli:
  go install github.com/spf13/cobra-cli@latest
  export PATH=$PATH:$GOPATH

run: install-deps
  go run .

clean:
  rm -rf bin/
  sudo rm -rf /usr/libexec/docker/cli-plugins/docker-apps

build: install-deps clean
  go build -o bin/docker-apps

install: build
  sudo cp bin/docker-apps /usr/libexec/docker/cli-plugins
  sudo chmod +x /usr/libexec/docker/cli-plugins/docker-apps