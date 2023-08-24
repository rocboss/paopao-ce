// Copyright 2023 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package sakila

import (
	"time"

	"github.com/rocboss/paopao-ce/internal/core"
	"github.com/rocboss/paopao-ce/internal/core/ms"
	"github.com/rocboss/paopao-ce/internal/dao/jinzhu/dbr"
	"github.com/rocboss/paopao-ce/internal/dao/sakila/auto/pgc"
)

var (
	_ core.MessageService = (*pgcMessageSrv)(nil)
)

type pgcMessageSrv struct {
	*messageSrv
	p *pgc.Message
}

func (s *pgcMessageSrv) CreateMessage(r *ms.Message) (*ms.Message, error) {
	r.Model = &dbr.Model{
		CreatedOn: time.Now().Unix(),
	}
	if err := s.p.CreateMessage.Get(&r.ID, r); err != nil {
		return nil, err
	}
	return r, nil
}
