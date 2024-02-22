package runtime

import (
	"context"

	"github.com/heroiclabs/nakama-common/rtapi"
)

type (
	Peer interface {
		Send(node string, msg *rtapi.NakamaPeer_Frame, reliable bool) error
		Broadcast(msg *rtapi.NakamaPeer_Frame, reliable bool, includeSelf bool)
		Call(ctx context.Context, node string, msg *rtapi.NakamaPeer_Frame) (*rtapi.NakamaPeer_Frame, error)
		SessionToken(tk string) (userID, username string, vars map[string]string, exp int64, online bool, err error)
	}
)
