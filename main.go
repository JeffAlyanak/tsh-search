package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Println("Usage: tsh-search [searchterm …]")
		os.Exit(0)
	}

	cmd := exec.Command("tsh", "ls")
	output, err := cmd.Output()
	if err != nil {
		log.Fatal("", output, err)
	}

	scanner := bufio.NewScanner(bytes.NewReader(output))
	for scanner.Scan() {
		for _, search := range os.Args[1:] {
			if strings.Contains(scanner.Text(), search) {
				fmt.Println(scanner.Text())
			}
		}
	}
	os.Exit(0)
}
