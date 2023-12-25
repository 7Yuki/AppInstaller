package main

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
	"strings"
	"time"
)

const (
	videoCardzUrl = "https://videocardz.com/sections/drivers"
)

func getDriverVersions() (nvidiaVersion, amdVersion, intelVersion string, err error) {

	nvidiaVersion, err = getLatestVersionNumber(getLatestVersionNumberUrl(2))
	if err != nil {
		return "", "", "", err
	}

	amdVersion, err = getLatestVersionNumber(getLatestVersionNumberUrl(1))
	if err != nil {
		return "", "", "", err
	}

	intelVersion, err = getLatestVersionNumber(getLatestVersionNumberUrl(3))
	if err != nil {
		return "", "", "", err
	}

	return nvidiaVersion, amdVersion, intelVersion, nil
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

func getLatestVersionNumberUrl(divChild int) string {

	bodyBytes, err := httpGet(videoCardzUrl)
	if err != nil {
		return err.Error()
	}

	doc, err := goquery.NewDocumentFromReader(bytes.NewReader(bodyBytes))
	if err != nil {
		log.Fatal(err)
	}

	selectorString := fmt.Sprintf("#vc-maincontainer-related > div:nth-child(%d) > a:nth-child(2)", divChild)

	link := doc.Find(selectorString).Nodes[0].Attr[0].Val

	return link
}
