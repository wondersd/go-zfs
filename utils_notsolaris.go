// +build !solaris

package zfs

import (
	"strings"
)

// List of ZFS properties to retrieve from zfs list command on a non-Solaris platform
var dsPropList = []string{"name", "origin", "used", "available", "mountpoint", "compression", "type", "volsize", "quota", "referenced", "written", "logicalused", "usedbydataset"}

var dsPropListOptions = strings.Join(dsPropList, ",")

// List of Zpool properties to retrieve from zpool list command on a non-Solaris platform
var zpoolPropList = []string{"name", "health", "allocated", "size", "free", "readonly", "dedupratio", "fragmentation", "freeing", "leaked"}
var zpoolPropListOptions = strings.Join(zpoolPropList, ",")
var zpoolArgs = []string{"get", "-p", zpoolPropListOptions}

func getDataset(name string) (*Dataset, error) {
	out, err := zfs("list", "-Hp", "-o", dsPropListOptions, name)
	if err != nil {
		return nil, err
	}

	ds := &Dataset{Name: name}
	for _, line := range out {
		if err := ds.parseListLine(line); err != nil {
			return nil, err
		}
	}

	return ds, nil
}