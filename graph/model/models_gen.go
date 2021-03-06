// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type Language struct {
	ID         string  `json:"id"`
	Lang       string  `json:"lang"`
	Percentage float64 `json:"percentage"`
}

type PasswordInput struct {
	Time string `json:"time"`
	Pass string `json:"pass"`
}

type PostDto struct {
	Title string `json:"title"`
	Body  string `json:"body"`
}

type PostNode struct {
	ID   string `json:"id"`
	Type string `json:"type"`
	Leaf string `json:"leaf"`
}

type Profile struct {
	ID              string `json:"id"`
	AvatarURL       string `json:"avatarURL"`
	GithubURL       string `json:"githubURL"`
	Name            string `json:"name"`
	Location        string `json:"location"`
	Bio             string `json:"bio"`
	TwitterUsername string `json:"twitterUsername"`
	ReposCount      int    `json:"reposCount"`
	Followers       int    `json:"followers"`
	Following       int    `json:"following"`
}

type Repo struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	RepoName    string  `json:"repoName"`
	URL         string  `json:"url"`
	Description string  `json:"description"`
	Language    *string `json:"language"`
}
