package main

import (
	"fmt"
	"sync"
)

/*
TODO:
- Add function that downloads files from a website or link (these are usually optimizer applications)
- Add function that will install applications via chocolatey
- Add function that allows the user to add custom applications to install (indexed from chocolatey's packages)
- Add base applications to install like good browsers, drivers, etc.
*/

/*
 url for NVCleaninstall: https://us5-dl.techpowerup.com/files/5NXZG0gGbMCH6RZG_rpV-A/1702545789/NVCleanstall_1.16.0.exe

*/

var applications = map[string]application{

	"nvidiaProfileInspector": {
		"nvidiaProfileInspector",
		"https://github.com/Orbmu2k/nvidiaProfileInspector",
		"",
		true,
	},
}

// Downloads a file from URL and updates progress bar

/*func main() {
	wg := new(sync.WaitGroup)
	wg.Add(1)
	for {
		//clearScreen()
		fmt.Println("Here!")
		showMenu()

		// get choice

		var choiceVal int
		fmt.Print("> ")
		fmt.Scan(&choiceVal)

		go handleChoice(choiceVal, wg)
		wg.Wait()

		//clearScreen()
		showMenu() // show menu again after choice finishes

	}
}*/

func main() {
	for {
		clearScreen()

		showMenu()

		var choice int
		fmt.Scanln(&choice)

		var wg sync.WaitGroup

		if choice == 1 {
			clearScreen()
			wg.Add(1)
			go func() {
				defer wg.Done()
				if downloadErr := downloadNvidiaDriver("546.33"); downloadErr != nil {
					fmt.Println("INTEL DRIVER DONE!")
					fmt.Println(downloadErr)
				}
			}()
			wg.Wait()

		} else if choice == 2 {
			clearScreen()
			wg.Add(1)
			go func() {
				defer wg.Done()
				if downloadErr := downloadNvidiaDriver("546.33"); downloadErr != nil {
					fmt.Println("AMD DRIVER DONE!")
					fmt.Println(downloadErr)
				}
			}()
			wg.Wait()

		} else if choice == 3 {
			clearScreen()
			wg.Add(1)
			go func() {
				defer wg.Done()
				if downloadErr := downloadNvidiaDriver("546.33"); downloadErr != nil {
					fmt.Println(downloadErr)
				}
			}()
			wg.Wait()
			fmt.Println("Done!")
		} else if choice == 0 {
			fmt.Println("Exiting...")
			break
		}
		clearScreen()
	}

}
