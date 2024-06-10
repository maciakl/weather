# weather

Show weather forecast on the command line.

This tool queries the national weather service to get current forecast for your current location which is determined by ip address geolocation.

    Usage:
      -version
            display version number and exit
      -location
            show the location of the forecast
      -week
            show the forecast for the entire week

Use the `-week` switch to show forecast for the entire week. Use `-location` to see the location for which the forecast is being fetched.

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
