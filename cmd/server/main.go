package main

import (
	"fmt"

	//"github.com/auth_service/cmd/console"
	"github.com/auth_service/internal/initialize"
)

func main() {
	//go console.Console()
	fmt.Println("Auth Service is running...")
	initialize.Run()
}
