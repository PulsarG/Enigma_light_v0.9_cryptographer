package crypt

import (
	"fmt"
	"image/color"
	"reflect"
	"strconv"
	/* "strings" */

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"

	"test/anna"
	"test/consts"
)

type Cryptor struct {
	textField *widget.Entry
	label     widget.Label
	button    widget.Button
}

func NewCryptor() *Cryptor {
	return &Cryptor{
		textField: createMultiLineEntry(),
	}
}

func createMultiLineEntry() *widget.Entry {
	multiLineTextField := widget.NewMultiLineEntry()
	return multiLineTextField
}

func (c *Cryptor) SetWidgetsInCryptor() {
	c.label.Text = consts.LABEL_TEXT
	c.label.Wrapping = fyne.TextWrapWord
	c.textField.SetPlaceHolder("Enter here...")
	c.textField.Wrapping = fyne.TextWrapWord
	c.button.Text = consts.BUTTON_TEXT
	c.button.OnTapped = func() { c.testCrypt() }
}

func (c *Cryptor) setTextFromButton() {
	c.label.SetText(c.textField.Text)
}

func (c *Cryptor) converToFloat() {
	if s, err := strconv.ParseFloat(c.textField.Text, 64); err == nil {
		fmt.Println(s)
		fmt.Println(reflect.TypeOf(s))
	} else {
		fmt.Println("Не верно введены цифры", err)
	}
}



func (c *Cryptor) testCrypt() {
	anna.StartCrypt(c.textField.Text)
}

func (c *Cryptor) GetTextFild() *widget.Entry {
	return c.textField
}

func (c *Cryptor) GetLabel() *widget.Label {
	return &c.label
}

func (c *Cryptor) GetButton() *widget.Button {
	return &c.button
}

func (c *Cryptor) GetColorButton() *fyne.Container {
	color := color.RGBA{11, 78, 150, 1}
	btn := container.New(
		layout.NewMaxLayout(),
		canvas.NewRectangle(color),
		&c.button,
	)
	return btn
}