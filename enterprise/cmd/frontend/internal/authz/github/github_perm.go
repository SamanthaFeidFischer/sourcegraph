package github

// // permModel is responsible for computing the set of accessible repositories for a given user given facts
// // received from the GitHub API. This data structure is agnostic to where these facts come from (in practice
// // they can either come directly from an API call or from cache).
// //
// // This is not threadsafe.
// type permModel struct {
// 	// set of public repos
// 	publicRepos map[api.RepoName]struct{}

// 	// set of repos the user explicitly has access to, verified by an explicit fetch of the repo
// 	// from the GitHub API with the user's auth credentials.
// 	userRepos map[api.RepoName]struct{}

// 	// // set of repos affiliated with the user, verified by a
// 	// // `/user/repos?affiliation=owner,collaborator,organization_member&visibility=private` API call
// 	// // to the GitHub API.
// 	// affiliatedRepos map[api.RepoName]struct{}

// 	// the repositories for which permissions remain to be computed;
// 	// this is synonymous with the set of repositories that we haven't yet
// 	// determined to be accessible
// 	remaining map[authz.Repo]struct{}

// 	// the (potentially partially constructed) permissions
// 	perms map[api.RepoName]map[authz.Perm]bool
// }

// func newPermModel(reposToVerify map[authz.Repo]struct{}) *PermModel {
// 	return &PermModel{
// 		// TODO
// 	}
// }

// func (p *PermModel) updatePublicRepos(publicRepos []authz.Repo) {
// 	// TODO
// }

// // TODO(beyang): should hit the GraphQL API for this...
// func (p *PermModel) updateUserRepos(userRepos []authz.Repo) {
// 	// TODO
// }

// // func (p *PermModel) updateAffiliatedRepos(affiliatedRepos []authz.Repo) {
// // 	// TODO
// // }

// func (p *PermModel) getPerms() (perms map[api.RepoName]map[authz.Perm]bool, complete bool) {
// 	return p.perms, len(p.remaining) == 0
// }
