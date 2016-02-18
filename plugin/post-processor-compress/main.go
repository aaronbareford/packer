package main

import (
	"github.com/aaronbareford/packer/packer/plugin"
	"github.com/mitchellh/packer/post-processor/compress"
)

func main() {
	server, err := plugin.Server()
	if err != nil {
		panic(err)
	}
	server.RegisterPostProcessor(new(compress.PostProcessor))
	server.Serve()
}
