package main

import (
	"github.com/aaronbareford/packer/packer/plugin"
	"github.com/mitchellh/packer/post-processor/docker-import"
)

func main() {
	server, err := plugin.Server()
	if err != nil {
		panic(err)
	}
	server.RegisterPostProcessor(new(dockerimport.PostProcessor))
	server.Serve()
}
