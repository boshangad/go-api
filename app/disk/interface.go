package disk

type UploadInterface interface {
	Upload(filename string, path string) (err error)
}
