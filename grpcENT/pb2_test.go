package main

import (
	"grpcENT/ent/category"
	"grpcENT/ent/proto/entpb"
	"grpcENT/ent/user"
	"testing"

	"context"

	_ "github.com/mattn/go-sqlite3"

	"grpcENT/ent/enttest"
)

func TestServiceWithEdges(t *testing.T) {
	// start by initializing an ent client connected to an in memory sqlite instance
	ctx := context.Background()
	client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
	defer client.Close()

	// next, initialize the UserService. Notice we won't be opening an actual port and
	// creating a gRPC server and instead we are just calling the library code directly.
	svc := entpb.NewUserService(client)

	// next, we create a category directly using the ent client.
	// Notice we are initializing it with no relation to a User.
	cat := client.Category.Create().SetName("cat_1").SaveX(ctx)

	// next, we invoke the User service's `Create` method. Notice we are
	// passing a list of entpb.Category instances with only the ID set.
	create, err := svc.Create(ctx, &entpb.CreateUserRequest{
		User: &entpb.User{
			Name:         "user",
			EmailAddress: "user@service.code",
			Administered: []*entpb.Category{
				{Id: int64(cat.ID)},
			},
		},
	})
	if err != nil {
		t.Fatal("failed creating user using UserService", err)
	}

	// to verify everything worked correctly, we query the category table to check
	// we have exactly one category which is administered by the created user.
	count, err := client.Category.
		Query().
		Where(
			category.HasAdminWith(
				user.ID(int(create.Id)),
			),
		).
		Count(ctx)
	if err != nil {
		t.Fatal("failed counting categories admin by created user", err)
	}
	if count != 1 {
		t.Fatal("expected exactly one group to managed by the created user")
	}
}
