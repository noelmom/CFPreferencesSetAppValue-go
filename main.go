package main

import (
	"bytes"
	"flag"
	"fmt"
	"log"
	"os/exec"
)

// ShellToUse set to bash
const ShellToUse = "bash"

// Shellout output from bash to grab output
func Shellout(command string) (error, string, string) {
	var stdout bytes.Buffer
	var stderr bytes.Buffer
	cmd := exec.Command(ShellToUse, "-c", command)
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	err := cmd.Run()
	return err, stdout.String(), stderr.String()
}

func main() {
	// Flag section
	key := flag.String("key", "HomePage", "The preference key whose value to obtain.")
	appid := flag.String("appid", "com.apple.Safari", "The identifier of the application whose preferences to search.")

	flag.Parse()

	err, out, errout := Shellout(`python -c "from Foundation import CFPreferencesCopyAppValue; print CFPreferencesCopyAppValue('` + (*key) + `', '` + (*appid) + `')"`)
	if err != nil {
		log.Printf("error: %v\n", err)
	} else {

	}

	// Output
	fmt.Println(out)
	fmt.Println(errout)
}
