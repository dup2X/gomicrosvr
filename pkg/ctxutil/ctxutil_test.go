package ctxutil

import (
	"context"
	"math/rand"
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestErrNotExist(t *testing.T) {
	ctx := context.Background()
	var err error
	ass := assert.New(t)
	for _, key := range builtInKeys {
		_, err = getString(ctx, key)
		ass.EqualError(err, ErrNotExist.Error())
		_, err = getInt64(ctx, key)
		ass.EqualError(err, ErrNotExist.Error())
		_, err = getTimeDuration(ctx, key)
		ass.EqualError(err, ErrNotExist.Error())
	}
}

func TestErrDurationNotSetted(t *testing.T) {
	ctx := context.Background()
	var err error
	ass := assert.New(t)
	_, err = incTimeDuration(ctx, dbTimeKey, time.Second)
	ass.EqualError(err, ErrDurationNotSetted.Error())
	_, err = incTimeDuration(ctx, redisTimeKey, time.Hour)
	ass.EqualError(err, ErrDurationNotSetted.Error())
}

func TestTraceID(t *testing.T) {
	ass := assert.New(t)
	for i := 0; i < 1000; i++ {
		s := randString(32)
		ctx := SetTraceID(context.Background(), s)
		r, err := GetTraceID(ctx)
		ass.NoError(err)
		ass.Equal(s, r)
	}
}

var letterRunes = []byte("~!@#$%^&*()_+0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func randString(n int) string {
	length := len(letterRunes)
	s := make([]byte, length)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < length; i++ {
		s[i] = letterRunes[rand.Intn(length)]
	}
	return string(s)
}

func TestSetPassengerID(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	ass := assert.New(t)
	for i := 0; i <= 1000; i++ {
		pid := rand.Int63()
		ctx := SetPassengerID(context.Background(), pid)
		r, err := GetPassengerID(ctx)
		ass.NoError(err)
		ass.Equal(pid, r)
	}
}

func TestAccumulation(t *testing.T) {
	ctx := context.Background()
	ass := assert.New(t)
	_, err := IncDBTime(ctx, time.Second)
	ass.EqualError(err, ErrDurationNotSetted.Error())
	ctx = SetDBTime(ctx)
	var wg sync.WaitGroup
	wg.Add(5)
	for i := int64(1); i <= 5; i++ {
		go func(s int64) {
			defer wg.Done()
			_, err = IncDBTime(ctx, time.Duration(s*time.Second.Nanoseconds()))
			ass.NoError(err)
		}(i)
	}
	wg.Wait()
	r, err := GetDBTime(ctx)
	ass.NoError(err)
	ass.Equal(15*time.Second, r)
}

func TestFormat(t *testing.T) {
	var err error
	ass := assert.New(t)
	ctx := context.Background()
	ctx = SetTraceID(ctx, keyName[traceIDKey])
	ctx = SetSpanID(ctx, keyName[spanIDKey])
	ctx = SetDBTime(ctx)
	ctx, err = IncDBTime(ctx, 5*time.Second)
	ass.NoError(err)
	ctx = SetRedisTime(ctx)
	ctx, err = IncRedisTime(ctx, 2*time.Minute)
	ass.NoError(err)
	var driverID int64 = 123456
	var passengerID int64 = 12345678
	var orderID int64 = 123456789
	ctx = SetDriverID(ctx, driverID)
	ctx = SetPassengerID(ctx, passengerID)
	ctx = SetOrderID(ctx, orderID)
	ctx = SetHintCode(ctx, 23456)
	ctx = SetHintContent(ctx, keyName[hintContentKey])
	ctx = SetDistrict(ctx, "010")
	s, err := Format(ctx)
	ass.NoError(err)
	expected := `traceID=traceID||spanID=spanID||redisTime=2m0s||dbTime=5s||passengerID=12345678||driverID=123456||orderID=123456789||district=010||hintCode=23456||hintContent=hintContent`
	ass.Equal(expected, s)
}
