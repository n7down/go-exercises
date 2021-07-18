package trip

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetConnectionsWithFourConnections(t *testing.T) {
	var (
		expectedConnections = 4
	)

	trip := NewTrip()
	trip.AddEdge("san franscico", "san diego")
	trip.AddEdge("san diego", "las vegas")
	trip.AddEdge("san diego", "los angeles")
	trip.AddEdge("las vegas", "salt lake city")
	trip.AddEdge("las vegas", "phoenix")
	trip.AddEdge("phoenix", "san antonio")
	trip.AddEdge("salt lake city", "denver")
	trip.AddEdge("denver", "indianapolis")
	trip.AddEdge("las vegas", "boise")
	trip.AddEdge("boise", "alberta")

	actualConnections := trip.GetConnections("san franscico", "denver")

	assert.Equal(t, expectedConnections, actualConnections, fmt.Sprintf("%d should return %d", actualConnections, expectedConnections))
}
