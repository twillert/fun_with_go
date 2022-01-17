package main

import (
	"bytes"
	"fmt"
	"log"
	"os"

	"golang.org/x/crypto/ssh"
)

func create_session(c *ssh.Client, cmd string) *ssh.Session {
	fmt.Println("inside create_session")
	s, err := c.NewSession()
	if err != nil {
		log.Fatal("Failed to create session: ", err)
	}
	defer s.Close()
	s.Run(cmd)
	return s
}

func main() {
	// var hostKey ssh.PublicKey
	// An SSH client is represented with a ClientConn.
	//
	// To authenticate with the remote server you must pass at least one
	// implementation of AuthMethod via the Auth field in ClientConfig,
	// and provide a HostKeyCallback.
	config := &ssh.ClientConfig{
		User: os.Getenv("USER"),
		Auth: []ssh.AuthMethod{
			ssh.Password("qq"),
		},
		// HostKeyCallback: ssh.FixedHostKey(hostKey),
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}
	client, err := ssh.Dial("tcp", "localhost:22", config)
	if err != nil {
		log.Fatal("Failed to dial: ", err)
	}
	defer client.Close()

	// Each ClientConn can support multiple interactive sessions,
	// represented by a Session.
	session, err := client.NewSession()
	if err != nil {
		log.Fatal("Failed to create session: ", err)
	}
	defer session.Close()

	// Once a Session is created, you can execute a single command on
	// the remote side using the Run method.
	var b bytes.Buffer
	session.Stdout = &b
	if err := session.Run("/usr/bin/ls -l"); err != nil {
		log.Fatal("Failed to run: " + err.Error())
	}
	fmt.Println(b.String())

	session2, err := client.NewSession()
	if err != nil {
		log.Fatal("Failed to create session2: ", err)
	}
	defer session2.Close()
	var b2 bytes.Buffer
	session2.Stdout = &b2
	if err := session2.Run("echo blah"); err != nil {
		log.Fatal("Failed to run: " + err.Error())
	}
	fmt.Println(b2.String())
}
