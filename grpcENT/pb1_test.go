package main

import (
	"grpcENT/ent/proto/entpb"
	"testing"

	_ "github.com/mattn/go-sqlite3"
)

func TestUserProto(t *testing.T) {
	user := entpb.User{
		Name:         "rotemtam",
		EmailAddress: "rotemtam@example.com",
	}
	if user.GetName() != "rotemtam" {
		t.Fatal("expected user name to be rotemtam")
	}
	if user.GetEmailAddress() != "rotemtam@example.com" {
		t.Fatal("expected email address to be rotemtam@example.com")
	}
}
