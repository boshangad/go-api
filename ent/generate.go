package ent

//go:generate go run -mod=mod entgo.io/ent/cmd/ent generate --template glob=../app/model/templates/*.tmpl --feature sql/modifier --feature sql/lock,sql/schemaconfig,sql/upsert ./schema
