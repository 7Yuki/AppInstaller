package main

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/vbauerster/mpb/v8"
	"github.com/vbauerster/mpb/v8/decor"
	"io"
	"math/rand"
	"strings"
	"time"
)

func makeStream(limit int) func() (int, error) {
	return func() (int, error) {
		if limit <= 0 {
			return 0, io.EOF
		}
		limit--
		return rand.Intn(1024) + 1, nil
	}
}

func determineBrandColor(brand string) *color.Color {
	var brandColor *color.Color
	switch strings.ToLower(brand) {
	case "intel":
		brandColor = color.New(color.FgHiBlue).Add(color.Bold)
	case "amd":
		brandColor = color.New(color.FgHiRed).Add(color.Bold)
	case "nvidia":
		brandColor = color.New(color.FgHiGreen).Add(color.Bold)
	}

	return brandColor
}

func createDynamicProgressBar(brand string, versionNumber string) *mpb.Bar {
	p := mpb.New(mpb.WithWidth(64))
	brandString := determineBrandColor(brand).Sprint(brand)
	// new bar with 'trigger complete event' disabled, because total is zero
	bar := p.AddBar(0,
		mpb.PrependDecorators(decor.Name(fmt.Sprintf("Downloading latest %v Driver: %v", brandString, versionNumber))),
		mpb.AppendDecorators(decor.Percentage()),
	)

	maxSleep := 100 * time.Millisecond
	read := makeStream(200)
	for {
		n, err := read()
		if err == io.EOF {
			// triggering complete event now
			bar.SetTotal(-1, true)
			break
		}
		// increment methods won't trigger complete event because bar was constructed with total = 0
		bar.IncrBy(n)
		// following call is not required, it's called to show some progress instead of an empty bar
		bar.SetTotal(bar.Current()+2048, false)
		time.Sleep(time.Duration(rand.Intn(10)+1) * maxSleep / 10)
	}

	p.Wait()

	return bar
}
