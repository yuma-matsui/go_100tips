// 変数のシャドーイング
// -> 内側のブロック内で変数名を再宣言した際に発生

package no1

import (
	"fmt"
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
func NG() {
	var client http.Client

	// ブロック内で変数をシャドーイングしてしまっている
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

func SolutionA() {
	var client http.Client

	// 一時変数の活用
	if tracing {
		c, err := createClientWithTracing()
		if err != nil {
			log.Fatal(err)
		}
		client = c
		log.Panicln(client)
	} else {
		c, err := createDefaultClient()
		if err != nil {
			log.Fatal(err)
		}
		client = c
		log.Println(client)
	}
}

func SolutionB() {
	// client, errを変数宣言して直接代入
	var client http.Client
	var err error

	if tracing {
		client, err = createClientWithTracing()
	} else {
		client, err = createDefaultClient()
	}

	// client, errに関する処理を共通化可能
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(client)
}
