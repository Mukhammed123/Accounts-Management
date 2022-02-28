package user

import (
	"apulse.ai/tzuchi-upmp/server/model"
	"apulse.ai/tzuchi-upmp/server/utils/argon2id"
)

type (
	addUserRequest struct {
		Username string `json:"username" validate:"username"`
		Password string `json:"password" validate:"password"`
		FullName string `json:"fullName" validate:"required"`
		IDNumber string `json:"idNumber" validate:"id_number"`
		Email    string `json:"email" validate:"omitempty,email"`
		Role     string `json:"role" validate:"must=VS CR HN RN NP"`
	}

	addUsersRequest []addUserRequest

	getUserByUserIDRequest struct {
		UserID string `param:"userID" validate:"uuid4"`
	}

	updateUserByUserIDRequest struct {
		UserID string `param:"userID" validate:"uuid4"`

		FullName *string `json:"fullName"`
		// IDNumber *string `json:"idNumber" validate:"omitempty,id_number"`
		Email *string `json:"email" validate:"omitempty,email"`
		Role  *string `json:"role" validate:"omitempty,must=VS CR HN RN NP"`
	}

	deleteUserByUserIDRequest struct {
		UserID string `param:"userID" validate:"uuid4"`
	}

	/*
		getRolesOfUserByUserIDRequest struct {
			UserID string `param:"userID" validate:"uuid4"`
		}

		setRolesOfUserByUserIDRequest struct {
			UserID string `param:"userID" validate:"uuid4"`

			Roles []string `validate:"must=admin"`
		}
	*/

	checkUserExistsRequest struct {
		Username string `json:"username" validate:"required"` // validate:"username"
		Password string `json:"password" validate:"required"`
	}

	signInRequest struct {
		Username string `json:"username" validate:"required"` // validate:"username"
		Password string `json:"password" validate:"password"`
	}

	updateCurrentUserRequest struct {
		FullName *string `json:"fullName"`
		// IDNumber *string `json:"idNumber" validate:"omitempty,id_number"`
		Email *string `json:"email" validate:"omitempty,email"`
		Role  *string `json:"role" validate:"omitempty,must=VS CR HN RN NP"`
	}
)

func (r *addUserRequest) toModel() *model.User {
	password := argon2id.Hash(r.Username, r.Password)
	return &model.User{
		Username: r.Username,
		Password: password,
		FullName: r.FullName,
		IDNumber: r.IDNumber,
		Email:    r.Email,
		Role:     r.Role,
	}
}

func (rs *addUsersRequest) toModels() []model.User {
	data := make([]model.User, len(*rs))
	for i, r := range *rs {
		data[i] = *r.toModel()
	}
	return data
}

func (r *updateUserByUserIDRequest) toPatch() map[string]interface{} {
	patch := make(map[string]interface{})
	if r.FullName != nil {
		patch["full_name"] = *r.FullName
	}
	// if r.IDNumber != nil {
	// 	patch["id_number"] = *r.IDNumber
	// }
	if r.Email != nil {
		patch["email"] = *r.Email
	}
	if r.Role != nil {
		patch["role"] = *r.Role
	}
	return patch
}

func (r *updateCurrentUserRequest) toPatch() map[string]interface{} {
	patch := make(map[string]interface{})
	if r.FullName != nil {
		patch["full_name"] = *r.FullName
	}
	// if r.IDNumber != nil {
	// 	patch["id_number"] = *r.IDNumber
	// }
	if r.Email != nil {
		patch["email"] = *r.Email
	}
	if r.Role != nil {
		patch["role"] = *r.Role
	}
	return patch
}
