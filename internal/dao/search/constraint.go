//go:build constraint

package search

import "github.com/rocboss/paopao-ce/internal/core"

var (
	_ core.TweetSearchService = (*bridgeTweetSearchServant)(nil)

	_ core.TweetSearchService = (*meiliTweetSearchServant)(nil)
	_ core.VersionInfo        = (*meiliTweetSearchServant)(nil)

	_ core.TweetSearchService = (*zincTweetSearchServant)(nil)
	_ core.VersionInfo        = (*zincTweetSearchServant)(nil)
)
