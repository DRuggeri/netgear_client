package netgear_client

import (
	"encoding/xml"
	"fmt"
)

// Implements the DeviceInfo/GetInfo SOAP message
func (client *NetgearClient) GetDeviceInfo() (map[string]string, error) {
	const ACTION = "urn:NETGEAR-ROUTER:service:DeviceInfo:1#GetInfo"
	const REQUEST = `<?xml version="1.0" encoding="UTF-8" standalone="no"?>
<SOAP-ENV:Envelope
	xmlns:SOAPSDK1="http://www.w3.org/2001/XMLSchema"
	xmlns:SOAPSDK2="http://www.w3.org/2001/XMLSchema-instance"
	xmlns:SOAPSDK3="http://schemas.xmlsoap.org/soap/encoding/"
	xmlns:SOAP-ENV="http://schemas.xmlsoap.org/soap/envelope/">
	<SOAP-ENV:Header>
		<SessionID>%s</SessionID>
	</SOAP-ENV:Header>
</SOAP-ENV:Envelope>`

	response, err := client.send_request(ACTION, fmt.Sprintf(REQUEST, client.sessionid), true)
	if err != nil {
		return make(map[string]string), err
	}

	var inside Node
	err = xml.Unmarshal(response, &inside)
	if err != nil {
		return make(map[string]string), fmt.Errorf("failed to unmarshal response from inside SOAP body: %v", err)
	}

	var info = make(map[string]string)
	for _, node := range inside.Nodes {
		name := node.XMLName.Local
		info[name] = node.Content
	}
	return info, nil
}
