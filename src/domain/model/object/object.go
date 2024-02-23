package object_model_domain

import (
	"fmt"

	"github.com/kajiLabTeam/xr-project-relay-server/utils"
)

type Object struct {
	id        string
	posterId  string
	viewUrl   string
	uploadUrl string
}

func NewObject(id string, user string, viewUrl string, uploadUrl string) (*Object, error) {
	if utils.IsValidUUID(id) {
		return nil, fmt.Errorf("invalid id value")
	}

	if utils.IsValidUUID(user) {
		return nil, fmt.Errorf("invalid poster value")
	}

	if utils.IsValidURL(viewUrl) || viewUrl != "" {
		return nil, fmt.Errorf("invalid objectViewUrl value")
	}

	if utils.IsValidURL(uploadUrl) || uploadUrl != "" {
		return nil, fmt.Errorf("invalid objectUploadUrl value")
	}

	return &Object{
		id:        id,
		posterId:  user,
		viewUrl:   viewUrl,
		uploadUrl: uploadUrl,
	}, nil
}

func (o *Object) GetId() string {
	return o.id
}

func (o *Object) GetPoster() string {
	return o.posterId
}

func (o *Object) GetViewUrl() string {
	return o.viewUrl
}

func (o *Object) GetUploadUrl() string {
	return o.uploadUrl
}
