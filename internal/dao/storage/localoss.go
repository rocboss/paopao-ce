// Copyright 2022 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package storage

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/Masterminds/semver/v3"
	"github.com/cockroachdb/errors"
	"github.com/rocboss/paopao-ce/internal/core"
	"github.com/sirupsen/logrus"
)

var (
	_ core.ObjectStorageService = (*localossServant)(nil)
	_ core.OssCreateService     = (*localossCreateServant)(nil)
	_ core.OssCreateService     = (*localossCreateTempDirServant)(nil)
	_ core.VersionInfo          = (*localossServant)(nil)
)

type localossCreateServant struct {
	savePath string
	domain   string
}

type localossCreateTempDirServant struct {
	savePath string
	domain   string
	tempDir  string
}

type localossServant struct {
	core.OssCreateService

	savePath string
	domain   string
}

func (s *localossCreateServant) PutObject(objectKey string, reader io.Reader, objectSize int64, contentType string, _persistance bool) (string, error) {
	saveDir := s.savePath + filepath.Dir(objectKey)
	err := os.MkdirAll(saveDir, 0750)
	if err != nil && !os.IsExist(err) {
		return "", err
	}

	savePath := s.savePath + objectKey
	writer, err := os.Create(savePath)
	if err != nil {
		return "", err
	}
	defer writer.Close()

	written, err := io.Copy(writer, reader)
	if err != nil {
		return "", err
	}
	if written != objectSize {
		os.Remove(savePath)
		return "", errors.New("put object not complete")
	}

	return s.domain + objectKey, nil
}

func (s *localossCreateServant) PersistObject(_objectKey string) error {
	// empty
	return nil
}

func (s *localossCreateTempDirServant) PutObject(objectKey string, reader io.Reader, objectSize int64, contentType string, persistance bool) (string, error) {
	objectName := objectKey
	if !persistance {
		objectName = s.tempDir + objectKey
	}
	saveDir := s.savePath + filepath.Dir(objectName)
	err := os.MkdirAll(saveDir, 0750)
	if err != nil && !os.IsExist(err) {
		return "", err
	}

	savePath := s.savePath + objectName
	writer, err := os.Create(savePath)
	if err != nil {
		return "", err
	}
	defer writer.Close()

	written, err := io.Copy(writer, reader)
	if err != nil {
		return "", err
	}
	if written != objectSize {
		os.Remove(savePath)
		return "", errors.New("put object not complete")
	}

	return s.domain + objectKey, nil
}

func (s *localossCreateTempDirServant) PersistObject(objectKey string) error {
	fi, err := os.Stat(s.savePath + objectKey)
	if err == nil && !fi.IsDir() {
		logrus.Debugf("object exist so do nothing objectKey: %s", objectKey)
		return nil
	}

	saveDir := s.savePath + filepath.Dir(objectKey)
	if err = os.MkdirAll(saveDir, 0750); err != nil && !os.IsExist(err) {
		return err
	}

	tmpObjPath := s.savePath + s.tempDir + objectKey
	reader, err := os.Open(tmpObjPath)
	if err != nil {
		return err
	}
	needCloseReader := true
	defer func() {
		if needCloseReader {
			reader.Close()
		}
	}()

	writer, err := os.Create(s.savePath + objectKey)
	if err != nil {
		return err
	}
	defer writer.Close()
	if _, err = io.Copy(writer, reader); err != nil {
		return err
	}

	reader.Close()
	needCloseReader = false
	if err = os.Remove(tmpObjPath); err != nil {
		return err
	}

	return nil
}

func (s *localossServant) DeleteObject(objectKey string) error {
	return os.Remove(s.savePath + objectKey)
}

func (s *localossServant) DeleteObjects(objectKeys []string) (err error) {
	// 宽松处理删除动作，尽可能删除所有objectKey，如果出错，只返回最后一个错误
	for _, objectKey := range objectKeys {
		if e := os.Remove(s.savePath + objectKey); e != nil {
			err = e
		}
	}
	return
}

func (s *localossServant) IsObjectExist(objectKey string) (bool, error) {
	fi, err := os.Stat(s.savePath + objectKey)
	if err != nil {
		return false, err
	}
	return !fi.IsDir(), nil
}

func (s *localossServant) SignURL(objectKey string, expiredInSec int64) (string, error) {
	if expiredInSec < 0 {
		return "", fmt.Errorf("invalid expires: %d, expires must bigger than 0", expiredInSec)
	}
	expiration := time.Now().Unix() + expiredInSec

	// Fixed: Just make things simple and simple now so return an veiry simple sign url.
	// Maybe make another process logic for sign url in future but not now.
	uri := fmt.Sprintf("%s%s?expired=%d", s.domain, objectKey, expiration)
	return uri, nil
}

func (s *localossServant) ObjectURL(objetKey string) string {
	return s.domain + objetKey
}

func (s *localossServant) ObjectKey(objectUrl string) string {
	return strings.Replace(objectUrl, s.domain, "", -1)
}

func (s *localossServant) Name() string {
	return "LocalOSS"
}

func (s *localossServant) Version() *semver.Version {
	return semver.MustParse("v0.2.0")
}
