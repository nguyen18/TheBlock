package auth

import (
	"context"
	"errors"
	"testing"

	authpb "TheBlock/src/api/auth/authpb"
	mds "TheBlock/src/api/auth/internal/datastore/mocks"
	mockauthdatastore "TheBlock/src/api/auth/internal/datastore/mocks"
	util "TheBlock/src/api/auth/internal/util"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func Test_Login(t *testing.T) {
	// define test cases
	uuid := "testuuid"
	email := "test@example.com"
	password := "testpassword"
	hashedPassword, _ := util.HashPassword(password)
	token := "jwt"

	tests := []struct {
		name              string
		mockDatastoreFunc func(store *mds.MockAuthDatastore)
		expectedResp      *authpb.LoginResponse
		expectedErrMsg    string
	}{
		{
			name: "successful login",
			mockDatastoreFunc: func(store *mockauthdatastore.MockAuthDatastore) {
				store.EXPECT().GetUserPasswordByEmail(context.TODO(), email).Return(hashedPassword, nil)
				store.EXPECT().GetUserUuidByEmail(context.TODO(), email).Return(uuid, nil)
			},
			expectedResp: &authpb.LoginResponse{
				Success: true,
				Token:   token,
				Uuid:    uuid,
			},
			expectedErrMsg: "",
		},
		{
			name: "user doesn't exist",
			mockDatastoreFunc: func(store *mockauthdatastore.MockAuthDatastore) {
				store.EXPECT().GetUserPasswordByEmail(context.TODO(), email).Return(nil, errors.New("issue retrieving password"))
			},
			expectedResp: &authpb.LoginResponse{
				Success: false,
			},
			expectedErrMsg: "user doesn't exist",
		},
		{
			name: "password doesn't match",
			mockDatastoreFunc: func(store *mockauthdatastore.MockAuthDatastore) {
				store.EXPECT().GetUserPasswordByEmail(context.TODO(), email).Return(nil, errors.New("issue retrieving password"))
			},
			expectedResp: &authpb.LoginResponse{
				Success: false,
			},
			expectedErrMsg: "passwords don't match",
		},
	}

	store := mds.NewMockAuthDatastore(gomock.NewController(t))
	s := AuthServer{store: store}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// call the function with the test case input
			req := &authpb.LoginRequest{Email: email, Password: password}

			tc.mockDatastoreFunc(store)

			resp, err := s.Login(context.TODO(), req)

			if tc.expectedErrMsg != "" {
				assert.Error(t, err)
				assert.NotNil(t, resp)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, resp)
				assert.Equal(t, tc.expectedResp.Success, resp.Success)
				if tc.expectedResp.Token != "" {
					assert.NotEmpty(t, resp.Token)
				}
				if tc.expectedResp.Uuid != "" {
					assert.NotEmpty(t, resp.Uuid)
				}
			}
		})
	}
}
