package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

var (
	inputString = `{
		"best-effort":      "0",
			"background":       "1",
			"spare":            "2",
			"excellent-effort": "3",
			"controlled-load":  "4",
			"video":            "5",
			"voice":            "6",
			"network-control":  "7"

	}`
)

func main() {
	reverse(inputString)
}

func reverse(in string) {
	input := make(map[string]string)
	err := json.Unmarshal([]byte(in), &input)
	if err != nil {
		panic(err)
	}
	output := make(map[string]string)

	for k, v := range input {
		output[v] = k
	}

	data, err := json.MarshalIndent(output, "", "\t")
	if err != nil {
		panic(err)
	}
	log.Println(string(data))
}

func generateJavaMap() {

	data, err := os.ReadFile("input.json")
	if err != nil {
		panic(err)
	}
	input := make(map[string]string)
	err = json.Unmarshal(data, &input)
	if err != nil {
		panic(err)
	}
	outStrArr := make([]string, 0)
	var outStr string

	for k, v := range input {
		val, _ := strconv.Atoi(k)
		outStr = fmt.Sprintf(`intrusionIDToNameMap.put(%d, "%s")`, val, v)
		outStrArr = append(outStrArr, outStr)
	}
	dataToWrite := strings.Join(outStrArr, ";\n")
	logFileSize("output.txt", []byte(dataToWrite))
	log.Println(string(outStr))
}

func generatedIntrusionNameMap() {

	data, err := os.ReadFile("input.json")
	if err != nil {
		panic(err)
	}
	input := make(map[string]string)
	err = json.Unmarshal(data, &input)
	if err != nil {
		panic(err)
	}
	outStrArr := make([]string, 0)
	var outStr string

	for k, v := range input {
		val, _ := strconv.Atoi(k)
		outStr = fmt.Sprintf(`intrusionIDToNameMap.put(%d, "%s")`, val, v)
		outStrArr = append(outStrArr, outStr)
	}
	dataToWrite := strings.Join(outStrArr, ";\n")
	logFileSize("output.txt", []byte(dataToWrite))
	log.Println(string(outStr))
}

func logFileSize(fileName string, apiConfig []byte) {

	err := ioutil.WriteFile(fileName, apiConfig, 0644)
	if err != nil {
		log.Fatalf("error while writing apiJsonConfig=> %v", err)
	}
	file, err := os.Stat(fileName)
	if err != nil {
		log.Fatalf("error while reading apiJsonConfig=> %v", err)
	}
	log.Fatalf("apiConfig File size in KB=> %v", file.Size()/1024)
}
