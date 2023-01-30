package handlers

import "github.com/dayterr/gophkeeper_diploma/internal/storage"

type AsyncHandler struct{
	Storage storage.Storager
	JWT_Key string
}
