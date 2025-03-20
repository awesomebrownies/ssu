#!/bin/bash
sudo sed -i 's/^PermitRootLogin/PermitRootLogin no' /etc/ssh/sshd_config