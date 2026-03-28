# tl

SQLite-backed CLI for managing task lists. Built for AI agent orchestration — an orchestrator creates a list, sub-agents pull tasks via atomic `next`.

## Structure

- `cmd/` — Cobra CLI commands (create, add, next, done, fail, status, list, get, reset)
- `db/` — SQLite layer: schema, list/task CRUD, atomic claim logic

## Stack

Go, Cobra, modernc.org/sqlite (pure Go, no CGO)
