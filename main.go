package main

import (
	"bit-agent/bitwarden"
	"bit-agent/environment/installation"
	"bit-agent/ssh"
	"bit-agent/util/cli"
	"flag"
)

var version = "dev"

func main() {
	var installFlag = flag.Bool("install", false, "Install bit-agent on the system.")

	flag.Parse()

	cli.Text(cli.Logo(version))

	if *installFlag {
		install()
		return
	}

	run()
}

func run() {
	cli.Section("Authentication")
	session := bitwarden.Authenticate()
	cli.Success("The authentication was successfully completed.")

	cli.Section("Sync thread startup")
	bitwarden.StartSync(session)
	cli.Success("The sync thread has been started.")

	cli.Section("Key(s) retrieving")
	folder := bitwarden.RetrieveSshFolder(session)
	keys := bitwarden.RetrieveSshKeys(session, folder)
	cli.Success("The key(s) has been successully loaded.")

	cli.Section("Agent setup")
	ssh.StartAgent(keys)
	cli.Success("The agent has been successfully started.")

	select {}
}

func install() {
	cli.Section("Installation")
	installation.Install()
}
