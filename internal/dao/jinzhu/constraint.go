//go:build constraint

package jinzhu

import (
	"github.com/rocboss/paopao-ce/internal/core"
)

var (
	_ core.AuthorizationManageService = (*authorizationManageSrv)(nil)
	_ core.CommentService             = (*commentSrv)(nil)
	_ core.CommentManageService       = (*commentManageSrv)(nil)

	_ core.ContactManageService   = (*contactManageSrv)(nil)
	_ core.FollowingManageService = (*followingManageSrv)(nil)

	_ core.MessageService  = (*messageSrv)(nil)
	_ core.SecurityService = (*securitySrv)(nil)

	_ core.TopicService  = (*topicSrv)(nil)
	_ core.TopicServantA = (*topicSrvA)(nil)

	_ core.IndexPostsService = (*shipIndexSrv)(nil)
	_ core.IndexPostsService = (*simpleIndexPostsSrv)(nil)

	_ core.TweetService       = (*tweetSrv)(nil)
	_ core.TweetManageService = (*tweetManageSrv)(nil)
	_ core.TweetHelpService   = (*tweetHelpSrv)(nil)

	_ core.TweetServantA       = (*tweetSrvA)(nil)
	_ core.TweetManageServantA = (*tweetManageSrvA)(nil)
	_ core.TweetHelpServantA   = (*tweetHelpSrvA)(nil)

	_ core.UserManageService = (*userManageSrv)(nil)
	_ core.WalletService     = (*walletSrv)(nil)

	_ core.DataService = (*dataSrv)(nil)
	_ core.VersionInfo = (*dataSrv)(nil)

	_ core.WebDataServantA = (*webDataSrvA)(nil)
	_ core.VersionInfo     = (*webDataSrvA)(nil)
)
