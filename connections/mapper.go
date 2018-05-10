package connections

import (
	"reflect"
	"sort"
	"strings"
)

// Mapper struct
type Mapper struct{}

func (m *Mapper) getRelatedConnections(currentConnections map[string][]string, updated map[string]bool, needToCheckAgain *bool) {
	// Parse each key for connections
	for currentConnectionKeys, currentConnectionValues := range currentConnections {
		uniqueConnections := make([]string, 0)
		newConnections := make([]string, 0)

		// Check each current connection values for other values
		for _, currentConnectionValue := range currentConnectionValues {
			newConnections = append(newConnections, currentConnections[currentConnectionValue]...)
		}

		// Add existing connections
		newConnections = append(newConnections, currentConnections[currentConnectionKeys]...)

		// Create unique connections
		for _, v := range newConnections {
			if !m.isStringInSlice(v, uniqueConnections) {
				uniqueConnections = append(uniqueConnections, v)
			}
		}

		// Update the status
		updated[currentConnectionKeys] = len(uniqueConnections) != len(currentConnections[currentConnectionKeys])

		// Notify if a recheck is needed
		if updated[currentConnectionKeys] && !*needToCheckAgain {
			*needToCheckAgain = true
		}

		// Sort connections
		sort.Strings(uniqueConnections)

		// Save as new current connection
		currentConnections[currentConnectionKeys] = uniqueConnections
	}
}

func (m *Mapper) isArrayStringInArraySlice(v []string, list [][]string) bool {
	for _, item := range list {
		if reflect.DeepEqual(item, v) {
			return true
		}
	}
	return false
}

func (m *Mapper) isStringInSlice(v string, list []string) bool {
	for _, item := range list {
		if item == v {
			return true
		}
	}
	return false
}

func (m *Mapper) normalizeData(container map[string]string) map[string][]string {
	normalizedData := make(map[string][]string)

	// Start mapping the immediate connections
	for k, v1 := range container {
		// Parse values of key
		relatives := strings.Split(v1, ",")

		// Clean the values
		for i, v2 := range relatives {
			relatives[i] = strings.TrimSpace(v2)
		}

		// Place them in the connections container
		normalizedData[strings.TrimSpace(k)] = relatives
	}

	return normalizedData
}

func (m *Mapper) CreateConnections(container map[string]string) map[string][]string {
	var needToCheckAgain bool

	// Normalize data
	connections := m.normalizeData(container)

	// Create updated status mapper
	updated := make(map[string]bool)

	// Get all connections
	for ok := true; ok; ok = needToCheckAgain != false {
		needToCheckAgain = false
		m.getRelatedConnections(connections, updated, &needToCheckAgain)
	}

	return connections
}

func (m *Mapper) CreateGroups(connections map[string][]string) [][]string {
	groups := make([][]string, 0)
	for _, v := range connections {
		if !m.isArrayStringInArraySlice(v, groups) {
			groups = append(groups, v)
		}
	}

	return groups
}

func (m *Mapper) CreateMap(container map[string]string, key, value string) {
	// Get the values
	key = strings.TrimSpace(key)
	value = strings.TrimSpace(value)

	// Check if there are values
	if len(key) != 0 && len(value) != 0 {
		container[key] = value
	}
}
