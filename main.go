package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/dohzya/goprismic"
)

func main() {
	accessToken := flag.String("token", "", "The access token")
	utc := flag.Bool("utc", false, "Display dates in UTC")
	flag.Parse()

	args := flag.Args()
	if len(args) == 0 {
		fmt.Fprintf(os.Stderr, "Missing URL to call")
		os.Exit(1)
	}

	url := args[0]
	api, err := goprismic.Get(url, *accessToken, goprismic.DefaultConfig)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling %s: %v\n", url, err.Error())
		os.Exit(2)
	}
	fmt.Println("Releases:")
	for _, ref := range api.Data.Refs {
		var flag string
		if ref.IsMasterRef {
			flag = " {master}"
		}
		var label string
		if ref.Label != "" {
			label = fmt.Sprintf(" “%s”", ref.Label)
		}
		var scheduledAt string
		if ref.ScheduledAt != 0 {
			date := *ref.ScheduledTime()
			if *utc {
				date = date.UTC()
			}
			scheduledAt = fmt.Sprintf(" <%v>", date)
		}
		fmt.Printf("- release (id=%v ref=%v)%s%s%s\n", ref.Id, ref.Ref, scheduledAt, flag, label)
	}
}
