package database

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRedis(t *testing.T) {
	address := "127.0.0.1:6379"
	redisClient, cancel, err := NewRedisClient(address)
	assert.NoError(t, err)
	defer cancel()
	err = redisClient.SetWeather("mashhad", "sum", "sam")
	assert.NoError(t, err)
	summ, samm, err := redisClient.GetWeather("mashhad")
	assert.NoError(t, err)
	fmt.Println(summ, samm)
}
