package user

import (
	"apulse.ai/tzuchi-upmp/server/model"
	"apulse.ai/tzuchi-upmp/server/utils/jwt"
)

type (
	accessTokenResponse struct {
		AccessToken string `json:"accessToken"`
	}

	userResponse struct {
		Username string `json:"username"`
		FullName string `json:"fullName"`
		IDNumber string `json:"idNumber"`
		Email    string `json:"email"`
		Role     string `json:"role"`
	}

	usersResponse map[string]userResponse

	/*
		rolesOfUserResponse []string
	*/
)

func newAccessTokenResponse(datum *model.User) (accessTokenResponse, error) {
	response := accessTokenResponse{}
	if accessToken, err := jwt.Generate(datum.ID); err != nil {
		return response, err
	} else {
		response.AccessToken = accessToken
		return response, nil
	}
}

func newUserResponse(datum *model.User) userResponse {
	return userResponse{
		Username: datum.Username,
		FullName: datum.FullName,
		IDNumber: datum.IDNumber,
		Email:    datum.Email,
		Role:     datum.Role,
	}
}

func newUsersResponse(data []model.User) usersResponse {
	response := make(usersResponse)
	for _, datum := range data {
		response[datum.ID.String()] = newUserResponse(&datum)
	}
	return response
}

/*
func newRolesOfUserResponse(roles []string) rolesOfUserResponse {
	return roles
}
*/
