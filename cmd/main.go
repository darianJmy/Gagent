package main

import (
	"gagent/api"
	"gagent/cmd/options"
)

func main() {
	s := options.NewOptions().
		AddFlags().
		Container().
		LogClient()

	api.RegisterRoutes(s.WsContainer)

	s.Run()

	select {}
}
