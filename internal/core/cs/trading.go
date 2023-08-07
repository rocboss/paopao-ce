// Copyright 2023 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package cs

type AttachmentBill struct {
	ID         int64 `json:"id"`
	PostID     int64 `json:"post_id"`
	UserID     int64 `json:"user_id"`
	PaidAmount int64 `json:"paid_amount"`
}
