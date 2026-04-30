module github.com/ezhi/tdt2gpx/cmd

go 1.23.3

replace github.com/ezhi/tdt2gpx => ../

replace github.com/ezhi/rresults => ../../rresults

require (
	github.com/PuerkitoBio/goquery v1.10.1
	github.com/ezhi/rresults v0.0.0-00010101000000-000000000000
	github.com/ezhi/tdt2gpx v0.0.0-00010101000000-000000000000
	github.com/tkrajina/gpxgo v1.4.0
)

require (
	github.com/andybalholm/cascadia v1.3.3 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/gofrs/flock v0.12.1 // indirect
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mattn/go-isatty v0.0.20 // indirect
	github.com/rs/zerolog v1.33.0 // indirect
	golang.org/x/exp v0.0.0-20241108190413-2d47ceb2692f // indirect
	golang.org/x/net v0.39.0 // indirect
	golang.org/x/sys v0.32.0 // indirect
	golang.org/x/text v0.24.0 // indirect
)
