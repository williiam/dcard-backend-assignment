package service

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
	_ "github.com/lib/pq"
	dao "github.com/williiam/dcard-backend-assignment/internal/app/dao"
)

var (
	ctx = context.Background()
)

const CacheDuration = 6 * time.Hour

type RedisService interface {
	Save(shortURL string, urlMapping dao.URLMapping, expTime time.Duration) error
	Get(shortURL string) (dao.URLMapping, error)
	Update(shortURL string, urlMapping dao.URLMapping) error
}

type RedisServiceClient struct {
	RedisClient *redis.Client
}

func InitService() *redis.Client {
	redisClient := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	pong, err := redisClient.Ping(ctx).Result()
	if err != nil {
		panic(fmt.Sprintf("Error init Redis: %v", err))
	}

	fmt.Printf("\nRedis started successfully: pong message = {%s}", pong)
	return redisClient
}


/* We want to be able to save the mapping between the originalUrl
and the generated shortUrl url
*/
func (s RedisServiceClient) SaveUrlMapping (shortUrl string, urlMapping dao.URLMapping) (error) {
	originalUrl := urlMapping.OriginalURL
	err := s.RedisClient.Set(ctx, shortUrl, originalUrl, CacheDuration).Err()
	if err != nil {
		panic(fmt.Sprintf("Failed saving key url | Error: %v - shortUrl: %s - originalUrl: %s\n", err, shortUrl, originalUrl))
	}

	fmt.Printf("Saved shortUrl: %s - originalUrl: %s\n", shortUrl, originalUrl)
	return err
}

/*
We should be able to retrieve the initial long URL once the short
is provided. This is when users will be calling the shortlink in the
url, so what we need to do here is to retrieve the long url and
think about redirect.
*/
func (s RedisServiceClient) RetrieveInitialUrl (shortUrl string) (dao.URLMapping, error) {
	var urlMapping dao.URLMapping
	result, err := s.RedisClient.Get(ctx, shortUrl).Result()
	if err != nil {
		// panic(fmt.Sprintf("Failed RetrieveInitialUrl url | Error: %v - shortUrl: %s\n", err, shortUrl))
		fmt.Printf("Failed RetrieveInitialUrl url | Error: %v - shortUrl: %s\n", err, shortUrl)
	}
	urlMapping.OriginalURL = result
	return urlMapping,err
}
