package tests

import (
	"github.com/bozaro/tech-db-forum/generated/client"
	"github.com/bozaro/tech-db-forum/generated/client/operations"
	"strings"
)

func init() {
	Register(Checker{
		Name:        "forum_get_one_simple",
		Description: "",
		FnCheck:     CheckForumGetOneSimple,
		Deps: []string{
			"forum_create_simple",
		},
	})
	Register(Checker{
		Name:        "forum_get_one_notfound",
		Description: "",
		FnCheck:     CheckForumGetOneNotFound,
		Deps: []string{
			"forum_get_one_simple",
		},
	})
	Register(Checker{
		Name:        "forum_get_one_nocase",
		Description: "",
		FnCheck:     CheckForumGetOneNocase,
		Deps: []string{
			"forum_get_one_simple",
		},
	})
}

func CheckForumGetOneSimple(c *client.Forum) {
	forum := CreateForum(c, nil, nil)
	CheckForum(c, forum)
}

func CheckForumGetOneNotFound(c *client.Forum) {
	forum := RandomForum()
	_, err := c.Operations.ForumGetOne(operations.NewForumGetOneParams().
		WithSlug(forum.Slug).
		WithContext(Expected(404, nil, nil)))
	CheckIsType(operations.NewForumGetOneNotFound(), err)
}

func CheckForumGetOneNocase(c *client.Forum) {
	forum := CreateForum(c, nil, nil)
	_, err := c.Operations.ForumGetOne(operations.NewForumGetOneParams().
		WithSlug(strings.ToLower(forum.Slug)).
		WithContext(Expected(200, forum, nil)))
	CheckNil(err)
}
