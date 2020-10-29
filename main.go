package main

import (
	"fmt"
	"io/ioutil"
)

func main() {

	/// Example curday from Wed. October 28th, 2020 - In the Madison, WI area
	curday := NewCurday(6, true, "MSN", "Madison")

	cbs := NewChannel(3, "CBS003", "WISC")

	cbs.AddListing(NewListing(TimeSlotEightPM, "Big Brother"))
	cbs.AddListing(NewListing(TimeSlotTenPM, "News 3 Now at 10:00 p.m."))
	cbs.AddListing(NewListing(TimeSlotTenThirtyPM, "The Late Show with Stephen Colbert"))
	cbs.AddListing(NewListing(TimeSlotElevenThirtyPM, "The Late Late Show with James Corden"))
	cbs.AddListing(NewListing(TimeSlotTwelveThirtyAM, "Impractical Jokers"))
	cbs.AddListing(NewListing(TimeSlotOneAM, "Just for Laughs: Gags"))
	cbs.AddListing(NewListing(TimeSlotOneThirtyAM, "Paid Program"))
	cbs.AddListing(NewListing(TimeSlotTwoAM, "CBS Overnight News"))

	curday.AddChannel(cbs)

	nbc := NewChannel(15, "NBC015", "WMTV")

	nbc.AddListing(NewListing(TimeSlotEightPM, "American Ninja Warrior"))
	nbc.AddListing(NewListing(TimeSlotTenPM, "15 News at 10:00 p.m."))
	nbc.AddListing(NewListing(TimeSlotTenThirtyPM, "The Tonight Show Starring Jimmy Fallon"))
	nbc.AddListing(NewListing(TimeSlotElevenThirtyPM, "Late Night with Seth Meyers"))

	curday.AddChannel(nbc)

	err := ioutil.WriteFile("test.dat", curday.ToBytes(), 0644)

	if err != nil {
		fmt.Printf("An error ocurred: %s\n", err.Error())
		return
	}

	fmt.Println("Generated curday!")
}