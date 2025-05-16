package test

import (
	"context"
	"database/sql"
	"github.com/stretchr/testify/require"
	"github.com/yuki/simplebank/db/sqlc/gen"
	"github.com/yuki/simplebank/db/util"
	"testing"
	"time"

	_ "github.com/stretchr/testify/require"
)

func createRandomAccount(t *testing.T) db.Account {
	user := createRandomUser(t)

	arg := db.CreateAccountParams{
		Owner:    user.Username,
		Balance:  util.RandomBalance(),
		Currency: util.RandomCurrency(),
	}

	account, err := testQueries.CreateAccount(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, account)

	require.Equal(t, arg.Owner, account.Owner)
	require.Equal(t, arg.Balance, arg.Balance)
	require.Equal(t, arg.Currency, account.Currency)

	require.NotZero(t, account.ID)
	require.NotZero(t, account.CreatedAt)

	return account
}

func TestCreateAccount(t *testing.T) {
	createRandomAccount(t)
}

func TestGetAccount(t *testing.T) {
	acccount1 := createRandomAccount(t)
	account2, err := testQueries.GetAccount(context.Background(), acccount1.ID)

	require.NoError(t, err)
	require.NotEmpty(t, account2)

	require.Equal(t, acccount1.ID, account2.ID)
	require.Equal(t, acccount1.Owner, account2.Owner)
	require.Equal(t, acccount1.Balance, account2.Balance)
	require.Equal(t, acccount1.Currency, account2.Currency)

	require.WithinDuration(t, acccount1.CreatedAt, account2.CreatedAt, time.Second)
}

func TestUpdateAccount(t *testing.T) {
	account1 := createRandomAccount(t)

	arg := db.UpdateAccountParams{
		ID:      account1.ID,
		Balance: util.RandomBalance(),
	}

	account2, err := testQueries.UpdateAccount(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, account2)
	require.Equal(t, account1.ID, account2.ID)
	require.Equal(t, account1.Owner, account2.Owner)
	require.Equal(t, arg.Balance, account2.Balance)
	require.Equal(t, account1.Currency, account2.Currency)
	require.WithinDuration(t, account1.CreatedAt, account2.CreatedAt, time.Second)
}

func TestDeletAccount(t *testing.T) {
	account1 := createRandomAccount(t)

	deletedAccount, err := testQueries.DeleteAccount(context.Background(), account1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, deletedAccount)

	require.Equal(t, account1.ID, deletedAccount.ID)
	require.Equal(t, account1.Owner, deletedAccount.Owner)
	require.Equal(t, account1.Balance, deletedAccount.Balance)
	require.Equal(t, account1.Currency, deletedAccount.Currency)

	account2, err := testQueries.GetAccount(context.Background(), account1.ID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, account2)
}

func TestListAccounts(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomAccount(t)
	}

	arg := db.ListAccountsParams{
		Limit:  5,
		Offset: 5,
	}

	accounts, err := testQueries.ListAccounts(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, accounts, 5)
	for _, account := range accounts {
		require.NotEmpty(t, account)
		require.NotZero(t, account.ID)
		require.NotZero(t, account.Owner)
		require.NotZero(t, account.Balance)
		require.NotZero(t, account.Currency)
		require.NotZero(t, account.CreatedAt)
	}

}
