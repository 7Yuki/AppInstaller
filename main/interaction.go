package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"runtime"
	"sync"
)

func showMenu() {
	fmt.Println("Select an option:")
	fmt.Println("1. Download Intel driver")
	fmt.Println("2. Download AMD driver")
	fmt.Println("3. Download Nvidia driver")
	fmt.Println("4. Download MSI Afterburner")
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

func handleChoice(path string, choice int) {
	nvidiaVersion, amdVersion, intelVersion, _ := getDriverVersions()
	var wg sync.WaitGroup

	if choice == 1 {
		clearScreen()
		wg.Add(1)
		go func() {
			defer wg.Done()
			if downloadErr := downloadDriver(
				intelVersion,
				"intel",
				path,
			); downloadErr != nil {
				fmt.Println(downloadErr)
			}
		}()
		wg.Wait()

	} else if choice == 2 {
		clearScreen()
		wg.Add(1)
		go func() {
			defer wg.Done()
			if downloadErr := downloadDriver(
				amdVersion,
				"amd",
				path,
			); downloadErr != nil {
				fmt.Println(downloadErr)
			}
		}()
		wg.Wait()

	} else if choice == 3 {
		clearScreen()
		wg.Add(1)
		go func() {
			defer wg.Done()
			if downloadErr := downloadDriver(
				nvidiaVersion,
				"nvidia",
				path,
			); downloadErr != nil {
				fmt.Println(downloadErr)
			}
		}()
		wg.Wait()
		fmt.Println("Done!")
	} else if choice == 4 {
		clearScreen()
		wg.Add(1)
		go func() {
			defer wg.Done()
			downloadErr := downloadAppWithJS("https://www.msi.com/Landing/afterburner/graphics-cards", "body > main > section.kv > div > div.kv__btn > a")
			if downloadErr != nil {
				log.Fatal(downloadErr)
				return
			}

		}()
		wg.Wait()
		fmt.Println("Done!")
	}
}
