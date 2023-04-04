// Copyright 2023 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package conf

import (
	"time"

	"github.com/alimy/cfg"
	"github.com/getsentry/sentry-go"
	"github.com/rocboss/paopao-ce/pkg/version"
)

func initSentry() {
	cfg.Be("Sentry", func() {
		opts := sentry.ClientOptions{
			Dsn:              sentrySetting.Dsn,
			Debug:            sentrySetting.Debug,
			AttachStacktrace: sentrySetting.AttachStacktrace,
			TracesSampleRate: sentrySetting.TracesSampleRate,
		}
		_ = sentry.Init(opts)
		if sentrySetting.AttachLogrus {
			setupSentryLogrus(opts)
		}
		sentry.WithScope(func(scope *sentry.Scope) {
			scope.SetExtras(map[string]any{
				"version": version.VersionInfo(),
				"time":    time.Now().Local(),
			})
			sentry.CaptureMessage("paopao-ce sentry works!")
		})
	})
}
