package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/SoCloz/goprismic"
)

func main() {
	accessToken := flag.String("token", "", "The access token")
	only := flag.String("only", "all", "The infos to display (releases,forms,bookmarks,types,tags)")
	utc := flag.Bool("utc", false, "Display dates in UTC")
	flag.Parse()

	var display func(string) bool
	if *only == "all" {
		display = func(_ string) bool { return true }
	} else {
		fieldNames := strings.Split(*only, ",")
		fields := make(map[string]bool, len(fieldNames))
		for _, field := range fieldNames {
			fields[field] = true
		}
		display = func(field string) bool { return fields[field] }
	}

	args := flag.Args()
	if len(args) == 0 {
		fmt.Fprintf(os.Stderr, "Missing URL to call")
		os.Exit(1)
	}

	url := args[0]
	api, err := goprismic.Get(url, *accessToken, goprismic.DefaultConfig)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling %v: %v\n", url, err.Error())
		os.Exit(2)
	}
	if display("releases") {
		fmt.Println("Releases:")
		for _, ref := range api.Data.Refs {
			var flag string
			if ref.IsMasterRef {
				flag = " {master}"
			}
			var label string
			if ref.Label != "" {
				label = fmt.Sprintf(" “%v”", ref.Label)
			}
			var scheduledAt string
			if ref.ScheduledAt != 0 {
				date := *ref.ScheduledTime()
				if *utc {
					date = date.UTC()
				}
				scheduledAt = fmt.Sprintf(" <%v>", date)
			}
			fmt.Printf("- release (id=%v ref=%v)%v%v%v\n", ref.Id, ref.Ref, scheduledAt, flag, label)
		}
	}
	if display("forms") {
		fmt.Println("Forms:")
		for id, form := range api.Data.Forms {
			var name string
			if form.Name != "" {
				name = fmt.Sprintf(" “%v”", form.Name)
			}
			fmt.Printf("- form (id=%v)%v\n", id, name)
		}
	}
	if display("bookmarks") {
		fmt.Println("Bookmarks:")
		for name, ref := range api.Data.Bookmarks {
			fmt.Printf("- bookmark (ref=%v) “%v”\n", ref, name)
		}
	}
	if display("types") {
		fmt.Println("Types:")
		for id, name := range api.Data.Types {
			fmt.Printf("- type (id=%v) “%v”\n", id, name)
		}
	}
	if display("tags") {
		tags := strings.Join(api.Data.Tags, ", ")
		fmt.Printf("Tags: %v\n", tags)
	}
}
