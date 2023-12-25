package main

import (
	"flag"
	"fmt"
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

	path := flag.String("path", "/Users/smile/Downloads", "Path to download files to")
	flag.Parse()

	for {
		clearScreen()

		showMenu()

		fmt.Print("\n> ")
		var choice int
		fmt.Scanln(&choice)

		handleChoice(*path, choice)

		if choice == 0 {
			fmt.Println("Exiting...")
			break
		}
		clearScreen()
	}

}
