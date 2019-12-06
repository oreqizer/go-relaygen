# RelayGen

[![godoc](https://img.shields.io/badge/godoc-reference-blue.svg)](https://godoc.org/github.com/oreqizer/go-relaygen)
[![Build Status](https://travis-ci.org/oreqizer/go-relaygen.svg?branch=master)](https://travis-ci.org/oreqizer/go-relaygen)
[![codecov](https://codecov.io/gh/oreqizer/go-relaygen/branch/master/graph/badge.svg)](https://codecov.io/gh/oreqizer/go-relaygen)

A `go generate` based toolset for [Relay](https://facebook.github.io/relay/docs/en/graphql-server-specification.html) compliant **GraphQL** servers.

> An evolution of [go-relay](https://github.com/oreqizer/go-relay) project.

It ain't much but it's honest work.

## API

Types and functions are located in `github.com/oreqizer/go-relaygen/relay`.

### Connection

Run `go run github.com/oreqizer/go-relaygen <name> <local id>` where you want to generate your connection utilities.

> Supports nested ID fields, e.g. `Person.ID`

The only function you care about is `<Name>ConnectionFromArray`.

Make your types satisfy the `Node` interface and create the `ConnectionArgs` object, feed it into it and you'll get a `<Name>Connection`.

```go
//go:generate go run github.com/oreqizer/go-relaygen User LocalID

package example

import (
	"github.com/oreqizer/go-relaygen/relay"
)

const UserType = "User"

type User struct {
	LocalID string
	Name    string
}

func (u *User) ID() string {
	return relay.ToGlobalID(UserType, u.LocalID)
}
```

### IDs

There are two functions - `ToGlobalID` and `FromGlobalID`. They behave the same like the JS reference implementation.

```go
var global = relay.ToGlobalID("User", "asdf") // Returns a base64 encoded string

var local = relay.FromGlobalID(global) // local.Type == "User", local.ID == "asdf"
```

## License

MIT
