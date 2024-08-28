package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
)

var input = flag.String("input", "", "File to translate (required)")
var key = flag.String("key", "", "Name of the key field (required)")
var value = flag.String("value", "", "Name of the value field (required)")
var output = flag.String("output", "output.csv", "Outputted file")
var excludedKeys = flag.String("excluded-keys", "", "Comma separated list of keys to exclude")

func main() {
	flag.Parse()

	if *input == "" || *key == "" || *value == "" {
		fmt.Println("Missing required parameters")
		flag.PrintDefaults()
		os.Exit(1)
	}

	excludedKeys := strings.Split(*excludedKeys, ",")

	if _, err := os.Stat(*input); os.IsNotExist(err) {
		exitWithError("File does not exist")
	}

	jsonFile, err := os.Open(*input)
	if err != nil {
		exitWithError("Error opening file")
	}

	defer jsonFile.Close()

	byteValue, _ := io.ReadAll(jsonFile)

	var result *map[string]interface{}
	err = json.Unmarshal([]byte(byteValue), &result)
	if err != nil {
		exitWithError("Error parsing JSON")
	}

	newFile, err := os.Create(*output)
	if err != nil {
		exitWithError("Error creating new file")
	}
	defer newFile.Close()

	_, err = newFile.Write([]byte(fmt.Sprintf("%s,%s\n", *key, *value)))
	if err != nil {
		exitWithError("Error writing to file")
	}

	for k, v := range *result {
		if contains(excludedKeys, k) {
			continue
		}

		_, err = newFile.Write([]byte(fmt.Sprintf("\"%s\",\"%s\"\n", k, strings.ReplaceAll(fmt.Sprintf("%v", v), "\"", "'"))))
		if err != nil {
			exitWithError("Error writing to file")
		}
	}
}

func exitWithError(err string) {
	fmt.Println(err)
	os.Exit(1)
}

func contains(arr []string, key string) bool {
	for _, a := range arr {
		if a == key {
			return true
		}
	}
	return false
}
