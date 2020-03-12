#!/bin/bash

# This code created by Andrian Latif 02-18-2020
# compile your code before using this code
# happy code 

__CURRENT_DIRECTORY=$(pwd)

logic(){
    rm -rf /etc/systemd/system/monitorendpoint.service
    # for golang app
    echo "[Unit]
Description= monitorendpoint
After=monitorendpoint.service
Requires=monitorendpoint.service

[Service]
User=root
ExecStart=/usr/bin/monitorendpoint

[Install]
WantedBy=default.target 
    " > /etc/systemd/system/monitorendpoint.service
    
    shell_command
}

shell_command(){
    rm -rf /usr/bin/monitorendpoint
    rm -rf /etc/monitorendpoint
    mkdir /etc/monitorendpoint
    chmod 777 /etc/monitorendpoint
    cp $__CURRENT_DIRECTORY/app/monitorendpoint /usr/bin/
    cp $__CURRENT_DIRECTORY/config.toml /etc/monitorendpoint
    cp $__CURRENT_DIRECTORY/nameapps.conf /etc/monitorendpoint
    systemctl daemon-reload
    systemctl enable monitorendpoint
    systemctl start monitorendpoint
    echo "successfuly install"
    echo "for see the log service of monitorendpoint just type journalctl -f -u monitorendpoint.service"
    journalctl -f -u monitorendpoint.service
}

install_start(){
    rm -rf /etc/systemd/system/monitorendpoint.service
    # for golang app
    echo "[Unit]
Description= monitorendpoint
After=monitorendpoint.service
Requires=monitorendpoint.service

[Service]
User=root
ExecStart=/usr/bin/monitorendpoint

[Install]
WantedBy=default.target 
    " > /etc/systemd/system/monitorendpoint.service
    rm -rf /usr/bin/monitorendpoint
    rm -rf /etc/monitorendpoint
    mkdir /etc/monitorendpoint
    chmod 777 /etc/monitorendpoint
    cp $__CURRENT_DIRECTORY/build/monitorendpoint /usr/bin/
    cp $__CURRENT_DIRECTORY/config.toml /etc/monitorendpoint
    cp $__CURRENT_DIRECTORY/nameapps.conf /etc/monitorendpoint
    systemctl daemon-reload
    systemctl enable monitorendpoint
    systemctl start monitorendpoint
    echo "successfuly install"
    echo "for see the log service of monitorendpoint just type journalctl -f -u monitorendpoint.service"
    journalctl -f -u monitorendpoint.service
}


if [[ "$EUID" -ne 0 ]]; then
    echo "Sorry, you need to run this app as root"
    exit
else
    if [[ $@ == "start" ]];then
        install_start
    # elif [ $@ == "start" ]; then
    #     echo $@
    else
        logic
    fi  
fi