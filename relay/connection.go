package relay

/*
Node is an interface for relay satisfying Relay's Node interface

https://facebook.github.io/relay/graphql/objectidentification.htm#sec-Node-Interface
*/
type Node interface {
	ID() string
}

/*
PageInfo holds information about connection edges

https://facebook.github.io/relay/graphql/connections.htm#sec-Reserved-Types
*/
type PageInfo struct {
	HasNextPage     bool    `json:"hasNextPage"`
	HasPreviousPage bool    `json:"hasPreviousPage"`
	StartCursor     *string `json:"startCursor"`
	EndCursor       *string `json:"endCursor"`
}

/*
ConnectionArgs holds information about a connection arguments

https://facebook.github.io/relay/graphql/connections.htm#sec-Reserved-Types
*/
type ConnectionArgs struct {
	Before *string `json:"before"`
	After  *string `json:"after"`
	First  *int    `json:"first"`
	Last   *int    `json:"last"`
}
