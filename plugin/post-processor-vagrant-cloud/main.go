package main

import (
	"github.com/aaronbareford/packer/packer/plugin"
	"github.com/mitchellh/packer/post-processor/vagrant-cloud"
)

func main() {
	server, err := plugin.Server()
	if err != nil {
		panic(err)
	}
	server.RegisterPostProcessor(new(vagrantcloud.PostProcessor))
	server.Serve()
}
