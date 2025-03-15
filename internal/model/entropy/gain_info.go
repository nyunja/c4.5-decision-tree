package entropy

import (
	"math"

	"github.com/nyunja/c4.5-decision-tree/internal/model/counter"
	t "github.com/nyunja/c4.5-decision-tree/internal/model/types"
)

func GainInfoAndSplitInfo(counter *counter.ClassCounter, instances []t.Instance, infoGain float64, splitInfo float64) (float64, float64) {
	prob := float64(counter.Total) / float64(len(instances))
	if prob > 0 {
		infoGain -= prob * counter.GetEntropy()
		splitInfo -= prob * math.Log2(prob)
	}
	return infoGain, splitInfo
}

func GainRatio(leftProb float64, rightProb float64, infoGain float64) float64 {
	splitInfo := 0.0
	if leftProb > 0 {
		splitInfo -= leftProb * math.Log2(leftProb)
	}
	if rightProb > 0 {
		splitInfo -= rightProb * math.Log2(rightProb)
	}

	// Calculate gain ratio
	gainRatio := 0.0
	if splitInfo > 0 {
		gainRatio = infoGain / splitInfo
	}
	return gainRatio
}
