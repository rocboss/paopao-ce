// Copyright 2023 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package joint

import (
	stdJson "encoding/json"

	"github.com/rocboss/paopao-ce/pkg/json"
)

func RespMarshal(data any) (stdJson.RawMessage, error) {
	return json.Marshal(data)
}
