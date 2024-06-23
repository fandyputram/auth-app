package repository

import (
	c "auth-app/constants"
	d "auth-app/internal/domain"
	"context"
	"reflect"
	"testing"
)

func Test_userImpl_SaveUser(t *testing.T) {
	dummyCtx := context.TODO()
	u := NewUserRepository().(*userImpl)
	u.SaveUser(dummyCtx, &d.User{
		Username: "TEST1",
		Password: "password",
		Email:    "email@email.com",
	})

	type args struct {
		ctx  context.Context
		user *d.User
	}
	tests := []struct {
		name    string
		u       *userImpl
		args    args
		wantErr bool
	}{
		{
			name:    c.TestFailed + " username exist",
			u:       u,
			wantErr: true,
			args: args{
				ctx: context.TODO(),
				user: &d.User{
					Username: "TEST1",
					Password: "password",
					Email:    "email@email.com",
				},
			},
		},
		{
			name:    c.TestSuccess,
			u:       u,
			wantErr: false,
			args: args{
				ctx: context.TODO(),
				user: &d.User{
					Username: "TEST2",
					Password: "password",
					Email:    "email2@email.com",
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.u.SaveUser(tt.args.ctx, tt.args.user); (err != nil) != tt.wantErr {
				t.Errorf("userImpl.SaveUser() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_userImpl_GetUserByUsernamePassword(t *testing.T) {
	dummyCtx := context.TODO()
	u := NewUserRepository().(*userImpl)
	user := &d.User{
		Username: "TEST1",
		Password: "password",
		Email:    "email@email.com",
	}
	u.SaveUser(dummyCtx, user)

	type args struct {
		ctx      context.Context
		username string
		password string
	}
	tests := []struct {
		name    string
		u       *userImpl
		args    args
		want    *d.User
		wantErr bool
	}{
		{
			name:    c.TestFailed + " invalid username",
			u:       u,
			wantErr: true,
			args: args{
				ctx:      context.TODO(),
				username: "TEST2",
				password: "password",
			},
		},
		{
			name:    c.TestFailed + " invalid password",
			u:       u,
			wantErr: true,
			args: args{
				ctx:      context.TODO(),
				username: "TEST1",
				password: "password1",
			},
		},
		{
			name:    c.TestSuccess,
			u:       u,
			wantErr: false,
			args: args{
				ctx:      context.TODO(),
				username: "TEST1",
				password: "password",
			},
			want: user,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.u.GetUserByUsernamePassword(tt.args.ctx, tt.args.username, tt.args.password)
			if (err != nil) != tt.wantErr {
				t.Errorf("userImpl.GetUserByUsernamePassword() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("userImpl.GetUserByUsernamePassword() = %v, want %v", got, tt.want)
			}
		})
	}
}
