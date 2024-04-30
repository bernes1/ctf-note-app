package main

import (
	"fmt"
)

type Platform struct {
	name string
}

type Artist struct {
	name string
}

type DJset struct {
	name     string
	url      string
	platform Platform
	artist   Artist
}

func main() {
	fmt.Println("type shiiii")
	addNewDJSet("shii", "soundcloud", "shhii", "https://example.com")
}

func addNewDJSet(artistName string, platformName string, djsetName string, djsetUrl string) {
	artist := Artist{name: artistName}
	platform := Platform{name: platformName}
	fmt.Println(DJset{name: djsetName, url: djsetUrl, platform: platform, artist: artist})
}
