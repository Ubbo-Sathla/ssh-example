package main

import (
	"fmt"
	"golang.org/x/crypto/ssh"
	"log"
)

type sshConn struct {
	conn *ssh.Client
}

type Endpoint struct {
	Host string
	Port int
}

func (endpoint *Endpoint) String() string {
	return fmt.Sprintf("%s:%d", endpoint.Host, endpoint.Port)
}

func publicKeyAuthFunc() ssh.AuthMethod {

	key := []byte(``)
	// Create the Signer for this private key.
	signer, err := ssh.ParsePrivateKey(key)
	if err != nil {
		log.Fatal("ssh key signer failed", err)
	}
	return ssh.PublicKeys(signer)
}
func sshClientConn() (*ssh.Client, error) {
	serverEndpoint := &Endpoint{
		Host: "myserver.com",
		Port: 22,
	}
	sshConfig := &ssh.ClientConfig{
		User: "jon",
		Auth: []ssh.AuthMethod{
			publicKeyAuthFunc(),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	sconn, err := ssh.Dial("tcp", serverEndpoint.String(), sshConfig)
	if err != nil {
		return nil, err
	}
	return sconn, nil
}
