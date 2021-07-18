package trip

type Node struct {
	City    string
	Visited bool
}

type Trip struct {
	connectionMap map[string][]Node
}

func NewTrip() *Trip {
	return &Trip{
		connectionMap: make(map[string][]Node),
	}
}

func (t *Trip) AddEdge(a, b string) {
	if val, ok := t.connectionMap[a]; ok {
		n := Node{
			City: b,
		}
		val = append(val, n)
		t.connectionMap[a] = val
	} else {
		t.connectionMap[a] = []Node{
			Node{
				City: b,
			},
		}
	}
}

func (t *Trip) GetConnections(a, b string) int {
	found, numberConnections := t.findCity(a, b)
	if found {
		return numberConnections
	}
	return 0
}

func (t *Trip) findCity(a, b string) (bool, int) {
	if nodes, ok := t.connectionMap[a]; ok {
		for _, node := range nodes {
			if node.City == b {
				return true, 1
			} else if !node.Visited {
				node.Visited = true
				found, numberConnections := t.findCity(node.City, b)
				if found {
					return true, numberConnections + 1
				}
			}
		}
	}
	return false, 0
}
