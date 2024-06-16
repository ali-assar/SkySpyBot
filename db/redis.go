// database/redis.go
package database

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/redis/rueidis"
)

var (
	IconGetError    = errors.New("unable to icon error")
	WeatherGetError = errors.New("unable to weather error")
)

type RedisClient struct {
	Client rueidis.Client
	Ctx    context.Context
}

type CloseFunc func()

func NewRedisClient(address string) (RedisClient, CloseFunc, error) {
	client, err := rueidis.NewClient(rueidis.ClientOption{InitAddress: []string{address}})
	if err != nil {
		return RedisClient{}, nil, fmt.Errorf("failed to create client: %w", err)
	}

	ctx := context.Background()

	return RedisClient{Client: client, Ctx: ctx}, func() {
		client.Close()
	}, nil
}

func (r RedisClient) SetWeather(city, weather, icon string) error {
	// Set the weather and icon with an expiry of 1 hour
	err := r.Client.Do(r.Ctx, r.Client.B().Set().Key(city+"weather").Value(weather).Ex(time.Hour).Build()).Error()
	if err != nil {
		return fmt.Errorf("failed to set weather: %w", err)
	}

	err = r.Client.Do(r.Ctx, r.Client.B().Set().Key(city+"weathericon").Value(icon).Ex(time.Hour).Build()).Error()
	if err != nil {
		return fmt.Errorf("failed to set weather icon: %w", err)
	}

	return nil
}

func (r RedisClient) GetWeather(city string) ([]byte, []byte, error) {
	// Get the weather
	weather, err := r.Client.Do(r.Ctx, r.Client.B().Get().Key(city+"weather").Build()).AsBytes()
	if err != nil {
		return nil, nil, errors.Join(WeatherGetError, err)
	}
	
	icon, err := r.Client.Do(r.Ctx, r.Client.B().Get().Key(city+"weathericon").Build()).AsBytes()
	if err != nil {
		return nil, nil, errors.Join(IconGetError, err)
	}

	return weather, icon, nil
}
