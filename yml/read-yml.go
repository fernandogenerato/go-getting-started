package main

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"log"
	"os"
)

type Yml struct {
	Name         string   `yaml:"Name"`
	Address      string   `yaml:"Address"`
	Phone        int      `yaml:"Phone"`
	WorkFromHome bool     `yaml:"WorkFromHome"`
	Mail         string   `yaml:"Mail"`
	Internships  []string `yaml:"Internships"`
	Education    struct {
		MSc     float64 `yaml:"MSc"`
		BSc     float64 `yaml:"BSc"`
		XII     int     `yaml:"XII"`
		Courses struct {
			Udemy    int `yaml:"Udemy"`
			Coursera int `yaml:"Coursera"`
		} `yaml:"Courses"`
		References string `yaml:"References"`
	} `yaml:"Education"`
}

func main() {

	y := Yml{}
	file, err := os.ReadFile("yml/example.yml")
	if err != nil {
		log.Fatal(err)
	}
	if err := yaml.Unmarshal(file, &y); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("All documents decoded")
	fmt.Println(y)
}
