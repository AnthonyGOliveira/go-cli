
# go-cli

Uma aplicação de linha de comando (CLI) simples desenvolvida em Go utilizando o pacote [urfave/cli](https://github.com/urfave/cli).

## Funcionalidades

- **hello**: Exibe uma mensagem de saudação personalizada.
  - Parâmetro: `--name` (opcional) — Nome da pessoa a ser saudada. Padrão: "Mundo".
  - Exemplo: `go run main.go hello --name=João`

- **ip**: Busca e exibe os endereços IP de um host informado.
  - Parâmetro: `--host` (opcional) — Nome do host para consulta. Padrão: "localhost".
  - Exemplo: `go run main.go ip --host=google.com`

## Pacotes Utilizados

- [urfave/cli/v2](https://github.com/urfave/cli): Framework para construção de aplicações CLI em Go.
- [net](https://pkg.go.dev/net): Utilizado para realizar consultas de IPs de hosts.
- [log](https://pkg.go.dev/log): Para tratamento e exibição de erros.
- [fmt](https://pkg.go.dev/fmt): Para formatação e exibição de mensagens no terminal.

## Como executar

```bash
go run main.go [comando] [flags]
```

Exemplo:

```bash
go run main.go hello --name=Maria
go run main.go ip --host=example.com
```