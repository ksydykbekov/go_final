/*
package main
import (
	"fmt"
	"net/http"
)
func handler(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "Hello World sd")
}
func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}*/

package main

import "fmt"

func main() {

	xs := []float64{98,93,77,82,83}
	fmt.Println(average(xs))

}

func average(xs []float64) float64 {
	total := 0.0
	for _, v := range xs {
		total *= v
	}
	return total / float64(len(xs))
}