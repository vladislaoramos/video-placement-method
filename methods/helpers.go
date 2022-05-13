package methods

import (
	server "diploma/node"
)

func hasServerContent(node *server.CacheServerNode, video *server.VideoObject) bool {
	list := node.GetVideoLists()
	for _, item := range list {
		if item.Id == video.Id {
			return true
		}
	}
	return false
}

func calcReloadDistance(node server.CacheServerNode, video *server.VideoObject) int {
	res := 0
	temp := node
	for temp.GetUpLevel() != nil {
		neighbour := temp.GetUpLevel()
		res += neighbour.Distance
		if hasServerContent(neighbour.Neighbour, video) {
			return res
		}
		temp = *neighbour.Neighbour
	}
	return res
}

func calcAfterTopDistance(node *server.CacheServerNode, videos server.VideoList) int {
	res := 0
	for _, video := range videos {
		res += calcReloadDistance(*node, &video)
	}
	return res
}

func canServerContain(node *server.CacheServerNode) int {
	count := 0
	list := node.GetVideoLists()
	available := node.GetCapacity()

	for _, video := range list {
		if available-video.Size >= 0 {
			count += 1
			available -= video.Size
		}
	}

	return count
}

func swapBatchSize(available float64, afterTop server.VideoList) int {
	// available = top_video_size + free_space_in_server
	size := 0
	memory := available
	for _, video := range afterTop {
		memory -= video.Size
		if memory >= 0 {
			size += 1
		} else {
			return size
		}
	}

	return size
}

func isCentralServer(node *server.CacheServerNode) bool {
	return node.GetRegionName() == "DE"
}

func cachedVideoPreCS(node *server.CacheServerNode) server.VideoList {
	var result server.VideoList
	available := node.GetAvailable()

	for _, video := range node.GetVideoLists() {
		if available > 0 {
			available -= video.Size
			result = append(result, video)
		}
	}

	return result
}
