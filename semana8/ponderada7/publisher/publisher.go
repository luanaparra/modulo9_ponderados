package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	DefaultClient "ponderada7/common"
	"time"

	godotenv "github.com/joho/godotenv"
	"os"
)

const maxSensorRange = 1.5
const minSensorRange = 0.1

type Sensor struct {
	Name        string
	Latitude    float64
	Longitude   float64
	Measurement float64
	Rate        int
	Unit        string
}

func NewSensor(
	name string,
	latitude float64,
	longitude float64,
	measurement float64,
	rate int,
	unit string) *Sensor {

	s := &Sensor{
		Name:        name,
		Latitude:    latitude,
		Longitude:   longitude,
		Measurement: measurement,
		Rate:        rate,
		Unit:        unit,
	}

	return s

}

func (s *Sensor) ToJSON() (string, error) {
	jsonData, err := json.Marshal(s)
	if err != nil {
		return "", err
	}
	return string(jsonData), nil
}

func main() {

	err := godotenv.Load("../.env")

	if err != nil {
		fmt.Printf("Error loading .env file: %s", err)
	}

	newBroker := os.Getenv("BROKER_ADDR")
	user := os.Getenv("HIVE_USER")
	pswd := os.Getenv("HIVE_PSWD")

	sensor := NewSensor("SPS30", 48.0, 0.0, 0.0, 1, "μg/m³")

	sensor2 := NewSensor("MiCS-6814", 12.0, 2.0, 0.0, 1, "NO2 - ppm")

	var sensors []Sensor
	sensors = append(sensors, *sensor, *sensor2)

	client := DefaultClient.CreateClientWithAuth(newBroker, DefaultClient.IdPublisher, DefaultClient.Handler, user, pswd)

	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	for {
		for _, sensor := range sensors {

			topic := "sensors/"

			sensor.Measurement = (rand.Float64() * (maxSensorRange - minSensorRange)) + minSensorRange

			payload, _ := sensor.ToJSON()

			token := client.Publish(topic, 0, false, payload)

			token.Wait()

			fmt.Printf("Published message: %s\n", payload)

			time.Sleep(time.Duration(sensor.Rate) * time.Second)

		}
	}
}