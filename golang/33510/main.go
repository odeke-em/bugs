package main

import (
	"fmt"
	"os"
        "syscall"
	"time"
)

func main() {
	const (
		nsecShift              = 30
		secondsPerMinute       = 60
		secondsPerHour         = 60 * secondsPerMinute
		secondsPerDay          = 24 * secondsPerHour
		secondsPerWeek         = 7 * secondsPerDay
		unixToInternal   int64 = (1969*365 + 1969/4 - 1969/100 + 1969/400) * secondsPerDay
		wallToInternal   int64 = (1884*365 + 1884/4 - 1884/100 + 1884/400) * secondsPerDay
		internalToWall   int64 = -wallToInternal
	)
	println("unixToInternal ", -unixToInternal, "\nwallToInternal ", wallToInternal)

	wall := func(s int64) int64 {
		strip := s - wallToInternal
		return (strip << (nsecShift + 1))>>1
	}

        println("Wall times:\nNow: ", wall(1565151693), "\nFS:  ", wall(1565151692))
	return

	defer println("Exiting now")

	for {
		time.Sleep(900 * time.Millisecond)
		if err := os.Remove("tmp1"); err != nil {
			panic(err)
		}
		var now time.Time
		for {
			now = time.Now()
			if now.Nanosecond() == 0 {
				break
			}
		}
		file, _ := os.OpenFile("tmp1", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
		stat, _ := file.Stat()
		fmt.Println(now.Unix(), stat.ModTime().Unix(), time.Now().Unix())
		if now.Second() > stat.ModTime().Second() {
			return
		}
	}
}
