// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    curdayJSON, err := UnmarshalCurdayJSON(bytes)
//    bytes, err = curdayJSON.Marshal()

package internal

import "encoding/json"

func UnmarshalCurdayJSON(data []byte) (CurdayJSON, error) {
	var r CurdayJSON
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *CurdayJSON) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type CurdayJSON struct {
	Settings string    `json:"settings"`
	Timezone int       `json:"timezone"`
	Dst      bool      `json:"dst"`
	City     string    `json:"city"`
	Airport  string    `json:"airport"`
	Channels []Channel `json:"channels"`
}

type Channel struct {
	Number    int       `json:"number"`
	ID        string    `json:"id"`
	Callsign  string    `json:"callsign"`
	Listings  []Listing `json:"listings"`
	Summary   bool      `json:"summary"`
	Hilite    bool      `json:"hilite"`
	AltHilite bool      `json:"althilite"`
}

type Listing struct {
	Time string `json:"time"`
	Name string `json:"name"`
}
