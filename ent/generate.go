package ent

//go:generate go run -mod=mod ./ent/entc.go
//go:generate go run -mod=mod entgo.io/ent/cmd/ent generate --feature sql/versioned-migration ./schema
