package main

import (
	"os"

	"github.com/rivo/tview"
)

type FilePanel struct {
	files []os.FileInfo
	*tview.Table
}

func NewFilePanel() *FilePanel {
	p := &FilePanel{
		Table: tview.NewTable(),
	}

	p.SetBorder(true).
		SetTitle("files").
		SetTitleAlign(tview.AlignLeft)

	p.SetSelectable(true, false)

	return p
}

func (f *FilePanel) SetFiles(files []os.FileInfo) {
	f.files = files
}

func (f *FilePanel) SelectedFile() os.FileInfo {
	row, _ := f.GetSelection()
	if row > len(f.files)-1 || row < 0 {
		return nil
	}
	return f.files[row]
}

func (f *FilePanel) Keybinding(g *GUI) {
	f.SetSelectionChangedFunc(func(row, col int) {
		if file := f.SelectedFile(); file != nil {
			g.PreviewPanel.UpdateView(file.Name())
		}
	})
}

func (f *FilePanel) UpdateView() {
	table := f.Clear()

	for i, fi := range f.files {
		cell := tview.NewTableCell(fi.Name())
		table.SetCell(i, 0, cell)
	}
}
