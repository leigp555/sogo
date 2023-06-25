package bucket

import (
	"github.com/juju/ratelimit"
	"sogo/app/global/variable"
	"time"
)

func NewBucket() *ratelimit.Bucket {
	fillInterval := variable.Config.GetInt("bucket.fillInterval")
	quantum := variable.Config.GetInt64("bucket.quantum")
	capacity := variable.Config.GetInt64("bucket.capacity")
	// 创建指定填充速率、容量大小和每次填充的令牌数的令牌桶
	bucket := ratelimit.NewBucketWithQuantum(time.Second*time.Duration(fillInterval), capacity, quantum)
	return bucket
}
