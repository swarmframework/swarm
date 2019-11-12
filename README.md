# SWARM

SWARM is the framework that was built to allow a red team to demonstrate adversarial honeypots and the risk of automated discovery services on a network. Currently SWARM has two different types of honeypots that can be deployed: Drones and Hives.  

Drones are executables compiled for either Windows or Linux and are then uploaded to a compromised host. Hives are physical devices (Raspberry Pi 2 or 3 currently) that are configured to run a drone listener on startup. The hive has been designed to have a fully encrypted memory card and once the hive is running all files including the encryption keys are deleted from the disk. The GPIO header is disabled on the Raspberry Pi and only the Ethernet port is enabled. If a 3G modem is used with the hive then a USB interface is enabled to support the external modem.   

The framework currently supports SSH, SMB, Telnet, FTP and SNMP. SWARM is almost completely written in GoLang except for parts of the SMB module which were borrowed from the Impacket Python framework. Currently SWARM uses a C2 over TLS setup to allow exfiltration of credentials as soon as they are captured either on a drone or a hive. This allows for the red team to quickly begin lateral movement and further system compromise during an engagement. 

## Built With

* [impacket](https://github.com/SecureAuthCorp/impacket) - Used in the SMB listener  

* [gliderlabs/ssh](https://github.com/gliderlabs/ssh) - Used in the SSH listener

## Demo Videos

* [SSH Drone Demo](https://youtu.be/Ja-4nGJxB3M)

* [SNMP Drone Demo](https://youtu.be/AFABEulsDjc)

* [SMB Hive Image Demo](https://youtu.be/g0QPJVjhcl4)

* [SMB Hive Demo](https://youtu.be/4NWiAj4K7ZM)

## Presentation Slides

* [Presentation](https://docs.google.com/presentation/d/e/2PACX-1vR2o57Fep8XQYL-7Uh1SNRsZkAWkH1nbYNUVG1hWkKzCVDO945uT49AklS6dp_llkFu65kPVM1YSHW6/pub?start=false&loop=false&delayms=3000)

## Installation Instructions
Coming Soon

## Authors

* **Jacob Griffith** - *@maddhatt3r2*
* **Tim Wright** - *@redteam_hacker*

## Contact

* **swarmframework@gmail.com**

## Acknowledgments

* **@Th3M00se**
