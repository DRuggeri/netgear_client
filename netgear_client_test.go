package netgear_client

import (
	"os"
	"strings"
	"testing"
)

var username = ""
var password = ""
var testUrl = ""

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

func get_username() string {
	if username != "" {
		return username
	}

	if os.Getenv("NETGEAR_USERNAME") != "" {
		username = os.Getenv("NETGEAR_USERNAME")
	} else if b, err := os.ReadFile(".netgear_username"); err == nil {
		username = strings.TrimSuffix(string(b), "\n")
	} else {
		username = "admin"
	}
	return username
}

func get_password() string {
	if password != "" {
		return password
	}

	if os.Getenv("NETGEAR_PASSWORD") != "" {
		password = os.Getenv("NETGEAR_PASSWORD")
	} else if b, err := os.ReadFile(".netgear_password"); err == nil {
		password = strings.TrimSuffix(string(b), "\n")
	} else {
		panic("password must be set in either NETGEAR_PASSWORD env variable or in the .netgear_password file")
	}
	return password
}

func get_url() string {
	if testUrl != "" {
		return testUrl
	}

	if os.Getenv("NETGEAR_URL") != "" {
		testUrl = os.Getenv("NETGEAR_URL")
	} else if b, err := os.ReadFile(".netgear_url"); err == nil {
		testUrl = strings.TrimSuffix(string(b), "\n")
	} else {
		testUrl = "https://www.routerlogin.com"
	}

	return testUrl
}
