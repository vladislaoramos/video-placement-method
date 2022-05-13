package quality

import (
	mtd "diploma/methods"
	nw "diploma/network"
)

type Estimator struct {
	model  *nw.CDNModel
	method mtd.PlacementMethod
}

func NewEstimator(cdn *nw.CDNModel, pm mtd.PlacementMethod) *Estimator {
	return &Estimator{
		model:  cdn,
		method: pm,
	}
}

func calcAvg(values []float64) float64 {
	var sum float64 = 0
	for _, v := range values[:len(values)-1] {
		sum += v
	}
	return sum / float64(len(values)-1)
}

func (e Estimator) Estimate(weights []float64) (float64, float64) {
	e.model.ApplyMethod(e.method, weights)

	var (
		freeCapacity []float64
		popularity   []float64
	)

	nodes := e.model.GetNodes()

	for _, node := range nodes {
		freeCapacity = append(freeCapacity, node.GetAvailable()/node.GetCapacity()) // минимизируем
		top := node.GetTop20()
		cached := node.GetCachedVideos()
		percentage, _ := e.model.GetPopularityPercentage(top, cached)
		popularity = append(popularity, percentage) // максимизируем
	}

	freeCap := calcAvg(freeCapacity)
	pop := calcAvg(popularity)

	return freeCap, pop
}
