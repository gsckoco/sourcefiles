package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func getFilesInDirectory(path string, extensions []string) []string {
	var outputSlice []string
	err := filepath.Walk(path, func(filePath string, info os.FileInfo, err error) error {
		// Check if the file is a directory, if it is recurse and grab it's files
		// If its a file check its extension, if it matches one of the provided
		// file extensions, append it to the output array
		if info.IsDir() {
			if filePath != path {
				outputSlice = append(outputSlice, getFilesInDirectory(filePath, extensions)...)
			}
		} else {
			if contains(extensions, strings.Replace(filepath.Ext(filePath), ".", "", 1)) {
				outputSlice = append(outputSlice, filePath)
			}
		}
		return nil
	})

	if err != nil {
		panic(err)
	}
	return outputSlice
}

func main() {
	args := os.Args[1:]
	if len(args) < 2 {
		fmt.Printf("Not enough arguments\n")
		return
	}
	extensions := os.Args[2:]
	var files []string
	var filesNoDuplicates []string
	files = getFilesInDirectory(args[0], extensions)
	for i := 0; i < len(files); i++ {
		if contains(filesNoDuplicates, files[i]) == false {
			filesNoDuplicates = append(filesNoDuplicates, files[i])
		}
	}
	var output string
	for i := 0; i < len(filesNoDuplicates); i++ {
		output += filesNoDuplicates[i] + " "
	}
	fmt.Printf(output + "\n")
}
