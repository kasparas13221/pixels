//go:build windows

package incus

import "github.com/gorilla/websocket"

// winchControl returns nil on Windows where SIGWINCH does not exist.
func winchControl(_ int) func(*websocket.Conn) {
	return nil
}
