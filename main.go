package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"time"
)

func listarArquivos(caminho string) error {
	arquivos, err := os.ReadDir(caminho)
	if err != nil {
		return err
	}

	fmt.Printf("\nConteúdo do diretório %s:\n", caminho)
	fmt.Println("----------------------------------------")
	for _, arquivo := range arquivos {
		info, err := arquivo.Info()
		if err != nil {
			continue
		}

		tipo := "Arquivo"
		if arquivo.IsDir() {
			tipo = "Diretório"
		}

		fmt.Printf("Nome: %-20s | Tipo: %-10s | Tamanho: %-8d bytes | Modificado: %s\n",
			arquivo.Name(),
			tipo,
			info.Size(),
			info.ModTime().Format("02/01/2006 15:04:05"))
	}
	return nil
}

func criarDiretorio(caminho string) error {
	return os.MkdirAll(caminho, 0755)
}

func copiarArquivo(origem, destino string) error {
	arquivoOrigem, err := os.Open(origem)
	if err != nil {
		return err
	}
	defer arquivoOrigem.Close()

	arquivoDestino, err := os.Create(destino)
	if err != nil {
		return err
	}
	defer arquivoDestino.Close()

	_, err = io.Copy(arquivoDestino, arquivoOrigem)
	return err
}

func moverArquivo(origem, destino string) error {
	return os.Rename(origem, destino)
}

func deletarArquivo(caminho string) error {
	return os.Remove(caminho)
}

func exibirInformacoes(caminho string) error {
	info, err := os.Stat(caminho)
	if err != nil {
		return err
	}

	fmt.Printf("\nInformações de %s:\n", caminho)
	fmt.Println("----------------------------------------")
	fmt.Printf("Nome: %s\n", info.Name())
	fmt.Printf("Tamanho: %d bytes\n", info.Size())
	fmt.Printf("Permissões: %s\n", info.Mode())
	fmt.Printf("Última modificação: %s\n", info.ModTime().Format("02/01/2006 15:04:05"))
	fmt.Printf("É diretório: %t\n", info.IsDir())
	return nil
}

func main() {
	fmt.Println("Escolha o modo de execução:")
	fmt.Println("1. Interface de Terminal")
	fmt.Println("2. Interface Gráfica")
	fmt.Print("\nOpção: ")

	var modo int
	fmt.Scan(&modo)

	switch modo {
	case 1:
		executarModoTerminal()
	case 2:
		iniciarGUI()
	default:
		fmt.Println("Opção inválida!")
	}
}

func executarModoTerminal() {
	for {
		fmt.Println("\nGerenciador de Arquivos em Go")
		fmt.Println("1. Listar arquivos")
		fmt.Println("2. Criar diretório")
		fmt.Println("3. Copiar arquivo")
		fmt.Println("4. Mover arquivo")
		fmt.Println("5. Deletar arquivo")
		fmt.Println("6. Exibir informações")
		fmt.Println("0. Sair")
		fmt.Print("\nEscolha uma opção: ")

		var opcao int
		fmt.Scan(&opcao)

		switch opcao {
		case 0:
			fmt.Println("Saindo...")
			return
		case 1:
			fmt.Print("Digite o caminho do diretório: ")
			var caminho string
			fmt.Scan(&caminho)
			if err := listarArquivos(caminho); err != nil {
				log.Printf("Erro ao listar arquivos: %v\n", err)
			}
		case 2:
			fmt.Print("Digite o caminho do novo diretório: ")
			var caminho string
			fmt.Scan(&caminho)
			if err := criarDiretorio(caminho); err != nil {
				log.Printf("Erro ao criar diretório: %v\n", err)
			} else {
				fmt.Println("Diretório criado com sucesso!")
			}
		case 3:
			fmt.Print("Digite o caminho do arquivo de origem: ")
			var origem string
			fmt.Scan(&origem)
			fmt.Print("Digite o caminho do arquivo de destino: ")
			var destino string
			fmt.Scan(&destino)
			if err := copiarArquivo(origem, destino); err != nil {
				log.Printf("Erro ao copiar arquivo: %v\n", err)
			} else {
				fmt.Println("Arquivo copiado com sucesso!")
			}
		case 4:
			fmt.Print("Digite o caminho do arquivo de origem: ")
			var origem string
			fmt.Scan(&origem)
			fmt.Print("Digite o caminho do arquivo de destino: ")
			var destino string
			fmt.Scan(&destino)
			if err := moverArquivo(origem, destino); err != nil {
				log.Printf("Erro ao mover arquivo: %v\n", err)
			} else {
				fmt.Println("Arquivo movido com sucesso!")
			}
		case 5:
			fmt.Print("Digite o caminho do arquivo a ser deletado: ")
			var caminho string
			fmt.Scan(&caminho)
			if err := deletarArquivo(caminho); err != nil {
				log.Printf("Erro ao deletar arquivo: %v\n", err)
			} else {
				fmt.Println("Arquivo deletado com sucesso!")
			}
		case 6:
			fmt.Print("Digite o caminho do arquivo: ")
			var caminho string
			fmt.Scan(&caminho)
			if err := exibirInformacoes(caminho); err != nil {
				log.Printf("Erro ao exibir informações: %v\n", err)
			}
		default:
			fmt.Println("Opção inválida!")
		}

		time.Sleep(2 * time.Second)
	}
}
