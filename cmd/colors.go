package cmd

import (
	"encoding/json"
	"errors"
	"github.com/jedib0t/go-pretty/v6/table"
	"os"
	"time"
)

type color struct {
	RGB     string
	HEX     string
	Created string
}

type Colors []color

type ColorsCollection []Colors

func (cc *ColorsCollection) Check(colors *Colors) {
	if len(*colors) <= 10 {
		*cc = append(*cc, *colors)
	} else {
		*colors = (*colors)[len(*colors)-10 : len(*colors)]
		*cc = append(*cc, *colors)
	}
}

func (c *Colors) Add(rgb string, hex string) {
	clr := color{
		RGB:     rgb,
		HEX:     hex,
		Created: time.Now().Local().Format("2006-01-02 15:04:05"),
	}
	*c = append(*c, clr)
}

func (c *Colors) Delete(index int) error {
	ls := *c
	if index <= 0 || index > len(ls) {
		return nil
	}
	*c = append(ls[:index-1], ls[index:]...)
	return nil
}

func (c *Colors) Save(filename string) error {
	data, err := json.Marshal(c)
	if err != nil {
		return err
	}
	return os.WriteFile(filename, data, 0644)
}

func (c *Colors) Read(filename string) error {
	file, err := os.ReadFile(filename)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return nil
		}
	}
	if len(file) == 0 {
		err = errors.New("file is empty")
	}
	err = json.Unmarshal(file, c)
	if err != nil {
		return err
	}
	return nil
}

func (c *Colors) Clear() {
	*c = nil
}

func (c *Colors) Print() {
	total := len(*c)
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"#", "Rgb", "Hex", "Created"})
	for i, clr := range *c {
		t.AppendRows([]table.Row{
			{i + 1, clr.RGB, clr.HEX, clr.Created},
		})

	}
	//t.AppendSeparator()
	//t.AppendRow([]interface{}{"Notice", "rgb()means", "Lannister", 5000})
	t.AppendFooter(table.Row{"", "Total", total, ``})
	t.SetStyle(table.StyleColoredCyanWhiteOnBlack)
	t.Render()
}
