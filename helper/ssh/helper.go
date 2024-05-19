package helper

import (
	"bit-agent/service/ssh"
	"bit-agent/util/cli"
)

func StartAgent(keys []string) {
	cli.Debug("Starting SSH agent...")
	go ssh.StartSSHAgent(keys)
	cli.Notice("The agent is listening via the " + ssh.GetSocketPath() + " socket.")
}
