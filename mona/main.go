package main

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path"
	"strings"
)

const (
	// MountPoint is path for mounts.
	MountPoint = "/media"
	// Version is command version.
	Version = "0.2.10"
)

// ApprovedDevices returns list of
// devices approved to mount.
func ApprovedDevices() [][]string {
	return [][]string{
		{"hd320", "cfe1663e-7e74-494d-bbf0-eee64db9d70b"},
		{"inesmp3", "268F-11BA"},
		{"kindle", "386D-51FD"},
		{"phone", "0003-AE87"},
		{"usb4", "691E-8498"},
		{"usb8", "6C54-1412"},
		{"usb16", "A5A4-DAFF"},
	}
}

// device holds device's data.
type device struct {
	name    string
	label   string
	mounted bool
}

// String returns formated device output.
func (d *device) String() string {
	s := " "
	if d.mounted {
		s = "*"
	}
	return s + " " + d.label + "(" + d.name + ")"
}

// cmd returns command and args for device.
func cmd(d *device) []string {
	mp := path.Join(MountPoint, d.label)
	if d.mounted {
		return []string{"mount", "/dev/" + d.name, mp}
	}
	return []string{"umount", mp}
}

// deviceLabel returns device label.
func deviceLabel(in []string) string {
	var lbl string
	if len(in) > 1 {
		lbl = in[1]
	}
	return lbl
}

// devices returns list of valid devices.
func devices(ld [][]string, err error) ([]device, error) {
	var ds []device
	if err != nil {
		return ds, err
	}
	for _, ad := range ApprovedDevices() {
		for _, l := range ld {
			if ad[1] == l[1] {
				var mntd bool
				if len(l) == 3 {
					mntd = l[2] != ""
				}
				d := device{
					name:    l[0],
					label:   ad[0],
					mounted: mntd,
				}
				ds = append(ds, d)
			}
		}
	}
	return ds, nil
}

// lsblkData returns valid lsblk command data.
func lsblkData(in []byte, err error) ([][]string, error) {
	var out [][]string
	if err != nil {
		return out, err
	}
	if string(in) == "" {
		return out, errors.New("empty lsblk data")
	}
	s := bufio.NewScanner(bytes.NewReader(in))
	for s.Scan() {
		r := strings.Split(s.Text(), " ")
		if r[0] == "" || r[1] == "" {
			continue
		}
		out = append(out, r)
	}
	return out, s.Err()
}

// lsblkOutput returns lsblk command output.
func lsblkOutput() ([]byte, error) {
	// lsblk -r -o NAME,UUID,MOUNTPOINT
	args := []string{"-n", "-r", "-o", "NAME,UUID,MOUNTPOINT"}
	return exec.Command("lsblk", args...).Output()
}

func main() {
	fmt.Print("mona " + Version + "\n\n")
	// device label
	lbl := deviceLabel(os.Args)
	// devices
	ds, err := devices(lsblkData(lsblkOutput()))
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	// toggle device state
	if d := toggle(ds, lbl); d != nil {
		c := cmd(d)
		out, err := exec.Command(c[0], c[1:]...).CombinedOutput()
		if err != nil {
			fmt.Println(string(out))
			os.Exit(1)
		}
	}
	// list devices
	if len(ds) > 0 {
		fmt.Println("Devices:")
		for _, d := range ds {
			fmt.Println(&d)
		}
	}
}

// toggle change mounted value of device.
func toggle(ds []device, lbl string) *device {
	for i, d := range ds {
		if d.label == lbl {
			d.mounted = !d.mounted
			ds[i] = d
			return &d
		}
	}
	return nil
}
