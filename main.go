package main

import (
	"fmt"
	"os"
	"strconv"
)

type NetworkEntry struct {
	EntryNumber     string `column:"sl"`
	LocalAddress    string `column:"local_address"`
	RemoteAddress   string `column:"rem_address"`
	State           string `column:"st"`
	TxQueue         string `column:"tx_queue"`
	RxQueue         string `column:"rx_queue"`
	TimerAndJiffies string `column:"tr tm->when"`
	Retransmit      string `column:"retrnsmt"`
	Uid             string `column:"uid"`
	Timeout         string `column:"timeout"`
	Inode           string `column:"inode"`
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

	allFdLinks := []string{}

	for _, fd := range allFds {
		original, linkErr := os.Readlink(fd.Name())
		if linkErr != nil {
			fmt.Println(linkErr)
		} else {
			allFdLinks = append(allFdLinks, original)
		}
	}

	for _, link := range allFdLinks {
		fmt.Println(link)
	}
}
