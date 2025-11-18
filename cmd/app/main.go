package main

import "github.com/kastuell/gotodoapp/internal/app"

const configsDir = "configs"

func main() {
	app.Run(configsDir)
}
