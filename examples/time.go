package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()

	fmt.Println("now:					", now)

	fmt.Print("now.Clock():				 ")
	fmt.Println(now.Clock())

	fmt.Print("now.Date():				 ")
	fmt.Println(now.Date())

	fmt.Println("now.GoString():				", now.GoString())
	fmt.Println("now.Local():				", now.Local())
	fmt.Println("now.Location():				", now.Location())
	fmt.Println("now.String():				", now.String())
	fmt.Println("now.UTC():				", now.UTC())
	fmt.Println("now.Unix():				", now.Unix())
	fmt.Println("now.Format(time.UnixDate):		", now.Format(time.UnixDate))
}
