package hdss

import (
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/google/uuid"

	"apulse.ai/tzuchi-upmp/server/model"
	"apulse.ai/tzuchi-upmp/server/service/requester"
	"apulse.ai/tzuchi-upmp/server/store"
)

type Service struct {
	store              *store.Store
	username           string
	password           string
	accessToken        string
	idMapperOfUsers    map[uuid.UUID]uuid.UUID
	idMapperOfPatients map[uuid.UUID]uuid.UUID
}

func NewService(store *store.Store) *Service {
	return &Service{
		store:              store,
		username:           "hdsp.sync",
		password:           os.Getenv("JWT_SECRET_KEY"),
		idMapperOfUsers:    make(map[uuid.UUID]uuid.UUID),
		idMapperOfPatients: make(map[uuid.UUID]uuid.UUID),
	}
}

func (s *Service) makeRequest(method string, url string, requestData map[string]interface{},
	needAccessToken bool, queries map[string]string, allowedStatusCode []int) (interface{}, error) {
	var accessToken *string
	if needAccessToken {
		accessToken = &s.accessToken
	}
	return requester.MakeRequestToHDSS(method, url, requestData, accessToken, queries, allowedStatusCode)
}

func (s *Service) makeSignInRequest() (map[string]interface{}, error) {
	if responseData, err := s.makeRequest(http.MethodPost, "users/login/", map[string]interface{}{
		"username": s.username,
		"password": s.password,
	}, false, nil, []int{http.StatusOK}); err != nil {
		return nil, err
	} else {
		return responseData.(map[string]interface{}), nil
	}
}

func (s *Service) makeGetUsersRequest() ([]interface{}, error) {
	queries := map[string]string{
		"include_deleted": "True",
	}
	if responseData, err := s.makeRequest(http.MethodGet, "users/", nil, true, queries, []int{http.StatusOK}); err != nil {
		return nil, err
	} else {
		return responseData.([]interface{}), nil
	}
}

func (s *Service) makeGetPatientsRequest() ([]interface{}, error) {
	if responseData, err := s.makeRequest(http.MethodGet, "users/patients/", nil, true, nil, []int{http.StatusOK}); err != nil {
		return nil, err
	} else {
		return responseData.([]interface{}), nil
	}
}

func (s *Service) makeGetShiftSchedulesRequest(startAt *time.Time, endAt *time.Time) ([]interface{}, error) {
	queries := make(map[string]string)
	if startAt != nil {
		queries["duty_start_at"] = strconv.FormatInt(startAt.Unix(), 10)
	}
	if endAt != nil {
		queries["duty_end_at"] = strconv.FormatInt(endAt.Unix(), 10)
	}
	if responseData, err := s.makeRequest(http.MethodGet, "shift_schedule/", nil, true, queries, []int{http.StatusOK}); err != nil {
		return nil, err
	} else {
		return responseData.([]interface{}), nil
	}
}

func (s *Service) makeGetHDBedsRequest() ([]interface{}, error) {
	queries := map[string]string{
		"is_active": "1",
	}
	if responseData, err := s.makeRequest(http.MethodGet, "beds/", nil, true, queries, []int{http.StatusOK}); err != nil {
		return nil, err
	} else {
		return responseData.([]interface{}), nil
	}
}

func (s *Service) makeGetScheduleOfHDBedsRequest(startAt *time.Time, endAt *time.Time) ([]interface{}, error) {
	queries := make(map[string]string)
	if startAt != nil {
		queries["start"] = strconv.FormatInt(startAt.Unix(), 10)
	}
	if endAt != nil {
		queries["end"] = strconv.FormatInt(endAt.Unix(), 10)
	}
	if responseData, err := s.makeRequest(http.MethodGet, "beds/schedule/", nil, true, queries, []int{http.StatusOK}); err != nil {
		return nil, err
	} else {
		return responseData.([]interface{}), nil
	}
}

func (s *Service) SetAccessToken() error {
	if responseData, err := s.makeSignInRequest(); err != nil {
		return err
	} else {
		s.accessToken = responseData["access_token"].(string)
		return nil
	}
}

func (s *Service) SyncDataOfUser() error {
	if responseData, err := s.makeGetUsersRequest(); err != nil {
		return err
	} else {
		users := make([]model.User, 0, len(responseData))
		userIDs := make([]uuid.UUID, 0, len(responseData))
		countOfUsersByIDNumber := make(map[string]int)
		for _, userData := range responseData {
			// VS -> Visiting Staff：主治醫師
			// CR -> Chief Resident：總住院醫師
			// HN -> Head Nurse：護理長
			// RN -> Registered Nurse：護理師
			// NP -> Nurse Practitioner：專科護理師
			userMap := userData.(map[string]interface{})
			username := userMap["username"].(string)
			if username == s.username {
				continue
			}
			user := model.User{
				Username: username,
				FullName: userMap["full_name"].(string),
				IDNumber: userMap["id_number"].(string),
				Email:    userMap["email"].(string),
				Role:     userMap["role"].(string),
			}
			if !model.IsIDNumberValid(user.IDNumber) {
				user.IDNumber = ""
			}
			if user.IDNumber == "" {
				continue
			}
			countOfUsersByIDNumber[user.IDNumber] += 1
			users = append(users, user)
			userIDs = append(userIDs, uuid.MustParse(userMap["id"].(string)))
		}
		filteredUsers := make([]model.User, 0, len(responseData))
		filteredUserIDs := make([]uuid.UUID, 0, len(responseData))
		for i, user := range users {
			if countOfUsersByIDNumber[user.IDNumber] == 1 {
				filteredUsers = append(filteredUsers, user)
				filteredUserIDs = append(filteredUserIDs, userIDs[i])
			}
		}
		if err := s.store.User.SyncData(filteredUsers, false); err != nil {
			return err
		}
		for i, user := range filteredUsers {
			// TODO: 暫時先給admin
			if _, err := s.store.Enforcer.AddRolesForUser(user.ID.String(), []string{"admin"}); err != nil {
				return err
			}
			s.idMapperOfUsers[filteredUserIDs[i]] = user.ID
		}
	}
	return nil
}
