# create by Andrian Latif on 02-18-2020

# make
# for see the log just type on terminal
# journalctl -f -u monitorendpoint.service

install: 
	@./install.sh

# make compile
# this stage will compile the app from source code 
compile:
	@rm -rf app
	@mkdir app
	@go mod download
	@go build -o app/monitorendpoint main.go
	@chmod +x app/monitorendpoint

# for delete all of services
uninstall:
	@./uninstall.sh

start:
	@chmod +x install.sh
	@./install.sh start