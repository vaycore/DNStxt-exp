// Package dns Created by vaycore on 2021-12-31.
package dns

import (
	"fmt"
	"golang.org/x/net/dns/dnsmessage"
	"net"
)

type Server struct {
	serverName string
	port       int
	txtRecord  []string
}

func CreateServer(name string, port int, record []string) *Server {
	if len(name) == 0 {
		name = "google"
	}
	if port < 0 || port > 65535 {
		port = 53
	}
	if len(record) == 0 {
		record = []string{"default txt record"}
	}
	return &Server{
		serverName: name,
		port:       port,
		txtRecord:  record,
	}
}

func (s Server) StartDNSServer() {
	conn, err := net.ListenUDP("udp", &net.UDPAddr{Port: s.port})
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	fmt.Printf("Start Listing ... 0.0.0.0:%d, ServerName: %s\n", s.port, s.serverName)
	for {
		buf := make([]byte, 512)
		_, addr, _ := conn.ReadFromUDP(buf)

		var msg dnsmessage.Message
		if err := msg.Unpack(buf); err != nil {
			fmt.Println(err)
			continue
		}
		go s.queryHandler(addr, conn, msg)
	}
}

// dns record query handler
func (s Server) queryHandler(addr *net.UDPAddr, conn *net.UDPConn, msg dnsmessage.Message) {
	// query info
	if len(msg.Questions) < 1 {
		return
	}
	question := msg.Questions[0]
	var (
		queryTypeStr = question.Type.String()
		queryNameStr = question.Name.String()
		queryType    = question.Type
		queryName    = question.Name
	)
	fmt.Printf("[%s] queryName: [%s]\n", queryTypeStr, queryNameStr)
	// find record
	var resource dnsmessage.Resource
	switch queryType {
	case dnsmessage.TypeTXT:
		resource = NewTXTResource(queryName, s.txtRecord)
	case dnsmessage.TypePTR:
		resource = NewPTRResource(queryName, s.serverName+".")
	default:
		fmt.Printf("not support dns queryType: [%s] \n", queryTypeStr)
		return
	}

	// send response
	msg.Response = true
	msg.Answers = append(msg.Answers, resource)
	Response(addr, conn, msg)
}

// Response return
func Response(addr *net.UDPAddr, conn *net.UDPConn, msg dnsmessage.Message) {
	packed, err := msg.Pack()
	if err != nil {
		fmt.Println(err)
		return
	}
	if _, err := conn.WriteToUDP(packed, addr); err != nil {
		fmt.Println(err)
	}
}

// NewTXTResource TXT record
func NewTXTResource(query dnsmessage.Name, txt []string) dnsmessage.Resource {
	return dnsmessage.Resource{
		Header: dnsmessage.ResourceHeader{
			Name:  query,
			Class: dnsmessage.ClassINET,
			TTL:   600,
		},
		Body: &dnsmessage.TXTResource{
			TXT: txt,
		},
	}
}

// NewPTRResource PTR record
func NewPTRResource(query dnsmessage.Name, ptr string) dnsmessage.Resource {
	name, _ := dnsmessage.NewName(ptr)
	return dnsmessage.Resource{
		Header: dnsmessage.ResourceHeader{
			Name:  query,
			Class: dnsmessage.ClassINET,
		},
		Body: &dnsmessage.PTRResource{
			PTR: name,
		},
	}
}
