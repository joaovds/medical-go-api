package main

import (
	"fmt"
	"log"

	"github.com/joaovds/first-go-api/internal/configs"
)

func main() {
  app := configs.GetApp()
  
  fmt.Println("Server running on port", 3333)
  log.Fatal(app.Listen(":3333"))
}

