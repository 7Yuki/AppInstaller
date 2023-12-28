package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
)

type application struct {
	name            string
	downloadURL     string
	applicationName string
	isGithub        bool
}

func parseGitHubURL(githubURL string) (string, string, error) {
	u, err := url.Parse(githubURL)
	if err != nil {
		return "", "", err
	}

	pathComponents := strings.Split(u.Path, "/")

	owner := pathComponents[1]

	repoName := pathComponents[2]

	return owner, repoName, nil
}

const (
	githubApiBaseURL = "https://api.github.com/"
	githubAuthToken  = "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"
)

func downloadLatestAssetRelease(repoURL string, filePath string) error {

	parts := strings.Split(repoURL, "/")
	owner := parts[len(parts)-2]
	repo := parts[len(parts)-1]

	releasesURL := fmt.Sprintf("https://api.github.com/repos/%s/%s/releases", owner, repo)

	client := &http.Client{}

	req, err := http.NewRequest("GET", releasesURL, nil)
	if err != nil {
		return err
	}

	req.Header.Set("Authorization", githubAuthToken)

	// Send the request
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	latestAssetURL := parseLatestAssetURL(resp.Body)

	// Create a new GET request to download the latest asset
	req, err = http.NewRequest("GET", latestAssetURL, nil)
	if err != nil {
		return err
	}

	req.Header.Set("Authorization", githubAuthToken)

	resp, err = client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	file, err := os.Create(filePath + string(os.PathSeparator) + "latest_asset.zip")
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = io.Copy(file, resp.Body)
	if err != nil {
		return err
	}

	fmt.Println("Latest asset downloaded successfully")
	return nil
}

func parseLatestAssetURL(body io.Reader) string {
	// Implement your logic to parse the response body and extract the URL of the latest asset
	// This can involve parsing JSON, XML, or any other format the API returns
	// Again, this is just a simplified example
	// Make sure to handle error cases properly in your actual implementation
	// Return the URL of the latest asset
	return "https://example.com/latest_asset.zip"
}

func downloadApplicationFromLink(app *application, filePath string) error {
	log.Printf("Downloading %s from %s as '%s'", app.name, app.downloadURL, app.applicationName)

	response, err := http.Get(app.downloadURL)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	file, err := os.Create(filePath + string(os.PathSeparator) + app.applicationName)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = io.Copy(file, response.Body)
	if err != nil {
		return err
	}

	return nil
}

func downloadApplication(app *application, filePath string) error {
	if app.isGithub {
		return downloadLatestAssetRelease(app.downloadURL, filePath)
	} else {
		return downloadApplicationFromLink(app, filePath)
	}
}
