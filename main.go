package main

import (
    "os"
	"os/exec"
    "fmt"
    "flag"
    "strings"
    "strconv"
	"runtime"
    "net/http"
    "encoding/json"
    "path/filepath"

    "github.com/fatih/color"
)

const version = "0.2.0"

var debug bool

// the weather forecast struct
// the relvant data is in the properties.Periods array
// the 0 intex is the current forecast and it goes up to 13
type Forecast struct {
	Context  []any  `json:"@context"`
	Type     string `json:"type"`
	Geometry struct {
		Type        string        `json:"type"`
		Coordinates [][][]float64 `json:"coordinates"`
	} `json:"geometry"`
	Properties struct {
		Updated           string `json:"updated"`
		Units             string    `json:"units"`
		ForecastGenerator string    `json:"forecastGenerator"`
		GeneratedAt       string `json:"generatedAt"`
		UpdateTime        string `json:"updateTime"`
		ValidTimes        string `json:"validTimes"`
		Elevation         struct {
			UnitCode string  `json:"unitCode"`
			Value    float64 `json:"value"`
		} `json:"elevation"`
		Periods []struct {
			Number                     int    `json:"number"`
			Name                       string `json:"name"`
			StartTime                  string `json:"startTime"`
			EndTime                    string `json:"endTime"`
			IsDaytime                  bool   `json:"isDaytime"`
			Temperature                int    `json:"temperature"`
			TemperatureUnit            string `json:"temperatureUnit"`
			TemperatureTrend           any    `json:"temperatureTrend"`
			ProbabilityOfPrecipitation struct {
				UnitCode string `json:"unitCode"`
				Value    any    `json:"value"`
			} `json:"probabilityOfPrecipitation"`
			Dewpoint struct {
				UnitCode string  `json:"unitCode"`
				Value    float64 `json:"value"`
			} `json:"dewpoint"`
			RelativeHumidity struct {
				UnitCode string `json:"unitCode"`
				Value    int    `json:"value"`
			} `json:"relativeHumidity"`
			WindSpeed        string `json:"windSpeed"`
			WindDirection    string `json:"windDirection"`
			Icon             string `json:"icon"`
			ShortForecast    string `json:"shortForecast"`
			DetailedForecast string `json:"detailedForecast"`
		} `json:"periods"`
	} `json:"properties"`
}

func main() {

    var ver bool
    flag.BoolVar(&ver, "version", false, "display version number and exit")

	flag.BoolVar(&debug, "debug", false, "show debug messages")

    var show_location bool
    flag.BoolVar(&show_location, "location", false, "display the location of the forecast")

    var week bool
    flag.BoolVar(&week, "week", false, "show the forecast for the entire week")

    var zip string
    flag.StringVar(&zip, "zip", "none", "show the forecast for a specific US zip code")

    flag.Parse()

    // show version and exit
    if ver {
        fmt.Println(filepath.Base(os.Args[0]), "version", version)
        os.Exit(0)
    }
	
	dmsg("Debug mode enabled")
	dmsg("version " + version)


    var lat, lon float64
    // get the latitude and longitude
    if zip != "none" {
		dmsg("Getting location for zip code: " + zip)
        lat, lon = getLatLongFromZip(zip)
    } else {

		dmsg("Attempting to retreive location from OS location services")
		var err error
		lat, lon, err = getLatLongFromOS()
		if err != nil {
			// fallback to IP based location
			dmsg("Falling back to IP based location")
			lat, lon = getLatLong()
		}
    }
    dmsg(fmt.Sprintf("Latitude: %f Longitude: %f", lat, lon)) 

    forecast_url, place := getForecastInformation(lat, lon)
    forecast := getForecast(forecast_url)

	dmsg("Location: " + place)

    if show_location {
        fmt.Fprintln(color.Output, "🗺 ", color.GreenString(place))
    }

    day := 0
    if week { day = 13 }

    for i := 0; i <= day; i++ {
        printForecast(forecast, i)
    }

}

func dmsg(msg string) {
	if debug {
		fmt.Fprintln(os.Stderr, color.BlackString(msg))
	}
}

// get the latitude and longitude from
// the user's IP address via ip-api.com
func getLatLong() (float64, float64) {

    dmsg("Getting location from ip-api.com")
    // get http://ip-api.com/json/ and parse
    resp, err := http.Get("http://ip-api.com/json/")
    if err != nil {
        fmt.Println("Error getting location information:", err)
        os.Exit(1)
    }
    defer resp.Body.Close()

    // the location struct
    type Location struct {
        Status      string  `json:"status"`
        Country     string  `json:"country"`
        CountryCode string  `json:"countryCode"`
        Region      string  `json:"region"`
        RegionName  string  `json:"regionName"`
        City        string  `json:"city"`
        Zip         string  `json:"zip"`
        Lat         float64 `json:"lat"`
        Lon         float64 `json:"lon"`
        Timezone    string  `json:"timezone"`
        Isp         string  `json:"isp"`
        Org         string  `json:"org"`
        As          string  `json:"as"`
        Query       string  `json:"query"`
    }

    var loc Location

    // decode the JSON response 
    err = json.NewDecoder(resp.Body).Decode(&loc)
    if err != nil {
        fmt.Println("Error decoding location information:", err)
        os.Exit(1)
    }

    // get the latitude and longitude
    lat := loc.Lat
    lon := loc.Lon

    //return lat, lon
    return lat, lon
}

func getLatLongFromOS() (float64, float64, error) {

	os := runtime.GOOS
	
	if strings.Contains(strings.ToLower(os), "windows") {
		dmsg("Windows detected")
		return getLatLongWindows()
	} else if strings.Contains(strings.ToLower(os), "darwin") {
		dmsg("macOS detected")
		return getLatLongMac()
	} else if strings.Contains(strings.ToLower(os), "linux") {
		dmsg("Linux detected")
		return getLatLongLinux()
	} else {
		dmsg("Unsupported OS: " + os)
		return 0, 0, fmt.Errorf("unsupported OS")
	}

}

func getLatLongWindows() (float64, float64, error) {

	// quick and dirty PowerShell script to get location
	pscode := `
Add-Type -AssemblyName System.Device
$watcher = New-Object System.Device.Location.GeoCoordinateWatcher
$watcher.Start()
while (($watcher.Status -ne "Ready") -and ($watcher.Permission -ne "Denied")) { Start-Sleep -Milliseconds 100 }
$coord = $watcher.Position.Location
Write-Output "$($coord.Latitude),$($coord.Longitude)"
`
	dmsg("Running a powershell script to get location from Windows location services")
	out, err := exec.Command("powershell", "-Command", pscode).Output()
	if err != nil {
		dmsg("Error running PowerShell: " + err.Error())
		dmsg("Make sure you allow location services access for PowerShell")
		return 0, 0, err
	}

	parts := strings.Split(strings.TrimSpace(string(out)), ",")	
	if len(parts) != 2 {
		return 0, 0, fmt.Errorf("invalid output from PowerShell")
	}

	lat, err := strconv.ParseFloat(parts[0], 64)
	if err != nil {
		return 0, 0, err
	}

	lon, err := strconv.ParseFloat(parts[1], 64)
	if err != nil {
		return 0, 0, err
	}

	return lat, lon, nil

}

func getLatLongMac() (float64, float64, error) {

	dmsg("Checking if CoreLocationCLI is installed")
	coreLocationCLIPath := "/opt/homebrew/bin/CoreLocationCLI"

	// check if CoreLocationCLI is installed
	_, err := os.Stat(coreLocationCLIPath)
	if os.IsNotExist(err) {
		dmsg("CoreLocationCLI not found in " + coreLocationCLIPath)
		dmsg("Run 'brew install corelocationcli' to install it")
		return 0, 0, fmt.Errorf("CoreLocationCLI is not installed")
	}

	// run CoreLocationCLI on the command line, parse the output to get lat and long
	out, err := exec.Command(coreLocationCLIPath).Output()
	if err != nil {
		dmsg("Error running CoreLocationCLI: " + err.Error())
		dmsg("Make sure you allow location services access for the CoreLocationCLI app")
		return 0, 0, err
	}

	// parse the output
	parts := strings.Split(strings.TrimSpace(string(out)), " ")
	if len(parts) != 2 {
		return 0, 0, fmt.Errorf("invalid output from CoreLocationCLI")
	}

	lat, err := strconv.ParseFloat(parts[0], 64)
	if err != nil {
		return 0, 0, err
	}

	lon, err := strconv.ParseFloat(parts[1], 64)
	if err != nil {
		return 0, 0, err
	}

	return lat, lon, nil

}

func getLatLongLinux() (float64, float64, error) {
	dmsg("Linux location services not implemented yet...")
	return 0, 0, fmt.Errorf("not implemented")
}


// get latitude and longitude based on user provided US zip cdode
func getLatLongFromZip(zip string) (float64, float64) {
    
		dmsg("Fetching location from zippopotam.us")

        // get http://api.zippopotam.us/us/ and parse
        resp, err := http.Get("http://api.zippopotam.us/us/" + zip)
        if err != nil {
            fmt.Println("Error getting location information:", err)
            os.Exit(1)
        }
        defer resp.Body.Close()
    
        // the location struct
        type Location struct {
            PostCode            string `json:"post code"`
            Country             string `json:"country"`
            CountryAbbreviation string `json:"country abbreviation"`
            Places              []struct {
                PlaceName         string `json:"place name"`
                Longitude         string `json:"longitude"`
                State             string `json:"state"`
                StateAbbreviation string `json:"state abbreviation"`
                Latitude          string `json:"latitude"`
            } `json:"places"`
        }
    
        var loc Location
    
        // decode the JSON response
        err = json.NewDecoder(resp.Body).Decode(&loc)
        if err != nil {
            fmt.Println("Error decoding location information:", err)
            os.Exit(1)
        }

        if len(loc.Places) == 0 {
            fmt.Println("Error: Invalid zip code")
            os.Exit(1)
        }
    
        // get the latitude and longitude
        lat, laterr := strconv.ParseFloat(loc.Places[0].Latitude, 64)
        lon, lonerr := strconv.ParseFloat(loc.Places[0].Longitude, 64)

        if laterr != nil || lonerr != nil {
            fmt.Println("Error converting zip code to location data")
            os.Exit(1)
        }
    
        return lat, lon
    
}

// get weather.gov forecast point from latitude and longitude
// return the forecast url and location
func getForecastInformation(lat, lon float64) (string, string) {

	dmsg("Getting forecast information from weather.gov")

    lat_str := fmt.Sprintf("%f", lat)
    lon_str := fmt.Sprintf("%f", lon)


    resp, err := http.Get("https://api.weather.gov/points/" + lat_str + "," + lon_str)
    if err != nil {
        fmt.Println("Error getting forecast point:", err)
        os.Exit(1)
    }

    // the struct
    type WeatherResponse struct {
        Context  []any  `json:"@context"`
        ID       string `json:"id"`
        Type     string `json:"type"`
        Geometry struct {
            Type        string    `json:"type"`
            Coordinates []float64 `json:"coordinates"`
        } `json:"geometry"`
        Properties struct {
            ID                  string `json:"@id"`
            Type                string `json:"@type"`
            Cwa                 string `json:"cwa"`
            ForecastOffice      string `json:"forecastOffice"`
            GridID              string `json:"gridId"`
            GridX               int    `json:"gridX"`
            GridY               int    `json:"gridY"`
            Forecast            string `json:"forecast"`
            ForecastHourly      string `json:"forecastHourly"`
            ForecastGridData    string `json:"forecastGridData"`
            ObservationStations string `json:"observationStations"`
            RelativeLocation    struct {
                Type     string `json:"type"`
                Geometry struct {
                    Type        string    `json:"type"`
                    Coordinates []float64 `json:"coordinates"`
                } `json:"geometry"`
                Properties struct {
                    City     string `json:"city"`
                    State    string `json:"state"`
                    Distance struct {
                        UnitCode string  `json:"unitCode"`
                        Value    float64 `json:"value"`
                    } `json:"distance"`
                    Bearing struct {
                        UnitCode string `json:"unitCode"`
                        Value    int    `json:"value"`
                    } `json:"bearing"`
                } `json:"properties"`
            } `json:"relativeLocation"`
            ForecastZone    string `json:"forecastZone"`
            County          string `json:"county"`
            FireWeatherZone string `json:"fireWeatherZone"`
            TimeZone        string `json:"timeZone"`
            RadarStation    string `json:"radarStation"`
        } `json:"properties"`
    }

    var weather WeatherResponse

    // decode the JSON response
    err = json.NewDecoder(resp.Body).Decode(&weather)
    if err != nil {
        fmt.Println("Error decoding weather information:", err)
        os.Exit(1)
    }

    // get the forecast URL
    forecast_url := weather.Properties.Forecast

    // get location (city, state)
    location := weather.Properties.RelativeLocation.Properties.City + ", " + weather.Properties.RelativeLocation.Properties.State

    return forecast_url, location

}


// get forecast from forecast URL
func getForecast(forecast_url string) Forecast {

		dmsg("Getting forecast from " + forecast_url)	
    
        resp, err := http.Get(forecast_url)
        if err != nil {
            fmt.Println("Error getting forecast:", err)
            os.Exit(1)
        }
    
        var forecast Forecast
    
        // decode the JSON response
        err = json.NewDecoder(resp.Body).Decode(&forecast)
        if err != nil {
            fmt.Println("Error decoding forecast information:", err)
            os.Exit(1)
        }
    
        return forecast
}


// print out the forecast for the given day, 0 being the current day
func printForecast(forecast Forecast, day int) {

    // convert the temperature to a string
    temperature := getTempString(forecast.Properties.Periods[day].Temperature, forecast.Properties.Periods[day].TemperatureUnit)

    forecast_for := forecast.Properties.Periods[day].Name

    detailed_forecast := forecast.Properties.Periods[day].DetailedForecast
    short := forecast.Properties.Periods[day].ShortForecast

    icon := getIcon(short)

    fmt.Fprintf(color.Output, "%-16s %s %s %s\n", forecast_for+":", icon, temperature, detailed_forecast)
}

// return an icon based on short forecast string
func getIcon(short string) string {


    // map of short forecast strings to icons
    icons := map[string]string{
        "Sunny": "☀️",
        "Clear": "☀️",
        "Mostly Clear": "🌤️",
        "Mostly Sunny": "🌤️",
        "Partly Sunny": "⛅",
        "Partly Cloudy": "⛅",
        "Mostly Cloudy": "☁️",
        "Cloudy": "☁️",
        "Rain": "🌧️",
        "Showers": "🌧️",
        "Thunderstorms": "⛈️",
        "Snow": "❄️",
        "Fog": "🌫️",
        "Haze": "🌫️",
        "Mist": "🌫️",
        "Smoke": "🌫️",
        "Squalls": "💨",
        "Windy": "💨",
        "Tornado": "🌪️",
        "Hurricane": "🌀",
        "Tropical Storm": "🌀",
        "Blizzard": "❄️",
        "Ice": "❄️",
        "Freezing Rain": "🌨",
        "Freezing Drizzle": "🌨",
        "Drizzle": "🌧️",
        "Freezing Fog": "🌫️",
        "Heavy Rain": "🌧️",
        "Heavy Snow": "❄️",
        "Heavy Thunderstorms": "⛈️",
        "Heavy Showers": "🌧️",
        "Heavy Drizzle": "🌧️",
        "Heavy Freezing Rain": "🌨",
    }

    // loop through the map and return the icon
    for k, v := range icons {
        if strings.Contains(short, k) {
            return v
        }
    }

    // if nothing matches return an generic icon
    return "🌡️"

}


// formats the temperature string
func getTempString(temp int, unit string) string {
    
    temp_str := fmt.Sprintf("%dº%s", temp, unit)

    if temp >= 90 {
        return color.RedString(temp_str)
    } else if temp >= 80 {
        return color.YellowString(temp_str)
    } else if temp <= 50 {
        return color.CyanString(temp_str)
    } else if temp <= 32 {
        return color.BlueString(temp_str)
    } else {
        return temp_str
    }
}
