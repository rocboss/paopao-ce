// Copyright 2023 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package sakila

import (
	"time"

	"github.com/rocboss/paopao-ce/internal/core"
	"github.com/rocboss/paopao-ce/internal/core/ms"
	"github.com/rocboss/paopao-ce/internal/dao/sakila/auto/pgc"
)

var (
	_ core.UserManageService = (*pgcUserManageSrv)(nil)
)

type pgcUserManageSrv struct {
	*userManageSrv
	p *pgc.UserManage
}

func (s *pgcUserManageSrv) CreateUser(r *ms.User) (*ms.User, error) {
	r.Model = &ms.Model{CreatedOn: time.Now().Unix()}
	if err := s.p.CreateUser.Get(&r.ID, r); err != nil {
		return nil, err
	}
	return r, nil
}
