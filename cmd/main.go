package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/ezhi/rresults/fetch/utils"
	"github.com/ezhi/tdt2gpx"
	"github.com/tkrajina/gpxgo/gpx"
)

// https://tracedetrail.fr/en/trace/278421
func main() {
	if len(os.Args) != 2 {
		log.Fatalf("Usage: %s <url>", os.Args[0])
	}
	download(os.Args[1])
}

func download(url string) {
	resp := utils.Get(".cache", url, false)
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Fatalf("Error: Failed to fetch URL with status code %d", resp.StatusCode)
	}

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Fatalf("Failed to parse HTML: %v", err)
	}

	doc.Find("script").Each(func(i int, s *goquery.Selection) {
		if strings.Contains(s.Text(), "dataTrace:") {
			gpxFile, err := tdt2gpx.ParseScript(s.Text())
			if err != nil {
				log.Fatalf("Failed to parse script: %v", err)
			}
			xmlBytes, err := gpxFile.ToXml(gpx.ToXmlParams{Version: "1.1", Indent: true})
			if err != nil {
				log.Fatalf("Failed to convert to XML: %v", err)
			}

			fmt.Print(string(xmlBytes))
		}
	})
}
