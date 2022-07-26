package db

import (
	"context"
	"testing"

	"github.com/Grama-Check/Address-Check-Api/util"
	"github.com/stretchr/testify/require"
)

func CreateRandomPerson(t *testing.T) Person {
	args := CreatePersonParams{
		Nic:     util.RandomID(),
		Name:    util.RandomName(),
		Address: util.RandomAddress(),
	}
	person, err := testQueries.CreatePerson(context.Background(), args)

	require.NoError(t, err)
	require.NotEmpty(t, person)

	require.Equal(t, args.Nic, person.Nic)
	require.Equal(t, args.Name, person.Name)
	require.Equal(t, args.Address, person.Address)

	return person

}

func TestCreatePerson(t *testing.T) {
	CreateRandomPerson(t)
}

func TestGetPerson(t *testing.T) {
	person := CreateRandomPerson(t)

	person2, err := testQueries.GetPerson(context.Background(), person.Nic)

	require.NoError(t, err)
	require.NotEmpty(t, person2)

	require.Equal(t, person.Nic, person2.Nic)
	require.Equal(t, person.Name, person2.Name)
	require.Equal(t, person.Address, person2.Address)

}
