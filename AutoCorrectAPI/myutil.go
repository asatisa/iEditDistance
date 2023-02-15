package main

import (
	"log"
	"os"
	"path/filepath"

	"github.com/vaughan0/go-ini"
)

// Constant of Application
const AppIni string = "app.ini"

var isInitialize bool = false

func getExecDir() (exPath string) {
	path, err := os.Executable() //os.Getwd() //os.Executable()
	if err != nil {
		log.Println(err)
	}
	exPath = filepath.Dir(path)
	//fmt.Println("my path: " + exPath)
	return
}

func readINI(section string, key string) (value string) {
	initializeMyUtil()
	//var err error
	file, err := ini.LoadFile(AppIni)
	if err != nil {
		log.Println(err)
	}
	value, ok := file.Get(section, key)
	if !ok {
		panic("'" + key + "' variable missing from '" + section + "' section")
	}
	return
}

func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

// Check file is exist.
func fileExist(filename string) bool {
	if _, err := os.Stat(filename); err == nil {
		//fmt.Printf("File '" + filename + "' exists\n")
		return true
	} else {
		//fmt.Printf("File '" + filename + "' does not exist\n")
		return false
	}
}

// //////////////////////////////////////////////////////////////////////////////////////////////////////
// Initialization modules
// /////////////////////////////////////////////////////////////////////////////////////////////////////
func initializeMyUtil() bool {
	//fmt.Println("Init: " + strconv.FormatBool(isInitialize))
	if isInitialize {
		return true
	}
	if !fileExist(AppIni) {
		isInitialize = false
		panic("File INI '" + AppIni + "' is not exist!")
	} else {
		isInitialize = true
		return true
	}

}
