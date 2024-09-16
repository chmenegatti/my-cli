  ### Princípios SOLID em Go

  Os princípios SOLID são diretrizes fundamentais para o desenvolvimento de software que promovem a criação de sistemas mais escaláveis, compreensíveis e de fácil manutenção. Esses princípios são aplicáveis em qualquer linguagem de programação, incluindo Go, e podem ser aplicados no desenvolvimento de aplicações CLI com `Cobra` e `Viper`.

  #### Princípios SOLID

  1. **S - Single Responsibility Principle (SRP)**
    - **Princípio da Responsabilidade Única**: Cada classe, função ou módulo deve ter uma única responsabilidade. Em Go, isso significa que cada função ou estrutura deve fazer uma única tarefa.
    
  Exemplo:
  ```go
  type User struct {
    Name  string
    Email string
  }

  // Função responsável apenas por formatar o nome do usuário
  func FormatName(user User) string {
    return strings.ToUpper(user.Name)
  }
  ```

  2. **O - Open/Closed Principle (OCP)**
    - **Princípio do Aberto/Fechado**: Os módulos devem ser abertos para extensão, mas fechados para modificação. Isso quer dizer que o comportamento de uma função ou estrutura pode ser estendido sem alterar seu código.

  Exemplo:
  ```go
  type Shape interface {
      Area() float64
  }

  type Circle struct {
    Radius float64
  }

  func (c Circle) Area() float64 {
    return math.Pi * c.Radius * c.Radius
  }

  type Square struct {
    Side float64
  }

  func (s Square) Area() float64 {
    return s.Side * s.Side
  }
  ```

  3. **L - Liskov Substitution Principle (LSP)**
    - **Princípio da Substituição de Liskov**: Os objetos de uma classe derivada devem poder substituir os objetos da classe base sem alterar o comportamento do programa.

    Exemplo:
  ```go
  type Bird interface {
    Fly()
  }

  type Sparrow struct{}

  func (s Sparrow) Fly() {
    fmt.Println("Sparrow is flying")
  }

  type Penguin struct{}

  // Embora um pinguim seja um pássaro, ele não pode voar, então o design deve refletir isso
  // Não implementando a interface Fly para Penguin
  ```

  4. **I - Interface Segregation Principle (ISP)**
    - **Princípio da Segregação de Interfaces**: Nenhum cliente deve ser forçado a depender de métodos que não usa. Em Go, as interfaces devem ser pequenas e específicas.

    Exemplo:
  ```go
  type Printer interface {
    Print() string
  }

  type Scanner interface {
    Scan() string
  }

  type MultiFunctionDevice interface {
    Printer
    Scanner
  }
    ```

  5. **D - Dependency Inversion Principle (DIP)**
    - **Princípio da Inversão de Dependência**: Módulos de alto nível não devem depender de módulos de baixo nível; ambos devem depender de abstrações. Em Go, as interfaces ajudam a seguir esse princípio.

    Exemplo:
  ```go
  type Notifier interface {
    Send(message string) error
  }

  type EmailNotifier struct{}

  func (e EmailNotifier) Send(message string) error {
    fmt.Println("Sending email:", message)
    return nil
  }

  func NotifyUser(n Notifier, message string) {
    n.Send(message)
  }
  ```

  ---

  ### Aplicação CLI em Go com Cobra, Viper e SOLID

  Agora, vamos construir uma aplicação CLI que utiliza os princípios SOLID. A aplicação vai receber um parâmetro `-u` ou `--user` e buscar informações sobre um usuário no GitHub. Vamos usar as bibliotecas `Cobra` e `Viper` para configurar a CLI e aplicar os conceitos de SOLID no design.

  #### 1. Estrutura da Aplicação

  Primeiro, vamos estruturar o código seguindo os princípios SOLID.

  ```bash
  my-cli/
  ├── cmd/
  │   └── root.go
  ├── github/
  │   └── api.go
  ├── config/
  │   └── config.go
  └── main.go
  ```

  #### 2. Configurando o `Cobra` e `Viper` (Princípios SRP e DIP)

  No arquivo `cmd/root.go`, vamos configurar os comandos da CLI.

  ```go
  package cmd

  import (
    "fmt"
    "os"
    "github.com/spf13/cobra"
    "my-cli/github"
    "github.com/spf13/viper"
  )

  var rootCmd = &cobra.Command{
    Use:   "my-cli",
    Short: "Uma aplicação CLI para buscar usuários no GitHub",
    Run: func(cmd *cobra.Command, args []string) {
      user, _ := cmd.Flags().GetString("user")
      if user == "" {
        fmt.Println("É necessário informar um usuário com -u ou --user")
        os.Exit(1)
      }
      github.GetUser(user)
    },
  }

  func Execute() {
    if err := rootCmd.Execute(); err != nil {
      fmt.Println(err)
      os.Exit(1)
    }
  }

  func init() {
    cobra.OnInitialize(initConfig)
    rootCmd.PersistentFlags().StringP("user", "u", "", "Usuário do GitHub")
  }

  func initConfig() {
    viper.AutomaticEnv() 
  }
  ```

  #### 3. Criando a Lógica de Negócio em um Pacote Separado (Princípios SRP e DIP)

  Vamos agora criar o pacote `github/api.go`, onde implementamos a lógica de busca de usuários na API do GitHub.

  ```go
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
  ```

  #### 4. Arquivo `main.go`

  Por fim, no arquivo `main.go`, apenas chamamos o `cmd.Execute()`.

  ```go
  package main

  import "my-cli/cmd"

  func main() {
      cmd.Execute()
  }
  ```

  ---
  ### Executando a CLI

  Para testar a aplicação, execute o seguinte comando:

  ```bash
  go run main.go --user facebook
  ```

  A saída deve ser algo parecido com:

  ```bash

  Usuário: facebook
  Nome: Facebook
  Seguidores: 0
  Seguindo: 0

  ```
  ___

  ### Aplicando SOLID na CLI

  1. **SRP**: Cada pacote e função tem uma responsabilidade bem definida. A lógica de interação com o GitHub está no pacote `github`, enquanto a configuração da CLI está em `cmd`.
    
  2. **OCP**: A CLI está aberta para novas funcionalidades (como adicionar novos comandos) sem precisar modificar o código existente.

  3. **LSP**: Se criarmos uma nova função que implemente a mesma interface ou assinatura de `GetUser`, ela deve ser intercambiável sem quebrar o código.

  4. **ISP**: Mantivemos as interfaces pequenas e específicas, como o uso da função `GetUser`, que não depende de outras funcionalidades desnecessárias.

  5. **DIP**: O código depende de abstrações (como o uso da função `GetUser` e de pacotes externos como `cobra` e `viper`), e não de implementações concretas.

  ---

  ### Conclusão

  Neste artigo, exploramos os princípios SOLID com exemplos práticos em Go. Em seguida, construímos uma aplicação CLI simples usando `Cobra` e `Viper`, onde aplicamos os conceitos SOLID no design e organização do código. Essa abordagem garante que o código seja fácil de manter e estender, permitindo o desenvolvimento de software de alta qualidade em Go.

  Nos próximos artigos, continuaremos a explorar boas práticas e funcionalidades avançadas de Go!