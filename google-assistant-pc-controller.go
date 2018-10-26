package main

import (
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strings"
	"time"

	"github.com/fsnotify/fsnotify"
	"github.com/gonutz/w32"
)

func main() {
	console := w32.GetConsoleWindow()
	if console != 0 {
		_, consoleProcID := w32.GetWindowThreadProcessId(console)
		if w32.GetCurrentProcessId() == consoleProcID {
			w32.ShowWindowAsync(console, w32.SW_HIDE)
		}
	}

	// Uncomment it when want to log to file
	// f, err := os.OpenFile("testlogfile.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	// if err != nil {
	// 	log.Fatalf("error opening file: %v", err)
	// }
	// defer f.Close()
	// log.SetOutput(f)

	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	defer watcher.Close()

	commandsMap := GetCommands("configuration/configuration.yaml")
	PrintConfiguration(commandsMap)

	done := make(chan bool)
	go func() {
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}
				if event.Op&fsnotify.Write == fsnotify.Write {
					if strings.HasSuffix(event.Name, "txt") {
						content, _ := ioutil.ReadFile(event.Name)
						s := string(content)
						log.Println("Command ", s)
						ExecuteCommand(commandsMap[strings.ToLower(s)])
						if s != "" {
							time.Sleep(1 * time.Second)
							os.Remove(event.Name)
						}
					}
				}
			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				log.Println("error:", err)
			}
		}
	}()

	err = watcher.Add("D:\\dropbox\\Dropbox\\IFTTT")
	if err != nil {
		log.Fatal(err)
	}
	<-done
}

func ExecuteCommand(cmd string) {
	log.Println("Execute command ", cmd)
	c := exec.Command(cmd)
	if err := c.Start(); err != nil {
		log.Fatal(err)
	}
}
