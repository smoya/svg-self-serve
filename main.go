package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/kelseyhightower/envconfig"
	"github.com/smoya/svg-self-serve/svg"
)

type Config struct {
	Port int `default:"8080"`
}

func main() {
	var c Config
	if err := envconfig.Process("svgselfserve", &c); err != nil {
		log.Fatal(err.Error())
	}

	log.Printf("Listening at http://localhost:%v/generate.svg", c.Port)

	http.Handle("/generate.svg", http.HandlerFunc(generateSVG))
	if err := http.ListenAndServe(fmt.Sprintf(":%v", c.Port), nil); err != nil {
		log.Fatal(err)
	}
}

func generateSVG(w http.ResponseWriter, r *http.Request) {
	configMap := make(map[string]string)
	for k, mv := range r.URL.Query() {
		configMap[k] = strings.Join(mv, ",")
	}
	c := svg.NewConfigFromMap(configMap)
	svg.Generate(c, w)
}
