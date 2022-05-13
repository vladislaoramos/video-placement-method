package setup

import (
	nw "diploma/network"
	server "diploma/node"
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

type Setup struct {
	nodes     []*server.CacheServerNode
	adjacency map[*server.CacheServerNode]*server.Edge
}

func NewSetup(table map[string]*server.CacheServerNode, upAdjacency map[*server.CacheServerNode]*server.Edge) *Setup {
	return &Setup{
		nodes:     initNodes(table),
		adjacency: joinServers(RegionsAdjacencyPath, table, upAdjacency),
	}
}

func (s *Setup) GetSetup() ([]*server.CacheServerNode, map[*server.CacheServerNode]*server.Edge) {
	return s.nodes, s.adjacency
}

func getVideosList(path string) server.VideoList {
	var list server.VideoList

	jsonFile, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}

	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal(byteValue, &list)
	if err != nil {
		log.Fatal(err)
	}

	return list
}

func getNodesCapacity(path string) map[string]float64 {
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}
	var capacity map[string]float64
	err = json.Unmarshal(bytes, &capacity)
	if err != nil {
		log.Fatal(err)
	}
	return capacity
}

const (
	VideoSetPath         = "./setup/videoset/"
	NodesStoragePath     = "./setup/nodes.json"
	RegionsAdjacencyPath = "./setup/adjacency.json"
	MainPath             = "./setup/"
)

func getAdjacency(path string) map[string]*nw.Channel {
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}

	var data map[string]*nw.Channel
	err = json.Unmarshal(bytes, &data)
	if err != nil {
		log.Fatal(err)
	}

	return data
}

// set(all_videos) in GE

var regions = [...]string{
	"FI", "KZ", "UA", "EG", "IT", "PT",
	"GB", "BE", "NO", "LV", "RU", "TR",
	"ES", "FR", "SE", "PL", "NL", "DE",
}

func initNodes(table map[string]*server.CacheServerNode) []*server.CacheServerNode {
	var nodes []*server.CacheServerNode
	for _, region := range regions {
		node := initServer(region, table)
		nodes = append(nodes, node)
	}

	return nodes
}

func initServer(region string, table map[string]*server.CacheServerNode) *server.CacheServerNode {
	capacity := getNodesCapacity(NodesStoragePath)[region]  // get from file nodes.json
	videosPath := VideoSetPath + region + "/video_set.json" // get from file ./setup/videoset/<region>/videoset.json
	videos := getVideosList(videosPath)                     // get from file RU_videos.txt

	node := server.NewCacheServerNode(region, capacity, videos)

	table[region] = node

	return node
}

func joinServers(
	path string,
	table map[string]*server.CacheServerNode,
	upAdjacency map[*server.CacheServerNode]*server.Edge) map[*server.CacheServerNode]*server.Edge {
	adjacencyTable := getAdjacency(path) // get from file adjacency.json
	for startName, channel := range adjacencyTable {
		edge := &server.Edge{
			Neighbour: table[channel.Name],
			Distance:  channel.Weight,
		}

		table[startName].SetUpLevel(edge)
		startNode := table[startName]
		upAdjacency[startNode] = edge
	}

	return upAdjacency
}
