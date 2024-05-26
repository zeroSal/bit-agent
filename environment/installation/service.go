package installation

import (
	"fmt"
	"io/fs"
	"os"
	"os/exec"
)

const destinationFolder = "/opt/bit-agent"
const symlinkPath = "/usr/local/bin/bit-agent"

var destinationExecutable = destinationFolder + "/bit-agent"

func getExecutablePath() (errOut string, success bool) {
	p, err := os.Executable()
	if err != nil {
		return err.Error(), false
	}

	return p, true
}

func doesFolderExists() (exists bool) {
	return doesInodeExist(destinationFolder)
}

func doesInodeExist(path string) (exists bool) {
	_, err := os.Stat(path)
	return err == nil
}

func createDestinationFolder() (success bool) {
	err := os.MkdirAll(destinationFolder, os.ModePerm)
	return err == nil
}

func copyExecutable(executablePath string) (success bool) {
	cmd := exec.Command("cp", "-f", executablePath, destinationExecutable)
	err := cmd.Run()
	return err == nil
}

func setupPermission() (success bool) {
	err := os.Chmod(destinationExecutable, 0505)
	return err == nil
}

func doesSymlinkExist() (symlink bool, success bool) {
	if !doesInodeExist(symlinkPath) {
		return false, true
	}

	stat, err := os.Lstat(symlinkPath)
	if err != nil {
		fmt.Print(err.Error())
		return false, false
	}

	return stat.Mode()&fs.ModeSymlink != 0, true
}

func createSymlink() (success bool) {
	err := os.Symlink(destinationExecutable, symlinkPath)
	return err == nil
}

func isRoot() bool {
	return os.Geteuid() == 0
}
