// +build linux freebsd

package system

import (
	"os"
	"syscall"
	"testing"

	"github.com/stretchr/testify/require"
)

// TestFromStatT tests fromStatT for a tempfile
func TestFromStatT(t *testing.T) {
	file, _, _, dir := prepareFiles(t)
	defer os.RemoveAll(dir)

	stat := &syscall.Stat_t{}
	err := syscall.Lstat(file, stat)
	require.NoError(t, err)

	s, err := fromStatT(stat)
	require.NoError(t, err)

	if stat.Mode != s.Mode() {
		t.Fatal("got invalid mode")
	}
	if stat.Uid != s.UID() {
		t.Fatal("got invalid uid")
	}
	if stat.Gid != s.GID() {
		t.Fatal("got invalid gid")
	}
	if stat.Rdev != s.Rdev() {
		t.Fatal("got invalid rdev")
	}
	if stat.Mtim != s.Mtim() {
		t.Fatal("got invalid mtim")
	}
}
