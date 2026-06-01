# blog-aggregator

A CLI tool for aggregating and following RSS feeds, backed by a PostgreSQL database.

## Prerequisites

- Go 1.24+
- PostgreSQL
- [goose](https://github.com/pressly/goose) for migrations
- [sqlc](https://sqlc.dev) if regenerating database code

## Setup

1. Create a PostgreSQL database (e.g. `gator`).
2. Create a config file at `~/.gatorconfig.json`:
   ```json
   {
     "db_url": "postgres://user:password@localhost:5432/gator?sslmode=disable"
   }
   ```
3. Run migrations:
   ```bash
   goose -dir sql/schema postgres "postgres://user:password@localhost:5432/gator?sslmode=disable" up
   ```

## Installation

```bash
go install
```

Or run directly with `go run . <command>`.

## Commands

### User management

| Command | Arguments | Description |
|---|---|---|
| `register` | `<name>` | Create a new user and set them as the current user |
| `login` | `<name>` | Switch to an existing user |
| `users` | | List all users (current user is marked) |
| `reset` | | Delete all users and their data |

### Feeds

| Command | Arguments | Description |
|---|---|---|
| `addfeed` | `<name> <url>` | Add a new feed and automatically follow it |
| `feeds` | | List all feeds and their owners |

### Following

| Command | Arguments | Description |
|---|---|---|
| `follow` | `<url>` | Follow an existing feed by URL |
| `unfollow` | `<url>` | Unfollow a feed by URL |
| `following` | | List all feeds the current user is following |

### Aggregation

| Command | Arguments | Description |
|---|---|---|
| `agg` | | Fetch and display the latest posts |

## Example usage

```bash
# Register a user
go run . register alice

# Add and follow a feed
go run . addfeed "Hacker News" "https://hnrss.org/newest"

# Follow another feed
go run . follow "https://www.wagslane.dev/index.xml"

# List followed feeds
go run . following

# Unfollow a feed
go run . unfollow "https://hnrss.org/newest"
```
