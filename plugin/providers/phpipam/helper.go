package phpipam

import (
	"encoding/binary"
	"errors"
	"fmt"
	"net"

	"github.com/pavel-z1/phpipam-sdk-go/controllers/addresses"
)

// linearSearchSlice provides a []string with a helper search function.
type linearSearchSlice []string

// Has checks the linearSearchSlice for the string provided by x and returns
// true if it finds a match.
func (s *linearSearchSlice) Has(x string) bool {
	for _, v := range *s {
		if v == x {
			return true
		}
	}
	return false
}

func FilterUsedIPAddresses(CIDR string, ips *[]addresses.Address, max int) (*[]string, error) {
	fmt.Printf("[ERROR] CIDR is %s", CIDR)
	result := make([]string, 0)
	cidr_addrs, err := expandCIDR(CIDR)
	if err != nil {
		panic("foo")
		return nil, err
	}
out:
	for _, range_ip := range (*cidr_addrs)[1:] {
		found := false
		for _, address := range *ips {
			if address.IPAddress == range_ip.String() {
				found = true
				break
			}
		}

		if !found {
			result = append(result, range_ip.String())
			if max <= len(result) {
				break out
			}
		}
	}

	if len(result) < max {
		return nil, errors.New("cannot get the asked number of IPs")
	}

	return &result, nil
}

// from https://stackoverflow.com/a/60542265
func expandCIDR(CIDR string) (*[]net.IP, error) {
	_, ipv4Net, err := net.ParseCIDR(CIDR)
	if err != nil {
		return nil, err
	}

	// convert IPNet struct mask and address to uint32
	// network is BigEndian
	mask := binary.BigEndian.Uint32(ipv4Net.Mask)
	start := binary.BigEndian.Uint32(ipv4Net.IP)

	// find the final address
	finish := (start & mask) | (mask ^ 0xffffffff)

	ips := make([]net.IP, 0)

	// loop through addresses as uint32
	for i := start; i <= finish; i++ {
		// convert back to net.IP
		ip := make(net.IP, 4)
		binary.BigEndian.PutUint32(ip, i)
		ips = append(ips, ip)
	}

	return &ips, nil
}
