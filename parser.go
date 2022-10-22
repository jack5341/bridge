package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/user"
	"strings"
)

type keysAndIpAdresses struct {
	keys       []string
	ipAdresses []string
}

var sshKeys []string
var ipAdresses []string

func Parser() keysAndIpAdresses {
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
		keys := strings.Split(f.Name(), ".pub")

		if f.Name() == "known_hosts" {
			b, err := os.ReadFile(path + "/known_hosts")

			if err != nil {
				panic(err)
			}

			lines := strings.Split(string(b), "\n")

			for _, l := range lines {
				ipAdresses = append(ipAdresses, strings.Split(l, " ")[0])
			}
		}

		if len(keys) > 1 {
			sshKeys = append(sshKeys, keys...)
		}
	}

	return keysAndIpAdresses{
		keys:       sshKeys,
		ipAdresses: ipAdresses,
	}
}
