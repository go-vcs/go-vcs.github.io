// A Hugo site has no Go code, so `go mod tidy` -- which Renovate runs after
// every bump -- would find nothing that uses github.com/go-fleettools/landingkit
// and delete its requirement from go.mod. Importing it here keeps it. See the
// landingkit README.
//
// Unlike the sibling sites this file carries no `tools` build constraint: a
// constrained file is one no CI configuration compiles, which the buildscan gate
// refuses, and the blank import costs nothing to compile.
package site

import _ "github.com/go-fleettools/landingkit"
