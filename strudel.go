package strudel

import (
	"net"
	"errors"
	"encoding/json"
)

type StrudelConn struct {
	c *net.UDPConn
}

type StrudelEvent struct {
	Type	string
	Key	string
	Value	interface{}
}

func StrudelConnection(host string)(s *StrudelConn, err error){
	addr, err := net.ResolveUDPAddr("udp", host)
	if err != nil {
		return nil, err
	}
	conn, err := net.DialUDP("udp", nil, addr)
	if err != nil {
		return nil, err
	}
	return &StrudelConn{c: conn}, nil
}

func (s *StrudelConn) Strudel (key string, value float64)(err error){
//	buf := make([]byte, 1500)
	v := StrudelEvent{Type: "float64", Key: key, Value: value}
	buf, err := json.Marshal(v)
	if err != nil {
		return err
	}
	byteswritten, err := s.c.Write(buf)
	if err != nil {
		return
	}
	if byteswritten != len(buf){
		return errors.New("Short write, incomplete strudel")
	}
	return nil
}
