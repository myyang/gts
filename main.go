package main

import (
	"flag"
	"fmt"
	"strconv"
	"time"
)

const (
	nanosecond      = "ns"
	microsecond     = "us"
	millisecond     = "ms"
	second          = "s"
	defaultDateTime = "1970-01-01T00:00:00Z"
)

var precision = flag.String(
	"p", "ms",
	"Timestamp precision, default: 'ms'. Options: 's', 'ms', 'us', 'ns'")
var datetime = flag.String(
	"dt", defaultDateTime,
	"UTC datetime string to parse. Format: RFC3339")

func printTime(t time.Time) {
	p, err := time.ParseDuration(fmt.Sprintf("1%v", *precision))
	if err != nil {
		v := millisecond
		precision = &v
		p = time.Millisecond
	}

	t = t.UTC()
	fmt.Printf(
		"%v -> %v (%s)\n",
		t.Format(time.RFC3339),
		t.UnixNano()/int64(p),
		*precision,
	)
}

func main() {
	flag.Parse()

	dt, err := time.Parse(time.RFC3339, *datetime)
	if err == nil && dt.Unix() != 0 {
		printTime(dt)
		return
	}

	if len(flag.Args()) == 0 {
		printTime(time.Now())
		return
	}

	value, err := strconv.ParseInt(flag.Arg(0), 10, 64)
	if err != nil {
		fmt.Printf("Can't parse input value. Err: %v\n", err)
		return
	}

	for value < 1e18 {
		value = value * 10
	}
	printTime(time.Unix(0, value))

}
