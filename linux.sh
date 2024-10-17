#!/bin/sh

pause(){
   read -p "
Press [ENTER] to continue" placeholder
}
update_and_configure_settings(){
   echo sudo apt update -y
   wait
   sudo apt upgrade -y
}
configure_firewall(){
   sudo ufw enable
   wait
   sudo ufw status
}
verify_user_admin_list(){
   echo "Looking inside... /etc/passwd"
   cat /etc/passwd
   python3 user_admin_list.py
}
user_prompt(){
   read -p "Select from one of the following choices:
   [1] Add a new user
   [2] Remove a current user and their directories
  > " prompt
   case "${prompt}" in
      1 ) add_user;;
      2 ) remove_user;;
   esac
}
add_user(){
   read -p "Enter username to add: " username
   pass=$(perl -e 'print crypt("1Soup3rS*Cure!", "supersalter3000")')
   sudo useradd -m -p ${pass} ${username}
   echo "Username ${username} has been added. Password: 1Soup3rS*Cure!"
}
remove_user(){
   read -p "Enter username to remove: " username
   sudo userdel -r ${username}
   echo "Username ${username} has been deleted."
}
group_management(){
   read -p "Select from one of the following choices:
   [1] Create a group
   [2] Remove a group
   [3] Add a user to a group
   [4] Remove a user from a group
  > " prompt
  case "${prompt}" in
     1 ) read -p "Enter group name to create: " name; sudo groupadd ${name}; echo "Added group ${name}";;
     2 ) read -p "Enter a group name to remove: " name; sudo groupdel ${name}; echo "Removed group ${name}";;
     3 ) read -p "Enter a group name: " group; read -p "Enter a user to be added: " user; sudo adduser ${user} ${group};;
     4 ) read -p "Enter a group name: " group; read -p "Enter a user to be removed: " user; sudo deluser ${user} ${group};;
  esac
}
GUEST_CONFIG_FILE="/etc/lightdm/lightdm.conf.d/50-no-guest.conf"
REMOTE_LOGIN_CONFIG_FILE="/etc/lightdm/lightdm.conf.d/50-no-remote-login.conf"
disable_guest_and_remote_login(){
   #Write new file to disable guest access
   echo -e "[SeatDefaults]\nallow-guest=false\n" | sudo tee "$GUEST_CONFIG_FILE" > /dev/null
   #Write new file to disable remote login
   echo -e "[SeatDefaults]\ngreeter-show-remote-login=false\n" | sudo tee "$REMOTE_LOGIN_CONFIG_FILE" > /dev/null
}



while true; do
   read -p "Select from one of the following choices:
   [1] Check for Updates & Configure Update Settings
   [2] Verify User & Admin List
      [2.1] User Management
      [2.2] Group Management
   [3] Disable Guest Account & Greeter Remote Login
   [4] Update Password Requirements & Change All Passwords (not finished)
   [5] Disable Root Access (not finished)
   [6] Configure Firewall (UFW)
   [7] Check all file permissions (not finished)
   [8] Remove Malware & Hacking Tools (not finished)
  > " OPTION
   case "${OPTION}" in
       1 ) echo "Check for Updates & Configure Settings \n"; update_and_configure_settings;;
       2 ) echo "Verify User & Admin List \n"; verify_user_admin_list;;
       2.1 ) echo "User Management"; user_prompt;;
       2.2 ) echo "Group Management"; group_management;;
       3 ) echo "Disable Guest Account & Greeter Remote Login"; disable_guest_and_remote_login;;
       6 ) echo "Configure Firewall (UFW) \n"; configure_firewall;;
   esac
   pause
   echo ""
done
