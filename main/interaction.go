package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"runtime"
)

func showMenu() {
	fmt.Println("Select an option:")
	fmt.Println("1. Download Intel driver")
	fmt.Println("2. Download AMD driver")
	fmt.Println("3. Download Nvidia driver")
	fmt.Println("0. Exit")
}

func clearScreen() {

	var cmd string
	if runtime.GOOS == "windows" {
		cmd = "cls"
	} else {
		cmd = "clear"
	}

	c := exec.Command(cmd)
	c.Stdout = os.Stdout
	c.Run()
}

func handleChoice(choice int) {
	nvidiaVersion, _, _, err := getDriverVersions()
	if err != nil {
		log.Fatal(err)
	}

	//clearScreen()

	switch choice {
	case 1:
		// handle case 1
	case 2:
		// handle case 2
	case 3:
		go func() {
			if downloadErr := downloadNvidiaDriver(nvidiaVersion); downloadErr != nil {
				fmt.Println(downloadErr)
			}
		}()
		fmt.Println("Download complete")

	case 0:
		fmt.Println("Exiting...")
		os.Exit(0)
	default:
		fmt.Println("Invalid choice")
	}
}
