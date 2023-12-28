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

var applications = map[string]application{

	"steam": {
		"Steam",
		"https://cdn.akamai.steamstatic.com/client/installer/SteamSetup.exe",
		"SteamSetup.exe",
		false,
	},
	"npi": {
		"nvidiaProfileInspector",
		"https://github.com/Orbmu2k/nvidiaprofileinspector/releases/download/2.4.0.4/package.zip",
		"nvidiaProfileInspector.zip",
		false,
	},
	"discord": {
		"Discord",
		"https://discord.com/api/downloads/distributions/app/installers/latest?channel=stable&platform=win&arch=x86",
		"DiscordSetup.exe",
		false,
	},
	"vencord": {
		"Vencord",
		"https://github.com/Vencord/Installer/releases/latest/download/VencordInstaller.exe",
		"VencordInstaller.exe",
		false,
	},
	"betterdiscord": {
		"Better Discord",
		"https://github.com/BetterDiscord/Installer/releases/latest/download/BetterDiscord-Windows.exe",
		"BetterDiscordSetup.exe",
		false,
	},
}

/*
todo:
- add a way to add custom applications to install
- add downloads for:
	- 7zip/winrar (goquery) 7zip link: https://7-zip.org/download.html
	- msiafterburner (goquery)
    - intel drivers (goquery)
    - nvcleaninstall (goquery)
    - DDU
    - ungoogled chromium, brave, firefox (ungoogled chromium requires goquery)
    - scewin automation (is this possible????)
    - steam
    - discord
	- vencord/betterdiscord
    - apps for: media viewing. videos: mpc-hc images: imageglass
*/

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
