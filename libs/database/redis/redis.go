package redis

import (
	"context"
	json2 "encoding/json"
	"fmt"
	redis "github.com/redis/go-redis/v9"
	"log"
)

var Client *redis.Client

type IRedis interface {
	RedisClientConnection() (*redis.Client, error)
	Insert(key string, value map[string]interface{}) error
	Delete(key string) error
	Get(key string) (*Response, error)
}

type Redis struct{}

func NewRedisRepository() IRedis {
	return &Redis{}
}

func InitializeRedis() {
	Client = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	pong, err := Client.Ping(context.Background()).Result()
	if err != nil {
		log.Fatalln("\n[Redis Connection Error]\n", err)
	}

	fmt.Println("\n[Redis Connection Success]\n", pong)
}

func (r Redis) RedisClientConnection() (*redis.Client, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	_, err := client.Ping(context.Background()).Result()
	if err != nil {
		log.Fatalln("\n[Redis Connection Error]\n", err)
	}

	return client, nil
}

func (r Redis) Insert(key string, value map[string]interface{}) error {
	client, err := r.RedisClientConnection()
	if err != nil {
		return err
	}

	data, err := json2.Marshal(value)
	if err != nil {
		log.Println("Error serializing struct to JSON", err)
		return err
	}

	err = client.Set(context.Background(), key, data, 0).Err()
	if err != nil {
		log.Println("Error inserting data into redis", err)
		return err
	}

	fmt.Println("Data inserted into redis successfully")
	return nil
}

func (r Redis) Delete(key string) error {
	client, err := r.RedisClientConnection()
	if err != nil {
		return err
	}

	defer func(client *redis.Client) {
		err := client.Close()
		if err != nil {
			fmt.Println("Redis is not closed yet.")
		}
	}(client)

	err = client.Del(context.Background(), key).Err()
	if err != nil {
		log.Println("Error deleting data from redis: ", err)
		return err
	}

	fmt.Println("Data deleted from Redis successfully.")
	return nil
}

type Response struct {
	AccessToken  string
	RefreshToken string
}

func (r Redis) Get(key string) (*Response, error) {
	client, err := r.RedisClientConnection()
	if err != nil {
		return nil, err
	}

	defer func(client *redis.Client) {
		err := client.Close()
		if err != nil {
			fmt.Println("Redis is not closed yet.")
		}
	}(client)

	data, err := client.Get(context.Background(), key).Result()
	if err != nil {
		if err == redis.Nil {
			return nil, fmt.Errorf("key not found")
		}
		return nil, err
	}

	var result map[string]string
	if err := json2.Unmarshal([]byte(data), &result); err != nil {
		return nil, err
	}

	response := Response{
		AccessToken:  result["accessToken"],
		RefreshToken: result["refreshToken"],
	}

	return &response, nil
}
