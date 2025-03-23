package main

import (
	"fmt"
	"github.com/charmbracelet/bubbles/stopwatch"
	"os"
	"time"

	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type errMsg error

type model struct {
	spinner   spinner.Model
	stopwatch stopwatch.Model
	quitting  bool
	err       error
}

var quitKeys = key.NewBinding(
	key.WithKeys("q", "esc", "ctrl+c"),
	key.WithHelp("", "Press q to quit"),
)

func initialModel() model {
	s := spinner.New()
	s.Spinner = spinner.Dot
	s.Style = lipgloss.NewStyle().Foreground(lipgloss.Color("205"))
	return model{spinner: s, stopwatch: stopwatch.NewWithInterval(time.Millisecond)}
}

func (m model) Init() tea.Cmd {
	return tea.Batch(m.spinner.Tick, m.stopwatch.Start())
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.KeyMsg:
		if key.Matches(msg, quitKeys) {
			m.quitting = true
			return m, tea.Quit

		}
		return m, nil
	case errMsg:
		m.err = msg
		return m, nil

	default:
		var spinnerCmd tea.Cmd
		var stopwatchCmd tea.Cmd
		m.spinner, spinnerCmd = m.spinner.Update(msg)
		m.stopwatch, stopwatchCmd = m.stopwatch.Update(msg)
		return m, tea.Batch(spinnerCmd, stopwatchCmd)
	}
}

func (m model) View() string {
	if m.err != nil {
		return m.err.Error()
	}
	str := lipgloss.JoinVertical(lipgloss.Left,
		m.spinner.View()+"Loading forever...",
		quitKeys.Help().Desc,
		"Elapsed Time: "+m.stopwatch.View(),
	)
	//str := fmt.Sprintf("\n\n   %s Loading forever... %s\n\n", m.spinner.View(), quitKeys.Help().Desc)
	if m.quitting {
		return str
	}
	return str
}

func main() {
	m := initialModel()
	p := tea.NewProgram(m)
	if _, err := p.Run(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	println(m.quitting)
}
