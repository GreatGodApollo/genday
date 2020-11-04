package cmd

import (
	"fmt"
	"github.com/GreatGodApollo/genday/internal"
	genday "github.com/GreatGodApollo/genday/lib"
	cli "github.com/jawher/mow.cli"
	"github.com/ttacon/chalk"
	"io/ioutil"
)

func cmdJSON(cmd *cli.Cmd) {
	cmd.Spec = "INPUT [ -o=<output> ]"

	var (
		inputFile = cmd.StringArg("INPUT", "", "Input JSON")
		outputFile = cmd.StringOpt("o output", "curday.dat", "Output file")
	)
	cmd.Action = func() {
		v := *verbose

		if v {
			fmt.Println(internal.NewMessage(chalk.Yellow, "Opening file"))
		}
		data, err := ioutil.ReadFile(*inputFile)
		if err != nil {
			fmt.Println(internal.NewMessage(chalk.Red, "Failed to read file:").ThenColor(chalk.Yellow, "\n" + err.Error()))
			return
		}

		if v {
			fmt.Println(internal.NewMessage(chalk.Yellow, "Decoding JSON"))
		}
		curdayJSON, err := internal.UnmarshalCurdayJSON(data)
		if err != nil {
			fmt.Println(internal.NewMessage(chalk.Red, "Failed to decode JSON:").ThenColor(chalk.Yellow, "\n" + err.Error()))
			return
		}

		curday := genday.NewCurday(curdayJSON.Timezone, curdayJSON.Dst, curdayJSON.Airport, curdayJSON.City)
		curday.DiagnosticSettings = curdayJSON.Settings

		for _, channel := range curdayJSON.Channels {
			if v {
				fmt.Println(internal.NewMessage(chalk.Yellow, "Generating " + channel.Callsign))
			}
			var curListings []*genday.Listing
			for _, listing := range channel.Listings {
				if v {
					fmt.Println(internal.NewMessage(chalk.Yellow, " - "+listing.Name))
				}
				ts, err := internal.NearestTimeslot(listing.Time)
				if err != nil {
					fmt.Println(internal.NewMessage(chalk.Red, fmt.Sprintf("Invalid time: Channel: %s Listing: %s Time: %s", channel.Callsign, listing.Name, listing.Time)))
					continue
				}
				curListings = append(curListings, genday.NewListing(ts, listing.Name))
			}

			var cf genday.ChannelFlags
			if channel.Hilite {
				cf = cf.Set(genday.ChannelFlagHiliteSrc)
			}
			if channel.AltHilite {
				cf = cf.Set(genday.ChannelFlagAltHiliteSrc)
			}
			if channel.Summary {
				cf = cf.Set(genday.ChannelFlagSumbySrc)
			}

			c := genday.NewChannel(channel.Number, channel.ID, channel.Callsign, cf)
			for _, l := range curListings {
				c.AddListing(l)
			}
			curday.AddChannel(c)
		}
		if v {
			fmt.Println(internal.NewMessage(chalk.Yellow, "Saving to \"" + *outputFile + "\""))
		}
		err = internal.SaveCurday(*outputFile, curday)
		if err != nil {
			fmt.Println(internal.NewMessage(chalk.Red, "Failed to save to").ThenColor(chalk.Yellow, *outputFile).ThenColor(chalk.Red, "\n" + err.Error()))
		} else {
			fmt.Println(internal.NewMessage(chalk.Green, "Successfully saved to").ThenColor(chalk.Yellow, *outputFile))
			fmt.Println(internal.NewMessage(chalk.Yellow, fmt.Sprintf("Generated %d listings", curday.ListingCount())))
		}
	}
}