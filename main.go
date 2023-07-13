package main

import (
	"io"
	"log"
	"math"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gliderlabs/ssh"
)

type Tunnel struct {
	w io.Writer
	done chan struct{}
}

var tunnels = map[int]chan Tunnel{}

func main() {

	go func() {
		http.HandleFunc("/", handleRequest)
		log.Fatal(http.ListenAndServe(":3000", nil))
	}()
	ssh.Handle(func(s ssh.Session) {
        id := rand.Intn(math.MaxInt)
		tunnels[id] = make(chan Tunnel)
		log.Println("tunnel ID ->", id)
		tunnel := <-tunnels[id]
		log.Println("tunnel is READY")
		_, err := io.Copy(tunnel.w, s)
		if err != nil {
			log.Fatal(err)
		}
		close(tunnel.done)
		s.Write([]byte("We are DONE"))
     })  

     log.Fatal(ssh.ListenAndServe(":2222", nil))
}

func handleRequest(w http.ResponseWriter, r *http.Request) {
	idstr := r.URL.Query().Get("id")

	id, err := strconv.Atoi(idstr)
	if err != nil {
		w.Write([]byte("Bad Request"))
		return
	}
	tunnel, ok := tunnels[id]
	if !ok {
		w.Write([]byte("Tunnel not found"))
		return
	}
	done := make(chan struct{})
	tunnel <- Tunnel{
		w: w,
		done: done,
	}
	<-done
}