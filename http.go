package main

import (
	"io"
	"net/http"
	"time"
)

func main() {
	go client()

	time.Sleep(time.Second * 5)
	go server()
}

func client() {
	headers := make(map[string]any)

	headers["Content-Type"] = "application/json"

	req, _ := http.NewRequest("GET", "http://localhost:8088/modules", nil)
	cli := http.Client{
		Timeout: time.Second * 20,
	}
	resp, err := cli.Do(req)
	if err != nil {
		println(err)
		return
	}
	defer resp.Body.Close()

	println(resp.Status)
	msg, _ := io.ReadAll(resp.Body)
	println(string(msg))
}

func server() {
	http.HandleFunc("/modules", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		start := time.Now()
		for {
			if time.Since(start) > time.Second*5 {
				println("waiting...")
				break
			}
		}
		w.Write([]byte(`{"message": "Hello, World!"}`))
	})

	http.ListenAndServe(":8088", nil)
}
