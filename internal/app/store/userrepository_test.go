package store_test

import (
    "github.com/denieryd/http-rest-api/internal/app/model"
    "github.com/denieryd/http-rest-api/internal/app/store"
    "github.com/stretchr/testify/assert"
    "testing"
)

func TestUserRepository_Create(t *testing.T) {
    s, teardown := store.TestStore(t, databaseURL)

    defer teardown("users")
    u, err := s.User().Create(&model.User{Email: "user@example.org"})

    assert.NoError(t, err)
    assert.NotNil(t, u)
}
