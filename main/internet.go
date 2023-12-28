package main

import (
	"errors"
	"fmt"
	"github.com/geziyor/geziyor"
	"github.com/geziyor/geziyor/client"
	"github.com/vbauerster/mpb/v8"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

func downloadFileProgressBar(url string, outFile *os.File, bar *mpb.Bar) error {

	var resp *http.Response
	client := &http.Client{}
	req, reqErr := http.NewRequest("GET", url, nil)

	if reqErr != nil {
		return reqErr
	}

	if strings.Contains(url, "amd.com") {
		req.Header.Set("Referer", "https://www.amd.com/")
	}

	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	buffer := make([]byte, 1024)
	for {
		n, bufferErr := resp.Body.Read(buffer)
		if n > 0 {
			bar.IncrBy(n)

			if _, fileErr := outFile.Write(buffer[:n]); err != nil {
				return fileErr
			}
		}
		if bufferErr == io.EOF {
			break
		}
	}
	return err
}

func downloadDriver(version string, brand string, path string) error {

	bar := createDynamicProgressBar(strings.ToUpper(brand), version)

	fileName := fmt.Sprintf("%v/%v-%v-desktop-win10-win11-64bit-graphics-drivers.exe", path, brand, version)
	var url string

	if strings.ToLower(brand) == "nvidia" {
		url = fmt.Sprintf("https://us.download.nvidia.com/Windows/%v/%[1]v-desktop-win10-win11-64bit-international-dch-whql.exe", version)
	} else if strings.ToLower(brand) == "amd" {
		url = fmt.Sprintf("https://drivers.amd.com/drivers/whql-amd-software-adrenalin-edition-%v-win10-win11-dec5-rdna.exe", version)
	} else if strings.ToLower(brand) == "intel" {
		//todo: implement intel driver download
	} else {
		url = "unknown brand"
	}

	outFile, fileErr := os.Create(fileName)
	if fileErr != nil {
		panic(fileErr)
	}
	defer outFile.Close()

	fmt.Printf("Downloading to: %v via link: %v\n", path, url)

	err := downloadFileProgressBar(url, outFile, bar)
	if err != nil {
		return errors.New("download went wrong")
	}
	bar.Wait()

	return nil
}

func downloadApp(site string, selector string) error {

	geziyor.NewGeziyor(&geziyor.Options{
		StartURLs: []string{site},
		ParseFunc: func(g *geziyor.Geziyor, r *client.Response) {
			attr := r.HTMLDoc.Find(selector)
			value, exists := attr.Attr("href")
			if exists {
				log.Printf("attr: %v", value)
			}
		},
	}).Start()

	return nil
}

func downloadAppWithJS(site string, selector string) error {

	attrErr := errors.New("")
	downloadUrl := ""
	bar := createDynamicProgressBarf(fmt.Sprintf("Downloading file from %s", site))

	geziyor.NewGeziyor(&geziyor.Options{
		LogDisabled: true,
		StartRequestsFunc: func(g *geziyor.Geziyor) {
			g.GetRendered(site, g.Opt.ParseFunc)
		},
		ParseFunc: func(g *geziyor.Geziyor, r *client.Response) {
			attr := r.HTMLDoc.Find(selector)
			value, exists := attr.Attr("href")
			if exists {
				downloadUrl = value
				log.Printf("link: %s", value)
				return
			}
			attrErr = fmt.Errorf("couldnt find value from: %s", selector)
		},
	}).Start()

	if attrErr.Error() != "" {
		return attrErr
	}

	if downloadUrl == "" {
		return errors.New("download url blank")
	}

	fileName := fmt.Sprintf("%v/msiafterburnersetup.zip", "/Users/smile/Downloads")

	outFile, fileErr := os.Create(fileName)
	log.Println("created file")
	if fileErr != nil {
		panic(fileErr)
	}
	defer outFile.Close()

	downloadFileErr := downloadFileProgressBar(downloadUrl, outFile, bar)
	if downloadFileErr != nil {
		return downloadFileErr
	}
	log.Println("downloaded content to file")
	return nil
}
