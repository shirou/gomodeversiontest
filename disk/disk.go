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
		return 0, err
	}
	fmt.Println("Uptime %d", sysinfo.Uptime)
}
