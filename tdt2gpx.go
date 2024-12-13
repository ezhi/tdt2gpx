package tdt2gpx

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"math"
	"strings"

	"github.com/tkrajina/gpxgo/gpx"
)

type Point struct {
	Lat float64 `json:"lat"`
	Lon float64 `json:"lon"`
	X   float64 `json:"x"`
	Y   float64 `json:"y"`
}

type PI struct {
	Name   string   `json:"infobulleTitre"`
	Text   string   `json:"infobulleText"`
	Labels []string `json:"labelsDiv"`
	Abs    string   `json:"abs"`
	Ord    string   `json:"ord"`
	X      string   `json:"x"`
	Y      string   `json:"y"`
	Type   string   `json:"type"`
}

func ParseScript(script string) (gpx.GPX, error) {
	if !strings.Contains(script, "dataTrace:") {
		return gpx.GPX{}, fmt.Errorf("No dataTrace found")
	}

	lines := strings.Split(script, "\n")
	gpxSegment := gpx.GPXTrackSegment{}
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line_, ok := strings.CutPrefix(line, `dataPi:"`); ok {
			var pis []PI
			err := parseScriptLine(line_, &pis)
			if err != nil {
				return gpx.GPX{}, fmt.Errorf("Failed to unmarshal PIs: %v", err)
			}
		} else if line_, ok := strings.CutPrefix(line, `geometry:"`); ok {
			var points []Point
			err := parseScriptLine(line_, &points)
			if err != nil {
				return gpx.GPX{}, fmt.Errorf("Failed to unmarshal geometry: %v", err)
			}
			gpxPoints := make([]gpx.GPXPoint, len(points))
			for i, point := range points {
				lat, lon := mercatorToLatLon(point.Lon, point.Lat)
				elevation := gpx.NullableFloat64{}
				elevation.SetValue(point.Y)
				gpxPoints[i] = gpx.GPXPoint{
					Point: gpx.Point{
						Latitude:  lat,
						Longitude: lon,
						Elevation: elevation,
					},
				}
				gpxPoints[i].Extensions.Nodes = append(gpxPoints[i].Extensions.Nodes,
					gpx.ExtensionNode{XMLName: xml.Name{Local: "distance"}, Data: fmt.Sprintf("%f", point.X*1000)},
				)
			}
			gpxSegment.Points = gpxPoints
		}
	}

	return gpx.GPX{
		Tracks: []gpx.GPXTrack{{Segments: []gpx.GPXTrackSegment{gpxSegment}}},
	}, nil
}

func parseScriptLine(line string, data any) error {
	line = strings.ReplaceAll(line, `\"`, `"`)
	line = strings.ReplaceAll(line, `\\`, `\`)
	line, _ = strings.CutSuffix(line, `",`)
	err := json.Unmarshal([]byte(line), data)
	if err != nil {
		return fmt.Errorf("Failed to unmarshal data: %v", err)
	}
	return nil
}

const earthRadius = 6378137.0

func mercatorToLatLon(x, y float64) (float64, float64) {
	lon := (x / earthRadius) * (180 / math.Pi)
	lat := (math.Atan(math.Sinh(y / earthRadius))) * (180 / math.Pi)
	return lat, lon
}
