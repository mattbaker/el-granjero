package main

// InventoryResponse is a response to an inventory query
type InventoryResponse struct {
	*BaseResponse
	Response struct {
		Data struct {
			Items []InventoryItem
		}
	}
}
