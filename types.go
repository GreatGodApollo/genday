package main

import (
	"bytes"
	"fmt"
	"time"
)

type Curday struct {
	Channels []*Channel
	DiagnosticSettings string
	Timezone int
	DaylightSavings bool
	DREV int
	Airport string
	City string
}

func boolToYN(b bool) string {
	if b {
		return "Y"
	}
	return "N"
}

func NewCurday(timezone int, dst bool, airportICAO, cityName string) *Curday {
	return &Curday{
		Channels: []*Channel{},
		DiagnosticSettings: "AE3366N",
		Timezone: timezone,
		DaylightSavings: dst,
		DREV: 5,
		Airport: airportICAO,
		City: cityName,
	}
}

func (c *Curday) AddChannel(ch *Channel) *Curday {
	c.Channels = append(c.Channels, ch)
	return c
}

func (c *Curday) Header() []byte {
	var out bytes.Buffer

	out.WriteString(c.DiagnosticSettings)
	out.Write([]byte{0x03, 0x01})
	out.WriteString(fmt.Sprintf("%d%s", c.Timezone, boolToYN(c.DaylightSavings)))
	out.WriteString("YNNYYNNl")
	out.Write([]byte{0x00, 0x00, '0', 0x00})
	out.WriteString(fmt.Sprintf("DREV %d", c.DREV))
	out.WriteByte(0x00)
	out.WriteString(c.Airport)
	out.WriteByte(0x00)
	out.WriteString(c.City)
	out.WriteByte(0x00)

	jDay := time.Now().YearDay() - 1
	if jDay > 255 {
		jDay = jDay - 255
	}

	out.WriteString(fmt.Sprintf("%d", jDay))
	out.WriteByte(0x00)
	out.WriteString(fmt.Sprintf("%d", c.ListingCount()))
	out.WriteByte(0x00)
	out.WriteString("131")
	out.WriteByte(0x00)
	out.WriteString("1348")
	out.WriteByte(0x00)

	return out.Bytes()
}

func (c *Curday) ToBytes() []byte {
	var out bytes.Buffer

	out.Write(c.Header())

	for _, channel := range c.Channels {
		out.Write(channel.ToBytes())
		for _, listing := range channel.Listings {
			out.Write(listing.ToBytes(c))
		}
		out.WriteString("49")
		out.WriteByte(0x00)
	}

	return out.Bytes()
}

func (c *Curday) ListingCount() int {

	var i = 0

	for _, channel := range c.Channels {
		for range channel.Listings {
			i ++
		}
	}

	return i

}

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
	// [  {str(channel)} \x00\x00\x00\x00\x00\x00{id}\x00{callsign}\x00\x00\x00\x00\x81\xFF\xFF\xFF\xFF\xFF\xFF\x00\x00\x00\x00\x00\x00\x8A\xFF\xFF00\x00\x00\x03{id}\x00' + str(''.join(listings)) + '49\x00
	out.WriteString(fmt.Sprintf("[%4d ", c.Channel))
	out.Write([]byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00})
	out.WriteString(c.Id)
	out.WriteByte(0x00)
	out.WriteString(c.Callsign)
	out.Write([]byte{0x00, 0x00, 0x00, 0x00, 0x81, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0x00,0x00,0x00,0x00, 0x00, 0x00, 0x8A, 0xFF, 0xFF, 0x30, 0x30, 0x00, 0x00, 0x03})
	out.WriteString(c.Id)
	out.WriteByte(0x00)

	return out.Bytes()
}

type Listing struct {
	Timeslot Timeslot
	Name string
}

func NewListing(timeslot Timeslot, name string) *Listing {
	return &Listing{
		Timeslot: timeslot,
		Name: name,
	}
}

func (l *Listing) ToBytes(c *Curday) []byte {
	var out bytes.Buffer

	// {str(timeslot)}\x001\x0034\x000\x000\x00{name}\x00
	ts := int(l.Timeslot)
	ts -= (c.Timezone - 4) * 2

	if ts > 48 {
		ts -= 48
	} else if ts <= 0 {
		ts += 48
	}

	out.WriteString(fmt.Sprintf("%d", ts))
	out.WriteByte(0x00)
	out.WriteString("1")
	out.WriteByte(0x00)
	out.WriteString("34")
	out.WriteByte(0x00)
	out.WriteString("0")
	out.WriteByte(0x00)
	out.WriteString("0")
	out.WriteByte(0x00)
	out.WriteString(l.Name)
	out.WriteByte(0x00)

	return out.Bytes()
}