package gomodules

import "fmt"

func PrintSliceOfString(slc []string) {
	for _, item := range slc {
		fmt.Println(item)
	}
}