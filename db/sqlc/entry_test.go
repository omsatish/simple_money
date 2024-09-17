package db

import (
	"context"
	"testing"
	"time"

	"github.com/omsatish/simplebank/db/util"
	"github.com/stretchr/testify/require"
)

func creatRandomeEntry(t *testing.T) Entry {
	account := createRandomAccount(t)

	args := CreateEntryParams{
		AccountID: account.ID,
		Amount:    util.RamdomMoney(),
	}

	entry, err := testQueries.CreateEntry(context.Background(), args)

	require.NoError(t, err)
	require.NotEmpty(t, entry)

	require.Equal(t, args.AccountID, entry.AccountID)
	require.Equal(t, args.Amount, entry.Amount)
	return entry
}

func TestCreateEntry(t *testing.T) {
	creatRandomeEntry(t)
}

func TestGetEntry(t *testing.T) {
	createdEntry := creatRandomeEntry(t)

	getEntry, err := testQueries.GetEntry(context.Background(), createdEntry.ID)

	require.NoError(t, err)
	require.NotEmpty(t, getEntry)

	require.Equal(t, createdEntry.ID, getEntry.ID)
	require.Equal(t, createdEntry.AccountID, getEntry.AccountID)
	require.Equal(t, createdEntry.Amount, getEntry.Amount)
	require.WithinDuration(t, createdEntry.CreatedAt, getEntry.CreatedAt, time.Second)
}

func TestListEntry(t *testing.T) {
	for i := 0; i < 2; i++ {
		creatRandomeEntry(t)
	}

	args := ListEntriesParams{
		Limit:  2,
		Offset: 2,
	}

	entries, err := testQueries.ListEntries(context.Background(), args)

	require.NoError(t, err)
	require.NotEmpty(t, entries)
	require.Len(t, entries, 2)

	for _, entry := range entries {
		require.NotEmpty(t, entry)
	}

}
