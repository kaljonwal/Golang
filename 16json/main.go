package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	fmt.Println("Hello World")
	DecodeJson()
}

type Course struct {
	Name     string `json:"coursename"`
	Price    int
	Platform string   `json:"website"`
	Tags     []string `json:"tags, omitempty"`
}

func DecodeJson() {

	var jsonCourse Course
	jsonData := []byte(`
	{
		"coursename" : "kalyan",
		"Price" : 0,
		"Website": "https://github.com/kaljonwal",
		"tags": ["github","learning"]
	}
	`)

	json.Unmarshal(jsonData, &jsonCourse)

	fmt.Printf("Json data : %v\n", jsonCourse)

	var storeJsonInMap map[string]interface{}

	json.Unmarshal(jsonData, &storeJsonInMap)

	fmt.Printf("Json Data umarshled in map %#v", storeJsonInMap)
}
