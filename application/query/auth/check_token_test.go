package auth

import (
	"context"
	mockenv "ddd-sample/mocks/pkg/env"
	mocklocaltime "ddd-sample/mocks/pkg/localtime"
	"ddd-sample/pkg/env"
	"ddd-sample/pkg/localtime"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestNewCheckTokenQuery(t *testing.T) {
	type args struct {
		env       env.Env
		localTime localtime.LocalTime
	}

	t.Run("工廠函式", func(t *testing.T) {
		arg := args{
			env:       mockenv.NewEnv(t),
			localTime: mocklocaltime.NewLocalTime(t),
		}

		var expected CheckTokenQuery = &checkTokenQuery{
			env:       arg.env,
			localTime: arg.localTime,
		}

		actual := NewCheckTokenQuery(arg.env, arg.localTime)

		assert.Equal(t, expected, actual)
	})
}

func Test_checkTokenQuery_Execute(t *testing.T) {
	type args struct {
		ctx   context.Context
		input CheckTokenQueryInput
	}

	t.Run("檢查auth token成功", func(t *testing.T) {
		arg := args{
			ctx: context.TODO(),
			input: CheckTokenQueryInput{
				AuthToken: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiJ0ZXN0IiwiZXhwIjo4NjQwMCwibmJmIjowLCJpYXQiOjAsImp0aSI6IjEifQ.dUBSXogn3RE-6-JvhjITXfhfRnkO0f42mDEAkRwyVfQ",
			},
		}

		mockEnv := mockenv.NewEnv(t)
		mockEnv.On("GetValue", mock.Anything).Return("1234")

		mockLocalTime := mocklocaltime.NewLocalTime(t)
		mockLocalTime.On("NowTime").Return(time.Unix(0, 0))

		expected := CheckTokenQueryOutput{
			UID:      "1",
			Username: "test",
		}

		checkTokenQuery := NewCheckTokenQuery(mockEnv, mockLocalTime)

		actual, actualErr := checkTokenQuery.Execute(arg.ctx, arg.input)

		assert.Equal(t, expected, actual)
		assert.NoError(t, actualErr)
	})

	t.Run("檢查auth token輸入錯誤", func(t *testing.T) {
		arg := args{
			ctx: context.TODO(),
			input: CheckTokenQueryInput{
				AuthToken: "1234",
			},
		}

		mockEnv := mockenv.NewEnv(t)
		mockEnv.On("GetValue", mock.Anything).Return("1234")

		mockLocalTime := mocklocaltime.NewLocalTime(t)

		expected := CheckTokenQueryOutput{}

		checkTokenQuery := NewCheckTokenQuery(mockEnv, mockLocalTime)

		actual, actualErr := checkTokenQuery.Execute(arg.ctx, arg.input)

		assert.Equal(t, expected, actual)
		assert.Error(t, actualErr)
	})

	t.Run("檢查auth token未生效", func(t *testing.T) {
		arg := args{
			ctx: context.TODO(),
			input: CheckTokenQueryInput{
				AuthToken: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiJ0ZXN0IiwiZXhwIjo4NzQwMCwibmJmIjoxMDAwLCJpYXQiOjEwMDAsImp0aSI6IjEifQ.AuXDcczrsZP7_hVZc8HhFYSjQS_YiYcizH3XSCnZEjg",
			},
		}

		mockEnv := mockenv.NewEnv(t)
		mockEnv.On("GetValue", mock.Anything).Return("1234")

		mockLocalTime := mocklocaltime.NewLocalTime(t)
		mockLocalTime.On("NowTime").Return(time.Unix(0, 0))

		expected := CheckTokenQueryOutput{}

		checkTokenQuery := NewCheckTokenQuery(mockEnv, mockLocalTime)

		actual, actualErr := checkTokenQuery.Execute(arg.ctx, arg.input)

		assert.Equal(t, expected, actual)
		assert.Error(t, actualErr)
	})
}
