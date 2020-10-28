// +build linux

package host

import (
	"fmt"

	"golang.org/x/sys/unix"
)

func Foo() {
	fmt.Println("FOO")

	sysinfo := &unix.Sysinfo_t{}
	if err := unix.Sysinfo(sysinfo); err != nil {
		panic(err)
	}
	fmt.Printf("Uptime %v\n", sysinfo.Uptime)
}
