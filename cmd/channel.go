/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"
	"fmt"
	"log"
	"os"
	"text/tabwriter"

	c "github.com/black-gato/ytstats/db"
	m "github.com/black-gato/ytstats/db/sqlc"
	"github.com/spf13/cobra"
)

// channelCmd represents the channel command

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
		w := tabwriter.NewWriter(os.Stdout, 0, 0, 4, ' ', 0)

		fmt.Fprintln(w, "ChannelName\tViews\tIsSubbed")

		limit, _ := cmd.Flags().GetInt64("limit")
		isSubbed, _ := cmd.Flags().GetBool("isSubbed")
		name, _ := cmd.Flags().GetString("name")

		con, err := c.OpenConnection()

		if err != nil {
			log.Fatalf("Can't connect to db %s", err.Error())
		}

		defer con.Close()

		db := m.New(con)

		if name != "" {
			channel, err := db.GetMostWatchedChannels(context.Background(), m.GetMostWatchedChannelsParams{ChannelName: name})

			if err != nil {
				log.Fatal(err)
			}

			for _, ch := range channel {
				fmt.Fprintf(w, "%s\t%d\t%t\n", ch.ChannelName, ch.WatchCount, ch.IsSubbed)

			}

			w.Flush()
			return

		}
		channel, err := db.GetMostWatchedChannels(context.Background(), m.GetMostWatchedChannelsParams{Limit: limit, IsSubbed: isSubbed})

		if err != nil {
			log.Fatal(err)
		}

		for _, ch := range channel {
			fmt.Fprintf(w, "%s\t%d\t%t\n", ch.ChannelName, ch.WatchCount, ch.IsSubbed)

		}

		w.Flush()

	},
}

func init() {
	rootCmd.AddCommand(channelCmd)
	channelCmd.PersistentFlags().String("name", "", "channel name")
	channelCmd.PersistentFlags().Bool("isSubbed", false, "subs only")

	channelCmd.PersistentFlags().Int64("limit", 10, "number of entries retur")
	channelCmd.MarkFlagsMutuallyExclusive("name", "limit")
	channelCmd.MarkFlagsMutuallyExclusive("name", "isSubbed")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// channelCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// channelCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
