package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"unsafe"

	dapr "github.com/dapr/go-sdk/client"
	"github.com/second-state/WasmEdge-go/wasmedge"
)

type host struct {
	fetchResult []byte
	client      dapr.Client
	ctx         context.Context
}

// Host function for writting memory
func (h *host) writeMem(_ interface{}, mem *wasmedge.Memory, params []interface{}) ([]interface{}, wasmedge.Result) {
	// write source code to memory
	pointer := params[0].(int32)
	mem.SetData(h.fetchResult, uint(pointer), uint(len(h.fetchResult)))

	return nil, wasmedge.Result_Success
}

// Host function for NewClient
func (h *host) newClient(_ interface{}, mem *wasmedge.Memory, params []interface{}) ([]interface{}, wasmedge.Result) {
	ctx := context.Background()
	client, err := dapr.NewClient()
	if err != nil {
		panic(err)
	}
	h.client = client
	h.ctx = ctx
	return []interface{}{*(*int32)(unsafe.Pointer(&client))}, wasmedge.Result_Success
}

func (h *host) newClientWithPort(_ interface{}, mem *wasmedge.Memory, params []interface{}) ([]interface{}, wasmedge.Result) {
	port := params[0].(int32)
	client, err := dapr.NewClientWithPort(string(port))
	if err != nil {
		panic(err)
	}
	h.client = client
	return []interface{}{client}, wasmedge.Result_Success
}

func (h *host) InvokeMethodWithContent(_ interface{}, mem *wasmedge.Memory, params []interface{}) ([]interface{}, wasmedge.Result) {
	ctx := context.Background()
	appIdPointer := params[0].(int32)
	appIdSize := params[1].(int32)
	appidByte, _ := mem.GetData(uint(appIdPointer), uint(appIdSize))
	appid := make([]byte, appIdSize)
	copy(appid, appidByte)

	methodPointer := params[2].(int32)
	methodSize := params[3].(int32)
	methodByte, _ := mem.GetData(uint(methodPointer), uint(methodSize))
	method := make([]byte, methodSize)
	copy(method, methodByte)

	verbPointer := params[4].(int32)
	verbSize := params[5].(int32)
	verbByte, _ := mem.GetData(uint(verbPointer), uint(verbSize))
	verb := make([]byte, verbSize)
	copy(verb, verbByte)

	contentTyPointer := params[6].(int32)
	contentTySize := params[7].(int32)
	contentTyByte, _ := mem.GetData(uint(contentTyPointer), uint(contentTySize))
	contentTy := make([]byte, contentTySize)
	copy(contentTy, contentTyByte)

	contentPointer := params[8].(int32)
	contentSize := params[9].(int32)
	contentByte, _ := mem.GetData(uint(contentPointer), uint(contentSize))
	content := make([]byte, contentSize)
	copy(content, contentByte)

	daprContent := &dapr.DataContent{
		ContentType: string(contentTy),
		Data:        content,
	}
	resp, err := h.client.InvokeMethodWithContent(ctx, string(appid), string(method), string(verb), daprContent)
	if err != nil {
		panic(err)
	}

	h.fetchResult = resp
	return []interface{}{len(resp)}, wasmedge.Result_Success

}

func runWasmHandle(writer http.ResponseWriter, reader *http.Request) {
	fmt.Println("Go: Args:", os.Args)
	/// Expected Args[0]: program name (./externref)
	/// Expected Args[1]: wasm file (funcs.wasm)

	/// Set not to print debug info
	wasmedge.SetLogErrorLevel()

	conf := wasmedge.NewConfigure(wasmedge.WASI)
	vm := wasmedge.NewVMWithConfig(conf)
	obj := wasmedge.NewImportObject("env")

	h := host{}

	funcWriteType := wasmedge.NewFunctionType(
		[]wasmedge.ValType{
			wasmedge.ValType_I32,
		},
		[]wasmedge.ValType{})
	hostWrite := wasmedge.NewFunction(funcWriteType, h.writeMem, nil, 0)
	obj.AddFunction("write_mem", hostWrite)

	// Add host functions into the import object
	newClientType := wasmedge.NewFunctionType(
		[]wasmedge.ValType{},
		[]wasmedge.ValType{
			wasmedge.ValType_I32,
		})

	hostNewClient := wasmedge.NewFunction(newClientType, h.newClient, nil, 0)
	obj.AddFunction("new_client", hostNewClient)

	// Add host functions into the import object
	newClientWithPortType := wasmedge.NewFunctionType(
		[]wasmedge.ValType{
			wasmedge.ValType_I32,
		},
		[]wasmedge.ValType{
			wasmedge.ValType_I32,
		})

	hostNewClientWithPort := wasmedge.NewFunction(newClientWithPortType, h.newClientWithPort, nil, 0)
	obj.AddFunction("new_client_with_port", hostNewClientWithPort)

	// Add host functions into the import object
	invokeMethodWithContentType := wasmedge.NewFunctionType(
		[]wasmedge.ValType{
			wasmedge.ValType_I32,
			wasmedge.ValType_I32,
			wasmedge.ValType_I32,
			wasmedge.ValType_I32,
			wasmedge.ValType_I32,
			wasmedge.ValType_I32,
			wasmedge.ValType_I32,
			wasmedge.ValType_I32,
			wasmedge.ValType_I32,
			wasmedge.ValType_I32,
		},
		[]wasmedge.ValType{
			wasmedge.ValType_I32,
		})

	InvokeMethodWithContent := wasmedge.NewFunction(invokeMethodWithContentType, h.InvokeMethodWithContent, nil, 0)
	obj.AddFunction("invoke_method_with_content", InvokeMethodWithContent)

	vm.RegisterImport(obj)

	//wasm_file := os.Args[1]
	wasm_file := "./demo/target/wasm32-wasi/release/demo.wasm"

	vm.LoadWasmFile(wasm_file)
	vm.Validate()
	vm.Instantiate()

	r, _ := vm.Execute("run_dapr")
	fmt.Printf("The parameter len is: %d\n", r[0])

	obj.Release()
	vm.Release()
	conf.Release()

	writer.Header().Set("Content-Type", "plain/text")
	fmt.Fprintf(writer, "%s\n", "invoked wasm file")
}

func main() {
	http.HandleFunc("/api/hello", runWasmHandle)
	println("listen to 8080 ...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
