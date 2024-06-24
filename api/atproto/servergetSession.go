// Code generated by cmd/lexgen (see Makefile's lexgen); DO NOT EDIT.

package atproto

// schema: com.atproto.server.getSession

import (
	"context"

	"github.com/bluesky-social/indigo/xrpc"
)

// ServerGetSession_Output is the output of a com.atproto.server.getSession call.
type ServerGetSession_Output struct {
	Active          *bool        `json:"active,omitempty" cborgen:"active,omitempty"`
	Did             string       `json:"did" cborgen:"did"`
	DidDoc          *interface{} `json:"didDoc,omitempty" cborgen:"didDoc,omitempty"`
	Email           *string      `json:"email,omitempty" cborgen:"email,omitempty"`
	EmailAuthFactor *bool        `json:"emailAuthFactor,omitempty" cborgen:"emailAuthFactor,omitempty"`
	EmailConfirmed  *bool        `json:"emailConfirmed,omitempty" cborgen:"emailConfirmed,omitempty"`
	Handle          string       `json:"handle" cborgen:"handle"`
	// status: If active=false, this optional field indicates a possible reason for why the account is not active. If active=false and no status is supplied, then the host makes no claim for why the repository is no longer being hosted.
	Status *string `json:"status,omitempty" cborgen:"status,omitempty"`
}

// ServerGetSession calls the XRPC method "com.atproto.server.getSession".
func ServerGetSession(ctx context.Context, c *xrpc.Client) (*ServerGetSession_Output, error) {
	var out ServerGetSession_Output
	if err := c.Do(ctx, xrpc.Query, "", "com.atproto.server.getSession", nil, nil, &out); err != nil {
		return nil, err
	}

	return &out, nil
}
