package cli

import "fmt"

func Logo(version string) string {
	logo := `
┳┓•   ┏┓          [ • Unofficial Bitwarden SSH agent • ]
┣┫┓╋━━┣┫┏┓┏┓┏┓╋   [ • Luca Saladino, Lecco, Italy • ]
┻┛┗┗  ┛┗┗┫┗ ┛┗┗   [ • %s • ]
         ┛`

	return fmt.Sprintf(logo, version)
}
