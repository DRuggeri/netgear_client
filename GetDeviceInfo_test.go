package netgear_client

import (
	"fmt"
	"testing"
)

func TestGetInfo(t *testing.T) {
	debug := true

	if get_password() == "" {
		t.Fatal("Error: NETGEAR_PASSWORD environment variable is not set")
	}

	client, err := NewNetgearClient(get_url(), true, get_username(), get_password(), 5, debug)
	if err != nil {
		t.Fatalf("Error getting a client: %s", err)
	}

	res, err := client.GetDeviceInfo()
	if err != nil {
		t.Fatalf("Error getting device info: %s", err)
	}
	if len(res) == 0 {
		t.Fatalf("Unexpected empty response")
	}

	expected := [...]string{
		"Description",
		"DeviceMode",
		"DeviceModeCapability",
		"DeviceName",
		"DeviceNameUserSet",
		"FirewallVersion",
		"FirmwareDLmethod",
		"FirmwareLastChecked",
		"FirmwareLastUpdate",
		"Firmwareversion",
		"FirstUseDate",
		"ModelName",
		"Otherhardwareversion",
		"OthersoftwareVersion",
		"SerialNumber",
		"SignalStrength",
		"SmartAgentversion",
		"VPNVersion",
	}
	for _, key := range expected {
		if _, ok := res[key]; !ok {
			t.Errorf("Expected `%s` key in the response, but did not find it", key)
		}
	}

	if debug {
		for k, v := range res {
			fmt.Printf("%v => %v\n", k, v)
		}
	}
}
