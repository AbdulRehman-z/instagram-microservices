// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.22.0

package db

import (
	"context"

	"github.com/google/uuid"
)

type Querier interface {
	FollowUser(ctx context.Context, arg FollowUserParams) (Follower, error)
	GetFollowers(ctx context.Context, arg GetFollowersParams) ([]uuid.UUID, error)
	GetFollowersCount(ctx context.Context, leaderUniqueID uuid.UUID) (int64, error)
	GetFollowing(ctx context.Context, arg GetFollowingParams) ([]uuid.UUID, error)
	GetFollowingAndFollowersCount(ctx context.Context, leaderUniqueID uuid.UUID) (GetFollowingAndFollowersCountRow, error)
	GetFollowingCount(ctx context.Context, followerUniqueID uuid.UUID) (int64, error)
	UnfollowUser(ctx context.Context, arg UnfollowUserParams) (Follower, error)
}

var _ Querier = (*Queries)(nil)
