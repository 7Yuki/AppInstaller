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

	// The owner is the second component in the path
	owner := pathComponents[1]

	// The repo name is the third component in the path
	repoName := pathComponents[2]

	return owner, repoName, nil
}

const (
	githubApiBaseURL = "https://api.github.com/"
	githubAuthToken  = "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"
)

func downloadLatestAssetRelease(repoURL string, filePath string) error {
	// Extract the owner's name and repo's name from the URL
	parts := strings.Split(repoURL, "/")
	owner := parts[len(parts)-2]
	repo := parts[len(parts)-1]

	// Construct the GitHub API URL for the repository releases
	releasesURL := fmt.Sprintf("https://api.github.com/repos/%s/%s/releases", owner, repo)

	// Create a new HTTP client
	client := &http.Client{}

	// Create a new GET request
	req, err := http.NewRequest("GET", releasesURL, nil)
	if err != nil {
		return err
	}

	// Set the authentication header if an auth token is provided

	req.Header.Set("Authorization", githubAuthToken)

	// Send the request
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Parse the response JSON to extract the URL of the latest asset
	// Assuming the response is in JSON format
	// You might need to adjust this depending on the structure of the GitHub API response
	// For example, if the response is in XML format, you would need to use an XML parser instead
	// This is just a simplified example
	// Make sure to handle error cases properly in your actual implementation
	latestAssetURL := parseLatestAssetURL(resp.Body)

	// Create a new GET request to download the latest asset
	req, err = http.NewRequest("GET", latestAssetURL, nil)
	if err != nil {
		return err
	}

	// Set the authentication header if an auth token is provided

	req.Header.Set("Authorization", githubAuthToken)

	// Send the request
	resp, err = client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Create the file to save the asset
	file, err := os.Create(filePath + string(os.PathSeparator) + "latest_asset.zip")
	if err != nil {
		return err
	}
	defer file.Close()

	// Copy the response body to the file
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
