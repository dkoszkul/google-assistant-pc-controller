package main

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

type Command struct {
	Cmd string `yaml:"cmd"`
	Key string `yaml:"key"`
}

type Commands struct {
	Cmds []Command `yaml:"commands"`
}

func GetCommands(filepath string) map[string]string {
	data, err := ioutil.ReadFile(filepath)
	CheckError(err)

	commands := Commands{}
	err = yaml.Unmarshal(data, &commands)
	CheckError(err)

	cMap := make(map[string]string)
	for _, v := range commands.Cmds {
		cMap[v.Key] = v.Cmd
	}

	return cMap
}

func PrintConfiguration(c map[string]string) {
	log.Println("Configuration. Voice command : Command on PC")
	for k, v := range c {
		log.Println("[Mapping] ", k, " : ", v)
	}
}

func CheckError(e error) {
	if e != nil {
		log.Fatalf("ERROR: %v", e)
	}
}
