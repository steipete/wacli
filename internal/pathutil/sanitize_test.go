package pathutil

import "testing"

func TestSanitizeSegment(t *testing.T) {
	if got := SanitizeSegment(""); got != "unknown" {
		t.Fatalf("expected unknown, got %q", got)
	}
	if got := SanitizeSegment(" ../a/b:c@d "); got == "" || got == " ../a/b:c@d " {
		t.Fatalf("unexpected sanitize result: %q", got)
	}
	if got := SanitizeSegment("a/b"); got != "a_b" {
		t.Fatalf("expected a_b, got %q", got)
	}
}

func TestSanitizeFilename(t *testing.T) {
	if got := SanitizeFilename(""); got != "file" {
		t.Fatalf("expected file, got %q", got)
	}
	if got := SanitizeFilename(".."); got == ".." {
		t.Fatalf("expected .. to be sanitized, got %q", got)
	}
	if got := SanitizeFilename("a/b"); got != "a_b" {
		t.Fatalf("expected a_b, got %q", got)
	}
}
