package usecase

import (
	c "auth-app/constants"
	d "auth-app/internal/domain"
	mocks "auth-app/mocks/repository"
	"context"
	"errors"
	"reflect"
	"testing"

	"github.com/golang/mock/gomock"
)

func Test_authImpl_Login(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	dummyCtx := context.TODO()

	dummyUser := d.User{
		ID:       "ID1",
		Username: "username",
		Email:    "email",
		Password: "password",
	}

	mockUserRepo := mocks.NewMockUserRepository(ctrl)
	a := NewAuthUseCase(&Options{
		UserRepository: mockUserRepo,
	}).(*authImpl)

	type args struct {
		ctx     context.Context
		payload *LoginPayload
	}
	tests := []struct {
		name    string
		a       *authImpl
		args    args
		want    *d.User
		wantErr bool
		fun     func()
	}{
		{
			name: c.TestFailed + " empty payload",
			a:    a,
			args: args{
				ctx: dummyCtx,
				payload: &LoginPayload{
					Username: "",
					Password: "",
				},
			},
			wantErr: true,
			fun:     nil,
		},
		{
			name: c.TestFailed + " error login",
			a:    a,
			args: args{
				ctx: dummyCtx,
				payload: &LoginPayload{
					Username: "username",
					Password: "password",
				},
			},
			wantErr: true,
			fun: func() {
				mockUserRepo.EXPECT().GetUserByUsernamePassword(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil, errors.New("dummy error")).Times(1)
			},
		},
		{
			name: c.TestSuccess,
			a:    a,
			args: args{
				ctx: dummyCtx,
				payload: &LoginPayload{
					Username: "username",
					Password: "password",
				},
			},
			wantErr: false,
			fun: func() {
				mockUserRepo.EXPECT().GetUserByUsernamePassword(gomock.Any(), gomock.Any(), gomock.Any()).Return(&dummyUser, nil).Times(1)
			},
			want: &dummyUser,
		},
	}
	for _, tt := range tests {
		if tt.fun != nil {
			tt.fun()
		}
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.a.Login(tt.args.ctx, tt.args.payload)
			if (err != nil) != tt.wantErr {
				t.Errorf("authImpl.Login() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("authImpl.Login() = %v, want %v", got, tt.want)
			}
		})
	}
}
