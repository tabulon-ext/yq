package yqlib

import (
	"bufio"
	"strings"
	"testing"

	"github.com/mikefarah/yq/v4/test"
)

var evaluateNodesScenario = []expressionScenario{
	{
		document:   `a: hello`,
		expression: `.a`,
		expected: []string{
			"D0, P[a], (!!str)::hello\n",
		},
	},
	{
		document:   `a: hello`,
		expression: `.`,
		expected: []string{
			"D0, P[], (!!map)::a: hello\n",
		},
	},
	{
		document:   `- a: "yes"`,
		expression: `.[] | has("a")`,
		expected: []string{
			"D0, P[0], (!!bool)::true\n",
		},
	},
}

func TestAllAtOnceEvaluateNodes(t *testing.T) {
	var evaluator = NewAllAtOnceEvaluator()
	// logging.SetLevel(logging.DEBUG, "")
	for _, tt := range evaluateNodesScenario {
		decoder := NewYamlDecoder(ConfiguredYamlPreferences)
		reader := bufio.NewReader(strings.NewReader(tt.document))
		err := decoder.Init(reader)
		if err != nil {
			t.Error(err)
			return
		}
		candidateNode, errorReading := decoder.Decode()

		if errorReading != nil {
			t.Error(errorReading)
			return
		}

		list, _ := evaluator.EvaluateNodes(tt.expression, candidateNode)
		test.AssertResultComplex(t, tt.expected, resultsToString(t, list))
	}
}
