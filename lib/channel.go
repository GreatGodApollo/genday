package genday

import (
	"bytes"
	"fmt"
)

type Channel struct {
	Listings []*Listing
	Channel int
	Id string
	Callsign string
}

func NewChannel(ch int, id string, callsign string) *Channel {
	return &Channel{
		Listings: []*Listing{},
		Channel: ch,
		Id: id,
		Callsign: callsign,
	}
}

func (c *Channel) AddListing(l *Listing) *Channel {
	c.Listings = append(c.Listings, l)
	return c
}

func (c *Channel) ToBytes() []byte {
	var out bytes.Buffer
	out.WriteString(fmt.Sprintf("[%4d ", c.Channel))
	out.Write([]byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00})
	out.WriteString(c.Id)
	out.WriteByte(0x00)
	out.WriteString(fmt.Sprintf("%4s", c.Callsign))
	out.Write([]byte{0x00, 0x00, 0x00, 0x00, 0x81, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x8A, 0xFF, 0xFF, '0', '0', 0x00, 0x00, 0x03})
	out.WriteString(c.Id)
	out.WriteByte(0x00)

	return out.Bytes()
}