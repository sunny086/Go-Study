package main

import (
	"fmt"
	"gitlab.com/tingshuo/go-diskstate/diskstate"
)

// example
func main() {
	state := diskstate.DiskUsage("D://")
	fmt.Printf("All=%dM, Free=%dM, Available=%dM, Used=%dM, Usage=%d%%",
		state.All/diskstate.MB, state.Free/diskstate.MB, state.Available/diskstate.MB, state.Used/diskstate.MB, 100*state.Used/state.All)
}
