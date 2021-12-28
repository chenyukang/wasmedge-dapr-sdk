
package client

// todo finish the real abi

//export dapr_state
func daprState(action action, storeName string, key string, data []byte, consistency int32,metadata string) (res string, err string)
