package main

import (
	"fmt"
	"syscall"

	"time"
)

/* Реализовать собственную функцию sleep. */

func main() {

	Sleep(1)

}

func Sleep(seconds uint32) {
	start := time.Now()
	pathp, _ := syscall.UTF16PtrFromString("")
	h, _ := syscall.CreateFile(pathp, 0, 0, nil, syscall.OPEN_EXISTING, syscall.FILE_FLAG_BACKUP_SEMANTICS, 0)
	syscall.WaitForSingleObject(h, seconds*1000)
	elapsed := time.Since(start)
	fmt.Printf("Sleeping took %v", elapsed)
}
