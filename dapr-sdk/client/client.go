package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	dapr "github.com/dapr/go-sdk/client"
)

func daprClientSendV1(data []byte) {
	ctx := context.Background()

	// create the client
	client, err := dapr.NewClient()
	if err != nil {
		panic(err)
	}

	content := &dapr.DataContent{
		ContentType: "text/plain",
		Data:        data,
	}

	resp, err := client.InvokeMethodWithContent(ctx, "dapr-demo", "/api/hello", "post", content)
	if err != nil {
		panic(err)
	}
	log.Printf("dapr-wasmedge-go method api/image has invoked, response: %s", string(resp))
	fmt.Printf("Image classify result: %q\n", resp)
}

func daprClientSendV2(data []byte, w http.ResponseWriter) {
	ctx := context.Background()

	// create the client
	client, err := dapr.NewClient()
	if err != nil {
		panic(err)
	}

	content := &dapr.DataContent{
		ContentType: "text/plain",
		Data:        data,
	}

	resp, err := client.InvokeMethodWithContent(ctx, "dapr-demo", "/api/hello", "post", content)
	if err != nil {
		panic(err)
	}
	log.Printf("dapr-wasmedge-go method api/image has invoked, response: %s", string(resp))
	fmt.Printf("Image classify result: %q\n", resp)
	w.Header().Set("Content-Type", "plain/text")
	fmt.Fprintf(w, "%s", resp)
}

func testDaprClientV1() {
	time.Sleep(8 * time.Second)
	body := "hello"
	daprClientSendV1([]byte(body))
}

func testDaprClientV2(w http.ResponseWriter, r *http.Request) {
	body := "hello"
	daprClientSendV2([]byte(body), w)
}

func main() {
	testDaprClientV1()

	//http.HandleFunc("/api/hello", testDaprClientV2)
	//println("listen to 8080 ...")
	//log.Fatal(http.ListenAndServe(":8080", nil))
}
