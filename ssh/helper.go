package ssh

import (
	"bit-agent/util/cli"
)

func StartAgent(keys []string) {
	cli.Debug("Starting SSH agent...")
	go startSSHAgent(keys)
	cli.Notice("The agent is listening via the " + getSocketPath() + " socket.")
}
