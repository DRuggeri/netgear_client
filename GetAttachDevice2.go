package netgear_client

import (
	"encoding/xml"
	"fmt"
	"html"
)

func (client *NetgearClient) GetAttachDevice2() ([]map[string]string, error) {
	const ACTION = "urn:NETGEAR-ROUTER:service:DeviceInfo:1#GetAttachDevice2"
	const REQUEST = `<?xml version="1.0" encoding="UTF-8" standalone="no"?>
<SOAP-ENV:Envelope
  xmlns:SOAPSDK1="http://www.w3.org/2001/XMLSchema"
  xmlns:SOAPSDK2="http://www.w3.org/2001/XMLSchema-instance"
  xmlns:SOAPSDK3="http://schemas.xmlsoap.org/soap/encoding/"
  xmlns:SOAP-ENV="http://schemas.xmlsoap.org/soap/envelope/">
  <SOAP-ENV:Header>
    <SessionID>%s</SessionID>
  </SOAP-ENV:Header>
  <SOAP-ENV:Body>
    <M1:GetAttachDevice xsi:nil="true" />
  </SOAP-ENV:Body>
</SOAP-ENV:Envelope>`

	response, err := client.send_request(ACTION, fmt.Sprintf(REQUEST, client.sessionid), true)
	if err != nil {
		return make([]map[string]string, 0), err
	}

	var inside Node
	err = xml.Unmarshal(response, &inside)
	if err != nil {
		return make([]map[string]string, 0), fmt.Errorf("failed to unmarshal response from inside SOAP body: %v", err)
	}

	devices := make([]map[string]string, 0)
	for _, node := range inside.Nodes[0].Nodes {
		infoMap := make(map[string]string)
		var deviceInfo Node

		err = xml.Unmarshal([]byte(node.Content), &deviceInfo)
		if err != nil {
			return make([]map[string]string, 0), fmt.Errorf("failed to unmarshal response from inside SOAP body: %v", err)
		}

		for _, infoValue := range node.Nodes {
			infoMap[infoValue.XMLName.Local] = html.UnescapeString(infoValue.Content)
		}
		devices = append(devices, infoMap)
	}
	return devices, nil
}
