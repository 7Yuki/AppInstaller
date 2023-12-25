package main

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"time"
)

const (
	videoCardzUrl = "https://videocardz.com/sections/drivers"
)

func downloadLatestGraphicsDrivers(gpuBrand string) error {

	if strings.ToLower(gpuBrand) == "nvidia" {

		return downloadFile("", "", "")
	} else if strings.ToLower(gpuBrand) == "amd" {

		return downloadFile("", "", "")
	} else if strings.ToLower(gpuBrand) == "intel" {

		return downloadFile("", "", "")

	} else {
		log.Fatal("Unknown GPU brand")
	}
	return nil
}

func downloadNvidiaDriver(version string) error {

	// progress bar logic

	bar := createDynamicProgressBar("Nvidia", version)

	url := "https://us.download.nvidia.com/Windows/546.33/546.33-desktop-win10-win11-64bit-international-dch-whql.exe"

	outFile, fileErr := os.Create("/Users/smile/Downloads/546.33-desktop-win10-win11-64bit-international-dch-whql.exe")
	if fileErr != nil {
		panic(fileErr)
	}
	defer outFile.Close()

	err := downloadFileProgressBar(url, outFile, bar)
	if err != nil {
		return errors.New("download went wrong")
	}
	bar.Wait()

	return nil
}

func getDriverVersions() (nvidiaVersion, amdVersion, intelVersion string, err error) {

	nvidiaVersion, err = getLatestVersionNumber("https://videocardz.com/driver/nvidia-geforce-game-ready-546-33\n")
	if err != nil {
		return "", "", "", err
	}

	amdVersion, err = getLatestVersionNumber("https://videocardz.com/driver/amd-radeon-software-adrenalin-23-12-1")
	if err != nil {
		return "", "", "", err
	}

	intelVersion, err = getLatestVersionNumber("https://videocardz.com/driver/intel-arc-graphics-31-0-101-5084-5122")
	if err != nil {
		return "", "", "", err
	}

	return nvidiaVersion, amdVersion, intelVersion, nil
}

func downloadFile(url string, path string, filename string) error {

	// Create output file
	outputFile, err := os.Create(filepath.Join(path, filename))
	if err != nil {
		return err
	}
	defer outputFile.Close()

	// Get response
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Copy response body to file
	_, err = io.Copy(outputFile, resp.Body)
	return err
}

func getLatestVersionNumber(rawLink string) (string, error) {
	versionRegex := regexp.MustCompile(`\d+(?:[.-]\d+)*`)

	matches := versionRegex.FindStringSubmatch(rawLink)

	if len(matches) > 0 {
		return strings.ReplaceAll(matches[0], "-", "."), nil
	} else {
		return "", errors.New("no version found in link")
	}
}

type cachedResponse struct {
	url     string
	body    []byte
	fetched time.Time
}

var cache cachedResponse

func httpGet(url string) ([]byte, error) {

	if time.Since(cache.fetched) < 5*time.Minute {
		return cache.body, nil
	}

	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	bodyBytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	cache = cachedResponse{
		url:     url,
		body:    bodyBytes,
		fetched: time.Now(),
	}

	return bodyBytes, nil
}

func getLatestVersionNumberUrl(divChild int) (string, error) {

	bodyBytes, err := httpGet(videoCardzUrl)
	if err != nil {
		return "", err
	}

	doc, err := goquery.NewDocumentFromReader(bytes.NewReader(bodyBytes))
	if err != nil {
		log.Fatal(err)
	}

	selectorString := fmt.Sprintf("#vc-maincontainer-related > div:nth-child(%d) > a:nth-child(2)", divChild)

	link := doc.Find(selectorString).Nodes[0].Attr[0].Val

	fmt.Printf("link: %v\n", link)

	return link, nil
}

/*func getLatestVersionNumberUrl(divChild int) string {
	if time.Since(cache.fetched) < 5 * time.Minute {
		doc, err := goquery.NewDocumentFromReader(res.Body)
		if err != nil {
			log.Fatal(err)
		}
	}
	res, err := http.Get(videoCardzUrl)
	if err != nil {
		log.Fatal(err)
	}

	cache = cachedResponse{
		url: videoCardzUrl,
		body: res.Body.,
		fetched: time.Now(),
	}

	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	selectorString := fmt.Sprintf("#vc-maincontainer-related > div:nth-child(%d) > a:nth-child(2)", divChild)

	findings := doc.Find(selectorString).Nodes[0].Attr[0].Val

	fmt.Printf("link: %v\n", findings)

	return findings
}*/
