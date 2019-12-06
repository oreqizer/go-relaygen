// Code generated by github.com/oreqizer/go-relaygen, DO NOT EDIT.

package example

import (
	"testing"

	"github.com/oreqizer/go-relaygen/relay"
)

var (
	ONESTR   = "1"
	TWOSTR   = "2"
	THREESTR = "3"
	FOURSTR  = "4"
	FIVESTR  = "5"
)

var (
	ZERO = 0
	ONE  = 1
	TWO  = 2
	// THREE = 3
	FOUR = 4
	// FIVE  = 5
	SIX = 6
)

var nodesUser = []*User{
	&User{},
	&User{},
	&User{},
	&User{},
	&User{},
}

var edgesUser = []*UserEdge{
	{Node: &User{}, Cursor: ONESTR},
	{Node: &User{}, Cursor: TWOSTR},
	{Node: &User{}, Cursor: THREESTR},
	{Node: &User{}, Cursor: FOURSTR},
	{Node: &User{}, Cursor: FIVESTR},
}

func init() {
	nodesUser[0].Person.ID = ONESTR
	nodesUser[1].Person.ID = TWOSTR
	nodesUser[2].Person.ID = THREESTR
	nodesUser[3].Person.ID = FOURSTR
	nodesUser[4].Person.ID = FIVESTR

	edgesUser[0].Node.Person.ID = ONESTR
	edgesUser[1].Node.Person.ID = TWOSTR
	edgesUser[2].Node.Person.ID = THREESTR
	edgesUser[3].Node.Person.ID = FOURSTR
	edgesUser[4].Node.Person.ID = FIVESTR
}

var tableUserConnectionFromArray = []struct {
	nodes []*User
	args  *relay.ConnectionArgs
	out   *UserConnection
}{
	{
		nodes: []*User{},
		args:  &relay.ConnectionArgs{Before: nil, After: nil, First: &ZERO, Last: &ZERO},
		out: &UserConnection{
			Edges:    []*UserEdge{},
			PageInfo: relay.PageInfo{HasNextPage: false, HasPreviousPage: false, StartCursor: nil, EndCursor: nil},
		},
	},
	{
		nodes: nodesUser,
		args:  &relay.ConnectionArgs{Before: nil, After: nil, First: &ZERO, Last: &ZERO},
		out: &UserConnection{
			Edges:    edgesUser,
			PageInfo: relay.PageInfo{HasNextPage: false, HasPreviousPage: false, StartCursor: &ONESTR, EndCursor: &FIVESTR},
		},
	},
	{
		nodes: nodesUser,
		args:  &relay.ConnectionArgs{Before: nil, After: nil, First: &TWO, Last: &ZERO},
		out: &UserConnection{
			Edges:    edgesUser[:2],
			PageInfo: relay.PageInfo{HasNextPage: true, HasPreviousPage: false, StartCursor: &ONESTR, EndCursor: &TWOSTR},
		},
	},
	{
		nodes: nodesUser,
		args:  &relay.ConnectionArgs{Before: nil, After: nil, First: &ZERO, Last: &TWO},
		out: &UserConnection{
			Edges:    edgesUser[3:],
			PageInfo: relay.PageInfo{HasNextPage: false, HasPreviousPage: true, StartCursor: &FOURSTR, EndCursor: &FIVESTR},
		},
	},
	{
		nodes: nodesUser,
		args:  &relay.ConnectionArgs{Before: &FOURSTR, After: &TWOSTR, First: &ZERO, Last: &ZERO},
		out: &UserConnection{
			Edges:    edgesUser[1:4],
			PageInfo: relay.PageInfo{HasNextPage: false, HasPreviousPage: false, StartCursor: &TWOSTR, EndCursor: &FOURSTR},
		},
	},
}

func TestUserConnectionFromArray(t *testing.T) {
	empty := UserConnectionFromArray(tableUserConnectionFromArray[0].nodes, nil)
	if empty != nil {
		t.Errorf("Expected nil output for nil args")
	}

	for i, e := range tableUserConnectionFromArray {
		out := UserConnectionFromArray(e.nodes, e.args)
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

var tableUserEdgesToReturn = []struct {
	edges  []*UserEdge
	before *string
	after  *string
	first  *int
	last   *int
	out    *UserConnection
}{
	{edges: []*UserEdge{}, before: nil, after: nil, first: &ZERO, last: &ZERO, out: tableUserConnectionFromArray[0].out},
	{edges: edgesUser, before: nil, after: nil, first: &ZERO, last: &ZERO, out: tableUserConnectionFromArray[1].out},
	{edges: edgesUser, before: nil, after: nil, first: &TWO, last: &ZERO, out: tableUserConnectionFromArray[2].out},
	{edges: edgesUser, before: nil, after: nil, first: &ZERO, last: &TWO, out: tableUserConnectionFromArray[3].out},
	{edges: edgesUser, before: &FOURSTR, after: &TWOSTR, first: &ZERO, last: &ZERO, out: tableUserConnectionFromArray[4].out},
}

func TestUserEdgesToReturn(t *testing.T) {
	for i, e := range tableUserEdgesToReturn {
		out := userEdgesToReturn(e.edges, e.before, e.after, e.first, e.last)
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

var tableUserApplyCursorsToEdges = []struct {
	edges  []*UserEdge
	before *string
	after  *string
	len    int
}{
	{edges: []*UserEdge{}, before: nil, after: nil, len: 0},
	{edges: edgesUser, before: nil, after: nil, len: 5},
	{edges: edgesUser, before: nil, after: &TWOSTR, len: 4},
	{edges: edgesUser, before: &FOURSTR, after: nil, len: 4},
	{edges: edgesUser, before: &FOURSTR, after: &TWOSTR, len: 3},
}

func TestUserApplyCursorsToEdges(t *testing.T) {
	for i, e := range tableUserApplyCursorsToEdges {
		out := userApplyCursorsToEdges(e.edges, e.before, e.after)
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

var tableUserHasPreviousPage = []struct {
	edges  []*UserEdge
	before *string
	after  *string
	last   *int
	out    bool
}{
	{edges: []*UserEdge{}, before: nil, after: nil, last: nil, out: false},
	{edges: edgesUser, before: nil, after: nil, last: nil, out: false},
	{edges: edgesUser, before: nil, after: nil, last: &ZERO, out: false},
	{edges: edgesUser, before: nil, after: nil, last: &SIX, out: false},
	{edges: edgesUser, before: nil, after: nil, last: &FOUR, out: true},
	{edges: edgesUser, before: &FOURSTR, after: &TWOSTR, last: &FOUR, out: false},
	{edges: edgesUser, before: &FOURSTR, after: &TWOSTR, last: &ONE, out: true},
}

func TestUserHasPreviousPage(t *testing.T) {
	for i, e := range tableUserHasPreviousPage {
		out := userHasPreviousPage(e.edges, e.before, e.after, e.last)
		if out != e.out {
			t.Errorf("%d: got %v, want %v", i, out, e.out)
		}
	}
}

var tableUserHasNextPage = []struct {
	edges  []*UserEdge
	before *string
	after  *string
	first  *int
	out    bool
}{
	{edges: []*UserEdge{}, before: nil, after: nil, first: nil, out: false},
	{edges: edgesUser, before: nil, after: nil, first: nil, out: false},
	{edges: edgesUser, before: nil, after: nil, first: &ZERO, out: false},
	{edges: edgesUser, before: nil, after: nil, first: &SIX, out: false},
	{edges: edgesUser, before: nil, after: nil, first: &FOUR, out: true},
	{edges: edgesUser, before: &FOURSTR, after: &TWOSTR, first: &FOUR, out: false},
	{edges: edgesUser, before: &FOURSTR, after: &TWOSTR, first: &ONE, out: true},
}

func TestUserHasNextPage(t *testing.T) {
	for i, e := range tableUserHasNextPage {
		out := userHasNextPage(e.edges, e.before, e.after, e.first)
		if out != e.out {
			t.Errorf("%d: got %v, want %v", i, out, e.out)
		}
	}
}
