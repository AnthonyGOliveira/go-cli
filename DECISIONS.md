# DECISIONS.md

Este arquivo registra decisões técnicas, discussões e justificativas relevantes para o projeto.


## 2025-07-26
- Decisão de seguir o padrão Keep a Changelog para o arquivo CHANGELOG.md.
- Utilização do Go Modules para gerenciamento de dependências.
- Estrutura inicial baseada em projeto CLI simples.
- Adoção do pacote `urfave/cli` para facilitar a criação e manutenção de comandos e flags.
- Opção por utilizar os pacotes padrão `net`, `log` e `fmt` para garantir simplicidade, robustez e clareza na implementação dos comandos.
