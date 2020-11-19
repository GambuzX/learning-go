package main

import (
	"encoding/json"
	"fmt"
)

type SensorReading struct {
	Name     string `json:"name"`
	Capacity int    `json:"capacity"`
	Time     string `json:"time"`
}

func main() {
	jsonString := `{"name": "battery sensor", "capacity": 40, "time": "2019-01-21T19:07:28Z"}`

	var reading SensorReading
	err := json.Unmarshal([]byte(jsonString), &reading)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%+v\n", reading)

	jsonString2 := `{"name": "battery sensor", "capacity": 40, "time": "2019-01-21T19:07:28Z"}`

	var reading2 map[string]interface{}
	err2 := json.Unmarshal([]byte(jsonString2), &reading2)
	if err != nil {
		fmt.Println(err2)
	}
	fmt.Printf("%+v\n", reading2)
}
