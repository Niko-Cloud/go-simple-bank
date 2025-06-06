package test

import (
	"context"
	"github.com/stretchr/testify/require"
	db "github.com/yuki/simplebank/db/sqlc/gen"
	util2 "github.com/yuki/simplebank/util"
	"testing"
	"time"
)

func createRandomUser(t *testing.T) db.User {
	hashedPassword, err := util2.HashPassword(util2.RandomString(6))
	require.NoError(t, err)

	arg := db.CreateUserParams{
		Username:       util2.RandomOwner(),
		HashedPassword: hashedPassword,
		FullName:       util2.RandomOwner(),
		Email:          util2.RandomEmail(),
	}

	user, err := testQueries.CreateUser(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, user)

	require.Equal(t, arg.Username, user.Username)
	require.Equal(t, arg.Email, arg.Email)
	require.Equal(t, arg.FullName, user.FullName)
	require.Equal(t, arg.HashedPassword, user.HashedPassword)

	require.True(t, user.PasswordChangedAt.IsZero())
	require.NotZero(t, user.CreatedAt)

	return user
}

func TestCreateUser(t *testing.T) {
	createRandomUser(t)
}

func TestGetUser(t *testing.T) {
	user1 := createRandomUser(t)
	user2, err := testQueries.GetUser(context.Background(), user1.Username)

	require.NoError(t, err)
	require.NotEmpty(t, user2)

	require.Equal(t, user1.FullName, user2.FullName)
	require.Equal(t, user1.Username, user2.Username)
	require.Equal(t, user1.HashedPassword, user2.HashedPassword)
	require.Equal(t, user1.Email, user2.Email)

	require.WithinDuration(t, user1.CreatedAt, user2.CreatedAt, time.Second)
	require.WithinDuration(t, user1.PasswordChangedAt, user2.PasswordChangedAt, time.Second)
}
