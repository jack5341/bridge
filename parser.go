package main

import (
	"fmt"
	"os"
	"strings"
)

var SSHKeys []string
var IpAddresses []string

func init() {
	userHOME := os.ExpandEnv("$HOME")

	sshPath := fmt.Sprintf("%s/.ssh", userHOME)
	files, err := os.ReadDir(sshPath)
	if err != nil {
		panic(err)
	}

	for _, f := range files {
		GetSshKeys(f.Name())
	}

	ipdAddrs, err := GetIpAddresses(sshPath)
	if err != nil {
		return
	}

	IpAddresses = ipdAddrs
}

func GetSshKeys(fileName string) []string {
	keyName := strings.Split(fileName, ".pub")

	if len(keyName) > 1 {
		SSHKeys = append(SSHKeys, keyName[0])
	}

	return SSHKeys
}

func GetIpAddresses(sshPath string) ([]string, error) {
	var ips []string
	b, err := os.ReadFile(sshPath + "/known_hosts")

	if err != nil {
		return nil, err
	}

	lines := strings.Split(string(b), "\n")

	for _, l := range lines {
		ips = append(ips, strings.Split(l, " ")[0])
	}

	return ips, nil
}
