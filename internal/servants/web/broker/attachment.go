// Copyright 2022 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package broker

import (
	"github.com/rocboss/paopao-ce/internal/core"
)

func CreateAttachment(attachment *core.Attachment) (*core.Attachment, error) {
	return ds.CreateAttachment(attachment)
}
