package cmd

import (
	"go-tmdb-cli/api"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var movieType string

var rootCmd = &cobra.Command{
	Use:   "tmdb-app",
	Short: "TMDb CLI tool",
	Long:  `A CLI tool for fetching movies from the TMDb API.`,
	Run: func(cmd *cobra.Command, args []string) {
		switch movieType {
		case "playing", "popular", "top", "upcoming":
			api.FetchMovies(movieType)
		default:
			fmt.Println("Invalid type. Use one of: playing, popular, top, upcoming")
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.Flags().StringVarP(&movieType, "type", "t", "", "Type of movies to fetch (e.g., playing, popular, top, upcoming)")
	rootCmd.MarkFlagRequired("type")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
