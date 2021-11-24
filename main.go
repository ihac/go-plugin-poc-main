package main

import (
	"plugin"
	"time"
	"github.com/ihac/go-plugin-poc-main/log"
)

type API interface {
	Execute()
}

func loadAndExecute(name string) {
	plugin, err := plugin.Open(name)
	if err != nil {
		panic(err)
	}

	method, err := plugin.Lookup("BinderImpl")
	if err != nil {
		panic(err)
	}

	api, ok := method.(API)
	if !ok {
		panic("error in type assertion")
	}

	api.Execute()
}

func main() {
	log.Info("main program")
	time.Sleep(1 * time.Second)
	loadAndExecute("./plugins/plugin1.so")
	time.Sleep(1 * time.Second)
	loadAndExecute("./plugins/plugin2.so")
}
