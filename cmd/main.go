package main

import (
	"DNB-Discord/internal/config"
	"DNB-Discord/internal/handler"
	"fmt"
)

func main() {
	fmt.Println("Hello, World!")
	config.LoadConfig("configs/.secret.toml")
	config.LoadConfig("configs/settings.toml")
	fmt.Println(config.Global.Discord.Token)
	handler.HandleMessage()
}
