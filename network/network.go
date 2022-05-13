package network

import (
	mtd "diploma/methods"
	server "diploma/node"
)

type Channel struct {
	Weight int    `json:"weight"`
	Name   string `json:"destination"`
}

type CDNModel struct {
	nodes         []*server.CacheServerNode
	upAdjacency   map[*server.CacheServerNode]*server.Edge
	methodApplied bool
}

func NewCDNModel(
	servers []*server.CacheServerNode,
	adjacency map[*server.CacheServerNode]*server.Edge) *CDNModel {
	return &CDNModel{
		nodes:       servers,
		upAdjacency: adjacency,
	}
}

func (model *CDNModel) GetNodes() []*server.CacheServerNode {
	return model.nodes
}

func (model *CDNModel) ApplyMethod(pm mtd.PlacementMethod, weights []float64) {
	for _, node := range model.nodes {
		node = pm.Method(node, weights)
	}
	// model.methodApplied = true
}

func Clear(backup []*server.CacheServerNode) []*server.CacheServerNode {
	res := make([]*server.CacheServerNode, 0, len(backup))

	for _, s := range backup {
		node := new(server.CacheServerNode)
		*node = *s
		res = append(res, node)
	}

	return res
}

func (model *CDNModel) GetPopularityPercentage(videos server.VideoList, cached server.CachedList) (float64, error) {
	//if !model.methodApplied {
	//	return 0, fmt.Errorf("any method doesn't apply for the model")
	//}
	lenCached := len(cached)
	lenVideos := len(videos)
	return float64(lenCached) / float64(lenVideos), nil
}
