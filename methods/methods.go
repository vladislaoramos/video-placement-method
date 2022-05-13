package methods

import (
	server "diploma/node"
)

const (
	Views                         = 0
	ViewsLikes                    = 1
	ViewsLikesComments            = 2
	ViewsLikesCommentsCategory    = 3
	ViewsComments                 = 4
	ViewsCommentsCategory         = 5
	ViewsLikesCategory            = 6
	ViewsCategory                 = 7
	ViewsWMA                      = 8
	ViewsLikesWMA                 = 9
	ViewsLikesCommentsWMA         = 10
	ViewsLikesCommentsWMACategory = 11
	ViewsCommentsWMA              = 12
	ViewsCommentsWMACategory      = 13
	ViewsLikesWMACategory         = 14
	ViewsWMACategory              = 15
)

func firstMethod(node *server.CacheServerNode, weights []float64) *server.CacheServerNode {
	node.SortVideoListByFeature(Views, weights)
	node.SetTop20(nil)
	return baseMethod(node)
}

func secondMethod(node *server.CacheServerNode, weights []float64) *server.CacheServerNode {
	node.SortVideoListByFeature(ViewsLikes, weights)
	node.SetTop20(nil)
	return baseMethod(node)
}

func thirdMethod(node *server.CacheServerNode, weights []float64) *server.CacheServerNode {
	node.SortVideoListByFeature(ViewsLikesComments, weights)
	node.SetTop20(nil)
	return baseMethod(node)
}

func fourthMethod(node *server.CacheServerNode, weights []float64) *server.CacheServerNode {
	node.SortVideoListByFeature(ViewsLikesCommentsCategory, weights)
	node.SetTop20(nil)
	return baseMethod(node)
}

func fifthMethod(node *server.CacheServerNode, weights []float64) *server.CacheServerNode {
	node.SortVideoListByFeature(ViewsComments, weights)
	node.SetTop20(nil)
	return baseMethod(node)
}

func sixthMethod(node *server.CacheServerNode, weights []float64) *server.CacheServerNode {
	node.SortVideoListByFeature(ViewsCommentsCategory, weights)
	node.SetTop20(nil)
	return baseMethod(node)
}

func seventhMethod(node *server.CacheServerNode, weights []float64) *server.CacheServerNode {
	node.SortVideoListByFeature(ViewsLikesCategory, weights)
	node.SetTop20(nil)
	return baseMethod(node)
}

func eighthMethod(node *server.CacheServerNode, weights []float64) *server.CacheServerNode {
	node.SortVideoListByFeature(ViewsCategory, weights)
	node.SetTop20(nil)
	return baseMethod(node)
}

func ninthMethod(node *server.CacheServerNode, weights []float64) *server.CacheServerNode {
	node.SortVideoListByFeature(ViewsWMA, weights)
	node.SetTop20(nil)
	return baseMethod(node)
}

func tenthMethod(node *server.CacheServerNode, weights []float64) *server.CacheServerNode {
	node.SortVideoListByFeature(ViewsLikesWMA, weights)
	node.SetTop20(nil)
	return baseMethod(node)
}

func eleventhMethod(node *server.CacheServerNode, weights []float64) *server.CacheServerNode {
	node.SortVideoListByFeature(ViewsLikesCommentsWMA, weights)
	node.SetTop20(nil)
	return baseMethod(node)
}

func twelfthMethod(node *server.CacheServerNode, weights []float64) *server.CacheServerNode {
	node.SortVideoListByFeature(ViewsLikesCommentsWMACategory, weights)
	node.SetTop20(nil)
	return baseMethod(node)
}

func thirteenthMethod(node *server.CacheServerNode, weights []float64) *server.CacheServerNode {
	node.SortVideoListByFeature(ViewsCommentsWMA, weights)
	node.SetTop20(nil)
	return baseMethod(node)
}

func fourteenthMethod(node *server.CacheServerNode, weights []float64) *server.CacheServerNode {
	node.SortVideoListByFeature(ViewsCommentsWMACategory, weights)
	node.SetTop20(nil)
	return baseMethod(node)
}

func fifteenthMethod(node *server.CacheServerNode, weights []float64) *server.CacheServerNode {
	node.SortVideoListByFeature(ViewsLikesWMACategory, weights)
	node.SetTop20(nil)
	return baseMethod(node)
}

func sixteenthMethod(node *server.CacheServerNode, weights []float64) *server.CacheServerNode {
	node.SortVideoListByFeature(ViewsWMACategory, weights)
	node.SetTop20(nil)
	return baseMethod(node)
}
