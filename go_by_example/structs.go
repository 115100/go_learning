package main

import "fmt"

type weather struct {
	name          string
	precipitation bool
}

func main() {
	fmt.Println(weather{name: "sunny", precipitation: false})
	fmt.Println(weather{name: "snowing", precipitation: true})
	fmt.Println(weather{})

	sunny_day := weather{name: "sunny", precipitation: false}
	fmt.Println(sunny_day.name)
	fmt.Println(sunny_day.precipitation)

	fmt.Println((&sunny_day).name)
}
