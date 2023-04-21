# WeatherFlow Overlay for Axis cameras

This ACAP app for Axis cameras will create a text overlay with live-updating
rapid wind readings from a single
[WeatherFlow Tempest](https://weatherflow.com/tempest-home-weather-system/)
weather station, using
[WeatherFlow's WebSockets API](https://weatherflow.github.io/Tempest/api/ws.html).

## Requirements

- [WeatherFlow Tempest](https://weatherflow.com/tempest-home-weather-system/)
  weather station (your own or someone else's)
- The device ID for said station (not the station ID!)
- A [Personal Access Token](https://weatherflow.github.io/Tempest/api/) from
  the WeatherFlow account associated with the station
- [Axis](https://www.axis.com/) network camera (ARTPEC-6 or newer)

## Build

```bash
make all
```

## Install

1. Navigate to http://insert-your-camera-ip/camera/index.html#/apps
2. Click "Add app"
3. Upload the `.eap` file corresponding to your camera architecture.  (ARTPEC-6 and ARTPEC-7 are armv7hf; ARTPEC-8 is aarch64.)
4. Click the three dots in the corner of the app card to access Settings.
5. Configure device ID, display name, position and token.
6. Click Save.
7. Click the Start toggle to start the app.
