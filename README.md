# IP Checker

![GitHub repo size](https://img.shields.io/github/repo-size/iuricode/README-template?style=for-the-badge)
![GitHub language count](https://img.shields.io/github/languages/count/iuricode/README-template?style=for-the-badge)

Uma ferramenta web leve e minimalista desenvolvida em Go para varrimento, monitorização e verificação de disponibilidade de endereços IP em redes locais.

## Funcionalidades

- Varrimento automático a partir de um IP base/gama de rede.
- Interface web sóbria e responsiva (Tailwind CSS).
- Filtragem e pesquisa de resultados em tempo real (por IP ou estado).
- Atualização assíncrona sem recarregamento de página.

---

## Futuras funcionalidades

- [] Executável unico
- [] Latência por ip
- [] Pesquisa por ipv6 inclusa

## Opção 1: Executar via Binário (.exe) - Recomendado para Utilizadores

> [!WARNING]
> V1.0 com erro, para resolver ver secção [Resolução de Problemas](#resolução-de-problemas).

> [!NOTE]
> Esta opção não requer que tenha o Go instalado no teu sistema.

1. Aceder à secção de **Releases** do repositório.
2. Descarregar o ficheiro `ip-checker.exe` mais recente.
3. Dar duplo clique em `ip-checker.exe`.
4. Abrir o navegador e aceder a: `http://localhost:8080`

---

## Opção 2: Executar Código Localmente - Recomendado para Desenvolvedores

Esta opção requer que tenhas o **Go 1.21+** instalado na tua máquina.

### 1. Clonar o Repositório

```bash
git clone https://github.com/Its-SMAC/IpsChecker.git
```

## Resolução de Problemas

### 1. Executável não funciona

```bash
    panic: html/template: pattern matches no files: `web/templates/*`
```

#### Como Resolver

Transferir o [Source Code(.zip)](https://github.com/Its-SMAC/IpsChecker/archive/refs/tags/Latest.zip), descompactar, introduzir o [.exe](https://github.com/Its-SMAC/IpsChecker/releases/download/Latest/ip-checker) no Source Code descompactado, e executar o ip-checker

## Estrutura do projeto

```text
.
├── cmd
│ └── main.go
├── go.mod
├── go.sum
├── internal
│ └── checker.go
├── LICENSE
├── README.md
└── web
├── static
│ └── js
│ └── script.js
└── templates
└── index.tmpl
```

## Stack

| Tecnologia                                                                                                              |   Tipo   | Função no Projeto                 |
| :---------------------------------------------------------------------------------------------------------------------- | :------: | :-------------------------------- |
| ![Go](https://img.shields.io/badge/Go-00ADD8?style=for-the-badge&logo=go&logoColor=white)                               | Backend  | Servidor API e varrimento de rede |
| ![Tailwind CSS](https://img.shields.io/badge/Tailwind_CSS-38B2AC?style=for-the-badge&logo=tailwind-css&logoColor=white) | Frontend | Estilização sóbria da interface   |
| ![JavaScript](https://img.shields.io/badge/JavaScript-F7DF1E?style=for-the-badge&logo=javascript&logoColor=black)       | Frontend | Pedidos assíncronos e filtragem   |

## Licence

Esse projeto está sob licença. Veja o arquivo [LICENÇA](LICENSE.md) para mais detalhes.

> [!WARNING]
> **Disclaimer:** O IP Checker é um projeto pessoal em constante desenvolvimento. Podem ocorrer instabilidades ou falsos negativos dependendo das permissões de rede do sistema operativo utilizado. Sinta-se livre para modificar, alterar ou estudar o código.
>
> _Nota: A interface web visual foi inteiramente produzida com o auxílio da [IA Gemini](https://gemini.google.com)._
