package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

type Config struct {
	Port    string `json:"port"`
	CertPem string `json:"cert_pem"`
	KeyPem  string `json:"key_pem"`
	MyToken string `json:"my_token"`
}

var config Config

func main() {
	loadConfig()

	http.HandleFunc("/", saludar)
	http.HandleFunc("fb", fbbot)

	log.Printf("Servidor escuchando.... https://localhost%s", config.Port)
	println(config.CertPem)
	println(config.KeyPem)
	//http.ListenAndServe(":8080", nil)
	//uso de https
	/*err := http.ListenAndServeTLS(":443", "./certificates/cert.pem", "./certificates/key.pem", nil) //el handle es nil
	if err != nil {
		log.Println(err)
	}*/

	//err := http.ListenAndServeTLS(config.Port, config.Cert, config.Key, nil) //el handle es nil
	err := http.ListenAndServeTLS(config.Port, config.CertPem, config.KeyPem, nil) //el handle es nil

	if err != nil {
		log.Println(err)
	}
}

func saludar(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hola mundo"))
}
func fbbot(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		vt := r.URL.Query().Get("hub.verify_token")
		if vt == config.MyToken {
			hc := r.URL.Query().Get("hub.challenge")
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(hc))
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("token no valido"))
		return
	}
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
