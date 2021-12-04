package main

import (
	"io/ioutil"
	"log"

	yaml "gopkg.in/yaml.v3"
)

func main() {
	var labelsYAML Configuration
	labelfile, err := ioutil.ReadFile("labels.yaml")
	if err != nil {
		log.Fatal(err)
	}
	if err = yaml.Unmarshal(labelfile, &labelsYAML); err != nil {
		log.Fatal(err)
	}
	log.Printf("%v", labelsYAML.Default.Labels)

	var sig Context
	sigData, err := ioutil.ReadFile("sigs.yaml")
	if err != nil {
		log.Fatal(err)
	}
	if err = yaml.Unmarshal(sigData, &sig); err != nil {
		log.Fatal(err)
	}
	log.Printf("%v", sig.Sigs[0].LabelName(""))

	// traverse through sigs,wgs..... and append respective labels to label[]
	// we compare the label[] with
	// traverse labelsYAML and append names to another array
	// Perfrom diff comparison between both array

	//assume label names are enough cz they are self descriptive
}
