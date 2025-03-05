package yqlib

import (
	"testing"
)

var compareOperatorScenarios = []expressionScenario{
	// ints, not equal
	{
		description: "Compare numbers (>)",
		document:    "a: 5\nb: 4",
		expression:  ".a > .b",
		expected: []string{
			"D0, P[a], (!!bool)::true\n",
		},
	},
	{
		skipDoc:    true,
		expression: "(.k | length) >= 0",
		expected: []string{
			"D0, P[k], (!!bool)::true\n",
		},
	},
	{
		skipDoc:    true,
		expression: `"2022-01-30T15:53:09Z" > "2020-01-30T15:53:09Z"`,
		expected: []string{
			"D0, P[], (!!bool)::true\n",
		},
	},
	{
		skipDoc:    true,
		document:   "a: 5\nb: 4",
		expression: ".a < .b",
		expected: []string{
			"D0, P[a], (!!bool)::false\n",
		},
	},
	{
		skipDoc:     true,
		description: "Compare integers (>=)",
		document:    "a: 5\nb: 4",
		expression:  ".a >= .b",
		expected: []string{
			"D0, P[a], (!!bool)::true\n",
		},
	},
	{
		skipDoc:    true,
		document:   "a: 5\nb: 4",
		expression: ".a <= .b",
		expected: []string{
			"D0, P[a], (!!bool)::false\n",
		},
	},

	// ints, equal
	{
		skipDoc:     true,
		description: "Compare equal numbers (>)",
		document:    "a: 5\nb: 5",
		expression:  ".a > .b",
		expected: []string{
			"D0, P[a], (!!bool)::false\n",
		},
	},
	{
		skipDoc:    true,
		document:   "a: 5\nb: 5",
		expression: ".a < .b",
		expected: []string{
			"D0, P[a], (!!bool)::false\n",
		},
	},
	{
		description: "Compare equal numbers (>=)",
		document:    "a: 5\nb: 5",
		expression:  ".a >= .b",
		expected: []string{
			"D0, P[a], (!!bool)::true\n",
		},
	},
	{
		skipDoc:    true,
		document:   "a: 5\nb: 5",
		expression: ".a <= .b",
		expected: []string{
			"D0, P[a], (!!bool)::true\n",
		},
	},

	// floats, not equal
	{
		skipDoc:    true,
		document:   "a: 5.2\nb: 4.1",
		expression: ".a > .b",
		expected: []string{
			"D0, P[a], (!!bool)::true\n",
		},
	},
	{
		skipDoc:    true,
		document:   "a: 5.2\nb: 4.1",
		expression: ".a < .b",
		expected: []string{
			"D0, P[a], (!!bool)::false\n",
		},
	},
	{
		skipDoc:    true,
		document:   "a: 5.2\nb: 4.1",
		expression: ".a >= .b",
		expected: []string{
			"D0, P[a], (!!bool)::true\n",
		},
	},
	{
		skipDoc:    true,
		document:   "a: 5.5\nb: 4.1",
		expression: ".a <= .b",
		expected: []string{
			"D0, P[a], (!!bool)::false\n",
		},
	},

	// floats, equal
	{
		skipDoc:    true,
		document:   "a: 5.5\nb: 5.5",
		expression: ".a > .b",
		expected: []string{
			"D0, P[a], (!!bool)::false\n",
		},
	},
	{
		skipDoc:    true,
		document:   "a: 5.5\nb: 5.5",
		expression: ".a < .b",
		expected: []string{
			"D0, P[a], (!!bool)::false\n",
		},
	},
	{
		skipDoc:    true,
		document:   "a: 5.1\nb: 5.1",
		expression: ".a >= .b",
		expected: []string{
			"D0, P[a], (!!bool)::true\n",
		},
	},
	{
		skipDoc:    true,
		document:   "a: 5.1\nb: 5.1",
		expression: ".a <= .b",
		expected: []string{
			"D0, P[a], (!!bool)::true\n",
		},
	},

	// strings, not equal
	{
		description:    "Compare strings",
		subdescription: "Compares strings by their bytecode.",
		document:       "a: zoo\nb: apple",
		expression:     ".a > .b",
		expected: []string{
			"D0, P[a], (!!bool)::true\n",
		},
	},
	{
		skipDoc:    true,
		document:   "a: zoo\nb: apple",
		expression: ".a < .b",
		expected: []string{
			"D0, P[a], (!!bool)::false\n",
		},
	},
	{
		skipDoc:    true,
		document:   "a: zoo\nb: apple",
		expression: ".a >= .b",
		expected: []string{
			"D0, P[a], (!!bool)::true\n",
		},
	},
	{
		skipDoc:    true,
		document:   "a: zoo\nb: apple",
		expression: ".a <= .b",
		expected: []string{
			"D0, P[a], (!!bool)::false\n",
		},
	},

	// strings, equal
	{
		skipDoc:    true,
		document:   "a: cat\nb: cat",
		expression: ".a > .b",
		expected: []string{
			"D0, P[a], (!!bool)::false\n",
		},
	},
	{
		skipDoc:    true,
		document:   "a: cat\nb: cat",
		expression: ".a < .b",
		expected: []string{
			"D0, P[a], (!!bool)::false\n",
		},
	},
	{
		skipDoc:    true,
		document:   "a: cat\nb: cat",
		expression: ".a >= .b",
		expected: []string{
			"D0, P[a], (!!bool)::true\n",
		},
	},
	{
		skipDoc:    true,
		document:   "a: cat\nb: cat",
		expression: ".a <= .b",
		expected: []string{
			"D0, P[a], (!!bool)::true\n",
		},
	},

	// datetime, not equal
	{
		description:    "Compare date times",
		subdescription: "You can compare date times. Assumes RFC3339 date time format, see [date-time operators](https://mikefarah.gitbook.io/yq/operators/date-time-operators) for more information.",
		document:       "a: 2021-01-01T03:10:00Z\nb: 2020-01-01T03:10:00Z",
		expression:     ".a > .b",
		expected: []string{
			"D0, P[a], (!!bool)::true\n",
		},
	},
	{
		skipDoc:    true,
		document:   "a: 2021-01-01T03:10:00Z\nb: 2020-01-01T03:10:00Z",
		expression: ".a < .b",
		expected: []string{
			"D0, P[a], (!!bool)::false\n",
		},
	},
	{
		skipDoc:    true,
		document:   "a: 2021-01-01T03:10:00Z\nb: 2020-01-01T03:10:00Z",
		expression: ".a >= .b",
		expected: []string{
			"D0, P[a], (!!bool)::true\n",
		},
	},
	{
		skipDoc:    true,
		document:   "a: 2021-01-01T03:10:00Z\nb: 2020-01-01T03:10:00Z",
		expression: ".a <= .b",
		expected: []string{
			"D0, P[a], (!!bool)::false\n",
		},
	},

	// datetime, equal
	{
		skipDoc:    true,
		document:   "a: 2021-01-01T03:10:00Z\nb: 2021-01-01T03:10:00Z",
		expression: ".a > .b",
		expected: []string{
			"D0, P[a], (!!bool)::false\n",
		},
	},
	{
		skipDoc:    true,
		document:   "a: 2021-01-01T03:10:00Z\nb: 2021-01-01T03:10:00Z",
		expression: ".a < .b",
		expected: []string{
			"D0, P[a], (!!bool)::false\n",
		},
	},
	{
		skipDoc:    true,
		document:   "a: 2021-01-01T03:10:00Z\nb: 2021-01-01T03:10:00Z",
		expression: ".a >= .b",
		expected: []string{
			"D0, P[a], (!!bool)::true\n",
		},
	},
	{
		skipDoc:    true,
		document:   "a: 2021-01-01T03:10:00Z\nb: 2021-01-01T03:10:00Z",
		expression: ".a <= .b",
		expected: []string{
			"D0, P[a], (!!bool)::true\n",
		},
	},
	// both null
	{
		description: "Both sides are null: > is false",
		expression:  ".a > .b",
		expected: []string{
			"D0, P[a], (!!bool)::false\n",
		},
	},
	{
		skipDoc:    true,
		expression: ".a < .b",
		expected: []string{
			"D0, P[a], (!!bool)::false\n",
		},
	},
	{
		description: "Both sides are null: >= is true",
		expression:  ".a >= .b",
		expected: []string{
			"D0, P[a], (!!bool)::true\n",
		},
	},
	{
		skipDoc:    true,
		expression: ".a <= .b",
		expected: []string{
			"D0, P[a], (!!bool)::true\n",
		},
	},

	// one null
	{
		skipDoc:     true,
		description: "One side is null: > is false",
		document:    `a: 5`,
		expression:  ".a > .b",
		expected: []string{
			"D0, P[a], (!!bool)::false\n",
		},
	},
	{
		skipDoc:    true,
		document:   `a: 5`,
		expression: ".a < .b",
		expected: []string{
			"D0, P[a], (!!bool)::false\n",
		},
	},
	{
		skipDoc:     true,
		description: "One side is null: >= is false",
		document:    `a: 5`,
		expression:  ".a >= .b",
		expected: []string{
			"D0, P[a], (!!bool)::false\n",
		},
	},
	{
		skipDoc:    true,
		document:   `a: 5`,
		expression: ".a <= .b",
		expected: []string{
			"D0, P[a], (!!bool)::false\n",
		},
	},
	{
		skipDoc:    true,
		document:   `a: 5`,
		expression: ".b <= .a",
		expected: []string{
			"D0, P[b], (!!bool)::false\n",
		},
	},
	{
		skipDoc:    true,
		document:   `a: 5`,
		expression: ".b < .a",
		expected: []string{
			"D0, P[b], (!!bool)::false\n",
		},
	},
}

func TestCompareOperatorScenarios(t *testing.T) {
	for _, tt := range compareOperatorScenarios {
		testScenario(t, &tt)
	}
	documentOperatorScenarios(t, "compare", compareOperatorScenarios)
}

var minOperatorScenarios = []expressionScenario{
	{
		description: "Minimum int",
		document:    "[99, 16, 12, 6, 66]\n",
		expression:  `min`,
		expected: []string{
			"D0, P[3], (!!int)::6\n",
		},
	},
	{
		description: "Minimum string",
		document:    "[foo, bar, baz]\n",
		expression:  `min`,
		expected: []string{
			"D0, P[1], (!!str)::bar\n",
		},
	},
	{
		description: "Minimum of empty",
		document:    "[]\n",
		expression:  `min`,
		expected:    []string{},
	},
}

func TestMinOperatorScenarios(t *testing.T) {
	for _, tt := range minOperatorScenarios {
		testScenario(t, &tt)
	}
	documentOperatorScenarios(t, "min", minOperatorScenarios)
}

var maxOperatorScenarios = []expressionScenario{
	{
		description: "Maximum int",
		document:    "[99, 16, 12, 6, 66]\n",
		expression:  `max`,
		expected: []string{
			"D0, P[0], (!!int)::99\n",
		},
	},
	{
		description: "Maximum string",
		document:    "[foo, bar, baz]\n",
		expression:  `max`,
		expected: []string{
			"D0, P[0], (!!str)::foo\n",
		},
	},
	{
		description: "Maximum of empty",
		document:    "[]\n",
		expression:  `max`,
		expected:    []string{},
	},
}

func TestMaxOperatorScenarios(t *testing.T) {
	for _, tt := range maxOperatorScenarios {
		testScenario(t, &tt)
	}
	documentOperatorScenarios(t, "max", maxOperatorScenarios)
}
