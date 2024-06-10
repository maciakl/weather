# weather

Show weather forecast on the command line.

⚠ Note: internet connection required.

This tool queries the [national weather service](https://weather-gov.github.io/api/general-faqs) to get current forecast for your current location.

Your location is determined by based on ip address geolocation via [ip-api.com](https://ip-api.com/) and then passed to weather service API to get your forecasting area.

Note that using a VPN may skew the results. You can use the `-location` switch to check what location you are getting the forecast for. There is currently no way to specify a location.

    Usage:
      -version
            display version number and exit
      -location
            show the location of the forecast
      -week
            show the forecast for the entire week

Use the `-week` switch to show forecast for the entire week.

Sample output:

<img width="1118" alt="screenshot" src="https://github.com/maciakl/weather/assets/189576/015463ae-4f69-49a9-8421-027632de7e63">

## Installing

 On Windows, this tool is distributed via `scoop` (see [scoop.sh](https://scoop.sh)).

 First, you need to add my bucket:

    scoop bucket add maciak https://github.com/maciakl/bucket
    scoop update

 Next simply run:
 
    scoop install weather

If you don't want to use `scoop` you can simply download the executable from the release page and extract it somewhere in your path.
