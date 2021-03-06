// Code generated by github.com/oreqizer/go-relaygen, DO NOT EDIT.

package example

import (
	"testing"

	"github.com/oreqizer/go-relaygen/relay"
)

var (
	useronestr   = "1"
	usertwostr   = "2"
	userthreestr = "3"
	userfourstr  = "4"
	userfivestr  = "5"

	// Set later
	useroneid   = ""
	usertwoid   = ""
	userthreeid = ""
	userfourid  = ""
	userfiveid  = ""
)

var (
	userzero = 0
	userone  = 1
	usertwo  = 2
	// userthree = 3
	userfour = 4
	// userfive  = 5
	usersix = 6
)

var nodesUser = []*User{
	&User{},
	&User{},
	&User{},
	&User{},
	&User{},
}

var edgesUser = []*UserEdge{
	{Node: &User{}, Cursor: ""},
	{Node: &User{}, Cursor: ""},
	{Node: &User{}, Cursor: ""},
	{Node: &User{}, Cursor: ""},
	{Node: &User{}, Cursor: ""},
}

func init() {
	nodesUser[0].LocalID = useronestr
	nodesUser[1].LocalID = usertwostr
	nodesUser[2].LocalID = userthreestr
	nodesUser[3].LocalID = userfourstr
	nodesUser[4].LocalID = userfivestr

	useroneid = nodesUser[0].ID()
	usertwoid = nodesUser[1].ID()
	userthreeid = nodesUser[2].ID()
	userfourid = nodesUser[3].ID()
	userfiveid = nodesUser[4].ID()

	edgesUser[0].Node.LocalID = useronestr
	edgesUser[0].Cursor = useroneid
	edgesUser[1].Node.LocalID = usertwostr
	edgesUser[1].Cursor = usertwoid
	edgesUser[2].Node.LocalID = userthreestr
	edgesUser[2].Cursor = userthreeid
	edgesUser[3].Node.LocalID = userfourstr
	edgesUser[3].Cursor = userfourid
	edgesUser[4].Node.LocalID = userfivestr
	edgesUser[4].Cursor = userfiveid
}

var tableUserConnectionFromArray = []struct {
	nodes []*User
	args  *relay.ConnectionArgs
	out   *UserConnection
}{
	{
		nodes: []*User{},
		args:  &relay.ConnectionArgs{Before: nil, After: nil, First: &userzero, Last: &userzero},
		out: &UserConnection{
			Edges:    []*UserEdge{},
			PageInfo: relay.PageInfo{HasNextPage: false, HasPreviousPage: false, StartCursor: nil, EndCursor: nil},
		},
	},
	{
		nodes: nodesUser,
		args:  &relay.ConnectionArgs{Before: nil, After: nil, First: &userzero, Last: &userzero},
		out: &UserConnection{
			Edges:    edgesUser,
			PageInfo: relay.PageInfo{HasNextPage: false, HasPreviousPage: false, StartCursor: &useroneid, EndCursor: &userfiveid},
		},
	},
	{
		nodes: nodesUser,
		args:  &relay.ConnectionArgs{Before: nil, After: nil, First: &usertwo, Last: &userzero},
		out: &UserConnection{
			Edges:    edgesUser[:2],
			PageInfo: relay.PageInfo{HasNextPage: true, HasPreviousPage: false, StartCursor: &useroneid, EndCursor: &usertwoid},
		},
	},
	{
		nodes: nodesUser,
		args:  &relay.ConnectionArgs{Before: nil, After: nil, First: &userzero, Last: &usertwo},
		out: &UserConnection{
			Edges:    edgesUser[3:],
			PageInfo: relay.PageInfo{HasNextPage: false, HasPreviousPage: true, StartCursor: &userfourid, EndCursor: &userfiveid},
		},
	},
	{
		nodes: nodesUser,
		args:  &relay.ConnectionArgs{Before: &userfourid, After: &usertwoid, First: &userzero, Last: &userzero},
		out: &UserConnection{
			Edges:    edgesUser[1:4],
			PageInfo: relay.PageInfo{HasNextPage: false, HasPreviousPage: false, StartCursor: &usertwoid, EndCursor: &userfourid},
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
	{edges: []*UserEdge{}, before: nil, after: nil, first: &userzero, last: &userzero, out: tableUserConnectionFromArray[0].out},
	{edges: edgesUser, before: nil, after: nil, first: &userzero, last: &userzero, out: tableUserConnectionFromArray[1].out},
	{edges: edgesUser, before: nil, after: nil, first: &usertwo, last: &userzero, out: tableUserConnectionFromArray[2].out},
	{edges: edgesUser, before: nil, after: nil, first: &userzero, last: &usertwo, out: tableUserConnectionFromArray[3].out},
	{edges: edgesUser, before: &userfourid, after: &usertwoid, first: &userzero, last: &userzero, out: tableUserConnectionFromArray[4].out},
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
	{edges: edgesUser, before: nil, after: &usertwoid, len: 4},
	{edges: edgesUser, before: &userfourid, after: nil, len: 4},
	{edges: edgesUser, before: &userfourid, after: &usertwoid, len: 3},
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
	{edges: edgesUser, before: nil, after: nil, last: &userzero, out: false},
	{edges: edgesUser, before: nil, after: nil, last: &usersix, out: false},
	{edges: edgesUser, before: nil, after: nil, last: &userfour, out: true},
	{edges: edgesUser, before: &userfourid, after: &usertwoid, last: &userfour, out: false},
	{edges: edgesUser, before: &userfourid, after: &usertwoid, last: &userone, out: true},
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
	{edges: edgesUser, before: nil, after: nil, first: &userzero, out: false},
	{edges: edgesUser, before: nil, after: nil, first: &usersix, out: false},
	{edges: edgesUser, before: nil, after: nil, first: &userfour, out: true},
	{edges: edgesUser, before: &userfourid, after: &usertwoid, first: &userfour, out: false},
	{edges: edgesUser, before: &userfourid, after: &usertwoid, first: &userone, out: true},
}

func TestUserHasNextPage(t *testing.T) {
	for i, e := range tableUserHasNextPage {
		out := userHasNextPage(e.edges, e.before, e.after, e.first)
		if out != e.out {
			t.Errorf("%d: got %v, want %v", i, out, e.out)
		}
	}
}
