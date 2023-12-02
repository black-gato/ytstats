/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"bytes"
	"context"
	"database/sql"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"time"

	c "github.com/black-gato/ytstats/db"
	m "github.com/black-gato/ytstats/db/sqlc"
	"github.com/spf13/cobra"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		setup()

	},
}

type YoutubeEntry struct {
	Header           string      `json:"header"`
	Title            string      `json:"title"`
	TitleURL         string      `json:"titleUrl,omitempty"`
	Subtitles        []Subtitles `json:"subtitles,omitempty"`
	Time             time.Time   `json:"time"`
	Products         []string    `json:"products"`
	ActivityControls []string    `json:"activityControls"`
}
type Subtitles struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

func init() {
	rootCmd.AddCommand(initCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// initCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// initCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func setup() {

	f, err := os.Open("subscriptions.csv")

	if err != nil {
		log.Fatalf("Can't open file %s", err.Error())
	}

	defer f.Close()
	con, err := c.OpenConnection()

	if err != nil {
		log.Fatalf("Can't connect to db %s", err.Error())
	}

	defer con.Close()

	data, err := io.ReadAll(f)

	if err != nil {
		log.Fatalf("Can't Read file %s", err.Error())
	}

	reader := csv.NewReader(bytes.NewReader(data))

	if _, err := reader.Read(); err != nil {
		log.Fatalf("Can't Read first row of csv %s", err.Error())
	}

	subscriptions, err := reader.ReadAll()

	if err != nil {
		log.Fatalf("Can't Read file of csv %s", err.Error())
	}

	db := m.New(con)

	for _, row := range subscriptions {

		_, err := db.AddChannel(context.Background(), m.AddChannelParams{ID: row[0], ChannelUrl: row[1], ChannelName: row[2], IsSubbed: 1})

		if err != nil {
			continue
		}

	}

	j, err := os.Open("watch-history.json")
	if err != nil {
		log.Fatalf("couldn't read file %s", err.Error())
	}
	byteValue, err := io.ReadAll(j)

	if err != nil {
		log.Fatalf("couldn't read file %s", err.Error())
	}

	d := json.NewDecoder(bytes.NewReader(byteValue))

	_, err = d.Token()

	if err != nil {
		log.Fatalf("couldn't decode token %s", err.Error())
	}

	for d.More() {
		var y YoutubeEntry
		err := d.Decode(&y)

		if err != nil {
			log.Fatal(err)
		}

		if y.Subtitles != nil {
			vId := strings.Split(y.TitleURL, "watch?v\u003d")
			fmt.Println(vId[1])
			videoNullStr := sql.NullString{
				String: vId[1],
				Valid:  true,
			}
			title := strings.Split(y.Title, "Watched ")

			cId := strings.SplitAfterN(y.Subtitles[0].URL, "/", 5)

			channelNullStr := sql.NullString{
				String: cId[4],
				Valid:  true,
			}

			_, err = db.AddWatchHistory(context.Background(), m.AddWatchHistoryParams{VideoID: videoNullStr, WatchedAt: y.Time.String(), ChannelID: channelNullStr})

			if err != nil {
				fmt.Printf("hello %v \n", err)
				continue
			}

			_, err = db.AddVideo(context.Background(), m.AddVideoParams{ID: vId[1], VideoType: y.Header, VideoTitle: title[1], ChannelID: channelNullStr})

			if err != nil {
				fmt.Println(err)
				continue
			}

			_, err := db.AddChannel(context.Background(), m.AddChannelParams{ID: cId[4], ChannelUrl: y.Subtitles[0].URL, ChannelName: y.Subtitles[0].Name, IsSubbed: 0})

			if err != nil {
				fmt.Println(err)
				continue
			}

		}
	}

}
