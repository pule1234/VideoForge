package gapi

import (
	"context"
	"fmt"
	"github.com/pule1234/VideoForge/token"
	"google.golang.org/grpc/metadata"
	"strings"
)

const (
	authorizationHeader = "Authorization"
	authorizationBearer = "Bearer "
)

func (server *Server) authorizeUser(ctx context.Context) (*token.Payload, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, fmt.Errorf("missing metadata")
	}

	values := md.Get(authorizationHeader)
	if len(values) == 0 {
		return nil, fmt.Errorf("missing authorization header")
	}

	authHeader := values[0]
	fields := strings.Fields(authHeader)
	if len(fields) < 2 {
		return nil, fmt.Errorf("invalid authorization header format")
	}

	authType := fields[0]
	if authType != authorizationBearer {
		return nil, fmt.Errorf("invalid authorization type")
	}

	accessToken := fields[1]
	accessPayload, err := server.tokenMaker.VerifyToken(accessToken)
	if err != nil {
		return nil, fmt.Errorf("access token verification failed")
	}

	return accessPayload, nil
}
