
package client

// todo finish the real abi

//export dapr_state
func daprState(action action, storeName string, key string, data []byte, consistency int32,metadata string) (res string, err string)

//export dapr_bulk_state
func daprBulkState(action action) //todo