package node

import (
	"sort"
)

const TopVideoThreshold = 40

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

type Edge struct {
	Neighbour *CacheServerNode
	Distance  int
}

var backupTable map[string]CacheServerNode

type CacheServerNode struct {
	// init
	name        string
	capacity    float64
	videos      VideoList
	upLevelNode *Edge
	top20       VideoList // заполнить

	// calculate
	cachedVideos         VideoList
	popularityPercentage int
	available            float64
}

func NewCacheServerNode(name string, cap float64, videos VideoList) *CacheServerNode {
	node := &CacheServerNode{
		name:      name,
		capacity:  cap,
		available: cap,
		videos:    videos,
	}
	backupTable[name] = *node
	return node
}

func GetTableBackup() map[string]CacheServerNode {
	return backupTable
}

func InitBackupTable() {
	backupTable = make(map[string]CacheServerNode)
}

func (csn *CacheServerNode) GetAvailable() float64 {
	return csn.available
}

func (csn *CacheServerNode) GetVideoLists() VideoList {
	return csn.videos
}

func (csn *CacheServerNode) SetVideoLists(value VideoList) {
	csn.videos = value
}

func (csn *CacheServerNode) GetTop20() VideoList {
	return csn.top20
}

func (csn *CacheServerNode) SetTop20(value VideoList) {
	if value != nil {
		csn.top20 = value
	} else {
		csn.top20 = csn.videos[:TopVideoThreshold]
	}
}

func BackUp(node *CacheServerNode) {
	temp := backupTable[node.name]
	node = &temp
}

func (csn *CacheServerNode) GetUpLevel() *Edge {
	return csn.upLevelNode
}

func (csn *CacheServerNode) SetUpLevel(e *Edge) {
	if e.Neighbour != nil {
		if csn.upLevelNode == nil {
			csn.upLevelNode = &Edge{
				Neighbour: e.Neighbour,
				Distance:  e.Distance,
			}
		} else {
			csn.upLevelNode.Neighbour = e.Neighbour
			csn.upLevelNode.Distance = e.Distance
		}
	}
}

func (csn *CacheServerNode) GetCapacity() float64 {
	return csn.capacity
}

type CachedVideoObject struct {
	Id    string  `json:"video_id"`
	Title string  `json:"title"`
	Size  float64 `json:"size"`
}

type CachedList []*CachedVideoObject

func (csn *CacheServerNode) GetCachedVideos() CachedList {
	var result []*CachedVideoObject
	for _, video := range csn.cachedVideos {
		object := &CachedVideoObject{
			Id:    video.Id,
			Title: video.Title,
			Size:  video.Size,
		}
		result = append(result, object)
	}
	return result
}

func (csn *CacheServerNode) SetCachedVideos(cached VideoList) {
	cachedVideoCnt := 0

	for _, video := range cached {
		if csn.available-video.Size >= 0 {
			cachedVideoCnt += 1
			//csn.cachedVideos = append(csn.cachedVideos, video)
			csn.available -= video.Size
		}
	}

	csn.cachedVideos = cached[:cachedVideoCnt]

	neighbour := csn.GetUpLevel()
	if neighbour != nil {
		if neighbour.Neighbour.name != "DE" {
			neighbour.Neighbour.SetCachedVideos(cached[cachedVideoCnt:])
		}
	}
}

func (csn *CacheServerNode) GetRegionName() string {
	return csn.name
}

func (csn *CacheServerNode) SortVideoListByFeature(feature int, weights []float64) {
	// добавить остальные 4 метода

	switch feature {
	case Views:
		sort.Slice(csn.videos, func(i, j int) bool {
			return csn.videos[i].ViewsLastDay > csn.videos[j].ViewsLastDay
		})
	case ViewsLikes:
		sort.Slice(csn.videos, func(i, j int) bool {
			left, right := csn.videos[i], csn.videos[j]
			left.setViewsLikesRatio(weights)
			right.setViewsLikesRatio(weights)
			return csn.videos[i].ViewsLikesRatio > csn.videos[j].ViewsLikesRatio
		})
	case ViewsLikesComments:
		sort.Slice(csn.videos, func(i, j int) bool {
			left, right := csn.videos[i], csn.videos[j]
			left.setViewsLikesCommRatio(weights)
			right.setViewsLikesCommRatio(weights)
			return csn.videos[i].ViewsLikesCommRatio > csn.videos[j].ViewsLikesCommRatio
		})
	case ViewsLikesCommentsCategory:
		sort.Slice(csn.videos, func(i, j int) bool {
			left, right := csn.videos[i], csn.videos[j]
			left.setViewsLikesCommCatRatio(weights)
			right.setViewsLikesCommCatRatio(weights)
			return csn.videos[i].ViewsLikesCommCatRatio > csn.videos[j].ViewsLikesCommCatRatio
		})
	case ViewsComments:
		sort.Slice(csn.videos, func(i, j int) bool {
			left, right := csn.videos[i], csn.videos[j]
			left.setViewsCommRatio(weights)
			right.setViewsCommRatio(weights)
			return csn.videos[i].ViewsCommRatio > csn.videos[j].ViewsCommRatio
		})
	case ViewsCommentsCategory:
		sort.Slice(csn.videos, func(i, j int) bool {
			left, right := csn.videos[i], csn.videos[j]
			left.setViewsCommCatRatio(weights)
			right.setViewsCommCatRatio(weights)
			return csn.videos[i].ViewsCommCatRatio > csn.videos[j].ViewsCommCatRatio
		})
	case ViewsLikesCategory:
		sort.Slice(csn.videos, func(i, j int) bool {
			left, right := csn.videos[i], csn.videos[j]
			left.setViewsLikesCatRatio(weights)
			right.setViewsLikesCatRatio(weights)
			return csn.videos[i].ViewsLikesCatRatio > csn.videos[j].ViewsLikesCatRatio
		})
	case ViewsCategory:
		sort.Slice(csn.videos, func(i, j int) bool {
			left, right := csn.videos[i], csn.videos[j]
			left.setViewsCatRatio(weights)
			right.setViewsCatRatio(weights)
			return csn.videos[i].ViewsCatRatio > csn.videos[j].ViewsCatRatio
		})
	case ViewsWMA:
		sort.Slice(csn.videos, func(i, j int) bool {
			return csn.videos[i].ViewsWMA > csn.videos[j].ViewsWMA
		})
	case ViewsLikesWMA:
		sort.Slice(csn.videos, func(i, j int) bool {
			left, right := csn.videos[i], csn.videos[j]
			left.setViewsLikesWMA(weights)
			right.setViewsLikesWMA(weights)
			return csn.videos[i].ViewsLikesRatio > csn.videos[j].ViewsLikesRatio
		})
	case ViewsLikesCommentsWMA:
		sort.Slice(csn.videos, func(i, j int) bool {
			left, right := csn.videos[i], csn.videos[j]
			left.setViewsLikesCommWMA(weights)
			right.setViewsLikesCommWMA(weights)
			return csn.videos[i].ViewsLikesCommRatio > csn.videos[j].ViewsLikesCommRatio
		})
	case ViewsLikesCommentsWMACategory:
		sort.Slice(csn.videos, func(i, j int) bool {
			left, right := csn.videos[i], csn.videos[j]
			left.setViewsLikesCommCatWMA(weights)
			right.setViewsLikesCommCatWMA(weights)
			return csn.videos[i].ViewsLikesCommCatRatio > csn.videos[j].ViewsLikesCommCatRatio
		})
	case ViewsCommentsWMA:
		sort.Slice(csn.videos, func(i, j int) bool {
			left, right := csn.videos[i], csn.videos[j]
			left.setViewsCommWMA(weights)
			right.setViewsCommWMA(weights)
			return csn.videos[i].ViewsCommRatio > csn.videos[j].ViewsCommRatio
		})
	case ViewsCommentsWMACategory:
		sort.Slice(csn.videos, func(i, j int) bool {
			left, right := csn.videos[i], csn.videos[j]
			left.setViewsCommCatWMA(weights)
			right.setViewsCommCatWMA(weights)
			return csn.videos[i].ViewsCommCatRatio > csn.videos[j].ViewsCommCatRatio
		})
	case ViewsLikesWMACategory:
		sort.Slice(csn.videos, func(i, j int) bool {
			left, right := csn.videos[i], csn.videos[j]
			left.setViewsLikesCatWMA(weights)
			right.setViewsLikesCatWMA(weights)
			return csn.videos[i].ViewsLikesCatRatio > csn.videos[j].ViewsLikesCatRatio
		})
	case ViewsWMACategory:
		sort.Slice(csn.videos, func(i, j int) bool {
			left, right := csn.videos[i], csn.videos[j]
			left.setViewsCatWMA(weights)
			right.setViewsCatWMA(weights)
			return csn.videos[i].ViewsCatRatio > csn.videos[j].ViewsCatRatio
		})
	}
}
