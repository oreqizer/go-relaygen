// Code generated by github.com/oreqizer/go-relaygen, DO NOT EDIT.

package example

import (
	"testing"

	"github.com/oreqizer/go-relaygen/relay"
)

var (
	adminonestr   = "1"
	admintwostr   = "2"
	adminthreestr = "3"
	adminfourstr  = "4"
	adminfivestr  = "5"

	// Set later
	adminoneid   = ""
	admintwoid   = ""
	adminthreeid = ""
	adminfourid  = ""
	adminfiveid  = ""
)

var (
	adminzero = 0
	adminone  = 1
	admintwo  = 2
	// adminthree = 3
	adminfour = 4
	// adminfive  = 5
	adminsix = 6
)

var nodesAdmin = []*Admin{
	&Admin{},
	&Admin{},
	&Admin{},
	&Admin{},
	&Admin{},
}

var edgesAdmin = []*AdminEdge{
	{Node: &Admin{}, Cursor: ""},
	{Node: &Admin{}, Cursor: ""},
	{Node: &Admin{}, Cursor: ""},
	{Node: &Admin{}, Cursor: ""},
	{Node: &Admin{}, Cursor: ""},
}

func init() {
	nodesAdmin[0].User.LocalID = adminonestr
	nodesAdmin[1].User.LocalID = admintwostr
	nodesAdmin[2].User.LocalID = adminthreestr
	nodesAdmin[3].User.LocalID = adminfourstr
	nodesAdmin[4].User.LocalID = adminfivestr

	adminoneid = nodesAdmin[0].ID()
	admintwoid = nodesAdmin[1].ID()
	adminthreeid = nodesAdmin[2].ID()
	adminfourid = nodesAdmin[3].ID()
	adminfiveid = nodesAdmin[4].ID()

	edgesAdmin[0].Node.User.LocalID = adminonestr
	edgesAdmin[0].Cursor = adminoneid
	edgesAdmin[1].Node.User.LocalID = admintwostr
	edgesAdmin[1].Cursor = admintwoid
	edgesAdmin[2].Node.User.LocalID = adminthreestr
	edgesAdmin[2].Cursor = adminthreeid
	edgesAdmin[3].Node.User.LocalID = adminfourstr
	edgesAdmin[3].Cursor = adminfourid
	edgesAdmin[4].Node.User.LocalID = adminfivestr
	edgesAdmin[4].Cursor = adminfiveid
}

var tableAdminConnectionFromArray = []struct {
	nodes []*Admin
	args  *relay.ConnectionArgs
	out   *AdminConnection
}{
	{
		nodes: []*Admin{},
		args:  &relay.ConnectionArgs{Before: nil, After: nil, First: &adminzero, Last: &adminzero},
		out: &AdminConnection{
			Edges:    []*AdminEdge{},
			PageInfo: relay.PageInfo{HasNextPage: false, HasPreviousPage: false, StartCursor: nil, EndCursor: nil},
		},
	},
	{
		nodes: nodesAdmin,
		args:  &relay.ConnectionArgs{Before: nil, After: nil, First: &adminzero, Last: &adminzero},
		out: &AdminConnection{
			Edges:    edgesAdmin,
			PageInfo: relay.PageInfo{HasNextPage: false, HasPreviousPage: false, StartCursor: &adminoneid, EndCursor: &adminfiveid},
		},
	},
	{
		nodes: nodesAdmin,
		args:  &relay.ConnectionArgs{Before: nil, After: nil, First: &admintwo, Last: &adminzero},
		out: &AdminConnection{
			Edges:    edgesAdmin[:2],
			PageInfo: relay.PageInfo{HasNextPage: true, HasPreviousPage: false, StartCursor: &adminoneid, EndCursor: &admintwoid},
		},
	},
	{
		nodes: nodesAdmin,
		args:  &relay.ConnectionArgs{Before: nil, After: nil, First: &adminzero, Last: &admintwo},
		out: &AdminConnection{
			Edges:    edgesAdmin[3:],
			PageInfo: relay.PageInfo{HasNextPage: false, HasPreviousPage: true, StartCursor: &adminfourid, EndCursor: &adminfiveid},
		},
	},
	{
		nodes: nodesAdmin,
		args:  &relay.ConnectionArgs{Before: &adminfourid, After: &admintwoid, First: &adminzero, Last: &adminzero},
		out: &AdminConnection{
			Edges:    edgesAdmin[1:4],
			PageInfo: relay.PageInfo{HasNextPage: false, HasPreviousPage: false, StartCursor: &admintwoid, EndCursor: &adminfourid},
		},
	},
}

func TestAdminConnectionFromArray(t *testing.T) {
	empty := AdminConnectionFromArray(tableAdminConnectionFromArray[0].nodes, nil)
	if empty != nil {
		t.Errorf("Expected nil output for nil args")
	}

	for i, e := range tableAdminConnectionFromArray {
		out := AdminConnectionFromArray(e.nodes, e.args)
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

var tableAdminEdgesToReturn = []struct {
	edges  []*AdminEdge
	before *string
	after  *string
	first  *int
	last   *int
	out    *AdminConnection
}{
	{edges: []*AdminEdge{}, before: nil, after: nil, first: &adminzero, last: &adminzero, out: tableAdminConnectionFromArray[0].out},
	{edges: edgesAdmin, before: nil, after: nil, first: &adminzero, last: &adminzero, out: tableAdminConnectionFromArray[1].out},
	{edges: edgesAdmin, before: nil, after: nil, first: &admintwo, last: &adminzero, out: tableAdminConnectionFromArray[2].out},
	{edges: edgesAdmin, before: nil, after: nil, first: &adminzero, last: &admintwo, out: tableAdminConnectionFromArray[3].out},
	{edges: edgesAdmin, before: &adminfourid, after: &admintwoid, first: &adminzero, last: &adminzero, out: tableAdminConnectionFromArray[4].out},
}

func TestAdminEdgesToReturn(t *testing.T) {
	for i, e := range tableAdminEdgesToReturn {
		out := adminEdgesToReturn(e.edges, e.before, e.after, e.first, e.last)
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

var tableAdminApplyCursorsToEdges = []struct {
	edges  []*AdminEdge
	before *string
	after  *string
	len    int
}{
	{edges: []*AdminEdge{}, before: nil, after: nil, len: 0},
	{edges: edgesAdmin, before: nil, after: nil, len: 5},
	{edges: edgesAdmin, before: nil, after: &admintwoid, len: 4},
	{edges: edgesAdmin, before: &adminfourid, after: nil, len: 4},
	{edges: edgesAdmin, before: &adminfourid, after: &admintwoid, len: 3},
}

func TestAdminApplyCursorsToEdges(t *testing.T) {
	for i, e := range tableAdminApplyCursorsToEdges {
		out := adminApplyCursorsToEdges(e.edges, e.before, e.after)
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

var tableAdminHasPreviousPage = []struct {
	edges  []*AdminEdge
	before *string
	after  *string
	last   *int
	out    bool
}{
	{edges: []*AdminEdge{}, before: nil, after: nil, last: nil, out: false},
	{edges: edgesAdmin, before: nil, after: nil, last: nil, out: false},
	{edges: edgesAdmin, before: nil, after: nil, last: &adminzero, out: false},
	{edges: edgesAdmin, before: nil, after: nil, last: &adminsix, out: false},
	{edges: edgesAdmin, before: nil, after: nil, last: &adminfour, out: true},
	{edges: edgesAdmin, before: &adminfourid, after: &admintwoid, last: &adminfour, out: false},
	{edges: edgesAdmin, before: &adminfourid, after: &admintwoid, last: &adminone, out: true},
}

func TestAdminHasPreviousPage(t *testing.T) {
	for i, e := range tableAdminHasPreviousPage {
		out := adminHasPreviousPage(e.edges, e.before, e.after, e.last)
		if out != e.out {
			t.Errorf("%d: got %v, want %v", i, out, e.out)
		}
	}
}

var tableAdminHasNextPage = []struct {
	edges  []*AdminEdge
	before *string
	after  *string
	first  *int
	out    bool
}{
	{edges: []*AdminEdge{}, before: nil, after: nil, first: nil, out: false},
	{edges: edgesAdmin, before: nil, after: nil, first: nil, out: false},
	{edges: edgesAdmin, before: nil, after: nil, first: &adminzero, out: false},
	{edges: edgesAdmin, before: nil, after: nil, first: &adminsix, out: false},
	{edges: edgesAdmin, before: nil, after: nil, first: &adminfour, out: true},
	{edges: edgesAdmin, before: &adminfourid, after: &admintwoid, first: &adminfour, out: false},
	{edges: edgesAdmin, before: &adminfourid, after: &admintwoid, first: &adminone, out: true},
}

func TestAdminHasNextPage(t *testing.T) {
	for i, e := range tableAdminHasNextPage {
		out := adminHasNextPage(e.edges, e.before, e.after, e.first)
		if out != e.out {
			t.Errorf("%d: got %v, want %v", i, out, e.out)
		}
	}
}
