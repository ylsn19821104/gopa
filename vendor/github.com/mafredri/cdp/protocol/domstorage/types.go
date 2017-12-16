// Code generated by cdpgen. DO NOT EDIT.

package domstorage

// StorageID DOM Storage identifier.
type StorageID struct {
	SecurityOrigin string `json:"securityOrigin"` // Security origin for the storage.
	IsLocalStorage bool   `json:"isLocalStorage"` // Whether the storage is local storage (not session storage).
}

// Item DOM Storage item.
type Item []string