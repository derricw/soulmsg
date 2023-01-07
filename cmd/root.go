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
var allowConjuctions bool
var conjunctionRate float32

func randBool(trueProportion float32) bool {
	return rand.Float32() < trueProportion
}

func printRandomMessage() {
	var doConjuction bool
	if allowConjuctions {
		doConjuction = randBool(conjunctionRate)
	}
	message := msg.RandomMessage(doConjuction)
	fmt.Printf("%s\n", message)
}

var rootCmd = &cobra.Command{
	Use:   "soulmsg",
	Short: "create a random souls message",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		rand.Seed(time.Now().UnixNano())
		if numMessages < 0 {
			for {
				printRandomMessage()
			}
		}
		for i := 0; i < numMessages; i++ {
			printRandomMessage()
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
	rootCmd.Flags().IntVarP(&numMessages, "num-messages", "n", 1, "number of messages to generate. -1 is infinite")
	rootCmd.Flags().BoolVarP(&allowConjuctions, "allow-conjunctions", "c", false, "allow conjuctions")
	rootCmd.Flags().Float32VarP(&conjunctionRate, "conjunction-rate", "r", 0.5, "proportion of messages with conjunctions")
}
