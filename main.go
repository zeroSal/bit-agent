package main

import (
	bitwardenHelper "bit-agent/helper/bitwarden"
	sshHelper "bit-agent/helper/ssh"
	"bit-agent/util/cli"
)

func main() {
	cli.Text(cli.Logo())

	session := bitwardenHelper.Authenticate()

	name := cli.Ask("Bitwarden item name: ")
	key := bitwardenHelper.RetrieveKey(session, name)

	sshHelper.StartAgent(key)

	select {}
}
