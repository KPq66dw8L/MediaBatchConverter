package models

type CustomError struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
}

type Folder struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Path string `json:"path"`
}
