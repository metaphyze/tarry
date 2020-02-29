package main

import (
	"flag"
	"fmt"
	"os"
	"time"
)

func main() {
	var (
		until = flag.String("until", "", "Required. Wait (tarry) until this specific time (hours, minutes, seconds)."+
			"  Format: 17:59:01\n"+
			"For example,\n\ttarry -until=17:59:01\nUse -whatTimeIsIt to find out what the current system time is.  "+
			"If you need millisecond precision, use the -plusMs flag.")
		plusMs = flag.Int("plusMs", 0, "Optional (defaults to 0).  Add a few milliseconds onto your time (0-999).  For example,\n"+
			"\ttarry -until=17:59:01 -plusMs=250\n"+
			"would wait until 17:59:01:250, that is, 1/4 second beyond 17:59:01")
		inDaysFromNow = flag.Int("inDaysFromNow", 0,
			"Optional (defaults to 0). Specifies the number of days (0-365) from today.  For example, "+
				"tomorrow would be -inDaysFromNow=1")
		whatTimeIsIt = flag.Bool("whatTimeIsIt", false, "Show the current system time.  You'll probably need "+
			"to use this to set -until.  For example,\n"+
			"\ttarry -whatTimeIsIt\n"+
			"\tCurrent time: 17:22:03")
		printTimeWhenDone = flag.Bool("printTimeWhenDone", false, "Print out the current time when done")
	)

	flag.Parse()

	if *whatTimeIsIt {
		nowInMs := time.Now().UnixNano() / 1000000
		timeStr := GetSimpleTimeString(nowInMs, false)
		fmt.Printf("Current time: %v\n", timeStr)
		os.Exit(1) // Exit with non-zero so commands after "&&" do not run
	}

	if *until == "" {
		fmt.Println("-until not specified.  Use -help for options.")
		os.Exit(1) // Exit with non-zero so commands after "&&" do not run
	}

	if *inDaysFromNow < 0 || *inDaysFromNow > 365 {
		fmt.Println("inDaysFromNow must be between 0 and 365 (inclusive)")
		os.Exit(1) // Exit with non-zero so commands after "&&" do not run
	}

	if *plusMs < 0 || *plusMs > 999 {
		fmt.Println("plusMs must between 0 and 999 (inclusive)")
		os.Exit(1) // Exit with non-zero so commands after "&&" do not run
	}

	timeInNs, err := GetTimeUntilInNs(*until, *plusMs, *inDaysFromNow)

	if err != nil {
		fmt.Printf("Error converting time/day: %v\nUse -help for options.", err)
		os.Exit(1) // Exit with non-zero so commands after "&&" do not run
	}

	duration := timeInNs - time.Now().UnixNano()

	if duration < 0 {
		fmt.Printf("Error: Time is in the past by %v seconds \n", -float64(duration)/1000000/1000)
		os.Exit(1) // Exit with non-zero so commands after "&&" do not run
	}

	time.Sleep(time.Duration(duration) * time.Nanosecond)

	if *printTimeWhenDone {
		nowInMs := time.Now().UnixNano() / 1000000
		fmt.Printf("Current time: %v\n", GetSimpleTimeString(nowInMs, true))
	}
}

func GetTimeUntilInNs(timeStr string, plusMs int, inDaysFromNow int) (int64, error) {
	t, err := time.Parse("15:04:05", timeStr)

	if err != nil {
		return -1, err
	}

	now := time.Now().Add(time.Duration(inDaysFromNow*24) * time.Hour)

	ns := plusMs * 1000000

	return time.Date(now.Year(), now.Month(), now.Day(), t.Hour(), t.Minute(), t.Second(), ns, now.Location()).UnixNano(), nil
}

func GetSimpleTimeString(timeInMs int64, showMs bool) string {
	seconds := timeInMs / 1000
	leftOverMs := timeInMs - seconds*1000
	ns := leftOverMs * 1000000
	if !showMs {
		return time.Unix(seconds, ns).Format("15:04:05")
	} else {
		return time.Unix(seconds, ns).Format("15:04:05.000")
	}
}
