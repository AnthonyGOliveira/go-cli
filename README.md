
# go-cli

Uma aplicação de linha de comando (CLI) simples desenvolvida em Go utilizando o pacote [urfave/cli](https://github.com/urfave/cli).



## Funcionalidades

- **hello**: Exibe uma mensagem de saudação personalizada.
  - Parâmetro: `--name` (opcional) — Nome da pessoa a ser saudada. Padrão: "Mundo".
  - Exemplo: `go run main.go hello --name=João`
  - Utiliza o pacote `urfave/cli` para definição do comando e flags, e `fmt` para exibição da mensagem.

- **ip**: Busca e exibe os endereços IP de um host informado.
  - Parâmetro: `--host` (opcional) — Nome do host para consulta. Padrão: "localhost".
  - Exemplo: `go run main.go ip --host=google.com`
  - Utiliza o pacote `net` (função `LookupIP`) para buscar os IPs do host informado.
  - Em caso de erro, utiliza o pacote `log` para exibir mensagens de erro.

- **server**: Lista os servidores de nomes (DNS) de um host informado.
  - Parâmetro: `--host` (opcional) — Nome do host para consulta. Padrão: "localhost".
  - Exemplo: `go run main.go server --host=google.com`
  - Utiliza o pacote `net` (função `LookupNS`) para buscar os servidores de nomes do host informado.
  - Em caso de erro, utiliza o pacote `log` para exibir mensagens de erro.



## Pacotes Utilizados

- [urfave/cli/v2](https://github.com/urfave/cli): Framework para construção de aplicações CLI em Go, responsável pelo gerenciamento de comandos, flags e execução da aplicação.
- [net](https://pkg.go.dev/net): Utilizado para realizar consultas de IPs de hosts (`LookupIP`) e servidores de nomes (`LookupNS`).
- [log](https://pkg.go.dev/log): Para tratamento e exibição de erros, garantindo robustez na execução dos comandos.
- [fmt](https://pkg.go.dev/fmt): Para formatação e exibição de mensagens no terminal, tornando a saída mais amigável ao usuário.


## Como executar

```bash
go run main.go [comando] [flags]
```

Exemplos:

```bash
go run main.go hello --name=Maria
go run main.go ip --host=example.com
go run main.go server --host=example.com
```