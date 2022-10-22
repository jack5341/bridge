package main

import (
	"fmt"
	"io/fs"
	"io/ioutil"
	"os"
	"os/user"
	"strings"
)

var SSHKeys []string
var IpAdresses []string

func init() {
	user, err := user.Lookup(os.ExpandEnv("$USER"))

	if err != nil {
		panic(err)
	}

	path := fmt.Sprintf("/Users/%s/.ssh", user.Username)
	files, err := ioutil.ReadDir(path)

	if err != nil {
		panic(err)
	}

	for _, f := range files {
		GetSshKeys(f, path)
		GetIpAdresses(f, path)
	}
}

func GetSshKeys(f fs.FileInfo, path string) []string {
	keys := strings.Split(f.Name(), ".pub")

	if len(keys) > 1 {
		SSHKeys = append(SSHKeys, keys...)
	}

	return SSHKeys
}

func GetIpAdresses(f fs.FileInfo, path string) []string {
	if f.Name() == "known_hosts" {
		b, err := os.ReadFile(path + "/known_hosts")

		if err != nil {
			panic(err)
		}

		lines := strings.Split(string(b), "\n")

		for _, l := range lines {
			IpAdresses = append(IpAdresses, strings.Split(l, " ")[0])
		}
	}

	return IpAdresses
}
