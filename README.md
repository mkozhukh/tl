# tl

SQLite-backed task list manager for AI agents.

An orchestrator creates a list, adds tasks, and sub-agents pull work with `tl next` — atomic claim, no double-processing.

## Install

```bash
go install tl@latest
```

Or build from source:

```bash
git clone <repo> && cd tl
go build -o tl .
```

## Quick start

```bash
# create a list
tl create "migrate-tables"
# {"id":1,"name":"migrate-tables","created_at":"...","updated_at":"..."}

# add tasks
tl add 1 "users table" --meta '{"file":"users.sql"}'
tl add 1 "orders table"

# sub-agent claims next available task (optionally identifying itself)
tl next 1 --owner agent-7
# {"id":1,"title":"users table","status":"active","file":"users.sql","owner":"agent-7"}

# or claim a specific task by ID
tl next 1 2 --owner agent-7

# mark done with result data
tl done 1 --result '{"rows":42}'

# or mark failed
tl fail 2 --reason "connection timeout"

# reset a failed task back to pending
tl reset 2
```

## Commands

```
tl create <name>                             Create a task list
tl add <list-id> <title> [--meta '{}']       Add a task
tl next <list-id> [task-id] [--owner <id>]   Claim next pending task, or a specific one (atomic)
tl done <task-id> [--result '{}']            Mark task done
tl fail <task-id> [--reason "..."]           Mark task failed
tl status <list-id>                          Task counts by state
tl list                                      All lists with counts
tl tasks <list-id> [--status <s>]            List tasks, optionally filtered by status
                   [--owner <id>]            Filter by owner (combinable with --status)
tl get <task-id>                             Get a single task
tl reset <task-id>                           Reset failed/active → pending
```

## Design

- All output is JSON. Agents parse stdout, errors go to stderr.
- Exit codes: `0` success, `1` error, `2` no tasks available (from `next`).
- `next` uses a single atomic SQL statement — safe with concurrent agents.
- Task states: `pending → active → done | failed`.
- DB defaults to `./tl.db`. Override with `TL_DB` env var.
- IDs are auto-increment integers — short and easy to pass around.

## License

MIT
