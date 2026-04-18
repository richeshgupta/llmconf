package cli

import (
	"os"
	"strings"
	"testing"
)

func TestExecutablePath(t *testing.T) {
	path := executablePath()

	if path == "" {
		t.Fatal("executablePath() returned empty string")
	}

	if !strings.HasPrefix(path, "/") {
		t.Fatalf("executablePath() returned non-absolute path: %s", path)
	}

	info, err := os.Stat(path)
	if err != nil {
		t.Fatalf("executablePath() returned path that does not exist: %s", path)
	}
	if info.Mode()&0111 == 0 {
		t.Fatalf("executablePath() returned non-executable file: %s", path)
	}
}
