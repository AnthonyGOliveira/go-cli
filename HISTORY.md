
# HISTORY.md

Este arquivo documenta o histórico de alterações importantes, eventos e marcos do projeto.


## 2025-07-26
- Criação do projeto inicial.
- Adição dos arquivos README.md e CHANGELOG.md.
- Adicionados comandos:
  - `hello`: Exibe saudação personalizada ao usuário, com opção de informar o nome.
  - `ip`: Busca e exibe os endereços IP de um host informado.
  - `server`: Lista os servidores de nomes (DNS) de um host informado.
- Utilização dos pacotes:
  - `urfave/cli` para estruturação dos comandos e flags.
  - `net` para consultas de IP (`LookupIP`) e servidores de nomes (`LookupNS`).
  - `log` para tratamento de erros.
  - `fmt` para exibição de mensagens no terminal.
