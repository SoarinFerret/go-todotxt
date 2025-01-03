package todo

import (
	"testing"

	"github.com/stretchr/testify/require"
)

// Data provider for TestTaskSegments.
//
//nolint:gochecknoglobals // test data
var testCasesTaskSegment = []struct {
	text string
	segs []*TaskSegment
}{
	{
		text: "2013-02-22 Pick up milk @GroceryStore",
		segs: []*TaskSegment{
			{
				Type:      SegmentCreatedDate,
				Originals: []string{"2013-02-22"},
				Display:   "2013-02-22",
			},
			{
				Type:      SegmentTodoText,
				Originals: []string{"Pick up milk @GroceryStore"},
				Display:   "Pick up milk @GroceryStore",
			},
			{
				Type:      SegmentContext,
				Originals: []string{"GroceryStore"},
				Display:   "@GroceryStore",
			},
		},
	},
	{
		text: "x Download Todo.txt mobile app @Phone",
		segs: []*TaskSegment{
			{
				Type:      SegmentIsCompleted,
				Originals: []string{"x"},
				Display:   "x",
			},
			{
				Type:      SegmentTodoText,
				Originals: []string{"Download Todo.txt mobile app @Phone"},
				Display:   "Download Todo.txt mobile app @Phone",
			},
			{
				Type:      SegmentContext,
				Originals: []string{"Phone"},
				Display:   "@Phone",
			},
		},
	},
	{
		text: "(B) 2013-12-01 Outline chapter 5 +Novel @Computer Level:5 private:false due:2014-02-17",
		segs: []*TaskSegment{
			{
				Type:      SegmentPriority,
				Originals: []string{"B"},
				Display:   "(B)",
			},
			{
				Type:      SegmentCreatedDate,
				Originals: []string{"2013-12-01"},
				Display:   "2013-12-01",
			},
			{
				Type:      SegmentTodoText,
				Originals: []string{"Outline chapter 5 +Novel @Computer"},
				Display:   "Outline chapter 5 +Novel @Computer",
			},
			{
				Type:      SegmentContext,
				Originals: []string{"Computer"},
				Display:   "@Computer",
			},
			{
				Type:      SegmentProject,
				Originals: []string{"Novel"},
				Display:   "+Novel",
			},
			{
				Type:      SegmentTag,
				Originals: []string{"Level", "5"},
				Display:   "Level:5",
			},
			{
				Type:      SegmentTag,
				Originals: []string{"private", "false"},
				Display:   "private:false",
			},
			{
				Type:      SegmentDueDate,
				Originals: []string{"due:2014-02-17"},
				Display:   "due:2014-02-17",
			},
		},
	},
	{
		text: "x 2014-01-02 (B) 2013-12-30 Create golang library test cases @Go +go-todotxt",
		segs: []*TaskSegment{
			{
				Type:      SegmentIsCompleted,
				Originals: []string{"x"},
				Display:   "x",
			},
			{
				Type:      SegmentCompletedDate,
				Originals: []string{"2014-01-02"},
				Display:   "2014-01-02",
			},
			{
				Type:      SegmentPriority,
				Originals: []string{"B"},
				Display:   "(B)",
			},
			{
				Type:      SegmentCreatedDate,
				Originals: []string{"2013-12-30"},
				Display:   "2013-12-30",
			},
			{
				Type:      SegmentTodoText,
				Originals: []string{"Create golang library test cases @Go +go-todotxt"},
				Display:   "Create golang library test cases @Go +go-todotxt",
			},
			{
				Type:      SegmentContext,
				Originals: []string{"Go"},
				Display:   "@Go",
			},
			{
				Type:      SegmentProject,
				Originals: []string{"go-todotxt"},
				Display:   "+go-todotxt",
			},
		},
	},
	{
		text: "x 2014-01-03 2014-01-01 Create some more golang library test cases @Go +go-todotxt",
		segs: []*TaskSegment{
			{
				Type:      SegmentIsCompleted,
				Originals: []string{"x"},
				Display:   "x",
			},
			{
				Type:      SegmentCompletedDate,
				Originals: []string{"2014-01-03"},
				Display:   "2014-01-03",
			},
			{
				Type:      SegmentCreatedDate,
				Originals: []string{"2014-01-01"},
				Display:   "2014-01-01",
			},
			{
				Type:      SegmentTodoText,
				Originals: []string{"Create some more golang library test cases @Go +go-todotxt"},
				Display:   "Create some more golang library test cases @Go +go-todotxt",
			},
			{
				Type:      SegmentContext,
				Originals: []string{"Go"},
				Display:   "@Go",
			},
			{
				Type:      SegmentProject,
				Originals: []string{"go-todotxt"},
				Display:   "+go-todotxt",
			},
		},
	},
}

func TestTask_Segments(t *testing.T) {
	t.Parallel()

	for _, test := range testCasesTaskSegment {
		task, err := ParseTask(test.text)

		require.NoError(t, err, "failed to parse task during test: %s", test.text)

		expectSegments := test.segs
		actualSegments := task.Segments()
		require.Equal(t, expectSegments, actualSegments, "segments do not match for task: %s", test.text)
	}
}
