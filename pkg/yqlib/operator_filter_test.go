package yqlib

import (
	"testing"
)

var filterOperatorScenarios = []expressionScenario{
	{
		description: "Filter array",
		document:    `[1,2,3]`,
		expression:  `filter(. < 3)`,
		expected: []string{
			"D0, P[], (!!seq)::[1, 2]\n",
		},
	},
	{
		description: "Filter array splat",
		skipDoc:     true,
		document:    `[1,2,3]`,
		expression:  `filter(. < 3)[]`,
		expected: []string{
			"D0, P[0], (!!int)::1\n",
			"D0, P[1], (!!int)::2\n",
		},
	},
	{
		description: "Filter map values",
		document:    `{c: {things: cool, frog: yes}, d: {things: hot, frog: false}}`,
		expression:  `filter(.things == "cool")`,
		expected: []string{
			"D0, P[], (!!seq)::[{things: cool, frog: yes}]\n",
		},
	},
	{
		skipDoc:    true,
		document:   `[1,2,3]`,
		expression: `filter(. > 1)`,
		expected: []string{
			"D0, P[], (!!seq)::[2, 3]\n",
		},
	},
	{
		skipDoc:     true,
		description: "Filter array to empty",
		document:    `[1,2,3]`,
		expression:  `filter(. > 4)`,
		expected: []string{
			"D0, P[], (!!seq)::[]\n",
		},
	},
	{
		skipDoc:     true,
		description: "Filter empty array",
		document:    `[]`,
		expression:  `filter(. > 1)`,
		expected: []string{
			"D0, P[], (!!seq)::[]\n",
		},
	},
}

func TestFilterOperatorScenarios(t *testing.T) {
	for _, tt := range filterOperatorScenarios {
		testScenario(t, &tt)
	}
	documentOperatorScenarios(t, "filter", filterOperatorScenarios)
}
