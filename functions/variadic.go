package main

import (
	"fmt"
	"os"
)

func main() {
	temp := os.TempDir()

	buildFiles(temp, "test1")
}

func buildFiles(dirBase string, files ...string) {

	for _, name := range files {

		path := fmt.Sprintf("%s/%s.%s", dirBase, name, "txt")
		file, err := os.Create(path)

		defer file.Close()

		if err != nil {
			fmt.Printf("Error to create file %s: %v\n", name, err)
			os.Exit(1)
		}
		fmt.Printf("File %s created. \n", file.Name())
	}

}
