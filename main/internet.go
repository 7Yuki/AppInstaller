package main

import (
	"errors"
	"fmt"
	"github.com/vbauerster/mpb/v8"
	"io"
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
