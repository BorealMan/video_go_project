package main

import (
	"log"

	"github.com/borealman/video_go_project/internal/server"
)

func main() {
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}
