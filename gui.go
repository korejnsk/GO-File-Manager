package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/gdamore/tcell/v2"
)

type FileItem struct {
	Nome     string
	Tamanho  int64
	EhDir    bool
	selected bool
}

func iniciarGUI() {
	screen, err := tcell.NewScreen()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Erro ao criar tela: %v\n", err)
		os.Exit(1)
	}

	err = screen.Init()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Erro ao inicializar tela: %v\n", err)
		os.Exit(1)
	}
	defer screen.Fini()

	caminhoAtual, err := os.Getwd()
	if err != nil {
		caminhoAtual = "."
	}
	var arquivos []FileItem
	selectedIndex := 0
	scrollOffset := 0

	atualizarLista := func() {
		entries, err := os.ReadDir(caminhoAtual)
		if err != nil {
			return
		}

		arquivos = nil
		if caminhoAtual != "/" && caminhoAtual != "\\" {
			arquivos = append(arquivos, FileItem{
				Nome:    "..",
				Tamanho: 0,
				EhDir:   true,
			})
		}

		for _, entry := range entries {
			info, err := entry.Info()
			if err != nil {
				continue
			}

			arquivos = append(arquivos, FileItem{
				Nome:    entry.Name(),
				Tamanho: info.Size(),
				EhDir:   entry.IsDir(),
			})
		}
	}

	desenharTela := func() {
		screen.Clear()
		style := tcell.StyleDefault.Background(tcell.ColorBlack).Foreground(tcell.ColorWhite)
		selectedStyle := tcell.StyleDefault.Background(tcell.ColorBlue).Foreground(tcell.ColorWhite)

		caminhoAbsoluto, _ := filepath.Abs(caminhoAtual)
		drawText(screen, 0, 0, style, fmt.Sprintf("Caminho: %s", caminhoAbsoluto))

		drawText(screen, 0, 1, style, "F1: Atualizar | F2: Nova Pasta | F3: Deletar | F4: Info | Enter: Abrir | Backspace: Voltar | Q: Sair")

		_, maxY := screen.Size()
		listY := 3
		for i := scrollOffset; i < len(arquivos) && i-scrollOffset < maxY-3; i++ {
			arquivo := arquivos[i]
			currentStyle := style
			if i == selectedIndex {
				currentStyle = selectedStyle
			}

			icon := "📄"
			if arquivo.EhDir {
				icon = "📁"
			}

			text := fmt.Sprintf("%s %s - %d bytes", icon, arquivo.Nome, arquivo.Tamanho)
			drawText(screen, 0, listY+i-scrollOffset, currentStyle, text)
		}

		screen.Show()
	}

	atualizarLista()
	desenharTela()

	for {
		ev := screen.PollEvent()
		switch ev := ev.(type) {
		case *tcell.EventKey:
			switch ev.Key() {
			case tcell.KeyEscape, tcell.KeyCtrlC:
				return
			case tcell.KeyRune:
				switch ev.Rune() {
				case 'q', 'Q':
					return
				}
			case tcell.KeyUp:
				if selectedIndex > 0 {
					selectedIndex--
					if selectedIndex < scrollOffset {
						scrollOffset--
					}
				}
			case tcell.KeyDown:
				if selectedIndex < len(arquivos)-1 {
					selectedIndex++
					_, maxY := screen.Size()
					if selectedIndex-scrollOffset >= maxY-3 {
						scrollOffset++
					}
				}
			case tcell.KeyEnter:
				if selectedIndex >= 0 && selectedIndex < len(arquivos) {
					arquivo := arquivos[selectedIndex]
					if arquivo.Nome == ".." {
						caminhoAtual = filepath.Dir(caminhoAtual)
						selectedIndex = 0
						scrollOffset = 0
						atualizarLista()
					} else if arquivo.EhDir {
						novoCaminho := filepath.Join(caminhoAtual, arquivo.Nome)
						if _, err := os.Stat(novoCaminho); err == nil {
							caminhoAtual = novoCaminho
							selectedIndex = 0
							scrollOffset = 0
							atualizarLista()
						}
					}
				}
			case tcell.KeyBackspace, tcell.KeyBackspace2:
				if caminhoAtual != "/" && caminhoAtual != "\\" {
					caminhoAtual = filepath.Dir(caminhoAtual)
					selectedIndex = 0
					scrollOffset = 0
					atualizarLista()
				}
			case tcell.KeyF1:
				atualizarLista()
			case tcell.KeyF2:
				screen.Clear()
				drawText(screen, 0, 0, tcell.StyleDefault, "Digite o nome da nova pasta: ")
				screen.Show()

				nome := ""
				for {
					ev := screen.PollEvent()
					switch ev := ev.(type) {
					case *tcell.EventKey:
						switch ev.Key() {
						case tcell.KeyEnter:
							if nome != "" {
								caminho := filepath.Join(caminhoAtual, nome)
								criarDiretorio(caminho)
								atualizarLista()
							}
							goto continueMainLoop
						case tcell.KeyEscape:
							goto continueMainLoop
						case tcell.KeyBackspace, tcell.KeyBackspace2:
							if len(nome) > 0 {
								nome = nome[:len(nome)-1]
								screen.Clear()
								drawText(screen, 0, 0, tcell.StyleDefault, "Digite o nome da nova pasta: "+nome)
								screen.Show()
							}
						case tcell.KeyRune:
							nome += string(ev.Rune())
							screen.Clear()
							drawText(screen, 0, 0, tcell.StyleDefault, "Digite o nome da nova pasta: "+nome)
							screen.Show()
						}
					}
				}
			case tcell.KeyF3:
				if selectedIndex >= 0 && selectedIndex < len(arquivos) {
					arquivo := arquivos[selectedIndex]
					if arquivo.Nome != ".." {
						caminho := filepath.Join(caminhoAtual, arquivo.Nome)
						deletarArquivo(caminho)
						atualizarLista()
						if selectedIndex >= len(arquivos) {
							selectedIndex = len(arquivos) - 1
						}
					}
				}
			case tcell.KeyF4:
				if selectedIndex >= 0 && selectedIndex < len(arquivos) {
					arquivo := arquivos[selectedIndex]
					if arquivo.Nome != ".." {
						caminho := filepath.Join(caminhoAtual, arquivo.Nome)
						if info, err := os.Stat(caminho); err == nil {
							infoText := fmt.Sprintf(
								"Nome: %s\nTamanho: %d bytes\nPermissões: %s\nModificado: %s\nÉ diretório: %t",
								info.Name(),
								info.Size(),
								info.Mode(),
								info.ModTime().Format("02/01/2006 15:04:05"),
								info.IsDir(),
							)
							mostrarInfo(screen, infoText)
						}
					}
				}
			}
		case *tcell.EventResize:
			screen.Sync()
		}
	continueMainLoop:
		desenharTela()
	}
}

func drawText(screen tcell.Screen, x, y int, style tcell.Style, text string) {
	for i, r := range text {
		screen.SetContent(x+i, y, r, nil, style)
	}
}

func mostrarInfo(screen tcell.Screen, text string) {
	screen.Clear()
	lines := strings.Split(text, "\n")
	style := tcell.StyleDefault

	for i, line := range lines {
		drawText(screen, 0, i, style, line)
	}
	drawText(screen, 0, len(lines)+1, style, "Pressione qualquer tecla para continuar...")
	screen.Show()

	for {
		ev := screen.PollEvent()
		switch ev.(type) {
		case *tcell.EventKey:
			return
		}
	}
}
