package pkg

import (
	"crypto/tls"
	"github.com/redis/go-redis/v9"
	"time"
)

const redisDefaultTTL = 24 * time.Hour

type RedisConfig struct {
	IsCluster bool     `json:"IsCluster,default=false"`
	Addrs     []string `json:"Addrs"`
	Password  string   `json:"Password,optional"`
	TLS       bool     `json:"TLS,default=false"`
}

// NewClient 新建客户端
func (x *RedisConfig) NewClient() redis.UniversalClient {
	if x.IsCluster {
		return x._clusterClient()
	}
	return x._client()
}

// 单机
func (x *RedisConfig) _client() *redis.Client {
	var tlsConfig *tls.Config
	if x.TLS {
		tlsConfig = &tls.Config{
			InsecureSkipVerify: true,
		}
	}
	return redis.NewClient(&redis.Options{
		Addr:      x.Addrs[0],
		Password:  x.Password,
		TLSConfig: tlsConfig,
	})
}

// 集群
func (x *RedisConfig) _clusterClient() *redis.ClusterClient {
	var tlsConfig *tls.Config
	if x.TLS {
		tlsConfig = &tls.Config{
			InsecureSkipVerify: true,
		}
	}
	return redis.NewClusterClient(&redis.ClusterOptions{
		Addrs:     x.Addrs,
		Password:  x.Password,
		TLSConfig: tlsConfig,
	})
}
