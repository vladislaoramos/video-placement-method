package main

import (
	mtd "diploma/methods"
	nw "diploma/network"
	server "diploma/node"
	"diploma/quality"
	config "diploma/setup"
	"fmt"
	"sort"
	"time"
)

func ApplyBestMethod(setup *config.Setup, method mtd.PlacementMethod, weights []float64) map[string]server.CachedList {
	servers, adjacency := setup.GetSetup()
	network := nw.NewCDNModel(servers, adjacency)
	network.ApplyMethod(method, weights)

	output := make(map[string]server.CachedList)

	for _, s := range servers {
		output[s.GetRegionName()] = s.GetCachedVideos()
	}

	return output
}

type MethodObject struct {
	Name                 string
	FreeCapacity         float64
	PopularityPercentage float64
	Weights              []float64
	Method               mtd.PlacementMethod
}

type MethodsTable []*MethodObject

func calcWeights() [][]float64 {
	var res [][]float64

	for views := 1; views < 10; views++ {
		var item []float64
		for likes := 0; likes < 11-views; likes++ {
			for comments := 0; comments < 11-views-likes; comments++ {
				for categories := 0; categories < 11-views-likes-comments; categories++ {
					if views+likes+comments+categories == 10 {
						item = []float64{
							float64(views) / 10,
							float64(likes) / 10,
							float64(comments) / 10,
							float64(categories) / 10,
						}
						res = append(res, item)
					}
				}
			}
		}
	}

	return res
}

func MethodsComparison(
	methods []mtd.PlacementMethod,
	setup *config.Setup) MethodsTable {
	servers, adjacency := setup.GetSetup()

	table := make(MethodsTable, 0, len(methods))
	weights := calcWeights()
	backup := nw.Clear(servers)

	for _, method := range methods {
		for _, w := range weights {
			network := nw.NewCDNModel(servers, adjacency)
			estimator := quality.NewEstimator(network, method)
			free, popularity := estimator.Estimate(w)

			//fmt.Printf("Free Capacity after applying method %s: %v\n", methodName, free)
			//fmt.Printf("Popularity percentage after applying method %s: %v\n\n", methodName, popularity)

			m := mtd.PlacementMethod{
				Name:   method.Name,
				Method: method.Method,
			}

			methodObject := &MethodObject{
				Name:                 method.Name,
				FreeCapacity:         free,
				PopularityPercentage: popularity,
				Weights:              w,
				Method:               m,
			}

			table = append(table, methodObject)

			servers = nw.Clear(backup)
		}
	}

	sort.Slice(table, func(i, j int) bool {
		return table[i].FreeCapacity < table[j].FreeCapacity &&
			table[i].PopularityPercentage > table[j].PopularityPercentage
	})

	return table
}

func displayBestWorstMethods(table MethodsTable) {
	fmt.Printf(
		"Best method is %s\nFree Capacity = %v\nPopularity percentage = %v\nWeights = %v\n",
		table[0].Name,
		table[0].FreeCapacity,
		table[0].PopularityPercentage,
		table[0].Weights)

	fmt.Printf(
		"Worst method is %s\nFree Capacity = %v\nPopularity percentage = %v\nWeights = %v\n",
		table[len(table)-1].Name,
		table[len(table)-1].FreeCapacity,
		table[len(table)-1].PopularityPercentage,
		table[len(table)-1].Weights)
}

func clearTable(table map[string]*server.CacheServerNode) {
	backup := server.GetTableBackup()
	for region, _ := range backup {
		item := backup[region]
		table[region] = &item
	}
}

func main() {
	start := time.Now()

	table := make(map[string]*server.CacheServerNode)
	upAdjacency := make(map[*server.CacheServerNode]*server.Edge)

	server.InitBackupTable()

	setup := config.NewSetup(table, upAdjacency)

	placementMethods := mtd.NewPlacementOpts()
	methodsSet := MethodsComparison(placementMethods, setup)
	displayBestWorstMethods(methodsSet)
	best, _ := methodsSet[0], methodsSet[len(methodsSet)-1]

	clearTable(table)
	clearSetup := config.NewSetup(table, upAdjacency)
	list := ApplyBestMethod(clearSetup, best.Method, best.Weights)
	mtd.ExportVideoPlacement(list)

	finish := time.Now()
	diff := finish.Sub(start)

	fmt.Printf("Time to calculate: %v min %v sec", int(diff.Minutes()), int(diff.Seconds())%60)

}
