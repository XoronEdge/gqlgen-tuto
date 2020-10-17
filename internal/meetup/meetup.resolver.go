package meetup

import (
	"context"
	"errors"

	"github.com/xoronedge/gqlgen-tuto/graph"
	"github.com/xoronedge/gqlgen-tuto/graph/models"
)

var meetups = []*models.Meetup{
	{
		ID:          "1",
		Name:        "A meetup in Newyork",
		Description: "Well",
		UserID:      "1",
	},
	{
		ID:          "2",
		Name:        "A meetup in NewJersy",
		Description: "Excellent",
		UserID:      "2",
	},
	{
		ID:          "3",
		Name:        "A meetup in Newhampshire",
		Description: "Amaze",
		UserID:      "1",
	},
}
var users = []models.User{
	{
		ID:       "1",
		Username: "lala",
		Email:    "lala@gmail.com",
	},
}

type meetupResolver struct{ *graph.Resolver }

func (r *meetupResolver) User(ctx context.Context, obj *models.Meetup) (*models.User, error) {
	user := new(models.User)

	for _, u := range users {
		if u.ID == obj.UserID {
			user = &u
			break
		}
	}

	if user == nil {
		return nil, errors.New("user with id not exist")
	}
	return user, nil
}
