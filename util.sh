#!/bin/sh

command=$(gum choose "run server" "build server")

case $command in
  "run server")
	  go run server/main.go
	  ;;
	"build server")
		tag=$(gum input --placeholder "tag")
		docker build -t ghcr.io/cloudybyte/shawty:latest -t ghcr.io/cloudybyte/shawty:$tag -f server/Dockerfile .
		;;
esac