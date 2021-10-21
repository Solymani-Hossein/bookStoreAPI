package main

import "log"

func init() {
	var cfg *config.Config

	if error := config.ReadFile(cfg); error != nil {
		log.Fatalln(error)
	}
	if error := config.ReadEnv(cfg); error != nil {
		log.Fatalln(error)
	}

	config.SetConfig(cfg)
}

func main() {

}
