package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

type Config struct {
	Port string `json: "port"`
	Cert string `json: "cert"`
	Key  string `json: "key"`
}

var config Config

func main() {
	loadConfig()

	http.HandleFunc("/", saludar)

	log.Printf("Servidor escuchando.... https://localhost:%s", config.Port)
	//http.ListenAndServe(":8080", nil)
	//uso de https
	/*err := http.ListenAndServeTLS(":443", "./certificates/cert.pem", "./certificates/key.pem", nil) //el handle es nil
	if err != nil {
		log.Println(err)
	}*/

	err := http.ListenAndServeTLS(config.Port, config.Cert, config.Key, nil) //el handle es nil
	if err != nil {
		log.Println(err)
	}
}

func saludar(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hola mundo"))
}

func loadConfig() {
	log.Println("Leyendo el archivo de configuración")

	b, err := ioutil.ReadFile("./config.json")
	if err != nil {
		log.Println("Ocurrio un error al intentar leer")
	}

	err = json.Unmarshal(b, &config)
	if err != nil {
		log.Println("Error de lectura")
	}

	log.Print(config.Port)

	log.Println("Leyendo el archivo de configuración leido")
}
