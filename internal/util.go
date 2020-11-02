package internal

import (
	genday "github.com/GreatGodApollo/genday/lib"
	"io/ioutil"
	"strconv"
	"strings"
)

func SaveCurday(filename string, curday *genday.Curday) error {
	return ioutil.WriteFile(filename, curday.ToBytes(), 0644)
}

func NearestTimeslot(t string) (genday.Timeslot, error) {

	split := strings.Split(t, ":")
	hour, err := strconv.Atoi(split[0])
	if err != nil  { return -1, err }
	minute, err := strconv.Atoi(split[1])
	if err != nil  { return -1, err }

	ts := (hour * 2) + 1

	if minute >= 15 && minute < 45 {
		ts += 1
	} else if minute >= 45 {
		ts += 2
	}

	if ts > 48 {
		ts -= 48
	} else if ts <= 0 {
		ts += 48
	}

	return genday.Timeslot(ts), nil
}