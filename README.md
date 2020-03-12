How to install this app from source code:

- download and install golang sdk https://dl.google.com/go/go1.14.linux-amd64.tar.gz
- add your list of service app name on nameapps.conf file after install it the configuration of nameapps.conf move to /etc/monitorendpoint/nameapps.conf
- add your information email server on config.toml file after install it the configuration of config.toml move to /etc/monitorendpoint/config.toml
- $ make compile
- $ sudo make install 

Stop the services
- $ sudo systemctl stop monitorendpoint.service


Start the services
- $ sudo systemctl start monitorendpoint.service


Check the services
- $ sudo systemctl status monitorendpoint.service

for see logs just type:
- $ journalctl -f -u monitorendpoint.service


install app not from source
- sudo make start

for uninstall is just simple :
- $ sudo make uninstall
