package cmd

import (
	"bufio"
	"errors"
	"fmt"
	"github.com/clybs/comms/connections"
	"github.com/spf13/cobra"
	"os"
	"strings"
	"time"
)

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "comms",
	Short: "Program id usage counter",
	Long:  `Program id usage counter that counts the number of comms using the same program id`,
	Run:   start,
}

var file string
var pipe string
var cnx connections.Mapper

func init() {
	RootCmd.PersistentFlags().StringVar(&file, "file", "input.txt", "File input")
	RootCmd.PersistentFlags().StringVar(&pipe, "pipe", "<->", "Pipe symbol to use")
}

func check(e error) {
	if e != nil {
		fmt.Println(e)
		os.Exit(1)
	}
}

func readLines(filePath string) ([]string, error) {
	fileDetails, err := os.Open(filePath)
	check(err)
	defer fileDetails.Close()

	var lines []string
	scanner := bufio.NewScanner(fileDetails)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func readFile() map[string]string {

	lines, err := readLines(file)
	check(err)

	entries := make(map[string]string)
	for _, v := range lines {
		entry := strings.Split(v, pipe)

		if len(entry) > 1 {
			key, value := entry[0], entry[1]
			entries[key] = value
		}
	}

	// Check if data was found
	if len(entries) == 0 {
		err := errors.New("No data found in file")
		check(err)
	}

	return entries
}

func start(cmd *cobra.Command, args []string) {
	start := time.Now()

	// Read the file
	lines := readFile()

	// Create connections
	entries := cnx.CreateConnections(lines)

	// Identify groups
	groups := cnx.CreateGroups(entries)

	elapsed := time.Since(start)
	fmt.Printf("Data loaded and processed in %s", elapsed)
	fmt.Println("")

	// Scan for user input
	scanner := bufio.NewScanner(os.Stdin)

	// Scan user input
	for scanner.Scan() {
		input := scanner.Text()
		commandParam := getCommandParams(input)

		for _, v := range commandParam {
			// Get program id summary
			programsInTheGroup := "None"

			if len(entries[v]) > 0 {
				programsInTheGroup = strings.Join(entries[v][:], ", ")
			}

			fmt.Println("Program ID:*", v)
			fmt.Println("   * Programs count that are in the group that contains program ID", v, ":", len(entries[v]))
			fmt.Println("   * Programs in the group:", programsInTheGroup)
			fmt.Println("   * Group count total:", len(groups))
		}
	}

	// Gotcha errors
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "Reading standard input:", err)
	}
}

// Execute adds all child commands to the root command
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
