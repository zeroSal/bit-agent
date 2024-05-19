package helper

import (
	"bit-agent/service/ssh"
	"bit-agent/util/cli"
)

func StartAgent(key string) {
	cli.Debug("Starting SSH agent...")
	go ssh.StartSSHAgent(key)

	cli.Success("SSH agent started.")
	cli.Notice("The agent is listening via the " + ssh.GetSocketPath() + " socket.")
}
