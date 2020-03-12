package main

/**
* Created by Andrian Latif 09-03-2020
**/

import (
	"bufio"
	"fmt"
	"log"
	"net/smtp"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/BurntSushi/toml"
	"github.com/fatih/color"
	resty "gopkg.in/resty.v1"
)

type Config struct {
	DataName dataname
}

type dataname struct {
	CONFIG_SMTP_HOST string
	CONFIG_SMTP_PORT int
	CONFIG_EMAIL     string
	CONFIG_PASSWORD  string
	SEND_TO          string
	SEND_CC1         string
	SEND_CC2         string
}

func main() {
	mainProgram()
}

func getClient(service string, number_loop int) {
	// GET request
	resp, err := resty.R().Get(service)

	// convert to string
	a := strconv.Itoa(int(number_loop))
	b := strconv.Itoa(int(resp.StatusCode()))

	// explore response object
	// fmt.Printf("\nError: %v", err)
	// fmt.Printf("\nResponse Status Code: %v", resp.StatusCode())
	// fmt.Printf("\nResponse Status: %v", resp.Status())
	// fmt.Printf("\nResponse Time: %v", resp.Time())
	// fmt.Printf("\nResponse Received At: %v", resp.ReceivedAt())
	// fmt.Printf("\nResponse Body: %v", resp) // or resp.String() or string(resp.Body())
	// // more...

	// untuk POST
	// resp, err := resty.R().
	// 	SetHeader("Content-Type", "application/json").
	// 	SetBody(`"datakirims":"testuser"}`).
	// 	// or SetResult(AuthSuccess{}).
	// 	Post("https://192.168.20.6:8080/coba/kirimdata")

	if err == nil {
		if resp.StatusCode() == 200 {
			fmt.Println(number_loop, "status code", resp.StatusCode(), "on service ", service)
		} else {
			color.Red(a + " error endpoint with status code : " + b + " on service " + service)
			Mail(service, resp.StatusCode(), "")
		}
	} else {
		color.Red(a + " endpoint failur : " + err.Error())
		Mail(service, resp.StatusCode(), err.Error())
	}
}

func mainProgram() {
	for true {
		time.Sleep(2 * time.Second)

		file, err := os.Open("/etc/monitorendpoint/nameapps.conf")

		if err != nil {
			log.Fatalf("failed opening file: nameapps.conf at /etc/monitorendpoint directory")
		}

		scanner := bufio.NewScanner(file)
		scanner.Split(bufio.ScanLines)
		var txtlines []string

		for scanner.Scan() {
			txtlines = append(txtlines, scanner.Text())
		}

		file.Close()
		for num_app, eachline_name_of_app := range txtlines {
			getClient(eachline_name_of_app, num_app)
			time.Sleep(1 * time.Second)
		}
	}
}

func Mail(s string, status_code int, resp string) {
	var conf Config
	sts_code := strconv.Itoa(status_code)
	if _, err := toml.DecodeFile("/etc/monitorendpoint/config.toml", &conf); err != nil {
		fmt.Println(err)
	}

	to := []string{conf.DataName.SEND_TO, conf.DataName.SEND_TO}

	//if you want to send cc
	cc := []string{conf.DataName.SEND_CC1, conf.DataName.SEND_CC2}

	subject := "Service " + s + " is Offline"
	message := "Service " + s + " is offline" + "\nOffline time is " + __getTime().String() + "\nStatus Code : " + sts_code + "\nLog Error : " + resp

	err := sendMail(to, cc, subject, message)
	if err != nil {
		log.Fatal(err.Error())
	}

	log.Println("Mail sent!")
}

func sendMail(to []string, cc []string, subject, message string) error {
	var conf Config
	if _, err := toml.DecodeFile("/etc/monitorendpoint/config.toml", &conf); err != nil {
		fmt.Println(err)
	}

	fmt.Println("running .... ")
	body := "From: " + conf.DataName.CONFIG_EMAIL + "\n" +
		"To: " + strings.Join(to, ",") + "\n" +
		"Cc: " + strings.Join(cc, ",") + "\n" +
		"Subject: " + subject + "\n\n" +
		message

	auth := smtp.PlainAuth("", conf.DataName.CONFIG_EMAIL, conf.DataName.CONFIG_PASSWORD, conf.DataName.CONFIG_SMTP_HOST)
	smtpAddr := fmt.Sprintf("%s:%d", conf.DataName.CONFIG_SMTP_HOST, conf.DataName.CONFIG_SMTP_PORT)

	fmt.Println("check ", smtpAddr)
	err := smtp.SendMail(smtpAddr, auth, conf.DataName.CONFIG_EMAIL, append(to, cc...), []byte(body))

	if err != nil {
		return err
	}
	return nil
}

// get hostname instance
func ___getHostname() string {
	hostname, err := os.Hostname()
	if err != nil {
		fmt.Println(err)
	}
	return hostname
}

// get time
func __getTime() time.Time {
	var time_not_utc = time.Now()
	return time_not_utc
}
