package helper

import (
	"bit-agent/service/bitwarden"
	"bit-agent/util/cli"
	"os"
)

func Authenticate() (session string) {
	cli.Debug("Checking for authentication...")
	authenticated, errOut, success := bitwarden.IsAuthenticated()
	if !success {
		cli.Error("Cannot check for authentication.\n" + errOut)
		os.Exit(1)
	}

	if authenticated {
		return unlock()
	}

	return login()
}

func RetrieveKey(session string, name string) (key string) {
	cli.Debug("Retrieving the \"" + name + "\" key...")
	sshKey := bitwarden.GetNotesItem(session, name)
	if sshKey == "" {
		cli.Warning("The SSH key retrieved seems empty.")
	}

	return sshKey
}

func login() (session string) {
	email := cli.Ask("Bitwarden email: ")
	password := cli.AskHidden("Bitwarden master password: ")

	cli.Debug("Logging into Bitwarden...")
	session, errOut, success := bitwarden.Login(email, string(password))
	if !success {
		cli.Error("Login failed.\n" + errOut)
		os.Exit(1)
	}

	return session
}

func unlock() (session string) {
	password := cli.AskHidden("Bitwarden master password: ")

	cli.Debug("Unlocking the vault...")
	session, errOut, success := bitwarden.Unlock(string(password))
	if !success {
		cli.Error("Unlock failed.\n" + errOut)
		os.Exit(1)
	}

	return session
}
