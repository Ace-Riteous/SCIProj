package utils

import (
	"math/rand"
	"time"
)

func GenerateRandomString(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyz" +
		"ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	seed := rand.NewSource(time.Now().UnixNano())
	generator := rand.New(seed)

	b := make([]byte, length)
	for i := range b {
		b[i] = charset[generator.Intn(len(charset))]
	}
	return string(b)
}

func GenerateCpuUsage(start int64, usage float64) float64 {
	startTime := time.Unix(start, 0)
	if time.Since(startTime)%time.Hour == 0 {
		return rand.Float64()*100 + 1
	}
	return usage
}
