/*
Copyright © 2023 Derric Williams <balls@balls.balls>

*/
package cmd

import (
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/derricw/soulmsg/msg"
	"github.com/spf13/cobra"
)

var numMessages int

var rootCmd = &cobra.Command{
	Use:   "soulmsg",
	Short: "create a random souls message",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		rand.Seed(time.Now().UnixNano())
		for i := 0; i < numMessages; i++ {
			message := msg.RandomMessage(false)
			fmt.Printf("%s\n", message)
		}
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().IntVarP(&numMessages, "num-messages", "n", 1, "number of messages to generate")
}
