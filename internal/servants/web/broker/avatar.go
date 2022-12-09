// Copyright 2022 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package broker

import (
	"math/rand"
	"time"
)

var defaultAvatars = []string{
	"https://assets.paopao.info/public/avatar/default/zoe.png",
	"https://assets.paopao.info/public/avatar/default/william.png",
	"https://assets.paopao.info/public/avatar/default/walter.png",
	"https://assets.paopao.info/public/avatar/default/thomas.png",
	"https://assets.paopao.info/public/avatar/default/taylor.png",
	"https://assets.paopao.info/public/avatar/default/sophia.png",
	"https://assets.paopao.info/public/avatar/default/sam.png",
	"https://assets.paopao.info/public/avatar/default/ryan.png",
	"https://assets.paopao.info/public/avatar/default/ruby.png",
	"https://assets.paopao.info/public/avatar/default/quinn.png",
	"https://assets.paopao.info/public/avatar/default/paul.png",
	"https://assets.paopao.info/public/avatar/default/owen.png",
	"https://assets.paopao.info/public/avatar/default/olivia.png",
	"https://assets.paopao.info/public/avatar/default/norman.png",
	"https://assets.paopao.info/public/avatar/default/nora.png",
	"https://assets.paopao.info/public/avatar/default/natalie.png",
	"https://assets.paopao.info/public/avatar/default/naomi.png",
	"https://assets.paopao.info/public/avatar/default/miley.png",
	"https://assets.paopao.info/public/avatar/default/mike.png",
	"https://assets.paopao.info/public/avatar/default/lucas.png",
	"https://assets.paopao.info/public/avatar/default/kylie.png",
	"https://assets.paopao.info/public/avatar/default/julia.png",
	"https://assets.paopao.info/public/avatar/default/joshua.png",
	"https://assets.paopao.info/public/avatar/default/john.png",
	"https://assets.paopao.info/public/avatar/default/jane.png",
	"https://assets.paopao.info/public/avatar/default/jackson.png",
	"https://assets.paopao.info/public/avatar/default/ivy.png",
	"https://assets.paopao.info/public/avatar/default/isaac.png",
	"https://assets.paopao.info/public/avatar/default/henry.png",
	"https://assets.paopao.info/public/avatar/default/harry.png",
	"https://assets.paopao.info/public/avatar/default/harold.png",
	"https://assets.paopao.info/public/avatar/default/hanna.png",
	"https://assets.paopao.info/public/avatar/default/grace.png",
	"https://assets.paopao.info/public/avatar/default/george.png",
	"https://assets.paopao.info/public/avatar/default/freddy.png",
	"https://assets.paopao.info/public/avatar/default/frank.png",
	"https://assets.paopao.info/public/avatar/default/finn.png",
	"https://assets.paopao.info/public/avatar/default/emma.png",
	"https://assets.paopao.info/public/avatar/default/emily.png",
	"https://assets.paopao.info/public/avatar/default/edward.png",
	"https://assets.paopao.info/public/avatar/default/clara.png",
	"https://assets.paopao.info/public/avatar/default/claire.png",
	"https://assets.paopao.info/public/avatar/default/chloe.png",
	"https://assets.paopao.info/public/avatar/default/audrey.png",
	"https://assets.paopao.info/public/avatar/default/arthur.png",
	"https://assets.paopao.info/public/avatar/default/anna.png",
	"https://assets.paopao.info/public/avatar/default/andy.png",
	"https://assets.paopao.info/public/avatar/default/alfred.png",
	"https://assets.paopao.info/public/avatar/default/alexa.png",
	"https://assets.paopao.info/public/avatar/default/abigail.png",
}

func GetRandomAvatar() string {
	rand.Seed(time.Now().UnixMicro())
	return defaultAvatars[rand.Intn(len(defaultAvatars))]
}
