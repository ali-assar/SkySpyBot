// database/redis.go
package database

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/rueidis"
)

type RedisClient struct {
	Client rueidis.Client
	Ctx    context.Context
}

func NewRedisClient(address string) (RedisClient, error) {
	client, err := rueidis.NewClient(rueidis.ClientOption{InitAddress: []string{address}})
	if err != nil {
		return RedisClient{}, fmt.Errorf("failed to create client: %w", err)
	}

	ctx := context.Background()

	return RedisClient{Client: client, Ctx: ctx}, nil
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

func (r RedisClient) GetWeather(city string) (map[string]string, map[string]string, error) {
	// Get the weather
	weather, err := r.Client.Do(r.Ctx, r.Client.B().Get().Key(city+"weather").Build()).AsStrMap()
	if err != nil {
		return nil, nil, fmt.Errorf("failed to get weather: %w", err)
	}

	// Get the icon
	icon, err := r.Client.Do(r.Ctx, r.Client.B().Get().Key(city+"weathericon").Build()).AsStrMap()
	if err != nil {
		return nil, nil, fmt.Errorf("failed to get weather icon: %w", err)
	}

	return weather, icon, nil
}
