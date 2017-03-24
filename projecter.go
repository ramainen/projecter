package main

import (
	"fmt"
	"html"
	"log"
	"net"
	"net/http"
	"os"
	"runtime"
	"time"

	"github.com/micmonay/keybd_event"
)

func main() {

	http.HandleFunc("/send", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		kb, err := keybd_event.NewKeyBonding()
		if err != nil {
			panic(err)
		}

		if runtime.GOOS == "linux" {
			time.Sleep(2 * time.Second)
		}

		switch r.FormValue("key") {
		case "RIGHT":
			kb.SetKeys(keybd_event.VK_RIGHT)
		case "LEFT":
			kb.SetKeys(keybd_event.VK_LEFT)
		case "UP":
			kb.SetKeys(keybd_event.VK_UP)
		case "DOWN":
			kb.SetKeys(keybd_event.VK_DOWN)
		case "ESC":
			kb.SetKeys(keybd_event.VK_ESC)
		case "F5":
			kb.SetKeys(keybd_event.VK_F5)
		case "SPACE":
			kb.SetKeys(keybd_event.VK_SPACE)
		}

		//	kb.HasSHIFT(true) //set shif is pressed

		err = kb.Launching()
		if err != nil {
			panic(err)
		}
		fmt.Fprintf(w, "sent, %q", html.EscapeString(r.FormValue("key")))
	})

	host, _ := os.Hostname()
	addrs, _ := net.LookupIP(host)
	for _, addr := range addrs {
		if ipv4 := addr.To4(); ipv4 != nil {
			fmt.Println("Adress: " + ipv4.String() + ":8085")
		}
	}

	log.Fatal(http.ListenAndServe(":8085", nil))

}
