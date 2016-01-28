// generated by gen-mocks; DO NOT EDIT

package mockstore

import (
	"golang.org/x/net/context"
	"src.sourcegraph.com/sourcegraph/store"
)

type RepoPerms struct {
	Add_           func(ctx context.Context, uid int32, repo string) error
	Update_        func(ctx context.Context, uid int32, repos []string) error
	Delete_        func(ctx context.Context, uid int32, repo string) error
	ListUserRepos_ func(ctx context.Context, uid int32) ([]string, error)
	DeleteUser_    func(ctx context.Context, uid int32) error
}

func (s *RepoPerms) Add(ctx context.Context, uid int32, repo string) error {
	return s.Add_(ctx, uid, repo)
}

func (s *RepoPerms) Update(ctx context.Context, uid int32, repos []string) error {
	return s.Update_(ctx, uid, repos)
}

func (s *RepoPerms) Delete(ctx context.Context, uid int32, repo string) error {
	return s.Delete_(ctx, uid, repo)
}

func (s *RepoPerms) ListUserRepos(ctx context.Context, uid int32) ([]string, error) {
	return s.ListUserRepos_(ctx, uid)
}

func (s *RepoPerms) DeleteUser(ctx context.Context, uid int32) error { return s.DeleteUser_(ctx, uid) }

var _ store.RepoPerms = (*RepoPerms)(nil)
