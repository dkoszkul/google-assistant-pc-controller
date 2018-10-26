package main

import (
        "log"
        "fmt"
        "io/ioutil"
        "gopkg.in/yaml.v2"
)

type Command struct {
        Cmd string `yaml:"cmd"`
        Key string `yaml:"key"`
}

type Commands struct {
        Cmds []Command `yaml:"commands"`
}

func main() {
        log.Println("Start application")
        data, err := ioutil.ReadFile("file.yaml")
        CheckError(err)

        commands := Commands{}
        err = yaml.Unmarshal([]byte(data), &commands)
        CheckError(err)

        c := commands.Cmds
        fmt.Println(c)

        cMap := map[string]string{}
        for _, v := range c {
                cMap[v.Key] = v.Cmd
        }

        fmt.Println(cMap)
        fmt.Println("chosen aaaaaaaaaaaaaaa: Command: ",cMap["aaaaaaaaaaaaaaa"])
}

func CheckError(e error) {
        if e != nil {
                log.Fatalf("error: %v", e)
        }
}
