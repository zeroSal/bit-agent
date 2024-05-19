package ssh

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"

	"golang.org/x/crypto/ssh"
	"golang.org/x/crypto/ssh/agent"
)

const socketPath string = "~/.bitwarden_ssh_agent.sock"

func GetSocketPath() string {
	return socketPath
}

func StartSSHAgent(privateKey string) {
	sockPath, err := expandPath(socketPath)
	if err != nil {
		fmt.Println("Failed to expand the socket path.")
		os.Exit(1)
	}

	if _, err := os.Stat(sockPath); err == nil {
		fmt.Println("Socket found. Removing it...")
		os.Remove(sockPath)
	}

	l, err := net.Listen("unix", sockPath)
	if err != nil {
		log.Fatalf("Failed to listen on socket: %v", err)
	}
	defer l.Close()

	sshAgent := agent.NewKeyring()
	key, err := ssh.ParseRawPrivateKey([]byte(privateKey))
	if err != nil {
		log.Fatalf("Failed to parse private key: %v", err)
	}

	addedKey := agent.AddedKey{PrivateKey: key}
	if err := sshAgent.Add(addedKey); err != nil {
		log.Fatalf("Failed to add key to agent: %v", err)
	}

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-signalChan
		os.Remove(sockPath)
		os.Exit(0)
	}()

	for {
		conn, err := l.Accept()
		if err != nil {
			log.Printf("Failed to accept connection: %v", err)
			continue
		}
		go agent.ServeAgent(sshAgent, conn)
	}
}

func expandPath(path string) (string, error) {
	if len(path) > 0 && path[0] == '~' {
		home, err := os.UserHomeDir()
		if err != nil {
			return "", err
		}
		return filepath.Join(home, path[1:]), nil
	}
	return path, nil
}
