How to install this app from source code:

- download and install golang sdk https://dl.google.com/go/go1.14.linux-amd64.tar.gz
- add your service app on nameapps.conf after install it moves to /etc/monitorendpoint/nameapps.conf
- add your information email server on config.toml after install it moves to /etc/monitorendpoint/config.toml
- $ make compile
- $ sudo make install

Check the services
- $ sudo systemctl status dockermonitorgo
- $ sudo systemctl status dockermonitorbash
- $ sudo systemctl enable dockermonitorgo
- $ sudo systemctl enable dockermonitorbash

for see logs just type:
- $ journalctl -f -u monitorendpoint.service


install app not from source
- sudo make start

for uninstall is just simple :
- $ sudo make uninstall
