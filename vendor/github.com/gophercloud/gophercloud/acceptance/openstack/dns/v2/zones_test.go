// +build acceptance dns zones

package v2

import (
	"testing"

	"github.com/gophercloud/gophercloud/acceptance/clients"
	"github.com/gophercloud/gophercloud/acceptance/tools"
	"github.com/gophercloud/gophercloud/openstack/dns/v2/zones"
)

func TestZonesList(t *testing.T) {
	client, err := clients.NewDNSV2Client()
	if err != nil {
		t.Fatalf("Unable to create a DNS client: %v", err)
	}

	var allZones []zones.Zone
	allPages, err := zones.List(client, nil).AllPages()
	if err != nil {
		t.Fatalf("Unable to retrieve zones: %v", err)
	}

	allZones, err = zones.ExtractZones(allPages)
	if err != nil {
		t.Fatalf("Unable to extract zones: %v", err)
	}

	for _, zone := range allZones {
		tools.PrintResource(t, &zone)
	}
}
