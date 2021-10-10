package main

import (
	"fmt"
	"github.com/rclone/rclone/fs/config"
	"github.com/rclone/rclone/fs/config/obscure"
)

func main() {
	for _, section := range config.FileSections() {
		// can change "password" to "password2" if you want
		pass := config.FileGet(section, "password")
		fmt.PrintLn("showing section")
		fmt.PrintLn(section)
		if pass == "" {
			fmt.PrintLn("skipping section because empty pass")
		} else {
			fmt.PrintLn("password is:")
			fmt.Println(section, obscure.MustReveal(pass))
		}
	}
}
