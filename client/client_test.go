package client

import "testing"

func TestClient(t *testing.T) {

	assertCorrectMessage := func(t *testing.T, got ChatClient, want ChatClient) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	t.Run("Create new Chat client", func(t *testing.T) {
		got := NewChatClient("4711")
		want := ChatClient{ServerAddress: "4711"}
		assertCorrectMessage(t, got, want)
	})

}
