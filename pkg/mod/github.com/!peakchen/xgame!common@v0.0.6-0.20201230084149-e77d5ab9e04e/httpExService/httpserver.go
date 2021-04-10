// add by stefan

package httpExService

import (
	"net/http"
	//"fmt"
	"encoding/json"
	"log"
	"strings"
	"time"

	define "github.com/Peakchen/xgameCommon/define"
)

func StartHttpService(addr string, handler define.HandlerFunc) {
	log.Printf("http addr: %s.", addr)
	server := &http.Server{
		Addr:           addr,
		Handler:        handler,
		ReadTimeout:    30 * time.Second,
		WriteTimeout:   30 * time.Second,
		MaxHeaderBytes: 1 << 16,
	}

	err := server.ListenAndServe()
	if err != nil {
		log.Fatal("http ListenAndServe: ", err)
		return
	}
}

func Get(url string, jsonst interface{}) error {
	resp, err := http.Get(url)
	if err != nil {
		log.Println("htp Get, err :", err)
		return err
	}

	defer resp.Body.Close()
	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(jsonst)
	return err
}

func Post(url string, jsonst interface{}) error {
	resp, err := http.Post(url, "application/x-www-form-urlencoded", strings.NewReader("name=cjb"))
	if err != nil {
		log.Println("htp post, err :", err)
		return err
	}

	defer resp.Body.Close()
	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(jsonst)
	return err
}
