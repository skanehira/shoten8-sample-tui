package main

import (
	"os"

	"github.com/rivo/tview"
)

type GUI struct {
	App          *tview.Application
	Pages        *tview.Pages
	FilePanel    *FilePanel
	PreviewPanel *PreviewPanel
}

func NewGUI() *GUI {
	return &GUI{
		App:          tview.NewApplication(),
		Pages:        tview.NewPages(),
		FilePanel:    NewFilePanel(),
		PreviewPanel: NewPreviewPanel(),
	}
}

func (g *GUI) Run() error {
	cur, err := os.Getwd()
	if err != nil {
		return err
	}
	files, err := Files(cur)
	if err != nil {
		return err
	}

	g.FilePanel.SetFiles(files)
	g.FilePanel.UpdateView()

	file := g.FilePanel.SelectedFile()
	if file != nil {
		g.PreviewPanel.UpdateView(file.Name())
	}

	g.SetKeybinding()

	grid := tview.NewGrid().SetColumns(0, 0).
		AddItem(g.FilePanel, 0, 0, 1, 1, 0, 0, true).
		AddItem(g.PreviewPanel, 0, 1, 1, 1, 0, 0, true)

	g.Pages.AddAndSwitchToPage("main", grid, true)

	return g.App.SetRoot(g.Pages, true).Run()
}

func (g *GUI) SetKeybinding() {
	g.FilePanel.Keybinding(g)
}
