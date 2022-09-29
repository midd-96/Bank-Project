package db

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

// dckr_pat_2gPBWPywfWHXPlwWVU__s_IWPic

func TestTransferTxDeadLock(t *testing.T) {
	store := NewStore(testDB)

	account1 := CreateRandomAccount(t)
	account2 := CreateRandomAccount(t)

	fmt.Println(" >> before: ", account1.Balance, account2.Balance)

	errs := make(chan error)

	n := 10
	amount := int64(10)

	//run n concurrent transfer transactions
	for i := 0; i < n; i++ {
		fromAccountID := account1.ID
		toAccountID := account2.ID

		if i%2 == 1 {
			fromAccountID = account2.ID
			toAccountID = account1.ID
		}

		go func() {
			ctx := context.Background()
			_, err := store.TransferTx(ctx, TransferTxParams{
				FromAccountID: fromAccountID,
				ToAccountID:   toAccountID,
				Amount:        amount,
			})
			errs <- err

		}()
	}
	//check results

	for i := 0; i < n; i++ {
		err := <-errs
		require.NoError(t, err)

	}
	// check the final updated balance
	updatedAccount1, err := store.GetAccount(context.Background(), account1.ID)
	require.NoError(t, err)

	updatedAccount2, err := store.GetAccount(context.Background(), account2.ID)
	require.NoError(t, err)

	fmt.Println(">> after:", updatedAccount1.Balance, updatedAccount2.Balance)
	require.Equal(t, account1.Balance, updatedAccount1.Balance)
	require.Equal(t, account2.Balance, updatedAccount2.Balance)

}
