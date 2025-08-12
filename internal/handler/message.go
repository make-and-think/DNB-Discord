package handler

import (
	"DNB-Discord/internal/config"
	"fmt"
)

func HandleMessage() {
	fmt.Println("inhandle message check", config.Global.Discord.Token)
}
