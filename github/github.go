package github

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type GitHubUser struct {
	Login     string `json:"login"`
	Name      string `json:"name"`
	Followers int    `json:"followers"`
	Following int    `json:"following"`
}

func GetUser(user string) {
	url := fmt.Sprintf("https://api.github.com/users/%s", user)
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Erro ao buscar o usuário:", err)
		return
	}
	defer resp.Body.Close()

	var gitHubUser GitHubUser
	if err := json.NewDecoder(resp.Body).Decode(&gitHubUser); err != nil {
		fmt.Println("Erro ao decodificar resposta:", err)
		return
	}

	fmt.Printf("Usuário: %s\nNome: %s\nSeguidores: %d\nSeguindo: %d\n",
		gitHubUser.Login, gitHubUser.Name, gitHubUser.Followers, gitHubUser.Following)
}
