package phpipam

import (
	"fmt"
	"testing"

	"github.com/pavel-z1/phpipam-sdk-go/controllers/addresses"
	"github.com/stretchr/testify/assert"
)

func TestFilterUsedIPAddresses(t *testing.T) {
	ip_addresses := []addresses.Address{
		{
			ID:        2,
			IPAddress: "192.168.1.2",
		},
	}

	expectedResult := []string{
		"192.168.1.1",
		"192.168.1.3",
		"192.168.1.4",
	}

	output, err := FilterUsedIPAddresses("192.168.1.0/24", &ip_addresses, 3)
	assert.Nil(t, err)
	assert.ElementsMatch(t, expectedResult, *output)
}

func TestFilterUsedIPAddressesNotEnoughFreeIPs(t *testing.T) {
	ip_addresses := genIps(255)

	for _, ip := range ip_addresses {
		fmt.Printf("%s", ip.IPAddress)
	}

	output, err := FilterUsedIPAddresses("192.168.1.0/24", &ip_addresses, 3)
	assert.Nil(t, output)
	assert.EqualError(t, err, "cannot get the asked number of IPs")
}

func genIps(number int) []addresses.Address {
	ips := make([]addresses.Address, number-1)
	for i := range ips {
		ips[i] = addresses.Address{
			ID:        i + 1,
			IPAddress: fmt.Sprintf("192.168.1.%d", i+1),
		}
	}

	return ips
}
