package main

import (
	"fmt"
	"github.com/Dev-TempusFugit/API/handler"
	"github.com/Dev-TempusFugit/API/storage"
	"log"
	"net/http"
)

func main(){

	store := storage.NewMemory()
	mux := http.NewServeMux()
	handler.RoutePerson(mux, &store)

	log.Println("Servidor iniciado en el puerto 8080")

	err := http.ListenAndServe(":8080", mux)
	 if err != nil{
	 	fmt.Printf("Error en el servidor %v\n",err)
	 }

}
