#!/bin/bash

#CREATED BY ANDRIAN LATIF on 02-18-2020

logic(){
    systemctl stop monitorendpoint
    rm -rf /etc/systemd/system/monitorendpoint.service
    rm -rf /usr/bin/monitorendpoint
    rm -rf /etc/monitorendpoint
    echo "successfuly uninstall "
    systemctl daemon-reload
}

if [[ "$EUID" -ne 0 ]]; then
    echo "Sorry, you need to run this app as root"
    exit
else
    logic    
fi
