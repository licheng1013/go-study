package main

import (
	"fmt"
	"time"

	"github.com/shirou/gopsutil/net"
)

func main() {
	var oldTotalBytesSent, oldTotalBytesRecv uint64

	for {
		netIOCounters, err := net.IOCounters(true)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		var totalBytesSent, totalBytesRecv uint64
		for _, netIOCounter := range netIOCounters {
			totalBytesSent += netIOCounter.BytesSent
			totalBytesRecv += netIOCounter.BytesRecv
		}

		if oldTotalBytesSent != 0 {
			fmt.Printf("Upload speed: %d bytes/sec\n", totalBytesSent-oldTotalBytesSent)
		}

		if oldTotalBytesRecv != 0 {
			fmt.Printf("Download speed: %d bytes/sec\n", totalBytesRecv-oldTotalBytesRecv)
		}

		// 一起打印

		oldTotalBytesSent = totalBytesSent
		oldTotalBytesRecv = totalBytesRecv

		time.Sleep(1 * time.Second)
	}
}
