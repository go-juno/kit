package nosql

import (
	"context"
	"time"

	"github.com/bsm/redislock"
	"github.com/go-redis/cache/v8"
	"github.com/go-redis/redis/v8"
	"golang.org/x/xerrors"
)

type RedisClient interface {
	// Del 删除keys
	Del(ctx context.Context, keys ...string) (int64, error)
	// Exists key 是否存在
	Exists(ctx context.Context, keys ...string) bool
	// EXPIRE 给定 key 设置过期时间，以秒计。
	EXPIRE(ctx context.Context, key string, time time.Duration) bool
	// EXPIREAT  为 key 设置过期时间。 不同在于 EXPIREAT 命令接受的时间参数是 UNIX 时间戳(unix timestamp)。
	EXPIREAT(ctx context.Context, key string, time time.Time) bool
	// TTL 以秒为单位，返回给定 key 的剩余生存时间(TTL, time to live)。
	TTL(ctx context.Context, key string) (time.Duration, error)
	// PERSIST 移除 key 的过期时间，key 将持久保持。
	PERSIST(ctx context.Context, key string) (bool, error)
	// Set 设置指定 key 的值
	Set(ctx context.Context, key string, value interface{}, expiration time.Duration) (string, error)
	// SetNx 分布式常用
	SetNx(ctx context.Context, key string, value interface{}, expiration time.Duration) bool
	// Get 获取指定 key 的值。
	Get(ctx context.Context, key string) (string, error)
	// HDel 哈希hash
	HDel(ctx context.Context, key string, fields ...string) error
	HExists(ctx context.Context, key, field string) bool
	HGet(ctx context.Context, key, field string) string
	HGetAll(ctx context.Context, key string) map[string]string
	// BlPOP 列表
	BlPOP(ctx context.Context, timeout time.Duration, keys ...string) []string
	BrPOP(ctx context.Context, timeout time.Duration, keys ...string) []string
	BrPOPLPush(ctx context.Context, source, destination string, timeout time.Duration) string
	LIndex(ctx context.Context, key string, index int64) string
	LLen(ctx context.Context, key string) int64
	// SAdd 集合
	SAdd(ctx context.Context, key string, members ...interface{}) (bool, error)
	SCard(ctx context.Context, key string) int64
	SDiff(ctx context.Context, keys ...string) ([]string, error)
	SINTER(ctx context.Context, keys ...string) ([]string, error)
	SIsMember(ctx context.Context, key string, member interface{}) (bool, error)
	SMembers(ctx context.Context, key string) ([]string, error)
	SPop(ctx context.Context, key string) (string, error)
	SUnion(ctx context.Context, keys ...string) ([]string, error)

	// ZAdd 有序集合
	ZAdd(ctx context.Context, key string, members ...*Z) (bool, error)
	ZRank(ctx context.Context, key, member string) (int64, error)
	ZCard(ctx context.Context, key string) (int64, error)
	ZLexCount(ctx context.Context, key, min, max string) (int64, error)
	ZRangeByLex(ctx context.Context, key string, opt *ZRangeBy) ([]string, error)
	ZRangeByScore(ctx context.Context, key string, opt *ZRangeBy) ([]string, error)
	ZRemRangeByLex(ctx context.Context, key, min, max string) (int64, error)
	ZRem(ctx context.Context, key string, members ...interface{}) (int64, error)
	ZCount(ctx context.Context, key, min, max string) (int64, error)
	ZScore(ctx context.Context, key, member string) (float64, error)
}
type redisClient struct {
	client *redis.Client
	Cache  *cache.Cache
	Locker *redislock.Client
}

func (r *redisClient) Get(ctx context.Context, key string) (string, error) {
	result, err := r.client.Get(ctx, key).Result()
	if err != nil {
		err = xerrors.Errorf("%w", err)
		return "", err
	}
	return result, nil
}

// EXPIRE 给定 key 设置过期时间，以秒计。
func (r *redisClient) EXPIRE(ctx context.Context, key string, time time.Duration) bool {
	result, err := r.client.Expire(ctx, key, time).Result()
	if err != nil {
		return false
	}

	return result
}

// EXPIREAT  为 key 设置过期时间。 不同在于 EXPIREAT 命令接受的时间参数是 UNIX 时间戳(unix timestamp)。
func (r *redisClient) EXPIREAT(ctx context.Context, key string, time time.Time) bool {
	result, err := r.client.ExpireAt(ctx, key, time).Result()
	if err != nil {
		return false
	}

	return result
}

func (r *redisClient) PERSIST(ctx context.Context, key string) (bool, error) {
	duration, err := r.client.Persist(ctx, key).Result()
	if err != nil {
		err = xerrors.Errorf("%w", err)
		return false, err
	}
	return duration, nil
}
func (r *redisClient) TTL(ctx context.Context, key string) (time.Duration, error) {
	duration, err := r.client.TTL(ctx, key).Result()
	if err != nil {
		err = xerrors.Errorf("%w", err)
		return -1, err
	}
	return duration, nil
}
func (r *redisClient) ZScore(ctx context.Context, key, member string) (float64, error) {
	result, err := r.client.ZScore(ctx, key, member).Result()
	if err != nil {
		err = xerrors.Errorf("%w", err)
		return -1, err
	}
	return result, nil
}

func (r *redisClient) SAdd(ctx context.Context, key string, members ...interface{}) (bool, error) {
	err := r.client.SAdd(ctx, key, members).Err()
	if err != nil {
		err = xerrors.Errorf("%w", err)
		return false, err
	}
	return true, nil
}

// SCard 获取集合的成员数
func (r *redisClient) SCard(ctx context.Context, key string) int64 {
	result, err := r.client.SCard(ctx, key).Result()
	if err != nil {
		return -1
	}
	return result
}

func (r *redisClient) SDiff(ctx context.Context, keys ...string) ([]string, error) {
	result, err := r.client.SDiff(ctx, keys...).Result()
	if err != nil {
		err = xerrors.Errorf("%w", err)
		return nil, err
	}
	return result, nil
}
func (r *redisClient) SINTER(ctx context.Context, keys ...string) ([]string, error) {
	result, err := r.client.SInter(ctx, keys...).Result()
	if err != nil {
		err = xerrors.Errorf("%w", err)
		return nil, err
	}
	return result, nil
}

func (r *redisClient) SIsMember(ctx context.Context, key string, member interface{}) (bool, error) {
	result, err := r.client.SIsMember(ctx, key, member).Result()
	if err != nil {
		err = xerrors.Errorf("%w", err)
		return false, err
	}
	return result, nil
}

func (r *redisClient) SMembers(ctx context.Context, key string) ([]string, error) {
	result, err := r.client.SMembers(ctx, key).Result()
	if err != nil {
		err = xerrors.Errorf("%w", err)
		return nil, err
	}
	return result, nil
}

// SPop 移除并返回集合中的一个随机元素
func (r *redisClient) SPop(ctx context.Context, key string) (string, error) {
	result, err := r.client.SPop(ctx, key).Result()
	if err != nil {
		err = xerrors.Errorf("%w", err)
		return "", err
	}
	return result, nil
}

// SUnion 返回所有给定集合的并集
func (r *redisClient) SUnion(ctx context.Context, keys ...string) ([]string, error) {
	result, err := r.client.SUnion(ctx, keys...).Result()
	if err != nil {
		err = xerrors.Errorf("%w", err)
		return nil, err
	}
	return result, nil
}

func (r *redisClient) ZAdd(ctx context.Context, key string, members ...*Z) (bool, error) {
	redisMemBers := make([]*redis.Z, 0)
	for _, member := range members {
		rZ := &redis.Z{
			Score:  member.Score,
			Member: member.Member,
		}
		redisMemBers = append(redisMemBers, rZ)
	}
	result, err := r.client.ZAdd(ctx, key, redisMemBers...).Result()
	if err != nil {
		err = xerrors.Errorf("%w", err)
		return false, err
	}
	if result == 0 {
		return false, nil
	}
	return true, nil
}

// ZRank 返回有序集合中指定成员的索引
func (r *redisClient) ZRank(ctx context.Context, key, member string) (int64, error) {
	result, err := r.client.ZRank(ctx, key, member).Result()
	if err != nil {
		err = xerrors.Errorf("%w", err)
		return -1, err
	}
	return result, nil
}

// ZCard 获取有序集合的成员数
func (r *redisClient) ZCard(ctx context.Context, key string) (int64, error) {
	result, err := r.client.ZCard(ctx, key).Result()
	if err != nil {
		err = xerrors.Errorf("%w", err)
		return -1, err
	}
	return result, nil
}

// ZLexCount 在有序集合中计算指定字典区间内成员数量
func (r *redisClient) ZLexCount(ctx context.Context, key, min, max string) (int64, error) {
	result, err := r.client.ZLexCount(ctx, key, min, max).Result()
	if err != nil {
		err = xerrors.Errorf("%w", err)
		return -1, err
	}
	return result, nil
}

// ZRangeByLex 通过字典区间返回有序集合的成员
func (r *redisClient) ZRangeByLex(ctx context.Context, key string, opt *ZRangeBy) ([]string, error) {
	redisOpt := &redis.ZRangeBy{
		Min:    opt.Min,
		Max:    opt.Max,
		Offset: opt.Offset,
		Count:  opt.Count,
	}
	result, err := r.client.ZRangeByLex(ctx, key, redisOpt).Result()
	if err != nil {
		err = xerrors.Errorf("%w", err)
		return nil, err
	}
	return result, nil
}

func (r *redisClient) ZRangeByScore(ctx context.Context, key string, opt *ZRangeBy) ([]string, error) {
	redisOpt := &redis.ZRangeBy{
		Min:    opt.Min,
		Max:    opt.Max,
		Offset: opt.Offset,
		Count:  opt.Count,
	}
	result, err := r.client.ZRangeByScore(ctx, key, redisOpt).Result()
	if err != nil {
		err = xerrors.Errorf("%w", err)
		return nil, err
	}
	return result, nil
}

func (r *redisClient) ZRemRangeByLex(ctx context.Context, key, min, max string) (int64, error) {
	result, err := r.client.ZRemRangeByLex(ctx, key, min, max).Result()
	if err != nil {
		err = xerrors.Errorf("%w", err)
		return -1, err
	}
	return result, nil
}

func (r *redisClient) ZRem(ctx context.Context, key string, members ...interface{}) (int64, error) {
	result, err := r.client.ZRem(ctx, key, members...).Result()
	if err != nil {
		err = xerrors.Errorf("%w", err)
		return -1, err
	}
	return result, nil
}

func (r *redisClient) ZCount(ctx context.Context, key, min, max string) (int64, error) {
	result, err := r.client.ZCount(ctx, key, min, max).Result()
	if err != nil {
		err = xerrors.Errorf("%w", err)
		return -1, err
	}
	return result, nil
}

func (r *redisClient) Set(ctx context.Context, key string, value interface{}, expiration time.Duration) (string, error) {
	result, err := r.client.Set(ctx, key, value, expiration).Result()
	if err != nil {
		err = xerrors.Errorf("%w", err)
		return "", err
	}
	return result, nil
}

func (r *redisClient) Del(ctx context.Context, keys ...string) (int64, error) {
	result, err := r.client.Del(ctx, keys...).Result()
	if err != nil {
		err = xerrors.Errorf("%w", err)
		return -100, err
	}
	return result, nil
}

func (r *redisClient) Exists(ctx context.Context, keys ...string) bool {
	result, err := r.client.Exists(ctx, keys...).Result()
	if err != nil {
		return false
	}
	if result == 0 {
		return false
	}
	return true

}

func (r *redisClient) HDel(ctx context.Context, key string, fields ...string) error {
	err := r.client.HDel(ctx, key, fields...).Err()
	if err != nil {
		err = xerrors.Errorf("%w", err)
		return err
	}
	return nil
}

func (r *redisClient) HExists(ctx context.Context, key, field string) bool {
	result, err := r.client.HExists(ctx, key, field).Result()
	if err != nil {
		return false
	}
	return result
}

func (r *redisClient) HGet(ctx context.Context, key, field string) string {
	result, err := r.client.HGet(ctx, key, field).Result()
	if err != nil {
		return ""
	}
	return result
}

func (r *redisClient) HGetAll(ctx context.Context, key string) map[string]string {
	result, err := r.client.HGetAll(ctx, key).Result()
	if err != nil {
		return nil
	}
	return result
}

func (r *redisClient) BlPOP(ctx context.Context, timeout time.Duration, keys ...string) []string {
	result, err := r.client.BLPop(ctx, timeout, keys...).Result()
	if err != nil {
		return []string{}
	}
	return result
}

func (r *redisClient) BrPOP(ctx context.Context, timeout time.Duration, keys ...string) []string {
	result, err := r.client.BRPop(ctx, timeout, keys...).Result()
	if err != nil {
		return []string{}
	}
	return result
}

func (r *redisClient) BrPOPLPush(ctx context.Context, source, destination string, timeout time.Duration) string {
	result, err := r.client.BRPopLPush(ctx, source, destination, timeout).Result()
	if err != nil {
		return ""
	}
	return result
}

func (r *redisClient) LIndex(ctx context.Context, key string, index int64) string {
	result, err := r.client.LIndex(ctx, key, index).Result()
	if err != nil {
		return ""
	}
	return result
}
func (r *redisClient) LLen(ctx context.Context, key string) int64 {
	result, err := r.client.LLen(ctx, key).Result()
	if err != nil {
		return -1
	}
	return result
}

func (r *redisClient) SetNx(ctx context.Context, key string, value interface{}, expiration time.Duration) bool {
	result, err := r.client.SetNX(ctx, key, value, expiration).Result()
	if err != nil {
		return false
	}
	return result
}

// RedisNoSQLClient  连接数据库
func RedisNoSQLClient(op *Options) (RedisClient, error) {
	dsn := &redis.Options{
		Network:         op.Network,
		Addr:            op.Addr,
		Username:        op.Username,
		Password:        op.Password,
		DB:              op.DB,
		MaxRetries:      op.MaxRetries,
		MinRetryBackoff: op.MinRetryBackoff,
		MaxRetryBackoff: op.MaxRetryBackoff,
		DialTimeout:     op.DialTimeout,
		ReadTimeout:     op.ReadTimeout,
		WriteTimeout:    op.WriteTimeout,
		PoolFIFO:        op.PoolFIFO,
		PoolSize:        op.PoolSize,
		PoolTimeout:     op.PoolTimeout,
	}
	ctx, canFn := context.WithTimeout(context.Background(), time.Second*10)
	defer canFn()
	db := redis.NewClient(dsn)
	if _, err := db.Ping(ctx).Result(); err != nil {
		err = xerrors.Errorf("%w", err)
		return nil, err
	}
	c := cache.New(&cache.Options{
		Redis:      db,
		LocalCache: cache.NewTinyLFU(1000, time.Minute),
	})

	locker := redislock.New(db)

	return &redisClient{
		client: db,
		Cache:  c,
		Locker: locker,
	}, nil
}
