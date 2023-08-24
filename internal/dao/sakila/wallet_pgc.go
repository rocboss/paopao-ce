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
	_ core.WalletService = (*pgcWalletSrv)(nil)
)

type pgcWalletSrv struct {
	*walletSrv
	p *pgc.Wallet
}

func (s *pgcWalletSrv) CreateRecharge(userId, amount int64) (res *ms.WalletRecharge, err error) {
	err = stmtGet(s.q.CreateRecharge, res, userId, amount, time.Now().Unix())
	return
}
