package config

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestAlias(t *testing.T) {
	t.Run("default value", func(t *testing.T) {
		Configure(true)

		database := MongoDatabase.Get()
		require.Equal(t, MongoDatabase.GetDefault(), database)
	})
	t.Run("default is nil", func(t *testing.T) {
		Configure(true)

		uri := StorageUri.Get()
		require.Equal(t, "", uri)
	})
	t.Run("main only", func(t *testing.T) {
		t.Setenv("DECKARD_MONGO_DATABASE", "takenet")

		Configure(true)

		database := MongoDatabase.Get()
		require.Equal(t, "takenet", database)
	})
	t.Run("alias only", func(t *testing.T) {
		t.Setenv("DECKARD_MONGODB_DATABASE", "blip")

		Configure(true)

		database := MongoDatabase.Get()
		require.Equal(t, "blip", database)
	})
	t.Run("main and alias", func(t *testing.T) {
		t.Setenv("DECKARD_MONGO_DATABASE", "takenet")
		t.Setenv("DECKARD_MONGODB_DATABASE", "blip")

		Configure(true)

		database := MongoDatabase.Get()
		require.Equal(t, "takenet", database)
	})
	t.Run("multiple aliases", func(t *testing.T) {
		t.Setenv("DECKARD_MONGO_URI", "blip")
		t.Setenv("DECKARD_MONGODB_URI", "stilingue")

		Configure(true)

		uri := StorageUri.Get()
		require.Equal(t, "blip", uri)
	})
}
