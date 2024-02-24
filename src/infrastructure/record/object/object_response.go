package object_record

import object_model_domain "github.com/kajiLabTeam/xr-project-relay-server/domain/model/object"

type ObjectResponse struct {
	Id        string `json:"id"  binding:"required,uuid"`
	PosterId  string `json:"posterId"  binding:"required,uuid"`
	SpotId    string `json:"spotId"  binding:"required,uuid"`
	Extension string `json:"extension"  binding:"required"`
	UploadUrl string `json:"viewUrl,omitempty"  binding:"required"`
}

type GetObjectBySpotIdResponse struct {
	Id       string `json:"id"  binding:"required,uuid"`
	PosterId string `json:"posterId"  binding:"required,uuid"`
	SpotId   string `json:"spotId"  binding:"required,uuid"`
	ViewUrl  string `json:"viewUrl"  binding:"required,url"`
}

func (gobsir *GetObjectBySpotIdResponse) ToDomainObject() (*object_model_domain.Object, error) {
	domainObject, err := object_model_domain.NewObject(
		gobsir.Id,
		gobsir.PosterId,
		gobsir.SpotId,
		gobsir.ViewUrl,
	)
	if err != nil {
		return nil, err
	}

	return domainObject, nil
}

type GetObjectCollectionBySpotIdsResponse struct {
	Id      string           `json:"id"  binding:"required,uuid"`
	Objects []ObjectResponse `json:"objects"  binding:"required"`
}

func (gobsir *GetObjectCollectionBySpotIdsResponse) ToDomainObjectCollection() (object_model_domain.ObjectCollection, error) {
	var domainObjectCollection object_model_domain.ObjectCollection
	for _, object := range gobsir.Objects {
		domainObject, err := object_model_domain.NewObject(
			object.Id,
			object.PosterId,
			object.SpotId,
			object.UploadUrl,
		)
		if err != nil {
			return nil, err
		}

		domainObjectCollection.AddObject(domainObject)
	}

	return domainObjectCollection, nil
}

type CreateObjectResponse struct {
	Id        string `json:"id"  binding:"required,uuid"`
	PosterId  string `json:"posterId"  binding:"required,uuid"`
	SpotId    string `json:"spotId"  binding:"required,uuid"`
	UploadUrl string `json:"uploadUrl,omitempty"  binding:"required"`
}

func (cor *CreateObjectResponse) ToDomainObject() (*object_model_domain.Object, error) {
	domainObject, err := object_model_domain.NewObject(
		cor.Id,
		cor.PosterId,
		cor.SpotId,
		cor.UploadUrl,
	)
	if err != nil {
		return nil, err
	}
	return domainObject, nil
}
