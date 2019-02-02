package netgear_client

import (
	"os"
	"testing"
	"fmt"
)

func TestLogin(t *testing.T) {
	debug := false

	if get_password() == "" {
		t.Fatal("Error: NETGEAR_PASSWORD environment variable is not set")
	}

	client, err := NewNetgearClient(get_url(), true, get_username(), get_password(), 2, debug)
	if err != nil {
		t.Fatalf("Error getting a client: %s", err)
	}

	err = client.LogIn()
	if err != nil {
		t.Fatalf("Error logging in: %s", err)
	}
}

func TestGetTrafficMeterStatistics(t *testing.T) {
	debug := true

	if get_password() == "" {
		t.Fatal("Error: NETGEAR_PASSWORD environment variable is not set")
	}

	client, err := NewNetgearClient(get_url(), true, get_username(), get_password(), 2, debug)
	if err != nil {
		t.Fatalf("Error getting a client: %s", err)
	}

	res, err := client.GetTrafficMeterStatistics()
	if err != nil {
		t.Fatalf("Error getting traffic statistics: %s", err)
	}
	if len(res) < 1 {
		t.Fatalf("Result is empty... WTF!")
	}
	if debug {
		for k, v := range res {
			fmt.Printf("%v => %v\n", k, v)
		}
	}
}

func get_username() string {
	if os.Getenv("NETGEAR_USERNAME") != "" {
		return os.Getenv("NETGEAR_USERNAME")
	}
	return "admin"
}
func get_password() string {
	return os.Getenv("NETGEAR_PASSWORD")
}
func get_url() string {
	if os.Getenv("NETGEAR_URL") != "" {
		return os.Getenv("NETGEAR_URL")
	}
	return "https://www.routerlogin.com"
}
