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
	fmt.Printf("Empty Struct: \n %+v\n\n\n", Movie{})

	TestMarshal()
	TestUnmarshal()
}

func TestMarshal() {
	fmt.Printf("TestMarshal【go struct Marshal to []byte】: \n")
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
	fmt.Printf("go struct: %+v\n", movie)
	if data, err := json.Marshal(movie); err != nil {
		fmt.Println("Marshal err：", err)
	} else {
		fmt.Printf("Marshal data: \n%s\n\n", data)
	}
	if data, err := json.MarshalIndent(movie, "", "    "); err != nil {
		fmt.Printf("MarshalIndent err：", err)
	} else {
		fmt.Printf("MarshalIndent data:\n %s\n\n\n\n", data)
	}
}

func TestUnmarshal() {
	fmt.Printf("TestMarshal【[]byte Unmarshal to go struct】: \n")

	data := `{"Title": "123", "released":2000, "SUNZ":123}`
	fmt.Printf("[]byte: %s\n", data)
	var movieTwo Movie
	if err := json.Unmarshal([]byte(data), &movieTwo); err != nil {
		fmt.Println("Unmarshal err: ", err)
	} else {
		fmt.Printf("Unmarshal struct: %+v\n\n", movieTwo)
	}

	dataThird := `{"released":2000, "SUNZ":123}`
	fmt.Printf("[]byte: %s\n", dataThird)
	var movieThird Movie
	if err := json.Unmarshal([]byte(dataThird), &movieThird); err != nil {
		fmt.Println("Unmarshal err: ", err)
	} else {
		fmt.Printf("Unmarshal struct: %+v\n\n", movieThird)
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
// Tips
// go语言结构体内小写变量对json不可见。
