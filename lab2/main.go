package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func main() {
	http.HandleFunc("/lab2/", lab2)
	http.ListenAndServe(":8080", nil)
}

func lab2(w http.ResponseWriter, r *http.Request) {

	r.ParseForm()
	headerContentTtype := r.Header.Get("Content-Type")
	if headerContentTtype != "application/x-www-form-urlencoded" {
		w.WriteHeader(http.StatusUnsupportedMediaType)
		return
	}
	value1 := r.Form.Get("n")
	value2 := r.Form.Get("start")

	n, err := strconv.Atoi(value1)
	if err != nil {
		fmt.Println(value1, " не является целым числом.")
	}

	start, err := strconv.Atoi(value2)
	if err != nil {
		fmt.Println(value2, " не является целым числом.")
	}

	fmt.Println("Ответ: ", xorOperation(n, start))

}

func xorOperation(n int, start int) int {
	array := make([]int, n)
	var xor int
	xor = 0
	for i := 0; i < n; i++ {
		array[i] = start + 2*i
		xor = xor ^ array[i]
	}
	return xor
}
