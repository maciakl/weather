# 🌤 weather

Show weather forecast on the command line.

This tool queries the [National Weather Service](https://weather-gov.github.io/api/general-faqs) to get current forecast for your current location.

⚠ Note: internet connection required.

By default, your location is determined by based on ip address geolocation via [ip-api.com](https://ip-api.com/) and then passed to weather service API to get your forecasting area.

Note that using a VPN may skew the results.

You can override this behavior, by using the `-zip` switch and providing a US zip code to see a forecast for that specific location. In that case, your geographical location will be determined via [Zipppopotam.us](https://api.zippopotam.us/) API instead.

You can use the `-location` switch to check what location of forecast. Please note that the `-location` displays your forecast area returned by the National Weather Service API based on your latitude and longtitude. The forcast areas are sometimes larger than a zip codes, so you may not always see your exact town name displayed -- this is normal.

## Usage

    Usage:
      -version
            display version number and exit
      -location
            show the location of the forecast
      -week
            show the forecast for the entire week
      -zip <zip code>
            show the forecast a specific zip code

Example:

    weather -week -zip 12345 -location

## Look and Feel

Sample output:

<img width="702" alt="screenshot2" src="https://github.com/user-attachments/assets/fa272d27-055c-45f5-b2c4-35a878eca05a">

Weekly forecast for specific location:

<img width="1062" alt="screenshot" src="https://github.com/user-attachments/assets/765bc482-c8ce-4e28-8272-21fb7d670d12">


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
