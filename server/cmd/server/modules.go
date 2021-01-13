//+build wireinject

package main

import (
	"github.com/EugeniaKol/forums_system/server/forums"
	"github.com/google/wire"
)

//ComposeAPIServer will create an instance of CharApiServer according to providers defined in this file.
func ComposeAPIServer(port HTTPPortNumber) (*ForumAPIServer, error) {
	wire.Build(
		// DB connection provider (defined in main.go).
		NewDbConnection,

		// Add providers from channels package.
		forums.Providers,

		// Provide ChatApiServer instantiating the structure and injecting channels handler and port number.
		wire.Struct(new(ForumAPIServer), "Port", "ForumsHandler"),
	)
	return nil, nil
}
