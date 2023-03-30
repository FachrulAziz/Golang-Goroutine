package belajargolanggoroutine

import (
	"fmt"
	"testing"
	"time"
)

func TestCreateChannel(t *testing.T) {
	channel := make(chan string)
	defer close(channel) // <= untuk menutup channel jika sudah tidak digunakan

	go func() {
		time.Sleep(2 * time.Second) // <= sebelum menampilkan, di beri jeda 2 detik oleh sleep
		channel <- "Fachrul Aziz"   // <= mengirim data ke channel
		fmt.Println("selesai mengirim data ke channel")
	}()

	terima := <-channel // menerima data dari channel
	fmt.Println(terima) // menampilkan data dari variable (terima)

	time.Sleep(5 * time.Second) // sebelum menampilkan data yang ada di variable "terima", data di jeda oleh sleep disamping ini
}
