# weather

Show weather forecast on the command line.

This tool queries the [national weather service](https://weather-gov.github.io/api/general-faqs) to get current forecast for your current location.

⚠ Note: internet connection required.

Your location is determined by based on ip address geolocation via [ip-api.com](https://ip-api.com/) and then passed to weather service API to get your forecasting area.

Note that using a VPN may skew the results. You can use the `-location` switch to check what location you are getting the forecast for. There is currently no way to specify a location.

## Usage

    Usage:
      -version
            display version number and exit
      -location
            show the location of the forecast
      -week
            show the forecast for the entire week

Use the `-week` switch to show forecast for the entire week.

## Look and Feel

Sample output (Windows):

<img width="646" alt="screenshot3" src="https://github.com/maciakl/weather/assets/189576/ff372885-3f9e-4c1b-88c6-904b8c23fa7f">

Weekly forecast output:

<img width="1015" alt="screenshot2" src="https://github.com/maciakl/weather/assets/189576/42ebad4d-1ed8-4447-8296-4ae5be41d84f">


## Installing

 On Windows, this tool is distributed via `scoop` (see [scoop.sh](https://scoop.sh)).

 First, you need to add my bucket:

    scoop bucket add maciak https://github.com/maciakl/bucket
    scoop update

 Next simply run:
 
    scoop install weather

If you don't want to use `scoop` you can simply download the executable from the release page and extract it somewhere in your path.
