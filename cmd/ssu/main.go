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
	"os"
	"os/exec"
	"strings"

	"greeter/lib/display"

	"github.com/spf13/cobra"
)

func selectMenu(prompt string) {
	if prompt == "../" {
		prompt = "System Security Overview"
	}
	menu := display.NewMenu("Select an option", prompt)

	if prompt == "System Security Overview" || prompt == "../" {
		menu.AddItem("Firewall")
		menu.AddItem("Remote Access Points")
		menu.AddItem("Least Privilege")
	} else {
		menu.AddItem("../")
	}

	switch prompt {
	case "Firewall":
		menu.AddItem(":Configure UFW")
	case "Remote Access Points":
		menu.AddItem(":Configure SSH")
	case "Least Privilege":
		menu.AddItem(":Set File Permissions")
		menu.AddItem(":Disable Root Access")
		menu.AddItem(":Disable Guest Account & Greeter Remote Login")
	}

	result, err := menu.Display()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}

	if _, ok := result.(int); ok {
		fmt.Println("")
		fmt.Println("Exit")
	} else if _, ok := result.(string); ok {
		if strings.Contains(result.(string), ":") {
			executeCommand(result.(string))
		} else {
			selectMenu(result.(string))
		}
	} else {
		fmt.Printf("selected option of unexpected type: %T with value: %v\n", result, result)
	}
}

// Split up the spaces
// Process each command per \n
func runCommand(commands string) {
	for _, line := range strings.Split(commands, "\n") {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		args := strings.Fields(line)
		cmd := exec.Command(args[0], args[1:]...)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		err := cmd.Run()
		if err != nil {
			fmt.Printf("%s\n", line)
		}
	}
}

func executeCommand(prompt string) {
	switch prompt {
	case ":Configure UFW":
		fmt.Println("Configure IPv4 in /etc/sysctl.conf")
		runCommand(`sudo sed -i 's/^net.ipv4.tcp_syncookies/net.ipv4.tcp_syncookies=1' /etc/sysctl.conf
		sudo sed -i 's/^net.ipv4.ip_forward/net.ipv4.ip_forward=0' /etc/sysctl.conf
		sudo sysctl --system`)
		fmt.Println("Configure UFW")
		runCommand(`sudo ufw enable
		sudo ufw default deny incoming
		sudo ufw default allow outgoing`)
	case ":Configure SSH":
		fmt.Println("Modifying sshd_config")
		runCommand(`sudo sed -i 's/^PermitRootLogin/PermitRootLogin no' /etc/ssh/sshd_config`)
	case ":Set File Permissions":
		fmt.Println("Modifying core file permissions")
		runCommand(`sudo chmod 644 /etc/passwd
		sudo chmod 640 /etc/shadow
		sudo chmod 644 /etc/group
		sudo chmod 640 /etc/gshadow
		sudo chmod 440 /etc/sudoers
		sudo chmod 644 /etc/ssh/sshd_config
		sudo chmod 644 /etc/fstab
		sudo chmod 600 /boot/grub/grub.cfg
		sudo chmod 644 /etc/hostname
		sudo chmod 644 /etc/hosts
		sudo chmod 600 /etc/crypttab
		sudo chmod 640 /var/log/auth.log
		sudo chmod 644 /etc/apt/sources.list
		sudo chmod 644 /etc/systemd/system/*.service
		sudo chmod 644 /etc/resolv.conf`)
	case ":Disable Root Access":
		fmt.Println("Disabling sudo su")
		runCommand(`sudo sed -i '/^auth   	sufficient pam_rootok.so/ c\#auth   	sufficient pam_rootok.so/' /etc/pam.d/su`)
	case ":Disable Guest Account & Greeter Remote Login":
		fmt.Println("Disabling guest account & greeter remote login")
		runCommand(`sudo touch /usr/share/lightdm/lightdm.conf.d/50-no-guest.conf
		sudo sed -i '/^\[SeatDefaults\]/!b;n;c\allow-guest=false' /usr/share/lightdm/lightdm.conf.d/50-no-guest.conf
		sudo touch /etc/lightdm/lightdm.conf.d/50-no-remote-login.conf
		sudo sed -i '/^\[SeatDefaults\]/!b;n;c\greeter-show-remote-login=false' /etc/lightdm/lightdm.conf.d/50-no-remote-login.conf`)
	}
}

func main() {
	rootCmd := &cobra.Command{
		Use:   "ssu",
		Short: "Performs a system security overview",
		Long:  "A CLI tool for running system security checks on your computer",
		Run: func(cmd *cobra.Command, args []string) {
			selectMenu("System Security Overview")
		},
	}

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
	}
}
