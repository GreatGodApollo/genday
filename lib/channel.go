package genday

import (
	"bytes"
	"fmt"
)

type ChannelFlags uint8

const (
	ChannelFlagNone = 1 << iota
	ChannelFlagHiliteSrc
	ChannelFlagSumbySrc
	ChannelFlagVideoTagDisable
	ChannelFlagCAF_PPV
	ChannelFlagDitto
	ChannelFlagAltHiliteSrc
	ChannelFlagStereo
)

func (cf ChannelFlags) Set(flag ChannelFlags) ChannelFlags    { return cf | flag }
func (cf ChannelFlags) Clear(flag ChannelFlags) ChannelFlags  { return cf &^ flag }
func (cf ChannelFlags) Toggle(flag ChannelFlags) ChannelFlags { return cf ^ flag }
func (cf ChannelFlags) Has(flag ChannelFlags) bool            { return *&flag != 0 }

type Channel struct {
	Listings []*Listing
	Channel  int
	Id       string
	Callsign string
	Flags    ChannelFlags
}

func NewChannel(ch int, id string, callsign string, flags ChannelFlags) *Channel {
	return &Channel{
		Listings: []*Listing{},
		Channel:  ch,
		Id:       id,
		Callsign: callsign,
		Flags:    flags,
	}
}

func (c *Channel) AddListing(l *Listing) *Channel {
	c.Listings = append(c.Listings, l)
	return c
}

func (c *Channel) ToBytes() []byte {
	var out bytes.Buffer
	out.WriteString(fmt.Sprintf("[%5d", c.Channel))
	out.Write([]byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00})
	out.WriteString(fmt.Sprintf("%6s", c.Id))
	out.WriteByte(0x00)
	out.WriteString(fmt.Sprintf("%6s", c.Callsign))
	out.Write([]byte{0x00, 0x00, byte(c.Flags), 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0xFF, 0xFF, '0', '0', 0x00, 0x00, 0x03})
	out.WriteString(fmt.Sprintf("%6s", c.Id))
	out.WriteByte(0x00)

	return out.Bytes()
}
