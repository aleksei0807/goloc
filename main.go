package main

import (
	"github.com/spf13/cobra"
	"fmt"
	"regexp"
	"github.com/aleksei0807/goloc/counter"
	"log"
)

var (
	exclude []string
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
					c, err := counter.New()
					if err != nil {
						log.Fatalf("counter.New() error: %s", err)
					}
					c.Exclude = exclude
					err = c.Local(path)
					if err != nil {
						log.Fatalf("Error: %s\n", err)
					}
				}
			}
		},
	}

	RootCmd.Flags().StringSliceVar(&exclude, "exclude", []string{".git"}, "paths to exclude")
	RootCmd.Execute()
}
