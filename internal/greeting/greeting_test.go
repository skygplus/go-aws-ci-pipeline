package greeting

import "testing"

func TestMessageWithName(t *testing.T) {
	got := Message("AWS")
	want := "Hello, AWS! Your Go CI pipeline is working."

	if got != want {
		t.Errorf("Message() = %q; want %q", got, want)
	}
}

func TestMessageWithEmptyName(t *testing.T) {
	got := Message("")
	want := "Hello, DevOps Engineer! Your Go CI pipeline is working."

	if got != want {
		t.Errorf("Message() = %q; want %q", got, want)
	}
}
