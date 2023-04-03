package main

import (
	"fmt"
	"os"
	"strings"

	color "github.com/fatih/color"
	flag "github.com/ogier/pflag"
)

// Available flags
var (
	user string
)

// Initialize default flags at app startup
func init() {
	flag.StringVarP(&user, "user", "u", "", "Search Users")
}

// Parse flags
func main() {
	flag.Parse()

	// if the user does not supply flags, print app usage
	if flag.NFlag() == 0 {
		printOptions()
	}

	var users = strings.Split(user, ",")
	fmt.Printf("Searching user(s): %s\n", users)
	for _, u := range users {
		result := getUsers(u)
		color.Cyan(`Username:		%s`, result.Login)
		color.Blue(`Name:			%s`, result.Name)
		color.Green(`Email:			%s`, result.Email)
		color.HiMagenta(`Bio:		%s`, result.Bio)
		fmt.Println("")
	}
}

func printOptions() {
	fmt.Printf("Usage: %s [options]\n", os.Args[0])
	fmt.Println("Options:")
	flag.PrintDefaults()
	os.Exit(1)
}
