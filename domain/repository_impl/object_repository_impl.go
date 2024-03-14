package repository_impl

import (
	application_models_domain "github.com/kajiLabTeam/xr-project-relay-server/domain/models/application"
	object_models_domain "github.com/kajiLabTeam/xr-project-relay-server/domain/models/object"
	object_collection_models_domain "github.com/kajiLabTeam/xr-project-relay-server/domain/models/object_collection"
	user_models_domain "github.com/kajiLabTeam/xr-project-relay-server/domain/models/user"
)

type ObjectRepositoryImpl interface {
	FindForSpotId(
		spotId string,
		user *user_models_domain.User,
		application *application_models_domain.Application,
	) (*object_models_domain.Object, error)

	FindForSpotIds(
		spotIds []string,
		user *user_models_domain.User,
		application *application_models_domain.Application,
	) (*object_collection_models_domain.ObjectCollection, error)

	Save(
		spotId string,
		user *user_models_domain.User,
		object *object_models_domain.Object,
		application *application_models_domain.Application,
	) (*object_models_domain.Object, error)
}
