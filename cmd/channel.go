/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"
	"fmt"
	"log"

	c "github.com/black-gato/ytstats/db"
	m "github.com/black-gato/ytstats/db/sqlc"
	"github.com/spf13/cobra"
)

// channelCmd represents the channel command

var mostWatched bool
var channelCmd = &cobra.Command{
	Use:   "channel",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	// Args: func(cmd *cobra.Command, args []string) error {
	// 	if len(args) < 1 {
	// 		return errors.New("requires at least one arg")
	// 	}
	// 	return nil
	// },
	Run: func(cmd *cobra.Command, args []string) {

		mostWatched, _ := cmd.Flags().GetBool("most-watched")

		if mostWatched {
			con, err := c.OpenConnection()

			if err != nil {
				log.Fatalf("Can't connect to db %s", err.Error())
			}

			defer con.Close()

			db := m.New(con)

			tst, err := db.GetMostWatched(context.Background())

			if err != nil {
				log.Fatal(err)
			}

			fmt.Println(tst)
		}

		fmt.Println(mostWatched)

	},
}

func init() {
	rootCmd.AddCommand(channelCmd)
	channelCmd.PersistentFlags().Bool("most-watched", false, "most watched channel")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// channelCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// channelCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
