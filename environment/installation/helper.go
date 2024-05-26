package installation

import (
	"bit-agent/bitwarden"
	"bit-agent/util/cli"
	"os"
)

func Install() {
	cli.Debug("Checking for root privileges...")
	if !isRoot() {
		cli.Error("The installation must be done as root.")
		os.Exit(1)
	}

	bitwarden.CheckInstallation()

	cli.Debug("Retrieving the executable path...")
	path, success := getExecutablePath()
	if !success {
		cli.Error("Cannot retrieve the executable path.")
		os.Exit(1)
	}

	cli.Debug("Checking for the destination folder existance...")
	if !doesFolderExists() {
		cli.Debug("Creating the destination folder...")
		if !createDestinationFolder() {
			cli.Error("Cannot create the destination folder.")
			os.Exit(1)
		}
	}

	cli.Debug("Copying the executable...")
	if !copyExecutable(path) {
		cli.Error("Failed to copy the executable.")
		os.Exit(1)
	}

	cli.Debug("Setting up the executable permissions...")
	if !setupPermission() {
		cli.Error("Failed to setup executable permissions...")
		os.Exit(1)
	}

	cli.Debug("Checking for the symlink existance...")
	symlink, success := doesSymlinkExist()
	if !success {
		cli.Warning("Failed to check for the symlink existance.")
	}

	if !symlink {
		cli.Debug("Creating the symlink...")
		if !createSymlink() {
			cli.Error("Error creating the symlink.")
			os.Exit(1)
		}
	}

	cli.Success("The installation was successful.")
}
