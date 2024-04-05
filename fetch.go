package main

import (
	//	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

type PortfolioPage struct {
	HeroTitle     string `json:"heroTitle"`
	Subtitle      string `json:"subtitle"`
	AboutTitle    string `json:"aboutTitle"`
	AboutBio      string `json:"aboutBio"`
	ProjectsTitle string `json:"projectsTitle"`
	Projects      []Project
	BlogTitle     string `json:"blogTitle"`
	BlogPosts     []BlogPost
}

type Project struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	SiteLink    string `json:"siteLink"`
	GithubLink  string `json:"githubLink"`
	Image       string `json:"image"`
}

type BlogPost struct {
	Title   string `json:"title"`
	Date    string `json:"date"`
	Content string `json:"content"`
}

var buf bytes.Buffer

func fetchPage() PortfolioPage {
	env, _ := os.ReadFile("./.env")
	cmsId, cmsToken := splitEnv(string(env))

	jsonData := map[string]string{
		"query": `
        {
            pagesConnection {
                edges {
                    node {
                        id
                        heroTitle
                        subtitle
                        aboutTitle
                        aboutBio
                        project {
                            title
                            description
                            siteLink
                            githubLink
                            image
                            
                        }
                    }
                }
            }
        }

        `,
	}
	jsonValue, _ := json.Marshal(jsonData)
	url := fmt.Sprintf("https://content.tinajs.io/1.4/content/%s/github/main", cmsId)

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonValue))
	if err != nil {
		fmt.Printf("The HTTP request failed %s\n", err)
	}
	req.Header.Set("X-API-KEY", cmsToken)
	req.Header.Set("Content-Type", "application/json")

	client := http.DefaultClient
	response, err := client.Do(req)
	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
	}
	defer response.Body.Close()

	_, err = io.Copy(&buf, response.Body)
	if err != nil {
		fmt.Println("print")
	}

	data := buf.Bytes()
    fmt.Println(string(data))
    page := unmarshallPage(data) 
    fmt.Println(page)
	return page
}

func splitEnv(envB string) (string, string) {
	envN := strings.Split(string(envB), "\n")
	envS := []string{}

	for _, env := range envN {
		envStr := strings.Split(string(env), "=")
		for _, envSp := range envStr {
			e := envSp
			envS = append(envS, e)
		}
	}
	return envS[1], envS[3]
}
var res Data
func unmarshallPage (b []byte) PortfolioPage{
    err := json.Unmarshal(b,&res)
    if err != nil {
        fmt.Println(err)
    }
    return res.Data.PagesConnection.Edges[0].Node
} 
