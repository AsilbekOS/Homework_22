package git

import (
	"bytes"
	"fmt"
	"os/exec"
)

func GetUserName() string {
	username := exec.Command("git", "config", "user.name")
	var cmd bytes.Buffer
	username.Stdout = &cmd
	err := username.Run()
	if err != nil {
		fmt.Println("Usernameni o'qishda xatolik!!!", err)
	}

	return cmd.String()
}