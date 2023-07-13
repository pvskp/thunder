/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"net/http"
	"os"
	"sync"

	"github.com/gleich/logoru"
	"github.com/spf13/cobra"
)

var (
	Host             string
	Workers          int
	NumberOfRequests int
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "thunder",
	Short: "A simple load test tool for simple tasks",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		doRequest(Host)
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func fillChan(requestQtd chan<- int) {
	for i := 0; i < NumberOfRequests; i++ {
		requestQtd <- i
	}
	close(requestQtd)
}

func worker(notStarted <-chan int, host string, wg *sync.WaitGroup) {
	for range notStarted {
		res, err := http.Head(host)
		if err != nil {
			msg := fmt.Sprintf("Error making http request: %s\n", err)
			logoru.Error(msg)
			continue
		}
		if res.StatusCode != 200 {
			logoru.Error("Got status code", res.StatusCode)
			continue
		}
		logoru.Success("Request returned 200")
	}
	wg.Done()
}

func doRequest(host string) {
	notStarted := make(chan int, NumberOfRequests)
	var wg sync.WaitGroup
	wg.Add(Workers)
	defer wg.Wait()

	logoru.Info("Making HEAD request to", host)
	for i := 0; i < Workers; i++ {
		go worker(notStarted, host, &wg)
	}

	fillChan(notStarted)
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.thunder.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	// Optional flags
	rootCmd.Flags().IntVarP(&Workers, "workers", "w", 1, "How many workers (goroutines) to use.")
	rootCmd.Flags().IntVarP(&NumberOfRequests, "requests", "n", 1, "How many requests to do to the target host.")

	// Required flags
	rootCmd.Flags().StringVarP(&Host, "host", "s", "", "Target host.")
	rootCmd.MarkFlagRequired("host")
}
