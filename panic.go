package main


import (
	"net/http"
	"log"
)
func main(){
	http.HandlerFunc("/", func(w http.ResponseWriter, r *http.Request){
		w.Write([]byte("Hello Go!"))
	})
	err := http.ListenAndServe(":9090")
	defer func(){
		if err := recover(); err != nil{
			log.Println("Error:", err)
		}
	}()
}