# system security utilities
This tool only supports Linux. The code inside uses executes console commands for each executable item
## Installation
Install `Go` on Linux
`snap install go`
Navigate to the correct directory
`cd cmd/ssu`
Run the program
`go run main.go`
## Modules

### Firewall
* Configure UFW (uncomplicated firewall)
* Configure System CTL IPv4
### Remote Access Points
* Modify SSH Configuration
### Least Privilege
* Disable Root (sudo su)
* Disable Guest User & Greeter Remote Login
* System File Permissions