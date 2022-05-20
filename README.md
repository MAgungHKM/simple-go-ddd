# Domain Driven Design on Go

This is just to pique my interest of DDD using GO

## Trying Yourself

- git clone this repo
- `cd` to the cloned repo directory
- run `go run .` or `go run main.go` on terminal and you should see some outputs from my random testing

## Folder Structure

```
├── core
│   ├── books
│   │   ├── book.go
│   │   └── book_repo.go
│   ├── borrow
│   │   ├── borrow.go
│   │   ├── borrow_repo.go
│   │   └── borrow_service.go
│   ├── member
│   │   ├── member.go
│   │   └── member_repo.go
│   └── status
│       └── status.go
├── infra # could've renamed it to db. your preferences.
│   └── mongodb # you could add other db implementation here.
│   └── inmemory # for the sake of simplicity.
│       ├── book_repo.go
│       ├── borrow_repo.go
│       └── member_repo.go
├── go.mod
└── main.go
```