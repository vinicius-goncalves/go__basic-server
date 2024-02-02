package main

import "net/http"

var m map[string]string = map[string]string{"k": "v", "a": "b"}

func GETHandler(w http.ResponseWriter, r *http.Request) {

	data, _ := parser(m)
	body := []byte(data)

	w.Write(body)
}

func main() {

	http.Handle("/endpoint", http.HandlerFunc(GETHandler))

	http.ListenAndServe(":8080", nil)
}
