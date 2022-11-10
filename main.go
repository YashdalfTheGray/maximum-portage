package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"

	mapset "github.com/deckarep/golang-set/v2"
	"github.com/o1egl/fwencoder"
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

func main() {
	f, _ := os.Open("/proc/1/net/tcp")
	defer f.Close()

	var entries []NetworkEntry
	unmarshalErr := fwencoder.UnmarshalReader(f, &entries)
	if unmarshalErr != nil {
		fmt.Println(unmarshalErr)
		os.Exit(1)
	}

	portSet := mapset.NewSet[int]()

	for _, entry := range entries {
		parts := strings.Split(entry.LocalAddress, ":")
		portNumber, parseErr := strconv.ParseUint(parts[1], 16, 16)
		if parseErr != nil {
			fmt.Println(parseErr)
			os.Exit(1)
		}
		portSet.Add(int(portNumber))
	}

	ports := portSet.ToSlice()
	sort.Ints(ports)

	fmt.Println(ports)
}
