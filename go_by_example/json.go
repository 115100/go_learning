package main

import (
	"encoding/json"
	"fmt"
)

type Auction struct {
	Floor    float64 `json:"floor"`
	ID       string  `json:"id"`
	Segments []int   `json:"segments,omitempty"`
}

func main() {
	someAuction := `{"floor": 0.01,
	                 "id": "DEADBEEF",
					 "segments": [1,2,3]}`

	a := &Auction{}
	json.Unmarshal([]byte(someAuction), a)
	fmt.Println(*a)

	someAuction2 := `{"floor": 0.01,
	                 "id": "DEADBEEF"}`

	a2 := &Auction{}
	json.Unmarshal([]byte(someAuction2), a2)
	fmt.Println(*a2)
}
