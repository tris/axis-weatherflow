module github.com/tris/axis-weatherflow

go 1.24

require (
	github.com/godbus/dbus v4.1.0+incompatible
	github.com/tris/weatherflow v0.3.2
	gopkg.in/ini.v1 v1.67.0
)

require (
	github.com/stretchr/testify v1.10.0 // indirect
	nhooyr.io/websocket v1.8.17 // indirect
)

//replace github.com/tris/weatherflow => /Users/tristan/weatherflow
