package methods

import (
	server "diploma/node"
	"encoding/json"
	"io/ioutil"
	"log"
)

type PlacementMethod struct {
	Name        string
	Description string
	Method      func(*server.CacheServerNode, []float64) *server.CacheServerNode
}

func NewPlacementOpts() []PlacementMethod {
	return []PlacementMethod{
		{
			Name:        "ViewsMethod",
			Description: "",
			Method:      firstMethod,
		},
		{
			Name:        "ViewsLikeMethod",
			Description: "",
			Method:      secondMethod,
		},
		{
			Name:        "ViewsLikesCommentsMethod",
			Description: "",
			Method:      thirdMethod,
		},
		{
			Name:        "ViewsLikesCommentsCategoryMethod",
			Description: "",
			Method:      fourthMethod,
		},
		{
			Name:        "ViewsComments",
			Description: "",
			Method:      fifthMethod,
		},
		{
			Name:        "ViewsCommentsCategory",
			Description: "",
			Method:      sixthMethod,
		},
		{
			Name:        "ViewsLikesCategory",
			Description: "",
			Method:      seventhMethod,
		},
		{
			Name:        "ViewsCategory",
			Description: "",
			Method:      eighthMethod,
		},
		{
			Name:        "ViewsWMA",
			Description: "",
			Method:      ninthMethod,
		},
		{
			Name:        "ViewsLikesWMA",
			Description: "",
			Method:      tenthMethod,
		},
		{
			Name:        "ViewsLikesCommentsWMA",
			Description: "",
			Method:      eleventhMethod,
		},
		{
			Name:        "ViewsLikesCommentsWMACategory",
			Description: "",
			Method:      twelfthMethod,
		},
		{
			Name:        "ViewsCommentsWMA",
			Description: "",
			Method:      thirteenthMethod,
		},
		{
			Name:        "ViewsCommentsWMACategory",
			Description: "",
			Method:      fourteenthMethod,
		},
		{
			Name:        "ViewsLikesWMACategory",
			Description: "",
			Method:      fifteenthMethod,
		},
		{
			Name:        "ViewsWMACategory",
			Description: "",
			Method:      sixteenthMethod,
		},
	}
}

func baseMethod(node *server.CacheServerNode) *server.CacheServerNode {
	if isCentralServer(node) {
		return node
	}

	howVideos := canServerContain(node)

	node.SetTop20(node.GetVideoLists()[:howVideos])

	if howVideos >= server.TopVideoThreshold {
		node.SetCachedVideos(node.GetTop20())
		return node
	}

	var (
		top      = node.GetTop20()
		afterTop = node.GetVideoLists()[howVideos:]
		steps    = howVideos
		cached   server.VideoList
		unloaded server.VideoList
	)

	i := 0
	available := node.GetAvailable()
	for i < steps {
		size := swapBatchSize(top[i].Size+available, afterTop)

		swap := false

		if size > 1 {
			headDistance := calcReloadDistance(*node, &top[i])
			afterDistance := calcAfterTopDistance(node, afterTop[:size])

			if headDistance < afterDistance {
				swap = true
				unloaded = append(unloaded, top[i])
				cached = append(cached, afterTop[:size]...)

				available += top[i].Size
				available -= calcBatchSize(cached)

				afterTop = afterTop[size:]
			}
		}

		if !swap {
			cached = append(cached, top[i])
			available -= top[i].Size
		}

		i += 1
	}

	node.SetCachedVideos(cached)
	neighbour := node.GetUpLevel().Neighbour
	neighbour.SetCachedVideos(unloaded)

	return node
}

const OutputPath = "./setup/output/cached.json"

func ExportVideoPlacement(object map[string]server.CachedList) {
	jsonObj, _ := json.Marshal(object)
	err := ioutil.WriteFile(OutputPath, jsonObj, 0644)
	if err != nil {
		log.Fatal(err)
	}
}

func calcBatchSize(cached server.VideoList) float64 {
	var sum float64 = 0
	for _, video := range cached {
		sum += video.Size
	}
	return sum
}
