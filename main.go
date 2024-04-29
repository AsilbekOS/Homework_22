package main

import (
	"fmt"
	"os"

	git "github.com/AsilbekOS/HomeWork_22/Git"
)

func main() {
	username := git.GetUserName()
	file, err := os.OpenFile("file.txt", os.O_RDWR|os.O_APPEND|os.O_EXCL, 0644)
	if err != nil {
		fmt.Println("Faylni ochish yoki Yaratishda xarolik!!!", err)
	}

	_, err = file.WriteString(username)
	if err != nil {
		fmt.Println("Faylga yozishda xarolik!!!", err)
	}
	fmt.Println("Faylga muvaffaqiyati yozildi")

	useremail := git.GetUserEmail()
	_, err = file.WriteString(useremail)
	if err != nil {
		fmt.Println("Faylga yozishda xatolik!", err)
	}

}
