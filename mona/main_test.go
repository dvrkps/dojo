package main

import (
	"errors"
	"reflect"
	"testing"
)

// fakeLsblk holds fake output
// from lsblk command.
const fakeLsblk = `sda  
sda1 70bf9fca-2654-40ca-b06b-befda95d252b /
sda2 cc5b3f5e-4a99-4fae-bae9-1410361bb151 /home
sda3 5ab2f66d-641f-4885-b6e0-5a30b72bb043 /data
sdc  
sdc1 0003-AE87 /media/phone
sdd 691E-8498
sde  
sde1 6C54-1412 
sdf A5A4-DAFF 
sdg  
sdg1 386D-51FD`

func TestApprovedDevices(t *testing.T) {
	if got := len(ApprovedDevices()); got != 7 {
		t.Errorf("len(ApprovedDevices()) = %v; want 7", got)
	}
}

func TestCmd(t *testing.T) {
	// mounted
	in := &device{
		name:    "name",
		label:   "label",
		mounted: true,
	}
	want := []string{"mount", "/dev/name", "/media/label"}
	if got := cmd(in); !reflect.DeepEqual(got, want) {
		t.Errorf("cmd(mounted)= %v; want %v", got, want)
	}
	// umounted
	in = &device{
		name:    "name",
		label:   "label",
		mounted: false,
	}
	want = []string{"umount", "/media/label"}
	if got := cmd(in); !reflect.DeepEqual(got, want) {
		t.Errorf("cmd(umounted)= %v; want %v", got, want)
	}
}

func TestDeviceLabel(t *testing.T) {
	in := []string{"cmd", "arg1"}
	want := in[1]
	if got := deviceLabel(in); got != want {
		t.Errorf("deviceLabel(%v) = %v; want %v", in, got, want)
	}
	// no args
	in = []string{"cmd"}
	want = ""
	if got := deviceLabel(in); got != want {
		t.Errorf("deviceLabel(%v) = %v; want %v", in, got, want)
	}
}

func TestDeviceString(t *testing.T) {
	d := &device{
		name:    "name",
		label:   "label",
		mounted: true,
	}
	want := "* label(name)"
	if got := d.String(); got != want {
		t.Errorf("device.String() = %v; want %v", got, want)
	}
}

func TestDevices(t *testing.T) {
	_, err := devices([][]string{}, errors.New("device error"))
	if err == nil {
		t.Errorf("devices(error) = _, %v; want error", err)
	}
	lo, _ := lsblkData([]byte(fakeLsblk), nil)
	want := 5
	if got, _ := devices(lo, nil); len(got) != want {
		t.Errorf("devices(valid): len(got) = %v; want %v",
			len(got), want)
	}
}

func TestLsblkData(t *testing.T) {
	// input error
	in := []byte("")
	if _, err := lsblkData(in, errors.New("lsblk error")); err == nil {
		t.Errorf("lsblkData(error) = _, %v; want error", err)
	}
	// empty
	in = []byte("")
	if _, err := lsblkData(in, nil); err == nil {
		t.Errorf("lsblkData(empty) = _, %v; want error", err)
	}
	// valid
	in = []byte(fakeLsblk)
	want := 8
	if got, _ := lsblkData(in, nil); len(got) != want {
		t.Errorf("lsblkData(fakeLsblk): len(got) = %v; want %v",
			len(got), want)
	}
}

func TestLsblkOutput(t *testing.T) {
	if got, err := lsblkOutput(); len(got) == 0 || err != nil {
		t.Errorf("len(lsblkOutput()) = %v, err = %v; want > 0", len(got), err)
	}
}

func TestToggle(t *testing.T) {
	lo, _ := lsblkData([]byte(fakeLsblk), nil)
	ds, _ := devices(lo, nil)
	// exists
	if got := toggle(ds, "kindle"); got == nil || !got.mounted {
		t.Errorf("toggle(ds, kindle) = %v; want device", got)
	}
	// nil
	if got := toggle(ds, "invalid"); got != nil {
		t.Errorf("toggle(ds, invalid) = %v; want <nil>", got)
	}
}
