package main

import (
	"io/ioutil"

	"github.com/rivo/tview"
)

type PreviewPanel struct {
	*tview.TextView
}

func NewPreviewPanel() *PreviewPanel {
	p := &PreviewPanel{
		TextView: tview.NewTextView(),
	}

	p.SetBorder(true).
		SetTitle("preview").
		SetTitleAlign(tview.AlignLeft)

	return p
}

func (p *PreviewPanel) UpdateView(name string) {
	var content string
	b, err := ioutil.ReadFile(name)
	if err != nil {
		content = err.Error()
	} else {
		content = string(b)
	}

	p.Clear().SetText(content)
}
