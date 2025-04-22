package redkacli

import "testing"

func TestClient(t *testing.T) {
	t.Run("New", func(t *testing.T) {
		options := NewOptions()
		options.path = "./test.db"
		_, err := New(options)
		if err != nil {
			t.Fatal(err)
		}
	})
}
