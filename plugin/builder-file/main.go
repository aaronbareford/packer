package main

import (
	"github.com/mitchellh/packer/builder/file"
	"github.com/aaronbareford/packer/packer/plugin"
)

func main() {
	server, err := plugin.Server()
	if err != nil {
		panic(err)
	}
	server.RegisterBuilder(new(file.Builder))
	server.Serve()
}
