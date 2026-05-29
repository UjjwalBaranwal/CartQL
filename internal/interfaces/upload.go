// package interfaces  -> give interface for the UploadProvider
package interfaces

import "mime/multipart"

// UploadProvider interface give interface function
type UploadProvider interface {
	UploadFile(file *multipart.FileHeader, path string) (string, error)
	DeleteFile(path string) error
}
