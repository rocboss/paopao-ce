//go:build constraint

package storage

import (
	"github.com/rocboss/paopao-ce/internal/core"
)

var (
	_ core.ObjectStorageService = (*aliossServant)(nil)
	_ core.OssCreateService     = (*aliossCreateServant)(nil)
	_ core.OssCreateService     = (*aliossCreateRetentionServant)(nil)
	_ core.OssCreateService     = (*aliossCreateTempDirServant)(nil)
	_ core.VersionInfo          = (*aliossServant)(nil)

	_ core.ObjectStorageService = (*cosServant)(nil)
	_ core.OssCreateService     = (*cosCreateServant)(nil)
	_ core.OssCreateService     = (*cosCreateTempDirServant)(nil)
	_ core.VersionInfo          = (*cosServant)(nil)

	_ core.ObjectStorageService = (*huaweiobsServant)(nil)
	_ core.OssCreateService     = (*hwobsCreateServant)(nil)
	_ core.OssCreateService     = (*hwobsCreateRetentionServant)(nil)
	_ core.OssCreateService     = (*hwobsCreateTempDirServant)(nil)
	_ core.VersionInfo          = (*huaweiobsServant)(nil)

	_ core.ObjectStorageService = (*localossServant)(nil)
	_ core.OssCreateService     = (*localossCreateServant)(nil)
	_ core.OssCreateService     = (*localossCreateTempDirServant)(nil)
	_ core.VersionInfo          = (*localossServant)(nil)

	_ core.ObjectStorageService = (*minioServant)(nil)
	_ core.OssCreateService     = (*minioCreateServant)(nil)
	_ core.OssCreateService     = (*minioCreateRetentionServant)(nil)
	_ core.OssCreateService     = (*minioCreateTempDirServant)(nil)
	_ core.VersionInfo          = (*minioServant)(nil)
)
