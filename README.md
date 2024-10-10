# weather

Show weather forecast on the command line.

This tool queries the [National Weather Service](https://weather-gov.github.io/api/general-faqs) to get current forecast for your current location.

⚠ Note: internet connection required.

By default, your location is determined by based on ip address geolocation via [ip-api.com](https://ip-api.com/) and then passed to weather service API to get your forecasting area.

You can override this behavior, by using the `-zip` switch and providing a US zip code to see a forecast for that specific location. In that case, your geographical location will be determined via [Zipppopotam.us](https://api.zippopotam.us/) API instead.

Note that using a VPN may skew the results. You can use the `-location` switch to check what location you are getting the forecast, and if it's incorrect, use the `-zip` switch to provide the correct zip code. Note that this is currently only supported for zip codes.

Please note that the `-location` displays your forecast area returned by the National Weather Service API based on your latitude and longtitude. The forcast arease are sometimes larger than a zip code, so you may not always see your exact town name displayed -- this is normal.

## Usage

    Usage:
      -version
            display version number and exit
      -location
            show the location of the forecast
      -week
            show the forecast for the entire week
      -zip
            show the forecast a specific zip code

Use the `-week` switch to show forecast for the entire week.

## Look and Feel

Sample output (Windows):

<img width="646" alt="screenshot3" src="https://github.com/maciakl/weather/assets/189576/ff372885-3f9e-4c1b-88c6-904b8c23fa7f">

Weekly forecast output:

<img width="1015" alt="screenshot2" src="https://github.com/maciakl/weather/assets/189576/42ebad4d-1ed8-4447-8296-4ae5be41d84f">


## Installing

Install via go:
 
    go install github.com/maciakl/weather@latest

On Windows, this tool is distributed via `scoop` (see [scoop.sh](https://scoop.sh)).

First, you need to add my bucket:

    scoop bucket add maciak https://github.com/maciakl/bucket
    scoop update

Next simply run:
 
    scoop install weather

If you don't want to use `scoop` you can simply download the executable from the release page and extract it somewhere in your path.
