package usecase

import (
	c "auth-app/constants"
	mocks "auth-app/mocks/repository"
	"context"
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
)

func Test_authImpl_Register(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	dummyCtx := context.TODO()

	mockUserRepo := mocks.NewMockUserRepository(ctrl)
	a := NewAuthUseCase(&Options{
		UserRepository: mockUserRepo,
	}).(*authImpl)

	type args struct {
		ctx     context.Context
		payload *RegisterPayload
	}
	tests := []struct {
		name    string
		a       *authImpl
		args    args
		wantErr bool
		fun     func()
	}{
		{
			name: c.TestFailed + "empty payload",
			a:    a,
			args: args{
				ctx: dummyCtx,
				payload: &RegisterPayload{
					Username: "",
					Email:    "",
					Password: "",
				},
			},
			wantErr: true,
		},
		{
			name: c.TestFailed + "invalid email format",
			a:    a,
			args: args{
				ctx: dummyCtx,
				payload: &RegisterPayload{
					Username: "username",
					Email:    "email",
					Password: "password",
				},
			},
			wantErr: true,
			fun: func() {
				mockUserRepo.EXPECT().SaveUser(gomock.Any(), gomock.Any()).Return(errors.New("dummy error")).Times(1)
			},
		},
		{
			name: c.TestFailed + "error save user",
			a:    a,
			args: args{
				ctx: dummyCtx,
				payload: &RegisterPayload{
					Username: "username",
					Email:    "email@email.com",
					Password: "password",
				},
			},
			wantErr: true,
			fun: func() {
				mockUserRepo.EXPECT().SaveUser(gomock.Any(), gomock.Any()).Return(nil).Times(1)
			},
		},
		{
			name: c.TestSuccess,
			a:    a,
			args: args{
				ctx: dummyCtx,
				payload: &RegisterPayload{
					Username: "username",
					Email:    "email@email.com",
					Password: "password",
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		if tt.fun != nil {
			tt.fun()
		}
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.a.Register(tt.args.ctx, tt.args.payload); (err != nil) != tt.wantErr {
				t.Errorf("authImpl.Register() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
