package main

//InventoryItem represents a single item in the InventoryItem
type InventoryItem struct {
	ItemHash       int
	ItemID         string
	Quantity       int
	TransferStatus int
	State          int
	BucketHash     int
}
