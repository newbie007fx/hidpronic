package main

import (
	"hidroponic/cmd/hidroponic/console"
	consoleService "hidroponic/internal/platform/console"
)

func main() {
	cs := consoleService.New()
	cs.Setup()

	console.InitApp(cs)

	cs.Run()
}
