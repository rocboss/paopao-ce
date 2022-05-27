package service

import (
	"math/rand"
	"time"
)

var defaultAvatars = []string{
	"https://paopao-assets.oss-cn-shanghai.aliyuncs.com/public/avatar/default/zoe.png",
	"https://paopao-assets.oss-cn-shanghai.aliyuncs.com/public/avatar/default/william.png",
	"https://paopao-assets.oss-cn-shanghai.aliyuncs.com/public/avatar/default/walter.png",
	"https://paopao-assets.oss-cn-shanghai.aliyuncs.com/public/avatar/default/thomas.png",
	"https://paopao-assets.oss-cn-shanghai.aliyuncs.com/public/avatar/default/taylor.png",
	"https://paopao-assets.oss-cn-shanghai.aliyuncs.com/public/avatar/default/sophia.png",
	"https://paopao-assets.oss-cn-shanghai.aliyuncs.com/public/avatar/default/sam.png",
	"https://paopao-assets.oss-cn-shanghai.aliyuncs.com/public/avatar/default/ryan.png",
	"https://paopao-assets.oss-cn-shanghai.aliyuncs.com/public/avatar/default/ruby.png",
	"https://paopao-assets.oss-cn-shanghai.aliyuncs.com/public/avatar/default/quinn.png",
	"https://paopao-assets.oss-cn-shanghai.aliyuncs.com/public/avatar/default/paul.png",
	"https://paopao-assets.oss-cn-shanghai.aliyuncs.com/public/avatar/default/owen.png",
	"https://paopao-assets.oss-cn-shanghai.aliyuncs.com/public/avatar/default/olivia.png",
	"https://paopao-assets.oss-cn-shanghai.aliyuncs.com/public/avatar/default/norman.png",
	"https://paopao-assets.oss-cn-shanghai.aliyuncs.com/public/avatar/default/nora.png",
	"https://paopao-assets.oss-cn-shanghai.aliyuncs.com/public/avatar/default/natalie.png",
	"https://paopao-assets.oss-cn-shanghai.aliyuncs.com/public/avatar/default/naomi.png",
	"https://paopao-assets.oss-cn-shanghai.aliyuncs.com/public/avatar/default/miley.png",
	"https://paopao-assets.oss-cn-shanghai.aliyuncs.com/public/avatar/default/mike.png",
	"https://paopao-assets.oss-cn-shanghai.aliyuncs.com/public/avatar/default/lucas.png",
	"https://paopao-assets.oss-cn-shanghai.aliyuncs.com/public/avatar/default/kylie.png",
	"https://paopao-assets.oss-cn-shanghai.aliyuncs.com/public/avatar/default/julia.png",
	"https://paopao-assets.oss-cn-shanghai.aliyuncs.com/public/avatar/default/joshua.png",
	"https://paopao-assets.oss-cn-shanghai.aliyuncs.com/public/avatar/default/john.png",
	"https://paopao-assets.oss-cn-shanghai.aliyuncs.com/public/avatar/default/jane.png",
	"https://paopao-assets.oss-cn-shanghai.aliyuncs.com/public/avatar/default/jackson.png",
	"https://paopao-assets.oss-cn-shanghai.aliyuncs.com/public/avatar/default/ivy.png",
	"https://paopao-assets.oss-cn-shanghai.aliyuncs.com/public/avatar/default/isaac.png",
	"https://paopao-assets.oss-cn-shanghai.aliyuncs.com/public/avatar/default/henry.png",
	"https://paopao-assets.oss-cn-shanghai.aliyuncs.com/public/avatar/default/harry.png",
	"https://paopao-assets.oss-cn-shanghai.aliyuncs.com/public/avatar/default/harold.png",
	"https://paopao-assets.oss-cn-shanghai.aliyuncs.com/public/avatar/default/hanna.png",
	"https://paopao-assets.oss-cn-shanghai.aliyuncs.com/public/avatar/default/grace.png",
	"https://paopao-assets.oss-cn-shanghai.aliyuncs.com/public/avatar/default/george.png",
	"https://paopao-assets.oss-cn-shanghai.aliyuncs.com/public/avatar/default/freddy.png",
	"https://paopao-assets.oss-cn-shanghai.aliyuncs.com/public/avatar/default/frank.png",
	"https://paopao-assets.oss-cn-shanghai.aliyuncs.com/public/avatar/default/finn.png",
	"https://paopao-assets.oss-cn-shanghai.aliyuncs.com/public/avatar/default/emma.png",
	"https://paopao-assets.oss-cn-shanghai.aliyuncs.com/public/avatar/default/emily.png",
	"https://paopao-assets.oss-cn-shanghai.aliyuncs.com/public/avatar/default/edward.png",
	"https://paopao-assets.oss-cn-shanghai.aliyuncs.com/public/avatar/default/clara.png",
	"https://paopao-assets.oss-cn-shanghai.aliyuncs.com/public/avatar/default/claire.png",
	"https://paopao-assets.oss-cn-shanghai.aliyuncs.com/public/avatar/default/chloe.png",
	"https://paopao-assets.oss-cn-shanghai.aliyuncs.com/public/avatar/default/audrey.png",
	"https://paopao-assets.oss-cn-shanghai.aliyuncs.com/public/avatar/default/arthur.png",
	"https://paopao-assets.oss-cn-shanghai.aliyuncs.com/public/avatar/default/anna.png",
	"https://paopao-assets.oss-cn-shanghai.aliyuncs.com/public/avatar/default/andy.png",
	"https://paopao-assets.oss-cn-shanghai.aliyuncs.com/public/avatar/default/alfred.png",
	"https://paopao-assets.oss-cn-shanghai.aliyuncs.com/public/avatar/default/alexa.png",
	"https://paopao-assets.oss-cn-shanghai.aliyuncs.com/public/avatar/default/abigail.png",
}

func (s *Service) GetRandomAvatar() string {
	rand.Seed(time.Now().UnixMicro())
	return defaultAvatars[rand.Intn(len(defaultAvatars))]
}
