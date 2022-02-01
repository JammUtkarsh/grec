package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
)

const fileLocation = "./.hidden/lists.json"

func readJson() (elements []string) {
	fileBytes, err := ioutil.ReadFile(fileLocation)
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(fileBytes, &elements)
	if err != nil {
		panic(err)
	}
	return elements
}

func clearJson() {
	clear := []byte("[]")
	//out, _ := json.Marshal(clear)
	_ = ioutil.WriteFile(fileLocation, clear, 0755)
}

func deleteObject() {
	delCMD := flag.NewFlagSet("delete", flag.ExitOnError)
	data := delCMD.Int("d", 0, "delete an entry")
	err := delCMD.Parse(os.Args[2:])
	if err != nil {
		panic(err)
	}
	jsonObjects := readJson()
	var newJsonObjects []string
	for i, entry := range jsonObjects {
		if *data-1 == i {
			continue
		}
		newJsonObjects = append(newJsonObjects, entry)
	}
	infoByte, err := json.Marshal(newJsonObjects)
	if err != nil {
		panic(err)
	}
	err = ioutil.WriteFile(fileLocation, infoByte, 0644)
}

func writeJson() {
	addCMD := flag.NewFlagSet("add", flag.ExitOnError)
	data := addCMD.String("d", "", "enter the data")
	err := addCMD.Parse(os.Args[2:])
	element := append(readJson(), *data)
	infoByte, err := json.Marshal(element)
	if err != nil {
		panic(err)
	}
	err = ioutil.WriteFile(fileLocation, infoByte, 0644)
	if err != nil {
		panic(err)
	}
}

func list() {
	readData := readJson()
	for i, data := range readData {
		fmt.Printf("|%d|\t%s\t|\n", i+1, data)
	}
}
