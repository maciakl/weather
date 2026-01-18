# 🌤 weather

Show weather forecast on the command line.

This tool queries the [National Weather Service](https://weather-gov.github.io/api/general-faqs) to get current forecast for your current location.

⚠ Note: internet connection required.

By default, your location is determined either by querying your devices location service, or based on ip address geolocation via [ip-api.com](https://ip-api.com/) and then passed to weather service API to get your forecasting area.

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
    -debug
            show debugging messages

Example:

    weather -week -zip 12345 -location

## Look and Feel

Sample output:

![weather1](https://github.com/user-attachments/assets/712adfcb-4ae5-4a6e-af6a-6c9fc3575756)

Weekly forecast for specific location:

![weather2](https://github.com/user-attachments/assets/251448fc-d81f-4291-a906-31e5939590eb)

## Dependencies

On macOS an optional dependency is the [CoreLocationCli](https://github.com/fulldecent/corelocationcli) tool. You can install it via brew:

```bash
brew install corelocationcli
```

This is completely optional, and `weather` will work perfectly fine without it. It will simply fall back on IP geolocation if it's not available.

⚠ Note: if you choose to use `corelocationcli` please make sure you test it first, and give it permission to access the location service.

## Installing

There are few different ways:

### Platform Independent

 Install via `go`:
 
    go install github.com/maciakl/weather@latest

### Mac & Linux

Install via [grab](https://github.com/maciakl/grab):

    grab maciakl/weather

### Windows

On Windows, this tool is distributed via `scoop` (see [scoop.sh](https://scoop.sh)).

 First, you need to add my bucket:

    scoop bucket add maciak https://github.com/maciakl/bucket
    scoop update

 Next simply run:
 
    scoop install weather

If you don't want to use `scoop` you can simply download the executable from the release page and extract it somewhere in your path.
