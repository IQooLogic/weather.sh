package main

/*
 TODO
 make units selectable as param
 make ascii art out of images to represent weather conditions
 make foracast as option
*/
import (
	"fmt"
	"os"
	"strings"

	"weather.sh/geoip"
	"weather.sh/weather"
)

func main() {
	location, err := geoip.GetLocation()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Printf("Found you in: %s, %s and you are using %s", location.City,
		location.Country, location.ISP)

	wi, err := weather.GetWeather(location.Lat, location.Lng)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	var description = wi.Weather[0].Description
	var icon = ""
	if strings.Contains(description, "cloud") {
		icon = "\u2601"
	} else if strings.Contains(description, "sun") {
		icon = "\u2600"
	} else if strings.Contains(description, "snow") {
		icon = "\u2603"
	} else { // rain
		icon = "\u2602"
	}
	fmt.Printf("\nTemperature is: %s %.1f\u00b0C (%s) \u2193%.1f \u2191%.1f", icon,
		wi.Main.Temp, description, wi.Main.TempMin, wi.Main.TempMax)
}
