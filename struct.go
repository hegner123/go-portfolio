package main

import "github.com/a-h/templ"

type Data struct {
	Data Pages `json:"data"`
}

type Pages struct {
	PagesConnection Edges `json:"pagesConnection"`
}

type Edges struct {
	Edges []Node `json:"edges"`
}

type Node struct {
	Node PortfolioPage `json:"node"`
}

type PortfolioPage struct {
	HeroTitle     string         `json:"heroTitle"`
	Subtitle      string         `json:"subtitle"`
	AboutTitle    string         `json:"aboutTitle"`
	AboutBio      MarkdownObject `json:"aboutBio"`
	ProjectsTitle string         `json:"projectsTitle"`
	Projects      []Project      `json:"project"`
	BlogTitle     string         `json:"blogTitle"`
	BlogPosts     []BlogPost
}

type Project struct {
	Title       string         `json:"title"`
	Description MarkdownObject `json:"description"`
	SiteLink    templ.SafeURL         `json:"siteLink"`
	GithubLink  templ.SafeURL         `json:"githubLink"`
	Image       string         `json:"image"`
}

type BlogPost struct {
	Title   string `json:"title"`
	Date    string `json:"date"`
	Content string `json:"content"`
}

type MarkdownObject struct {
	Type     string         `json:"type"`
	Children []MarkdownEdge `json:"children"`
}

type MarkdownEdge struct {
	Type     string         `json:"type"`
	Children []MarkdownLeaf `json:"children"`
}

type MarkdownLeaf struct {
	Type string `json:"type"`
	Text string `json:"text"`
}
