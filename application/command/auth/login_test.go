package auth

import (
	"context"
	"ddd-sample/internal/auth/aggregate"
	"ddd-sample/internal/auth/repository"
	"ddd-sample/pkg/env"
	"ddd-sample/pkg/localtime"
	mockauthrepository "ddd-sample/test/mocks/mockinternal/auth/repository"
	mockenv "ddd-sample/test/mocks/pkg/env"
	mocklocaltime "ddd-sample/test/mocks/pkg/localtime"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestNewLoginCommand(t *testing.T) {
	type args struct {
		identityRepository repository.IdentityRepository
		env                env.Env
		localtime          localtime.LocalTime
	}

	t.Run("工廠函式", func(t *testing.T) {
		arg := args{
			identityRepository: mockauthrepository.NewIdentityRepository(t),
			env:                mockenv.NewEnv(t),
			localtime:          mocklocaltime.NewLocalTime(t),
		}

		var expected LoginCommand = &loginCommand{
			identityRepository: arg.identityRepository,
			env:                arg.env,
			localtime:          arg.localtime,
		}

		actaul := NewLoginCommand(arg.identityRepository, arg.env, arg.localtime)

		assert.Equal(t, expected, actaul)
	})
}

func Test_loginCommand_Execute(t *testing.T) {
	type args struct {
		ctx   context.Context
		input LoginCommandInput
	}

	t.Run("登入成功", func(t *testing.T) {
		arg := args{
			ctx: context.TODO(),
			input: LoginCommandInput{
				Username: "test",
				Password: "1234",
			},
		}

		// mocks
		mockAuthRepository := mockauthrepository.NewIdentityRepository(t)
		mockAuthRepository.On("Find", mock.Anything).
			Return(aggregate.NewIdenetity(arg.input.Username, arg.input.Password), nil)
		mockAuthRepository.On("SaveLoginRecord", mock.Anything, mock.Anything).
			Return(nil)

		mockEnv := mockenv.NewEnv(t)
		mockEnv.On("GetValue", mock.Anything).Return("1234")

		mockLocaltime := mocklocaltime.NewLocalTime(t)
		mockLocaltime.On("NowTime").Return(time.Unix(0, 0))

		loginCommand := NewLoginCommand(mockAuthRepository, mockEnv, mockLocaltime)
		actual, actualErr := loginCommand.Execute(arg.ctx, arg.input)

		assert.True(t, actual.IsLogin)
		assert.NoError(t, actualErr)
	})

	t.Run("登入失敗 - 取會員資料失敗", func(t *testing.T) {
		arg := args{
			ctx: context.TODO(),
			input: LoginCommandInput{
				Username: "test",
				Password: "1234",
			},
		}

		// mocks
		mockAuthRepository := mockauthrepository.NewIdentityRepository(t)
		mockAuthRepository.On("Find", mock.Anything).
			Return(nil, assert.AnError)

		mockEnv := mockenv.NewEnv(t)

		mockLocaltime := mocklocaltime.NewLocalTime(t)

		expected := LoginCommandOutput{}

		loginCommand := NewLoginCommand(mockAuthRepository, mockEnv, mockLocaltime)
		actual, actualErr := loginCommand.Execute(arg.ctx, arg.input)

		assert.Equal(t, expected, actual)
		assert.Error(t, actualErr)
	})

	t.Run("登入失敗 - 帳號密碼錯誤", func(t *testing.T) {
		arg := args{
			ctx: context.TODO(),
			input: LoginCommandInput{
				Username: "test",
				Password: "123456",
			},
		}

		// mocks
		mockAuthRepository := mockauthrepository.NewIdentityRepository(t)
		mockAuthRepository.On("Find", mock.Anything).
			Return(aggregate.NewIdenetity("bar", "foo"), nil)
		mockAuthRepository.On("SaveLoginFailedRecord", mock.Anything).
			Return(nil)

		mockEnv := mockenv.NewEnv(t)

		mockLocaltime := mocklocaltime.NewLocalTime(t)

		expected := LoginCommandOutput{}

		loginCommand := NewLoginCommand(mockAuthRepository, mockEnv, mockLocaltime)
		actual, actualErr := loginCommand.Execute(arg.ctx, arg.input)

		assert.Equal(t, expected, actual)
		assert.NoError(t, actualErr)
	})

	t.Run("登入失敗 - 保存登入紀錄失敗", func(t *testing.T) {
		arg := args{
			ctx: context.TODO(),
			input: LoginCommandInput{
				Username: "test",
				Password: "1234",
			},
		}

		// mocks
		mockAuthRepository := mockauthrepository.NewIdentityRepository(t)
		mockAuthRepository.On("Find", mock.Anything).
			Return(aggregate.NewIdenetity(arg.input.Username, arg.input.Password), nil)
		mockAuthRepository.On("SaveLoginRecord", mock.Anything, mock.Anything).
			Return(assert.AnError)

		mockEnv := mockenv.NewEnv(t)
		mockEnv.On("GetValue", mock.Anything).Return("1234")

		mockLocaltime := mocklocaltime.NewLocalTime(t)
		mockLocaltime.On("NowTime").Return(time.Unix(0, 0))

		expected := LoginCommandOutput{}

		loginCommand := NewLoginCommand(mockAuthRepository, mockEnv, mockLocaltime)
		actual, actualErr := loginCommand.Execute(arg.ctx, arg.input)

		assert.Equal(t, expected, actual)
		assert.Error(t, actualErr)
	})
}
