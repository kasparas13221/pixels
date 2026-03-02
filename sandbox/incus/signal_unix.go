//go:build !windows

package incus

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/gorilla/websocket"
	"github.com/lxc/incus/v6/shared/api"
	"golang.org/x/term"
)

func winchControl(fd int) func(*websocket.Conn) {
	return func(conn *websocket.Conn) {
		ch := make(chan os.Signal, 1)
		signal.Notify(ch, syscall.SIGWINCH)
		defer signal.Stop(ch)

		for range ch {
			w, h, err := term.GetSize(fd)
			if err != nil {
				continue
			}
			msg := api.InstanceExecControl{
				Command: "window-resize",
				Args: map[string]string{
					"width":  fmt.Sprintf("%d", w),
					"height": fmt.Sprintf("%d", h),
				},
			}
			_ = conn.WriteJSON(msg)
		}
	}
}
