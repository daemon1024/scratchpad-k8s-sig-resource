package main

import (
	"io/ioutil"
	"log"
	yaml "gopkg.in/yaml.v3"
)

func main() {
	// for labels array
	set := make(map[string]int)
	// for teams array
	

	var arr[]string

	// loading labels from label.yaml
	var labelsYAML Configuration
	labelfile, err := ioutil.ReadFile("labels.yaml")
	if err != nil {
		log.Fatal(err)
	}
	if err = yaml.Unmarshal(labelfile, &labelsYAML); err != nil {
		log.Fatal(err)
	}

	//loading labels from sig.yaml
	var sig Context
	sigData, err := ioutil.ReadFile("sigs.yaml")
	if err != nil {
		log.Fatal(err)
	}
	if err = yaml.Unmarshal(sigData, &sig); err != nil {
		log.Fatal(err)
	}


	//loading individual teams
	var teams TeamList
	t, err := ioutil.ReadFile("sigs.yaml")
	if err != nil {
		log.Fatal(err)
	}
	if err = yaml.Unmarshal(t, &teams); err != nil {
		log.Fatal(err)
	}

	// getting the labels from sigs.yaml and setting up the map 
	// for _, j := range sig.Sigs {
	// 	set[j.Label] = 1
	// }
	
	// getting teams from sig.yaml and storing it in array
	for i, _ := range sig.Sigs {
		if(sig.Sigs[i].Contact.GithubTeams == nil){
			continue;
		}
		for j, _ := range sig.Sigs[i].Contact.GithubTeams {
			log.Println(sig.Sigs[i].Contact.GithubTeams[j].Name)
		}
	}

	 
	// getting the labels from label.yaml and adding it to a array
	 for _, s := range labelsYAML.Default.Labels {
		var check = s.Name
		var temp = s.Name
		if string(check[0:3]) == "sig" {
			var subtemp = string(temp[4:])
			arr = append(arr, subtemp)
		}
	 }
	 // iterating through the array and remove the set whoes label matches
	for _, s := range arr {
		if set[s] == 1 {
			delete(set, s)
		}
	}

	// logging the final set which is the difference in both the file 
	log.Println(set)
		

	// traverse through sigs,wgs..... and append respective labels to label[]
	// we compare the label[] with
	// traverse labelsYAML and append names to another array
	// Perfrom diff comparison between both array
	//assume label names are enough cz they are self descriptive
}
