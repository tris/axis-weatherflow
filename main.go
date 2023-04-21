package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/tris/axis-weatherflow/axis"
	"github.com/tris/weatherflow"
	// "github.com/wcharczuk/go-chart"
)

var (
	apiToken    string
	deviceID    int
	displayName string
	position    string
	overlay     *axis.Overlay
)

func windToString(speed float64, degrees int) string {
	speedMph := speed * 2.23694
	cardinal := degreesToCardinal(float64(degrees))

	msg := fmt.Sprintf("%s: %3s %2.0fmph", displayName, cardinal, speedMph)
	if speedMph < 0.1 {
		msg = fmt.Sprintf("%s: calm", displayName)
	}

	return msg
}

func handleMessage(msg weatherflow.Message) {
	switch m := msg.(type) {
	case *weatherflow.MessageObsSt:
		fmt.Printf("Observation: %.0fg%.0f mph @ %d deg\n", m.Obs[0].WindAvg*2.23694, m.Obs[0].WindGust*2.23694, m.Obs[0].WindDirection)
	case *weatherflow.MessageRapidWind:
		fmt.Printf("Rapid wind: %.0f mph @ %d deg\n", m.Ob.WindSpeed*2.23694, m.Ob.WindDirection)
		err := overlay.Set(axis.OverlayProperties{
			Text: stringPtr(windToString(m.Ob.WindSpeed, m.Ob.WindDirection)),
		})
		if err != nil {
			fmt.Printf("Error setting overlay: %v\n", err)
		}
	}
	return
}

func main() {
	apiToken, deviceID, displayName, position = readConfig("conf/weatherflow.conf")

	fmt.Printf("Using device ID: %d\n", deviceID)

	client, err := weatherflow.NewClient(apiToken, []int{deviceID}, log.Printf)
	if err != nil {
		log.Fatalf("Error creating WeatherFlow client: %v", err)
	}

	overlay, err = axis.NewOverlay(axis.OverlayProperties{
		Position: stringPtr(position),
	})
	if err != nil {
		log.Fatalf("Error creating overlay: %v", err)
	}

	client.Start(handleMessage)

	// Wait for an interrupt signal
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt, syscall.SIGTERM)
	sig := <-sigCh
	fmt.Printf("Received %v, stopping the client...\n", sig)
	client.Stop()
	fmt.Println("Exiting.")
}
