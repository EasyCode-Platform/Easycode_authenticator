package controller

import (
	"github.com/EasyCode/Easycode-authenticator-backend/src/authenticator"
	"github.com/EasyCode/Easycode-authenticator-backend/src/model"
	"github.com/EasyCode/Easycode-authenticator-backend/src/utils/tokenvalidator"
)

type Controller struct {
	Storage               *model.Storage
	Cache                 *model.Cache
	Drive                 *model.Drive
	RequestTokenValidator *tokenvalidator.RequestTokenValidator
	Authenticator         *authenticator.Authenticator
}

func NewController(storage *model.Storage, cache *model.Cache, drive *model.Drive, validator *tokenvalidator.RequestTokenValidator, auth *authenticator.Authenticator) *Controller {
	return &Controller{
		Storage:               storage,
		Cache:                 cache,
		Drive:                 drive,
		RequestTokenValidator: validator,
		Authenticator:         auth,
	}
}
