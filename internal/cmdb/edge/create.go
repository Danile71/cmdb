// Copyright 2022 Listware

package edge

import (
	"context"
	"encoding/json"

	"git.fg-tech.ru/listware/cmdb/internal/arangodb/edge"
	"git.fg-tech.ru/listware/proto/sdk/pbcmdb"
)

func (s *Server) Create(ctx context.Context, request *pbcmdb.Request) (response *pbcmdb.Response, err error) {
	response = &pbcmdb.Response{}
	meta, resp, err := edge.Create(ctx, s.client, request.GetCollection(), request.GetPayload())
	if err != nil {
		return
	}
	response.Meta = &pbcmdb.Meta{
		Key: meta.Key,
		Id:  meta.ID.String(),
		Rev: meta.Rev,
	}
	response.Payload, err = json.Marshal(resp)
	return
}
