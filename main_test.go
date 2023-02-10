package main

import (
	"fmt"
	"os"
	"strings"
	"testing"
)

func TestQueryf(t *testing.T) {
	query := queryf(`{
		repository(owner: "excing", name: "find-roots-of-word") {
			discussionCategories(first: 10) {
				nodes {
					id 
					name
					emoji
					description
				}
				totalCount
			}
		}
		viewer {
			login
		}
	}`)
	fmt.Println(query)
	// {"query": "query { repository(owner: \"excing\", name: \"find-roots-of-word\") { discussionCategories(first: 10) { nodes { id name emoji description } totalCount } } viewer { login } }" }
}

func TestRender(t *testing.T) {
	err := render(
		testRepository(),
		"assets/theme",
		true,
		func(s string, b []byte) error {
			fmt.Println(s)
			_, err := os.Stdout.Write(b)
			return err
		},
	)
	if err != nil {
		t.Fatal(err)
	}
}

func TestOsStat(t *testing.T) {
	if _, err := os.Stat(""); os.IsNotExist(err) {
		t.Fatal(err)
	} else {
		t.Log("PASS")
	}
}

func TestGetEmoji4GEmoji(t *testing.T) {
	gemojiFormat := `<div></div>`
	gemoji := `<g-emoji class="g-emoji" alias="mega" fallback-src="https://github.githubassets.com/images/icons/emoji/unicode/1f4e3.png">📣</g-emoji>`
	t.Log(getGemoji(fmt.Sprintf(gemojiFormat, gemoji)) == gemoji)
}

func TestPref(t *testing.T) {
	str := "example.txt"
	suffix := "txt"
	t.Log(strings.HasSuffix(str, suffix))
}

func testRepository() *GithubData {
	labels := &LabelPage{}
	labels.Nodes = append(labels.Nodes, &Label{Name: "bug"})
	labels.TotalCount = len(labels.Nodes)

	categories := &CategoryPage{}
	categories.Nodes = append(categories.Nodes, &Category{Name: "Announcements"})
	categories.Nodes = append(categories.Nodes, &Category{Name: "General"})
	categories.Nodes = append(categories.Nodes, &Category{Name: "Ideas"})
	categories.Nodes = append(categories.Nodes, &Category{Name: "Polls"})
	categories.Nodes = append(categories.Nodes, &Category{Name: "Q&A"})
	categories.TotalCount = len(categories.Nodes)

	discussions := &DiscussionPage{}
	discussions.Nodes = append(discussions.Nodes, &Discussion{Title: "关于模板版本的一些思考", GitHubURL: "https://github.com/ThreeTenth/GitHub-Discussions-to-Blog/discussions/8", Category: &Category{Name: "Ideas"}, Comments: &CommentPage{}})
	discussions.TotalCount = len(discussions.Nodes)

	return &GithubData{
		&Repository{Labels: labels, Categories: categories, Discussions: discussions},
		&User{Login: "excing"},
	}
}
