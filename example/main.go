package main

import (
	"github.com/masanetes/acos-client-go/pkg/axapi/slb/virtualserver"
	"github.com/masanetes/acos-client-go/pkg/axapi/slb/virtualserverport"
	"log"

	"github.com/masanetes/acos-client-go/pkg/client"
)

func main() {

	config := client.Config{Host: "", User: "", Pass: "", Debug: false}
	c, err := client.New(config, client.InsecureSkipVerify(true))
	if err != nil {
		log.Fatal(err)
	}

	name := "masanetes-sample"
	ip := "192.168.0.10"

	virtualServer, err := c.Slb.VirtualServer.Create(&virtualserver.Body{
		Object: virtualserver.Object{Name: name, IPAddress: ip}})
	if err != nil {
		log.Fatal(err)
	}

	_, err = c.Slb.VirtualServerPort.CreateList(virtualServer.Name, &virtualserverport.ListBody{
		ListObjects: virtualserverport.ListObjects{
			virtualserverport.Object{PortNumber: 80, Protocol: "http"},
			virtualserverport.Object{PortNumber: 443, Protocol: "https"},
		}})
	if err != nil {
		log.Fatal(err)
	}

	err = c.Slb.VirtualServer.Delete(virtualServer.Name)
	if err != nil {
		log.Fatal(err)
	}
}
