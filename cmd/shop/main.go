package main

import (
	"errors"
	"fmt"
	"golearn/internal/utils"
	"strings"
	"time"

	"github.com/charmbracelet/huh"
	"github.com/charmbracelet/lipgloss"
)

const (
	BLUEPRINT_BLUE   = lipgloss.Color("#0063D3")
	TEMP_EMAIL_VALID = true
)

func main() {
	values := utils.Values{}

	onSubmit := func() {
		time.Sleep(2 * time.Second)
	}

	emailForm := huh.NewForm(huh.NewGroup(
		huh.NewNote().
			Title("Btw, I use Arch Linux").
			Description("cublueprint.org"),
		// Next(true).
		// NextLabel("Next"),
		huh.NewInput().
			Value(&values.Email).
			Title("Email").
			Placeholder("info@cublueprint.org").
			Validate(func(s string) error {
				if s == "Frank" {
					return errors.New("no franks, sorry")
				}
				return nil
			}),
	))

	codeForm := huh.NewForm(
		huh.NewGroup(
			huh.NewInput().
				Value(&values.Code).
				Title("2FA Code").
				Placeholder("123456").
				Validate(func(s string) error {
					if s == "Frank" {
						return errors.New("no franks, sorry")
					}
					return nil
				}).Description("Sent to your email inbox"),
		),
	)

	utils.Form(emailForm).OnSubmit(onSubmit).SpinnerText("Validating email...").Run()
	utils.Form(codeForm).OnSubmit(onSubmit).SpinnerText("Verifying 2FA code...").Run()

	{
		var sb strings.Builder
		keyword := func(s string) string {
			return lipgloss.NewStyle().Foreground(BLUEPRINT_BLUE).Render(s)
		}

		fmt.Fprintf(&sb,
			"%s\n\nEmail is %s, code is %s",
			lipgloss.NewStyle().Bold(true).Render("SUMMARY"),
			keyword(values.Email),
			keyword(values.Code),
		)

		fmt.Println(
			lipgloss.NewStyle().
				Width(40).
				BorderStyle(lipgloss.RoundedBorder()).
				BorderForeground(BLUEPRINT_BLUE).
				Padding(1, 2).
				Render(sb.String()),
		)
	}
}
