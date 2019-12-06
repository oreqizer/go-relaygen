package generator

import (
	"text/template"
)

var test = template.Must(template.New("generated").
	Funcs(template.FuncMap{
		"lcFirst": lcFirst,
	}).
	Parse(`
// Code generated by github.com/oreqizer/go-relaygen, DO NOT EDIT.

package {{ .Package }}

import (
	"github.com/oreqizer/go-relaygen/relay"
	"testing"
)

var (
	{{ .Name|lcFirst }}onestr   = "1"
	{{ .Name|lcFirst }}twostr   = "2"
	{{ .Name|lcFirst }}threestr = "3"
	{{ .Name|lcFirst }}fourstr  = "4"
	{{ .Name|lcFirst }}fivestr  = "5"

	// Set later
	{{ .Name|lcFirst }}oneid   = ""
	{{ .Name|lcFirst }}twoid   = ""
	{{ .Name|lcFirst }}threeid = ""
	{{ .Name|lcFirst }}fourid  = ""
	{{ .Name|lcFirst }}fiveid  = ""
)

var (
	{{ .Name|lcFirst }}zero = 0
	{{ .Name|lcFirst }}one  = 1
	{{ .Name|lcFirst }}two  = 2
	// {{ .Name|lcFirst }}three = 3
	{{ .Name|lcFirst }}four = 4
	// {{ .Name|lcFirst }}five  = 5
	{{ .Name|lcFirst }}six = 6
)

var nodes{{ .Name }} = []*{{ .Name }}{
	&{{ .Name }}{},
	&{{ .Name }}{},
	&{{ .Name }}{},
	&{{ .Name }}{},
	&{{ .Name }}{},
}

var edges{{ .Name }} = []*{{ .Name }}Edge{
	{{"{"}}Node: &{{ .Name }}{}, Cursor: ""},
	{{"{"}}Node: &{{ .Name }}{}, Cursor: ""},
	{{"{"}}Node: &{{ .Name }}{}, Cursor: ""},
	{{"{"}}Node: &{{ .Name }}{}, Cursor: ""},
	{{"{"}}Node: &{{ .Name }}{}, Cursor: ""},
}

func init() {
	nodes{{ .Name }}[0].{{ .ID }} = {{ .Name|lcFirst }}onestr
	nodes{{ .Name }}[1].{{ .ID }} = {{ .Name|lcFirst }}twostr
	nodes{{ .Name }}[2].{{ .ID }} = {{ .Name|lcFirst }}threestr
	nodes{{ .Name }}[3].{{ .ID }} = {{ .Name|lcFirst }}fourstr
	nodes{{ .Name }}[4].{{ .ID }} = {{ .Name|lcFirst }}fivestr

	{{ .Name|lcFirst }}oneid = nodes{{ .Name }}[0].ID()
	{{ .Name|lcFirst }}twoid = nodes{{ .Name }}[1].ID()
	{{ .Name|lcFirst }}threeid = nodes{{ .Name }}[2].ID()
	{{ .Name|lcFirst }}fourid = nodes{{ .Name }}[3].ID()
	{{ .Name|lcFirst }}fiveid = nodes{{ .Name }}[4].ID()

	edges{{ .Name }}[0].Node.{{ .ID }} = {{ .Name|lcFirst }}onestr
	edges{{ .Name }}[0].Cursor = {{ .Name|lcFirst }}oneid
	edges{{ .Name }}[1].Node.{{ .ID }} = {{ .Name|lcFirst }}twostr
	edges{{ .Name }}[1].Cursor = {{ .Name|lcFirst }}twoid
	edges{{ .Name }}[2].Node.{{ .ID }} = {{ .Name|lcFirst }}threestr
	edges{{ .Name }}[2].Cursor = {{ .Name|lcFirst }}threeid
	edges{{ .Name }}[3].Node.{{ .ID }} = {{ .Name|lcFirst }}fourstr
	edges{{ .Name }}[3].Cursor = {{ .Name|lcFirst }}fourid
	edges{{ .Name }}[4].Node.{{ .ID }} = {{ .Name|lcFirst }}fivestr
	edges{{ .Name }}[4].Cursor = {{ .Name|lcFirst }}fiveid
}

var table{{ .Name }}ConnectionFromArray = []struct {
	nodes []*{{ .Name }}
	args  *relay.ConnectionArgs
	out   *{{ .Name }}Connection
}{
	{
		nodes: []*{{ .Name }}{},
		args:  &relay.ConnectionArgs{Before: nil, After: nil, First: &{{ .Name|lcFirst }}zero, Last: &{{ .Name|lcFirst }}zero},
		out: &{{ .Name }}Connection{
			Edges:    []*{{ .Name }}Edge{},
			PageInfo: relay.PageInfo{HasNextPage: false, HasPreviousPage: false, StartCursor: nil, EndCursor: nil},
		},
	},
	{
		nodes: nodes{{ .Name }},
		args:  &relay.ConnectionArgs{Before: nil, After: nil, First: &{{ .Name|lcFirst }}zero, Last: &{{ .Name|lcFirst }}zero},
		out: &{{ .Name }}Connection{
			Edges:    edges{{ .Name }},
			PageInfo: relay.PageInfo{HasNextPage: false, HasPreviousPage: false, StartCursor: &{{ .Name|lcFirst }}oneid, EndCursor: &{{ .Name|lcFirst }}fiveid},
		},
	},
	{
		nodes: nodes{{ .Name }},
		args:  &relay.ConnectionArgs{Before: nil, After: nil, First: &{{ .Name|lcFirst }}two, Last: &{{ .Name|lcFirst }}zero},
		out: &{{ .Name }}Connection{
			Edges:    edges{{ .Name }}[:2],
			PageInfo: relay.PageInfo{HasNextPage: true, HasPreviousPage: false, StartCursor: &{{ .Name|lcFirst }}oneid, EndCursor: &{{ .Name|lcFirst }}twoid},
		},
	},
	{
		nodes: nodes{{ .Name }},
		args:  &relay.ConnectionArgs{Before: nil, After: nil, First: &{{ .Name|lcFirst }}zero, Last: &{{ .Name|lcFirst }}two},
		out: &{{ .Name }}Connection{
			Edges:    edges{{ .Name }}[3:],
			PageInfo: relay.PageInfo{HasNextPage: false, HasPreviousPage: true, StartCursor: &{{ .Name|lcFirst }}fourid, EndCursor: &{{ .Name|lcFirst }}fiveid},
		},
	},
	{
		nodes: nodes{{ .Name }},
		args:  &relay.ConnectionArgs{Before: &{{ .Name|lcFirst }}fourid, After: &{{ .Name|lcFirst }}twoid, First: &{{ .Name|lcFirst }}zero, Last: &{{ .Name|lcFirst }}zero},
		out: &{{ .Name }}Connection{
			Edges:    edges{{ .Name }}[1:4],
			PageInfo: relay.PageInfo{HasNextPage: false, HasPreviousPage: false, StartCursor: &{{ .Name|lcFirst }}twoid, EndCursor: &{{ .Name|lcFirst }}fourid},
		},
	},
}

func Test{{ .Name }}ConnectionFromArray(t *testing.T) {
	empty := {{ .Name }}ConnectionFromArray(table{{ .Name }}ConnectionFromArray[0].nodes, nil)
	if empty != nil {
		t.Errorf("Expected nil output for nil args")
	}

	for i, e := range table{{ .Name }}ConnectionFromArray {
		out := {{ .Name }}ConnectionFromArray(e.nodes, e.args)
		if out == nil {
			t.Errorf("Unexpected nil output")
			return
		}

		// PageInfo
		if out.PageInfo.HasNextPage != e.out.PageInfo.HasNextPage {
			t.Errorf("%d: Has next page: got %v, want %v", i, out.PageInfo.HasNextPage, e.out.PageInfo.HasNextPage)
		}

		if out.PageInfo.HasPreviousPage != e.out.PageInfo.HasPreviousPage {
			t.Errorf("%d: Has previous page: got %v, want %v", i, out.PageInfo.HasPreviousPage, e.out.PageInfo.HasPreviousPage)
		}

		if out.PageInfo.StartCursor != nil || e.out.PageInfo.StartCursor != nil {
			if *out.PageInfo.StartCursor != *e.out.PageInfo.StartCursor {
				t.Errorf("%d: Start cursor: got %v, want %v", i, *out.PageInfo.StartCursor, *e.out.PageInfo.StartCursor)
			}
		}

		if out.PageInfo.EndCursor != nil || e.out.PageInfo.EndCursor != nil {
			if *out.PageInfo.EndCursor != *e.out.PageInfo.EndCursor {
				t.Errorf("%d: End cursor: got %v, want %v", i, *out.PageInfo.EndCursor, *e.out.PageInfo.EndCursor)
			}
		}

		// Edges
		if len(out.Edges) != len(e.out.Edges) {
			t.Errorf("%d: Edges length: got %d, want %d", i, len(out.Edges), len(e.out.Edges))
			return
		}

		for j, eedge := range e.out.Edges {
			oedge := out.Edges[j]
			if eedge.Cursor != oedge.Cursor {
				t.Errorf("%d: Edge #%d: Cursor: got %s, want %s", i, j, oedge.Cursor, eedge.Cursor)
			}
		}
	}
}

var table{{ .Name }}EdgesToReturn = []struct {
	edges  []*{{ .Name }}Edge
	before *string
	after  *string
	first  *int
	last   *int
	out    *{{ .Name }}Connection
}{
	{edges: []*{{ .Name }}Edge{}, before: nil, after: nil, first: &{{ .Name|lcFirst }}zero, last: &{{ .Name|lcFirst }}zero, out: table{{ .Name }}ConnectionFromArray[0].out},
	{edges: edges{{ .Name }}, before: nil, after: nil, first: &{{ .Name|lcFirst }}zero, last: &{{ .Name|lcFirst }}zero, out: table{{ .Name }}ConnectionFromArray[1].out},
	{edges: edges{{ .Name }}, before: nil, after: nil, first: &{{ .Name|lcFirst }}two, last: &{{ .Name|lcFirst }}zero, out: table{{ .Name }}ConnectionFromArray[2].out},
	{edges: edges{{ .Name }}, before: nil, after: nil, first: &{{ .Name|lcFirst }}zero, last: &{{ .Name|lcFirst }}two, out: table{{ .Name }}ConnectionFromArray[3].out},
	{edges: edges{{ .Name }}, before: &{{ .Name|lcFirst }}fourid, after: &{{ .Name|lcFirst }}twoid, first: &{{ .Name|lcFirst }}zero, last: &{{ .Name|lcFirst }}zero, out: table{{ .Name }}ConnectionFromArray[4].out},
}

func Test{{ .Name }}EdgesToReturn(t *testing.T) {
	for i, e := range table{{ .Name }}EdgesToReturn {
		out := {{ .Name|lcFirst }}EdgesToReturn(e.edges, e.before, e.after, e.first, e.last)
		// PageInfo
		if out.PageInfo.HasNextPage != e.out.PageInfo.HasNextPage {
			t.Errorf("%d: Has next page: got %v, want %v", i, out.PageInfo.HasNextPage, e.out.PageInfo.HasNextPage)
		}

		if out.PageInfo.HasPreviousPage != e.out.PageInfo.HasPreviousPage {
			t.Errorf("%d: Has previous page: got %v, want %v", i, out.PageInfo.HasPreviousPage, e.out.PageInfo.HasPreviousPage)
		}

		if out.PageInfo.StartCursor != nil || e.out.PageInfo.StartCursor != nil {
			if *out.PageInfo.StartCursor != *e.out.PageInfo.StartCursor {
				t.Errorf("%d: Start cursor: got %v, want %v", i, *out.PageInfo.StartCursor, *e.out.PageInfo.StartCursor)
			}
		}

		if out.PageInfo.EndCursor != nil || e.out.PageInfo.EndCursor != nil {
			if *out.PageInfo.EndCursor != *e.out.PageInfo.EndCursor {
				t.Errorf("%d: End cursor: got %v, want %v", i, *out.PageInfo.EndCursor, *e.out.PageInfo.EndCursor)
			}
		}

		// Edges
		if len(out.Edges) != len(e.out.Edges) {
			t.Errorf("%d: Edges length: got %d, want %d", i, len(out.Edges), len(e.out.Edges))
			return
		}

		for j, eedge := range e.out.Edges {
			oedge := out.Edges[j]
			if eedge.Cursor != oedge.Cursor {
				t.Errorf("%d: Edge #%d: Cursor: got %s, want %s", i, j, oedge.Cursor, eedge.Cursor)
			}
		}
	}
}

var table{{ .Name }}ApplyCursorsToEdges = []struct {
	edges  []*{{ .Name }}Edge
	before *string
	after  *string
	len    int
}{
	{edges: []*{{ .Name }}Edge{}, before: nil, after: nil, len: 0},
	{edges: edges{{ .Name }}, before: nil, after: nil, len: 5},
	{edges: edges{{ .Name }}, before: nil, after: &{{ .Name|lcFirst }}twoid, len: 4},
	{edges: edges{{ .Name }}, before: &{{ .Name|lcFirst }}fourid, after: nil, len: 4},
	{edges: edges{{ .Name }}, before: &{{ .Name|lcFirst }}fourid, after: &{{ .Name|lcFirst }}twoid, len: 3},
}

func Test{{ .Name }}ApplyCursorsToEdges(t *testing.T) {
	for i, e := range table{{ .Name }}ApplyCursorsToEdges {
		out := {{ .Name|lcFirst }}ApplyCursorsToEdges(e.edges, e.before, e.after)
		if len(out) != e.len {
			t.Errorf("%d: Length: got %d, want %d", i, len(out), e.len)
		}

		if len(out) > 0 {
			if cursor := out[len(out)-1].Cursor; e.before != nil && cursor != *e.before {
				t.Errorf("%d: Before: got %s, want %s", i, cursor, *e.before)
			}

			if cursor := out[0].Cursor; e.after != nil && cursor != *e.after {
				t.Errorf("%d: After: got %s, want %s", i, cursor, *e.after)
			}
		}
	}
}

var table{{ .Name }}HasPreviousPage = []struct {
	edges  []*{{ .Name }}Edge
	before *string
	after  *string
	last   *int
	out    bool
}{
	{edges: []*{{ .Name }}Edge{}, before: nil, after: nil, last: nil, out: false},
	{edges: edges{{ .Name }}, before: nil, after: nil, last: nil, out: false},
	{edges: edges{{ .Name }}, before: nil, after: nil, last: &{{ .Name|lcFirst }}zero, out: false},
	{edges: edges{{ .Name }}, before: nil, after: nil, last: &{{ .Name|lcFirst }}six, out: false},
	{edges: edges{{ .Name }}, before: nil, after: nil, last: &{{ .Name|lcFirst }}four, out: true},
	{edges: edges{{ .Name }}, before: &{{ .Name|lcFirst }}fourid, after: &{{ .Name|lcFirst }}twoid, last: &{{ .Name|lcFirst }}four, out: false},
	{edges: edges{{ .Name }}, before: &{{ .Name|lcFirst }}fourid, after: &{{ .Name|lcFirst }}twoid, last: &{{ .Name|lcFirst }}one, out: true},
}

func Test{{ .Name }}HasPreviousPage(t *testing.T) {
	for i, e := range table{{ .Name }}HasPreviousPage {
		out := {{ .Name|lcFirst }}HasPreviousPage(e.edges, e.before, e.after, e.last)
		if out != e.out {
			t.Errorf("%d: got %v, want %v", i, out, e.out)
		}
	}
}

var table{{ .Name }}HasNextPage = []struct {
	edges  []*{{ .Name }}Edge
	before *string
	after  *string
	first  *int
	out    bool
}{
	{edges: []*{{ .Name }}Edge{}, before: nil, after: nil, first: nil, out: false},
	{edges: edges{{ .Name }}, before: nil, after: nil, first: nil, out: false},
	{edges: edges{{ .Name }}, before: nil, after: nil, first: &{{ .Name|lcFirst }}zero, out: false},
	{edges: edges{{ .Name }}, before: nil, after: nil, first: &{{ .Name|lcFirst }}six, out: false},
	{edges: edges{{ .Name }}, before: nil, after: nil, first: &{{ .Name|lcFirst }}four, out: true},
	{edges: edges{{ .Name }}, before: &{{ .Name|lcFirst }}fourid, after: &{{ .Name|lcFirst }}twoid, first: &{{ .Name|lcFirst }}four, out: false},
	{edges: edges{{ .Name }}, before: &{{ .Name|lcFirst }}fourid, after: &{{ .Name|lcFirst }}twoid, first: &{{ .Name|lcFirst }}one, out: true},
}

func Test{{ .Name }}HasNextPage(t *testing.T) {
	for i, e := range table{{ .Name }}HasNextPage {
		out := {{ .Name|lcFirst }}HasNextPage(e.edges, e.before, e.after, e.first)
		if out != e.out {
			t.Errorf("%d: got %v, want %v", i, out, e.out)
		}
	}
}
`))
