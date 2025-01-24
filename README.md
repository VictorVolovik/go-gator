# Go Gator

Go-Gator is a lightweight, command-line RSS feed aggregator written in Go, based on Boot.Dev's [Build a Blog Aggregator in Go](https://www.boot.dev/courses/build-blog-aggregator-golang) guided project. It is designed for managing feeds and aggregating updates.

## Prerequisites

- Go version 1.23.4 or later installed
- PostgreSQL version 15
- `$HOME/.gatorconfig.json` file with databe connection URL

```json
{
  "db_url": "postgres://user:password@localhost:5432/dbname"
}
```

## Getting Started

- `go run .`: Run the application
- `go build`: Build an executable (see [Go documentation](https://pkg.go.dev/cmd/go#hdr-Compile_and_install_packages_and_dependencies) for more info)
- `go install` - Installs the application binary as `go-gator` (see [Go documentation](https://pkg.go.dev/cmd/go#hdr-Compile_and_install_packages_and_dependencies) for more info)

## SQL Tools Used

[Goose](https://github.com/pressly/goose): Used for SQL migrations.

While in `sql/schema` directory:

- Up SQL migration

```sh
goose postgres <connection_string> up

```

- Down SQL migration

```sh
goose postgres <connection_string> down

```

[SQLC](https://github.com/sqlc-dev/sqlc): Used to generate Go code for handling database queries.

`sqlc generate` - generates Go code

### Commands

`go-gator register <username>` - Registers a new user with the provided username.

`go-gator login <username>` - Logs in as the specified user.

`go-gator users` - Lists all registered users.

`go-gator addfeed <name> <url>` - Adds a new feed with the given name and URL as the RSS feed source.

`go-gator feeds` - Lists all available feeds.

`go-gator follow <url>` - Follows a feed with the specified URL for the current user.

`go-gator unfollow <url>` - Unfollows a feed with the specified URL for the current user.

`go-gator following` - Lists all feeds currently followed by the user.

`go-gator agg <time_between_reqs>` - Starts the aggregation process, fetching updates for followed feeds at the specified interval (e.g., `10s`, `5m30s`, `8h`).

`go-gator browse <limit - optional>` - Displays aggregated posts from followed feeds, with an optional limit. The default limit is 2.

`go-gator reset` - Resets the application, deleting all stored data.
