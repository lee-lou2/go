package redis

import (
	"fmt"
	"time"
)

// Pub Redis Publish
func (c *NewClient) Pub(channel string, message interface{}) error {
	err := c.Publish(channel, message).Err()
	if err != nil {
		return err
	}

	return nil
}

// Sub Redis Subscribe
func (c *NewClient) Sub(channel string, message chan<- string) error {
	sub := c.Subscribe(channel)

	defer sub.Close()

	for {
		msg, err := sub.ReceiveMessage()
		if err != nil {
			return err
		}

		message <- msg.Payload
	}
}

// Produce Redis Producer
func (c *NewClient) Produce(key string, value interface{}) error {
	_, err := c.LPush(key, value).Result()
	if err != nil {
		return err
	}

	return nil
}

// Consume Redis Consumer
func (c *NewClient) Consume(key string) (string, error) {
	result, err := c.BRPop(0, key).Result()
	if err != nil {
		return "", err
	}

	return result[1], nil
}

// GetValue Redis Get Value
func (c *NewClient) GetValue(key string) (string, error) {
	return c.Get(key).Result()
}

// SetValue Redis Set Value
func (c *NewClient) SetValue(key string, value string, expiration int) error {
	if err := c.Set(key, value, time.Duration(expiration)*time.Second).Err(); err != nil {
		return fmt.Errorf("redis set value error: %w", err)
	}
	return nil
}

// PushValue 캐시 값 추가
func (c *NewClient) PushValue(key string, value string) error {
	if err := c.LPush(key, value).Err(); err != nil {
		return err
	}
	return nil
}

// PopValue 캐시 값 제거
func (c *NewClient) PopValue(key string) (string, error) {
	value, err := c.RPop(key).Result()
	if err != nil {
		return "", err
	}
	return value, nil
}

// GetRange 값 리스트 조회
func (c *NewClient) GetRange(key string) ([]string, error) {
	values, err := c.LRange(key, 0, -1).Result()
	if err != nil {
		return nil, err
	}
	return values, nil
}

// GetLength 전체 카운트 조회
func (c *NewClient) GetLength(key string) (int, error) {
	count, err := c.LLen(key).Result()
	if err != nil {
		return 0, err
	}
	return int(count), nil
}

// PushValueSetLimit 최대 제한을 두고 추가
func (c *NewClient) PushValueSetLimit(key string, value string, limit int) error {
	var err error

	// 데이터 수 조회
	count, err := c.LLen(key).Result()
	if err != nil {
		return err
	}

	// 최대 제한 수를 초과하는 경우 초과되는 값들은 제거
	if int(count) >= limit {
		for i := 0; i <= int(count)-limit; i++ {
			c.RPop(key)
		}
	}
	return c.LPush(key, value).Err()
}

// ExistsValue 리스트 내 해당 데이터가 있는지 확인
func (c *NewClient) ExistsValue(key string, value string) (bool, error) {
	// 리스트 조회
	values, err := c.LRange(key, 0, -1).Result()
	if err != nil {
		return false, err
	}

	// 리스트내 존재 여부 확인
	for _, v := range values {
		if v == value {
			return true, nil
		}
	}
	return false, nil
}
