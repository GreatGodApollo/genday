package genday

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