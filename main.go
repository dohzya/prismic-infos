package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/SoCloz/goprismic"
)

func main() {
	accessToken := flag.String("token", "", "The access token")
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
	refs := api.Data.Refs
	for _, ref := range refs {
		var flag string
		if ref.IsMasterRef {
			flag = " [master]"
		}
		var label string
		if ref.Label != "" {
			label = fmt.Sprintf(" “%s”", ref.Label)
		}
		fmt.Printf("- (id=%v ref=%v)%s%s\n", ref.Id, ref.Ref, flag, label)
	}
}
