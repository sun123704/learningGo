package main

import (
	"encoding/json"
	"fmt"
)

type Movie struct {
	Title string
	Year  int    `json:"released"`
	Color bool   `json:"color,omitempty"`
	actor string `json:"actor"`
	Test  int    `json:"Test"`
}

func main() {
	var movie = []Movie{
		{
			Title: "Alice",
			Year:  2000,
			Color: true,
			actor: "A",
			Test:  1,
		},
		{
			Title: "Bob",
			Year:  2000,
			actor: "B",
			Test:  0,
		},
	}

	if data, err := json.Marshal(movie); err != nil {
		fmt.Println("Marshal err：", err)
	} else {
		fmt.Printf("Marshal data: \n%s\n\n", data)
	}

	if data, err := json.MarshalIndent(movie, "", "    "); err != nil {
		fmt.Println("MarshalIndent err：", err)
	} else {
		fmt.Printf("MarshalIndent data:\n %s\n", data)
	}
}

/*
output:
Marshal data:
[{"Title":"Alice","released":2000,"color":true,"Test":1},{"Title":"Bob","released":2000,"Test":0}]

MarshalIndent data:
 [
    {
        "Title": "Alice",
        "released": 2000,
        "color": true,
        "Test": 1
    },
    {
        "Title": "Bob",
        "released": 2000,
        "Test": 0
    }
]
*/
