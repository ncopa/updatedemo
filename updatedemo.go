package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
	"syscall"
	"time"
)

var Version = "0.1"

func main() {
	printVersion := flag.Bool("version", false, "print version and exit")
	flag.Parse()

	if *printVersion {
		fmt.Println(Version)
		os.Exit(0)
	}

	path, err := os.Executable()
	if err != nil {
		log.Fatal("failed to find executable path: %s\n", err)
	}

	fmt.Printf("Hello world. updatedemo version %s\n", Version)

	updateLoop(path)
}

func updateLoop(path string) {
	for {
		output, err := exec.Command(path, "--version").Output()
		if err != nil {
			log.Printf("Failed to get version from %s", path)
		} else {
			newVersion := strings.TrimSuffix(string(output), "\n")
			if newVersion != Version {
				fmt.Printf("It was fun but now its time for next generation to take over.. Make me proud %s!\n", newVersion)
				err := syscall.Exec(path, os.Args, os.Environ())
				if err != nil {
					log.Fatal(err)
				}
			}
		}
		time.Sleep(1 * time.Second)
	}
}
