package aggregate

import (
	"ddd-sample/internal/auth/entity"
	"ddd-sample/internal/auth/enum"
	"ddd-sample/internal/auth/valueobject"
	mockaggregate "ddd-sample/test/mocks/mockinternal/core/aggregate"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestIdentity_CheckIdentity(t *testing.T) {
	type args struct {
		username string
		password string
	}

	t.Run("檢查帳號密碼 - 成功", func(t *testing.T) {
		// 參數
		arg := args{
			username: "test",
			password: "1234",
		}

		// 期望值
		expected := true

		identity := Identity{
			account: entity.BuildAccount("1", "test", "03ac674216f3e15c761ee1a5e255f067953623c8b388b4459e13f978d7c846f4", "123", enum.AccountStatusNormal),
		}

		actaul := identity.CheckIdentity(arg.username, arg.password)

		assert.Equal(t, expected, actaul)
	})

	t.Run("檢查帳號密碼 - 失敗", func(t *testing.T) {
		// 參數
		arg := args{
			username: "test",
			password: "123456789",
		}

		// 期望值
		expected := false

		identity := Identity{
			account: entity.BuildAccount("1", "test", "03ac674216f3e15c761ee1a5e255f067953623c8b388b4459e13f978d7c846f4", "123", enum.AccountStatusNormal),
		}

		actaul := identity.CheckIdentity(arg.username, arg.password)

		assert.Equal(t, expected, actaul)
	})
}

func TestIdentity_CreateToken(t *testing.T) {
	type args struct {
		secretKey []byte
		nowTime   time.Time
	}

	t.Run("產生token", func(t *testing.T) {
		arg := args{
			secretKey: []byte("1234"),
			nowTime:   time.Unix(0, 0),
		}

		// mocks
		mocksCoreAggregate := mockaggregate.NewCoreAggregate(t)
		mocksCoreAggregate.On("AddEvent", mock.Anything)

		// 期望值
		expected := valueobject.Token{
			TokenString: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiJ0ZXN0IiwiZXhwIjo4NjQwMCwibmJmIjowLCJpYXQiOjAsImp0aSI6IjEifQ.dUBSXogn3RE-6-JvhjITXfhfRnkO0f42mDEAkRwyVfQ",
			CreateTime:  time.Unix(0, 0),
		}

		identity := Identity{
			account:       entity.BuildAccount("1", "test", "03ac674216f3e15c761ee1a5e255f067953623c8b388b4459e13f978d7c846f4", "123", enum.AccountStatusNormal),
			CoreAggregate: mocksCoreAggregate,
		}

		actual := identity.CreateToken(arg.secretKey, arg.nowTime)
		assert.Equal(t, expected, actual)
	})
}
