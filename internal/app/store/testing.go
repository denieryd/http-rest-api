package store

import (
    "fmt"
    "testing"
)

func TestStore(t *testing.T, databaseURL string) (*Store, func(string)) {
    t.Helper()
    config := NewConfig()
    config.DatabaseURL = databaseURL
    s := New(config)

    if err := s.Open(); err != nil {
        t.Fatal(err)
    }

    return s, func(tables string) {
        if _, err := s.db.Exec(fmt.Sprintf("TRUNCATE %s CASCADE", tables)); err != nil {

            t.Fatal(err)
        }

        s.Close()
    }
}
