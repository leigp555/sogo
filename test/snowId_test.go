package test

import (
	"fmt"
	sn "github.com/bwmarrin/snowflake"
	"sync"
	"testing"
	"time"
)

type InterfaceSnowFlake interface {
	GetId() int64
}

// 创建一个雪花算法生成器(生成工厂)
func CreateSnowflakeFactory() InterfaceSnowFlake {
	return &snowflake{
		timestamp: 0,
		machineId: 2,
		sequence:  0,
	}
}

func TestSnow(t *testing.T) {
	snow := CreateSnowflakeFactory()
	id := snow.GetId()
	fmt.Println(id)
	fmt.Println("================================")
	uid, _ := GetUid()
	fmt.Println(uid)
}

type snowflake struct {
	sync.Mutex
	timestamp int64
	machineId int64
	sequence  int64
}

// 生成分布式ID
func (s *snowflake) GetId() int64 {
	s.Lock()
	defer func() {
		s.Unlock()
	}()
	now := time.Now().UnixNano() / 1e6
	if s.timestamp == now {
		s.sequence = (s.sequence + 1) & int64(-1^(-1<<uint(10)))
		if s.sequence == 0 {
			for now <= s.timestamp {
				now = time.Now().UnixNano() / 1e6
			}
		}
	} else {
		s.sequence = 0
	}
	s.timestamp = now
	r := (now-int64(1483228800000))<<(uint(12)+uint(10)) | (s.machineId << uint(12)) | (s.sequence)
	return r
}

func GetUid() (uid string, err error) {
	//生成uid
	node, err := sn.NewNode(1)
	uid = node.Generate().String()
	return uid, err
}
