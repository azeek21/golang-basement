package breadthfirstsearch

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBreadthFirstSearch(t *testing.T) {
	// Test graph structure using string nodes
	type Graph map[string][]string

	t.Run("direct match in initial items", func(t *testing.T) {
		g := Graph{
			"A": {"B", "C"},
			"B": {},
			"C": {},
		}

		result, err := BreadthFirstSearch(g,
			func(g Graph) []string { return []string{"A", "B", "C"} },
			func(g Graph, item string) []string { return g[item] },
			func(item string) bool { return item == "B" },
		)

		assert.NoError(t, err)
		assert.Equal(t, "B", result)
	})

	t.Run("empty initial items", func(t *testing.T) {
		g := Graph{}
		_, err := BreadthFirstSearch(g,
			func(g Graph) []string { return []string{} },
			func(g Graph, item string) []string { return g[item] },
			func(item string) bool { return true },
		)

		assert.Error(t, err)
		assert.True(t, errors.Is(err, ERR_NOT_FOUND))
	})

	t.Run("no matching node", func(t *testing.T) {
		g := Graph{
			"A": {"B"},
			"B": {"C"},
			"C": {},
		}

		_, err := BreadthFirstSearch(g,
			func(g Graph) []string { return []string{"A"} },
			func(g Graph, item string) []string { return g[item] },
			func(item string) bool { return false },
		)

		assert.Error(t, err)
		assert.True(t, errors.Is(err, ERR_NOT_FOUND))
	})

	t.Run("linear graph traversal", func(t *testing.T) {
		g := Graph{
			"A": {"B"},
			"B": {"C"},
			"C": {"D"},
			"D": {},
		}

		result, err := BreadthFirstSearch(g,
			func(g Graph) []string { return []string{"A"} },
			func(g Graph, item string) []string { return g[item] },
			func(item string) bool { return item == "D" },
		)

		assert.NoError(t, err)
		assert.Equal(t, "D", result)
	})

	t.Run("cycle handling", func(t *testing.T) {
		g := Graph{
			"A": {"B"},
			"B": {"A", "C"},
			"C": {},
		}

		result, err := BreadthFirstSearch(g,
			func(g Graph) []string { return []string{"A"} },
			func(g Graph, item string) []string { return g[item] },
			func(item string) bool { return item == "C" },
		)

		assert.NoError(t, err)
		assert.Equal(t, "C", result)
	})

	t.Run("multiple valid nodes at different depths", func(t *testing.T) {
		g := Graph{
			"A": {"B", "D"},
			"B": {"C"},
			"C": {"E"},
			"D": {"E"},
			"E": {},
		}

		result, err := BreadthFirstSearch(g,
			func(g Graph) []string { return []string{"A"} },
			func(g Graph, item string) []string { return g[item] },
			func(item string) bool { return item == "D" || item == "E" },
		)

		assert.NoError(t, err)
		assert.Equal(t, "D", result)
	})

	t.Run("complex graph with multiple paths", func(t *testing.T) {
		g := Graph{
			"A": {"B", "C"},
			"B": {"D"},
			"C": {"D"},
			"D": {"E"},
			"E": {},
		}

		result, err := BreadthFirstSearch(g,
			func(g Graph) []string { return []string{"A"} },
			func(g Graph, item string) []string { return g[item] },
			func(item string) bool { return item == "E" },
		)

		assert.NoError(t, err)
		assert.Equal(t, "E", result)
	})

	t.Run("multiple initial items with match", func(t *testing.T) {
		g := Graph{
			"A": {},
			"B": {},
			"C": {},
		}

		result, err := BreadthFirstSearch(g,
			func(g Graph) []string { return []string{"A", "B", "C"} },
			func(g Graph, item string) []string { return g[item] },
			func(item string) bool { return item == "C" },
		)

		assert.NoError(t, err)
		assert.Equal(t, "C", result)
	})

	t.Run("custom type nodes", func(t *testing.T) {
		type Node struct {
			ID    int
			Value string
		}

		g := map[Node][]Node{
			{1, "A"}: {{2, "B"}, {3, "C"}},
			{2, "B"}: {{4, "D"}},
			{3, "C"}: {},
			{4, "D"}: {},
		}

		result, err := BreadthFirstSearch(g,
			func(g map[Node][]Node) []Node { return []Node{{1, "A"}} },
			func(g map[Node][]Node, item Node) []Node { return g[item] },
			func(item Node) bool { return item.ID == 4 },
		)

		assert.NoError(t, err)
		assert.Equal(t, Node{4, "D"}, result)
	})
}
