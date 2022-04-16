package graphql

import (
	"context"

	"github.com/hsblhsn/hn.hsblhsn.me/backend/graphql/internal/hackernews"
	"github.com/hsblhsn/hn.hsblhsn.me/backend/graphql/internal/relays"
	"github.com/hsblhsn/hn.hsblhsn.me/backend/graphql/model"
	"github.com/pkg/errors"
)

func (r *Resolver) NewRelayComments(ctx context.Context, ids []int) *relays.Resolver[int, *model.Comment] {
	return relays.NewResolver(ids, func(id int) (*model.Comment, error) {
		idStr := hackernews.NewID(id)
		out, err := r.hackerNews.GetTypedItem(ctx, hackernews.ItemTypeComment, idStr)
		if err != nil {
			return nil, errors.Wrap(err, "resolver: could not get comment")
		}
		return &model.Comment{ItemResponse: out}, nil
	})
}

func (r *Resolver) NewRelayPollOptions(ctx context.Context, ids []int) *relays.Resolver[int, *model.PollOption] {
	return relays.NewResolver(ids, func(id int) (*model.PollOption, error) {
		idStr := hackernews.NewID(id)
		out, err := r.hackerNews.GetTypedItem(ctx, hackernews.ItemTypePollOption, idStr)
		if err != nil {
			return nil, errors.Wrap(err, "graph: could not get poll option")
		}
		return &model.PollOption{ItemResponse: out}, nil
	})
}

func (r *Resolver) NewRelayStories(ctx context.Context, ids []int) *relays.Resolver[int, *model.Story] {
	return relays.NewResolver(ids, func(id int) (*model.Story, error) {
		idStr := hackernews.NewID(id)
		out, err := r.hackerNews.GetTypedItem(ctx, hackernews.ItemTypeStory, idStr)
		if err != nil {
			return nil, errors.Wrap(err, "graph: could not get story")
		}
		return &model.Story{ItemResponse: out}, nil
	})
}

func (r *Resolver) NewRelayJobs(ctx context.Context, ids []int) *relays.Resolver[int, *model.Job] {
	return relays.NewResolver(ids, func(id int) (*model.Job, error) {
		idStr := hackernews.NewID(id)
		out, err := r.hackerNews.GetTypedItem(ctx, hackernews.ItemTypeJob, idStr)
		if err != nil {
			return nil, errors.Wrap(err, "graph: could not get job")
		}
		return &model.Job{ItemResponse: out}, nil
	})
}
