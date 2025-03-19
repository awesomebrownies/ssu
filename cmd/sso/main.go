package main

/* Brainingstorming:

Firewall: Running Status Checks /
	Server Management:
Password and Authentication Modules: Ensuring basic requirements \
Remote Access Points: Reading details /
Least Privilege:
	User Management: Checking for insecure configurations -
	System CTL: Reading kernel settings
	System File Permissions: Reviewing files |
*/

//Optional action logging to files

import (
	"fmt"

	"greeter/lib/display"

	"github.com/spf13/cobra"
)

func main() {
	rootCmd := &cobra.Command{
		Use:   "sso",
		Short: "Performs a system security overview",
		Long:  "A CLI tool for running system status security checks on your computer",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Welcome to System Security Overview")
			menu := display.NewMenu("Select an option", "prompt")
			menu.AddItem("Option 1", "option1")

			menu.Display()
		},
	}

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
	}
}
