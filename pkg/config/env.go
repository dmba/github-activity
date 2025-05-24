package config

import "os"
import "github.com/joho/godotenv"

type Env struct {
	GithubToken string
}

func NewEnv() (*Env, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, err
	}
	return &Env{
		GithubToken: os.Getenv("GITHUB_TOKEN"),
	}, err
}
