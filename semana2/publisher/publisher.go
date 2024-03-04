package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"time"
	DefaultClient "mqtt/src/common"
)

const maxSensorRange = 1.0
const minSensorRange = 0.03

type Sensor struct {
	Name        string
	Latitude    float64
	Longitude   float64
	Rate        int
	Unit        string
	Measurement float64
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
		Rate:        rate,
		Unit:        unit,
		Measurement: measurement
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

	sensor := NewSensor("SPS30", 51.0, 0.0, 0.0, 1, "μg/m³") //sensor SPS30

	sensor2 := NewSensor("MiCS-6814", 10.0, 1.0, 0.0, 1, "NO2 - ppm") //sensor MiCS-6814

	var sensors []Sensor
	sensors = append(sensors, *sensor, *sensor2)

	client := DefaultClient.CreateClient(DefaultClient.Broker, DefaultClient.IdPublisher, DefaultClient.Handler)

	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	for {
		for _, sensor := range sensors {
			topic := "sensors/" + sensor.Name
			sensor.Measurement = (rand.Float64() * (maxSensorRange - minSensorRange)) + minSensorRange
			payload, _ := sensor.ToJSON()
			token := client.Publish(topic, 0, false, payload)
			token.Wait()
			fmt.Printf("Published message: %s\n", payload)
			time.Sleep(time.Duration(sensor.Rate) * time.Second)
		}
	}
}