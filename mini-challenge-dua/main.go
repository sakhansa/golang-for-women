package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	args := os.Args
	// fmt.Println(myName[1])
	if len(os.Args) == 1 {
		fmt.Println("Please enter name or id.")
	} else {
		arg := args[1]
		var participants = []Participant{
			{nama: "Khansa", alamat: "Jl. RE Martadinata", pekerjaan: "Accountant", alasan: "Menambah wawasan bahasa pemrograman"},
			{nama: "Erri", alamat: "Jl. Fatahilah", pekerjaan: "BE Developer", alasan: "Menunjang pekerjaan"},
			{nama: "Rizka", alamat: "Jl. Asia Afrika", pekerjaan: "Android Developer", alasan: "Mengisi waktu dengan kegiatan bermanfaat"},
			{nama: "Asrie", alamat: "Jl. Ahmad Yani", pekerjaan: "FE Developer", alasan: "Belajar bahasa pemrograman baru"},
			{nama: "Hazrina", alamat: "Jl. Stasiun Wonokromo", pekerjaan: "IOS Developer", alasan: "Untuk mencari freelance"},
		}

		id, participant, message := findParticipant(arg, participants)

		if message == "" {
			fmt.Printf("ID : %v\n", id)
			fmt.Printf("Nama : %s\n", participant.nama)
			fmt.Printf("Pekerjaan : %s\n", participant.pekerjaan)
			fmt.Printf("Alamat : %s\n", participant.alamat)
		} else {
			fmt.Println(message)
		}
	}
}

type Participant struct {
	nama      string
	alamat    string
	pekerjaan string
	alasan    string
}

func findParticipant(arg string, participants []Participant) (id int, participant Participant, errmessage string) {
	message := ""
	if len(participants) > 0 {
		for index, value := range participants {
			id = index + 1
			identity, errorParse := strconv.Atoi(arg)

			if errorParse == nil {
				if len(participants) >= identity {
					return identity, participants[identity-1], message
				} else {
					message = "ID not found."
				}
			} else if errorParse != nil {
				if strings.EqualFold(arg, value.nama) {
					return id, value, message
				} else {
					message = "Name not found."
				}
			} else {
				message = "Please enter character or number."
			}
		}
	} else {
		message = "Please enter character or number."
	}
	return id, participant, message
}
