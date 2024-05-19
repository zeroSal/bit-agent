package main

import (
	bitwardenHelper "bit-agent/helper/bitwarden"
	sshHelper "bit-agent/helper/ssh"
	"bit-agent/util/cli"
)

func main() {
	cli.Text(cli.Logo())

	cli.Section("Authentication")
	session := bitwardenHelper.Authenticate()
	cli.Success("The authentication was successfully completed.")

	cli.Section("Key(s) retrieving")
	folder := bitwardenHelper.RetrieveSshFolder(session)
	keys := bitwardenHelper.RetrieveSshKeys(session, folder)
	cli.Success("The key(s) has been successully loaded.")

	cli.Section("Agent setup")
	sshHelper.StartAgent(keys)
	cli.Success("The agent has been successfully started.")

	select {}
}
