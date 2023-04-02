// Copyright 2023 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package cache

import (
	"context"
	"fmt"
	"time"
	"unsafe"

	"github.com/Masterminds/semver/v3"
	"github.com/rocboss/paopao-ce/internal/core"
	"github.com/rueian/rueidis"
)

var (
	_ core.RedisCache = (*redisCache)(nil)
	_ tweetsCache     = (*redisCacheTweetsCache)(nil)
)

const (
	_cacheIndexKeyPattern = _cacheIndexKey + "*"
	_pushToSearchJobKey   = "paopao_push_to_search_job"
	_countLoginErrKey     = "paopao_count_login_err"
	_imgCaptchaKey        = "paopao_img_captcha:"
	_smsCaptchaKey        = "paopao_sms_captcha"
	_countWhisperKey      = "paopao_whisper_key"
	_rechargeStatusKey    = "paopao_recharge_status:"
)

type redisCache struct {
	c rueidis.Client
}

type redisCacheTweetsCache struct {
	expireDuration time.Duration
	expireInSecond int64
	c              rueidis.Client
}

func (s *redisCacheTweetsCache) getTweetsBytes(key string) ([]byte, error) {
	res, err := rueidis.MGetCache(s.c, context.Background(), s.expireDuration, []string{key})
	if err != nil {
		return nil, err
	}
	message := res[key]
	return message.AsBytes()
}

func (s *redisCacheTweetsCache) setTweetsBytes(key string, bs []byte) error {
	cmd := s.c.B().Set().Key(key).Value(rueidis.BinaryString(bs)).ExSeconds(s.expireInSecond).Build()
	return s.c.Do(context.Background(), cmd).Error()
}

func (s *redisCacheTweetsCache) delTweets(keys []string) error {
	cmd := s.c.B().Del().Key(keys...).Build()
	return s.c.Do(context.Background(), cmd).Error()
}

func (s *redisCacheTweetsCache) allKeys() ([]string, error) {
	cmd := s.c.B().Keys().Pattern(_cacheIndexKeyPattern).Build()
	return s.c.Do(context.Background(), cmd).AsStrSlice()
}

func (s *redisCacheTweetsCache) Name() string {
	return "RedisCacheIndex"
}

func (s *redisCacheTweetsCache) Version() *semver.Version {
	return semver.MustParse("v0.1.0")
}

func (r *redisCache) SetPushToSearchJob(ctx context.Context) error {
	return r.c.Do(ctx, r.c.B().Set().
		Key(_pushToSearchJobKey).Value("1").
		Nx().ExSeconds(3600).
		Build()).Error()
}

func (r *redisCache) DelPushToSearchJob(ctx context.Context) error {
	return r.c.Do(ctx, r.c.B().Del().Key(_pushToSearchJobKey).Build()).Error()
}

func (r *redisCache) SetImgCaptcha(ctx context.Context, id string, value string) error {
	return r.c.Do(ctx, r.c.B().Set().
		Key(_imgCaptchaKey+id).Value(value).
		ExSeconds(300).
		Build()).Error()
}

func (r *redisCache) GetImgCaptcha(ctx context.Context, id string) (string, error) {
	res, err := r.c.Do(ctx, r.c.B().Get().Key(_imgCaptchaKey+id).Build()).AsBytes()
	return unsafe.String(&res[0], len(res)), err
}

func (r *redisCache) DelImgCaptcha(ctx context.Context, id string) error {
	return r.c.Do(ctx, r.c.B().Del().Key(_imgCaptchaKey+id).Build()).Error()
}

func (r *redisCache) GetCountSmsCaptcha(ctx context.Context, phone string) (int64, error) {
	return r.c.Do(ctx, r.c.B().Get().Key(_smsCaptchaKey+phone).Build()).AsInt64()
}

func (r *redisCache) IncrCountSmsCaptcha(ctx context.Context, phone string) (err error) {
	if err = r.c.Do(ctx, r.c.B().Incr().Key(_smsCaptchaKey+phone).Build()).Error(); err == nil {
		currentTime := time.Now()
		endTime := time.Date(currentTime.Year(), currentTime.Month(), currentTime.Day(), 23, 59, 59, 0, currentTime.Location())
		err = r.c.Do(ctx, r.c.B().Expire().Key(_smsCaptchaKey+phone).Seconds(int64(endTime.Sub(currentTime)/time.Second)).Build()).Error()
	}
	return
}

func (r *redisCache) GetCountLoginErr(ctx context.Context, id int64) (int64, error) {
	return r.c.Do(ctx, r.c.B().Get().Key(fmt.Sprintf("%s:%d", _countLoginErrKey, id)).Build()).AsInt64()
}

func (r *redisCache) DelCountLoginErr(ctx context.Context, id int64) error {
	return r.c.Do(ctx, r.c.B().Del().Key(fmt.Sprintf("%s:%d", _countLoginErrKey, id)).Build()).Error()
}

func (r *redisCache) IncrCountLoginErr(ctx context.Context, id int64) error {
	err := r.c.Do(ctx, r.c.B().Incr().Key(fmt.Sprintf("%s:%d", _countLoginErrKey, id)).Build()).Error()
	if err == nil {
		err = r.c.Do(ctx, r.c.B().Expire().Key(fmt.Sprintf("%s:%d", _countLoginErrKey, id)).Seconds(3600).Build()).Error()
	}
	return err
}

func (r *redisCache) GetCountWhisper(ctx context.Context, uid int64) (int64, error) {
	return r.c.Do(ctx, r.c.B().Get().Key(fmt.Sprintf("%s:%d", _countWhisperKey, uid)).Build()).AsInt64()
}

func (r *redisCache) IncrCountWhisper(ctx context.Context, uid int64) (err error) {
	key := fmt.Sprintf("%s:%d", _countWhisperKey, uid)
	if err = r.c.Do(ctx, r.c.B().Incr().Key(key).Build()).Error(); err == nil {
		currentTime := time.Now()
		endTime := time.Date(currentTime.Year(), currentTime.Month(), currentTime.Day(), 23, 59, 59, 0, currentTime.Location())
		err = r.c.Do(ctx, r.c.B().Expire().Key(key).Seconds(int64(endTime.Sub(currentTime)/time.Second)).Build()).Error()
	}
	return
}

func (r *redisCache) SetRechargeStatus(ctx context.Context, tradeNo string) error {
	return r.c.Do(ctx, r.c.B().Set().
		Key(_rechargeStatusKey+tradeNo).Value("1").
		Nx().ExSeconds(5).Build()).Error()
}

func (r *redisCache) DelRechargeStatus(ctx context.Context, tradeNo string) error {
	return r.c.Do(ctx, r.c.B().Del().Key(_rechargeStatusKey+tradeNo).Build()).Error()
}
