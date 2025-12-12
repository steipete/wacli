package wa

import "testing"

func TestMediaTypeFromString(t *testing.T) {
	for _, tc := range []string{"image", "video", "audio", "document"} {
		if _, err := MediaTypeFromString(tc); err != nil {
			t.Fatalf("expected %s to be supported: %v", tc, err)
		}
	}
	if _, err := MediaTypeFromString("nope"); err == nil {
		t.Fatalf("expected error for unsupported type")
	}
}
