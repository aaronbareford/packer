package main

import (
	"github.com/aaronbareford/packer/packer/plugin"
	"github.com/mitchellh/packer/post-processor/atlas"
)

func main() {
	server, err := plugin.Server()
	if err != nil {
		panic(err)
	}
	server.RegisterPostProcessor(new(atlas.PostProcessor))
	server.Serve()
}
