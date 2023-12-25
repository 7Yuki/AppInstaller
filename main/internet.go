package main

import (
	"github.com/vbauerster/mpb/v8"
	"io"
	"net/http"
	"os"
)

func downloadFileProgressBar(url string, outFile *os.File, bar *mpb.Bar) error {

	// Download file
	resp, err := http.Get(url)
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
