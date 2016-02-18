package main

import (
	"github.com/mitchellh/packer/builder/null"
	"github.com/aaronbareford/packer/packer/plugin"
)

func main() {
	server, err := plugin.Server()
	if err != nil {
		panic(err)
	}
	server.RegisterBuilder(new(null.Builder))
	server.Serve()
}
