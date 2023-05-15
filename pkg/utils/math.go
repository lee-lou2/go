package utils

import (
	"math"
)

// Sum 리스트 내 값의 합
func Sum(data []float64) float64 {
	sum := 0.0
	for _, v := range data {
		sum += v
	}
	return sum
}

// PredictNext 다음 예측 값 추출
func PredictNext(data []float64) float64 {
	x := make([]float64, len(data))
	for i := range x {
		x[i] = float64(i)
	}
	y := data

	// x와 y의 평균값 산출
	xMean := Sum(x) / float64(len(x))
	yMean := Sum(y) / float64(len(y))

	// x와 y의 분산과 공분산 산출
	varianceX := 0.0
	for _, v := range x {
		a := v - xMean
		varianceX += math.Pow(a, 2)
	}
	varianceY := 0.0
	for _, v := range y {
		a := v - yMean
		varianceY += math.Pow(a, 2)
	}
	covarianceXY := 0.0
	for i := 0; i < len(x); i++ {
		covarianceXY += (x[i] - xMean) * (y[i] - yMean)
	}

	// 회귀계수와 절편을 계산
	b := covarianceXY / varianceX
	a := yMean - b*xMean
	nextX := len(data)
	return b*float64(nextX) + a
}
