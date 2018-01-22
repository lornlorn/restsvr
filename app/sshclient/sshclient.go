package sshclient

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net"
	"time"

	"golang.org/x/crypto/ssh"
)

type SSH struct {
	Host, User, Password string
	Port                 int
}

func NewSSH(host, user, password string, port int) *SSH {
	return &SSH{
		Host:     host,
		User:     user,
		Password: password,
		Port:     port,
	}
}

func (s *SSH) Connect() (*ssh.Session, error) {
	pemBytes, err := ioutil.ReadFile("C:\\Users\\Lorn\\.ssh\\id_rsa")
	if err != nil {
		return nil, err
	}
	log.Println(string(pemBytes[:]))
	signer, err := ssh.ParsePrivateKey(pemBytes)
	if err != nil {
		return nil, err
	}
	config := &ssh.ClientConfig{
		User: s.User,
		Auth: []ssh.AuthMethod{
			// ssh.Password(s.Password),
			ssh.PublicKeys(signer),
		},
		Timeout: time.Second * 10,
		HostKeyCallback: func(hostname string, remote net.Addr, key ssh.PublicKey) error {
			return nil
		},
	}

	addr := fmt.Sprintf("%s:%d", s.Host, s.Port)
	client, err := ssh.Dial("tcp", addr, config)
	if err != nil {
		return nil, err
	}

	return client.NewSession()
}

func (s *SSH) Run(cmd string) (stdout, stderr io.Reader, err error) {
	session, err := s.Connect()
	if err != nil {
		log.Println(err)
		return
	}
	defer session.Close()

	stdout, _ = session.StdoutPipe()
	stderr, _ = session.StderrPipe()

	err = session.Run(cmd)
	return stdout, stderr, err
}

func (s *SSH) PrintRun(cmd string) {
	var res []byte

	// initial session
	session, err := s.Connect()
	if err != nil {
		log.Println(err)
		return
	}
	defer session.Close()

	// set stdout, stderr
	stdout, _ := session.StdoutPipe()
	stderr, _ := session.StderrPipe()

	// run command
	err = session.Run(cmd + ";" + cmd)
	if err != nil {
		res, _ = ioutil.ReadAll(stderr)
		log.Printf("%s execute failed ~> \n\033[31m%s\033[0m", s.Host, string(res))
		return
	}
	res, _ = ioutil.ReadAll(stdout)
	log.Printf("%s execute ok ~> \n\033[32m%s\033[0m", s.Host, string(res))
}
