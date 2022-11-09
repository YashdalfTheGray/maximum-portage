package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	prodDirs, readErr := os.ReadDir("/host/proc")
	if readErr != nil {
		fmt.Printf("Error reading procfs, %v\n", readErr)
	}

	pidDirs := []os.DirEntry{}

	for _, dir := range prodDirs {
		if _, err := strconv.ParseUint(dir.Name(), 10, 16); err == nil {
			fmt.Printf("PID dir: %s\n", dir.Name())
			pidDirs = append(pidDirs, dir)
		}
	}

	allFds := []os.DirEntry{}

	for _, dir := range pidDirs {
		fds, fdReadErr := os.ReadDir(fmt.Sprintf("/proc/%s/fd", dir.Name()))
		if fdReadErr != nil {
			fmt.Println(fdReadErr)
		} else {
			allFds = append(allFds, fds...)
		}
	}

	for _, fd := range allFds {
		fmt.Println(fd.Name())
	}
}
