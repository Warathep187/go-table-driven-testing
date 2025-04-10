package models

type BadUsername struct {
	ID       string
	Username string
}

//go:generate mockery --name BadUsernameModelInterface --structname MockBadUsernameModel --filename=mock_bad_username_model.go --output ../tests/mocks/models
type BadUsernameModelInterface interface {
	GetBadUsernameByUsername(username string) (*BadUsername, bool)
}

type BadUsernameModel struct {
	badUsernames []BadUsername
}

func NewBadUsernameModel() *BadUsernameModel {
	return &BadUsernameModel{
		badUsernames: []BadUsername{
			{ID: "1", Username: "admin"},
			{ID: "2", Username: "root"},
			{ID: "3", Username: "test"},
			{ID: "4", Username: "user"},
			{ID: "5", Username: "guest"},
		},
	}
}

func (m *BadUsernameModel) GetBadUsernameByUsername(username string) (*BadUsername, bool) {
	for _, badUsername := range m.badUsernames {
		if badUsername.Username == username {
			return &badUsername, true
		}
	}
	return nil, false
}
