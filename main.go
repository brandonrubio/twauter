package main

import (
	"github.com/brandonrubio/twauter/service"
)

func main() {
	_, serviceCatalog := service.InitContainer(service.GetDependencies())
	serviceCatalog.InitServices()
	serviceCatalog.Run()
}
