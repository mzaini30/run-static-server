package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"os/exec"
	"runtime"
	"time"
)

var randomizer = rand.New(rand.NewSource(time.Now().UTC().UnixNano()))
var letters = []rune("123456789")

func randomString(length int) string {
	b := make([]rune, length)
	for i := range b {
		b[i] = letters[randomizer.Intn(len(letters))]
	}
	return string(b)
}

func openbrowser(url string) {
	var err error

	switch runtime.GOOS {
	case "linux":
		err = exec.Command("xdg-open", url).Start()
	case "windows":
		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	case "darwin":
		err = exec.Command("open", url).Start()
	default:
		err = fmt.Errorf("unsupported platform")
	}
	if err != nil {
		log.Fatal(err)
	}

}

func main() {
	lokasi := "./web"

	_, err_path := os.Stat("./web/path.txt")
	if err_path == nil {
		path_byte, err_path_ambil := os.ReadFile("./web/path.txt")
		if err_path_ambil != nil {
			log.Fatal(err_path_ambil)
		}
		path := string(path_byte)
		lokasi = "./web/" + path
	}

	fs := http.FileServer(http.Dir(lokasi))
	http.Handle("/", fs)

	dat, err := os.ReadFile("./web/port.txt")
	if err != nil {
		log.Fatal(err)
	}

	portnya := string(dat)

	log.Print("Tutup aplikasi ini jika sudah tidak digunakan")
	openbrowser("http://localhost:" + portnya)
	err = http.ListenAndServe(":"+portnya, nil)
	if err != nil {
		log.Fatal(err)
	}
}
