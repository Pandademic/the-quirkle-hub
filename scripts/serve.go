package main

import (
	"log"
	"os/exec"
	"time"
)

func main() {
	go func() {
		cmd := exec.Command("webserve", "dist", "1234")

		err := cmd.Run()

		if err != nil {
			log.Fatal(err.Error())
		}
	}()

	go func() {
		for {
			cmd := exec.Command("vite", "build")
			err := cmd.Run()
			if err != nil {
				log.Println(err)
			}
		}
	}()

	for {
		time.Sleep(1 * time.Second)
	}
}
