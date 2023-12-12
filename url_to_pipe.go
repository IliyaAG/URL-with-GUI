package main

import (
	"fmt"
	"strings"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
	"fyne.io/fyne/v2/container"
)

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("YouTube to Piped Link Converter")

	// Create the entry widget to input the URL
	urlEntry := widget.NewEntry()
	urlEntry.SetPlaceHolder("Enter YouTube URL...")

	// Create the label widget to display the Piped link
	outputLabel := widget.NewLabel("")

	// Button to convert the URL
	convertButton := widget.NewButton("Convert", func() {
		pipedLink := convertToPipedLink(urlEntry.Text)
		outputLabel.SetText(pipedLink)
	})

	// Add all the widgets to a container, then set it to the window content
	myWindow.SetContent(container.NewVBox(
		urlEntry,
		convertButton,
		outputLabel,
	))

	// Use `ShowAndRun` to start the application
	myWindow.ShowAndRun()
}

func convertToPipedLink(url string) string {
	videoID := extractVideoID(url)
	if videoID == "" {
		return "Invalid URL or Pattern not found."
	}
	pipedLink := fmt.Sprintf("https://piped.video/v/%s", videoID)
	return pipedLink
}

func extractVideoID(url string) string {
	startIndex := strings.Index(url, "v=")
	if startIndex == -1 {
		return ""
	}
	videoID := url[startIndex+2:]
	endIndex := strings.Index(videoID, "&")
	if endIndex != -1 {
		videoID = videoID[:endIndex]
	}
	return videoID
}
