package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/dapr/go-sdk/service/common"
	daprd "github.com/dapr/go-sdk/service/http"
)

func hello(_ context.Context, in *common.InvocationEvent) (out *common.Content, err error) {
	ans := string(in.Data) + " received"

	fmt.Printf("Response result: %q\n", ans)
	out = &common.Content{
		Data:        []byte(ans),
		ContentType: in.ContentType,
		DataTypeURL: in.DataTypeURL,
	}
	return out, nil
}

func main() {
	s := daprd.NewService(":9003")

	if err := s.AddServiceInvocationHandler("/api/hello", hello); err != nil {
		log.Fatalf("error adding invocation handler: %v", err)
	}

	if err := s.Start(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("error listenning: %v", err)
	}
}
