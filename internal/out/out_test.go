package out

import (
	"bytes"
	"encoding/json"
	"errors"
	"strings"
	"testing"
)

func TestWriteJSONEnvelope(t *testing.T) {
	var b bytes.Buffer
	if err := WriteJSON(&b, map[string]any{"ok": true}); err != nil {
		t.Fatalf("WriteJSON: %v", err)
	}
	var got map[string]any
	if err := json.Unmarshal([]byte(strings.TrimSpace(b.String())), &got); err != nil {
		t.Fatalf("unmarshal: %v", err)
	}
	if got["success"] != true {
		t.Fatalf("expected success=true, got %v", got["success"])
	}
	if got["error"] != nil {
		t.Fatalf("expected error=nil, got %v", got["error"])
	}
}

func TestWriteErrorJSONAndText(t *testing.T) {
	var b bytes.Buffer
	_ = WriteError(&b, true, errors.New("boom"))
	if !strings.Contains(b.String(), "\"success\":false") || !strings.Contains(b.String(), "boom") {
		t.Fatalf("unexpected json error output: %q", b.String())
	}

	b.Reset()
	_ = WriteError(&b, false, errors.New("boom"))
	if strings.TrimSpace(b.String()) != "boom" {
		t.Fatalf("unexpected text error output: %q", b.String())
	}
}
