# Gerenciador de Arquivos em Go

Este é um gerenciador de arquivos simples desenvolvido em Go que permite realizar operações básicas de gerenciamento de arquivos e diretórios.

## Funcionalidades

- Listar arquivos e diretórios
- Criar diretórios
- Copiar arquivos
- Mover arquivos
- Deletar arquivos
- Exibir informações detalhadas de arquivos

## Requisitos

- Go 1.16 ou superior
- Fyne (instalado automaticamente ao executar `go mod tidy`)

## Como executar

1. Certifique-se de ter o Go instalado em seu sistema
2. Clone este repositório
3. Execute `go mod tidy` para instalar as dependências
4. Execute o programa com o comando:
   ```bash
   go run .
   ```
5. Escolha o modo de execução:
   - Interface de Terminal: modo texto tradicional
   - Interface Gráfica: interface moderna com botões e lista de arquivos

## Modos de Uso

### Interface de Terminal

No modo terminal, você verá um menu com as seguintes opções:

1. **Listar arquivos**: Exibe o conteúdo de um diretório
2. **Criar diretório**: Cria um novo diretório no caminho especificado
3. **Copiar arquivo**: Copia um arquivo de uma origem para um destino
4. **Mover arquivo**: Move um arquivo de uma origem para um destino
5. **Deletar arquivo**: Remove um arquivo do sistema
6. **Exibir informações**: Mostra informações detalhadas sobre um arquivo
0. **Sair**: Encerra o programa

### Interface Gráfica

A interface gráfica oferece uma experiência mais moderna e intuitiva:

- **Campo de caminho**: Mostra o diretório atual
- **Lista de arquivos**: Exibe arquivos e pastas com ícones
- **Barra de ferramentas**:
  - Atualizar: Recarrega a lista de arquivos
  - Nova Pasta: Cria um novo diretório
  - Deletar: Remove o arquivo/pasta selecionado
  - Informações: Mostra detalhes do item selecionado

## Observações

- Os caminhos podem ser absolutos ou relativos
- Para diretórios com espaços no nome, use aspas ao inserir o caminho no modo terminal
- O programa trata erros e exibe mensagens informativas em caso de falha
- A interface gráfica requer um ambiente gráfico configurado no sistema 