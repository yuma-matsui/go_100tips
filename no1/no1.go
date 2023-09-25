// 変数のシャドーイング

package no1

import (
	"log"
	"net/http"
)

var tracing bool

func createClientWithTracing() (http.Client, error) {
	return http.Client{}, nil
}

func createDefaultClient() (http.Client, error) {
	return http.Client{}, nil
}

// NG例
func No1() {
	var client http.Client
	if tracing {
		client, err := createClientWithTracing()
		if err != nil {
			log.Fatal(err)
		}
		log.Println(client)
	} else {
		client, err := createDefaultClient()
		if err != nil {
			log.Fatal(err)
		}
		log.Println(client)
	}
}
