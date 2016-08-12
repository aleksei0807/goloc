package main

import (
	"github.com/spf13/cobra"
	"fmt"
	"regexp"
	"github.com/aleksei0807/goloc/counter"
)

func main() {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Printf("Error: %s\n", err)
		}
	}()

	RootCmd := &cobra.Command{
		Use: "goloc",
		Short: "goloc counts lines of code in given addres.",
		Long: `goloc counts lines of code in given addres.
		Supports local directories and git repositorires.`,
		Run: func(cmd *cobra.Command, args []string) {

			if len(args) > 0 {
				path := args[0]

				matched, err := regexp.MatchString("^https?://.+", path)
				if err != nil {
					fmt.Printf("Path parsing error: %s\n", err)
				}

				if matched {
					fmt.Println("repository")
				} else {
					err := counter.CountLocal(path)
					if err != nil {
						fmt.Printf("%#+v\n", err)
					}
				}
			}
		},
	}

	RootCmd.Execute()
}
