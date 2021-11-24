package main

import "github.com/CheckmarxDev/sast-correlation-engine/engine/cmd"

func main() {
	err := cmd.Execute()

	if err != nil {
		panic(err)
	}
}
