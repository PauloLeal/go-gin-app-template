package main

import (
	"PauloLeal/go-gin-app-template/app"
	"flag"
	"os"
	"strconv"
)

func getDefaultPort() int {
	defaultEnvPort := 8080
	envPort := os.Getenv("SERVER_PORT")
	if len(envPort) > 0 {
		p, err := strconv.Atoi(envPort)
		if err == nil {
			defaultEnvPort = p
		}
	}

	return defaultEnvPort
}

func main() {
	var serverPort = flag.Int("p", getDefaultPort(), "web server http port")
	flag.Parse()

	_ = app.App().RunServer(*serverPort)
}
