/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"math/rand"

	"github.com/spf13/cobra"
)

// helloCmd represents the hello command
var helloCmd = &cobra.Command{
	Use:   "hello",
	Short: "A simple but powerful message",
	Long:  `A simple but powerful message.`,
	Run: func(cmd *cobra.Command, args []string) {
		MESSAGES := []string{"Over here, stranger!", "What're ya buyin'?", "Not enough cash, stranger!", "Got somethin' that might interest ya'!"}
		randomIndex := rand.Intn(len(MESSAGES))

		fmt.Println(MESSAGES[randomIndex])
	},
}

func init() {
	rootCmd.AddCommand(helloCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// helloCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// helloCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
