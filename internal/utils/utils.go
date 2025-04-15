package utils

import (
	"log"

	"github.com/charmbracelet/huh"
	"github.com/charmbracelet/huh/spinner"
)

type Values struct {
	Email   string
	Code    string
	NewLink string
}

type FormRunner struct {
	huhForm     *huh.Form
	onSubmit    func()
	spinnerText string
}

func Form(form *huh.Form) *FormRunner {
	runner := FormRunner{
		huhForm: form,
	}
	return &runner
}

func (runner *FormRunner) SpinnerText(text string) *FormRunner {
	runner.spinnerText = text
	return runner
}

func (runner *FormRunner) OnSubmit(onSubmit func()) *FormRunner {
	runner.onSubmit = onSubmit
	return runner
}

func (runner *FormRunner) Run() {
	err := runner.huhForm.Run()
	if err != nil {
		log.Fatal(err)
	}

	_ = spinner.New().Title(runner.spinnerText).Action(runner.onSubmit).Run()
}
