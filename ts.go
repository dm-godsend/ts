package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("not enough args")
	}

	var out string
	if ts, err := strconv.ParseInt(os.Args[1], 0, 64); err != nil {
		format := time.DateOnly
		dateStr := os.Args[1]
		if len(os.Args) > 2 {
			format = time.DateTime
			dateStr += " " + os.Args[2]
		}
		t, err := time.Parse(format, dateStr)
		if err != nil {
			log.Fatal("date parse err: " + err.Error())
		}
		out = fmt.Sprint(t.UTC().Unix())
	} else {
		out = fmt.Sprint(time.Unix(ts, 0).UTC())
	}

	os.Stdout.WriteString(out + "\n")
	os.Exit(0)
}
