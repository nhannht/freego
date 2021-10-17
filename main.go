package main

import (
	"freego/service"
	"github.com/charmbracelet/bubbles/progress"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/lucasb-eyer/go-colorful"
	"log"
	"time"
)

type model struct {
	mem         *service.Memory
	cpu         *service.Stat
	memProgress progress.Model
	cpuProgress progress.Model
}

func (m model) View() string {
	borderStyle := lipgloss.NewStyle().Bold(true).Border(lipgloss.DoubleBorder()).PaddingLeft(5)
	memInfoText := lipgloss.NewStyle().Foreground(
		lipgloss.Color(colorful.WarmColor().Hex())).Render("Memory: ")
	quitInfoText := lipgloss.NewStyle().Foreground(
		lipgloss.Color(colorful.WarmColor().Hex())).Render("Press q to quit")
	cpuInfo := lipgloss.NewStyle().Foreground(lipgloss.Color(
		colorful.HappyColor().Hex(),
	)).Render("Cpu: ")
	mem := m.memProgress.ViewAs(m.mem.Percent())
	cpu := m.cpuProgress.ViewAs(m.cpu.CpuPercent())
	out := memInfoText + "\n" + mem + "\n" + cpuInfo + "\n" + cpu + "\n" + quitInfoText
	outWithStyle := borderStyle.Render(out)

	return outWithStyle
}
func initModel() model {
	mem := new(service.Memory)
	mem.UpdateMemory()
	cpu := new(service.Stat)
	cpu.UpdateCpu()
	mo := model{
		mem:         mem,
		cpu:         cpu,
		memProgress: progress.NewModel(progress.WithDefaultGradient()),
		cpuProgress: progress.NewModel(progress.WithDefaultGradient()),
	}
	return mo
}

func main() {

	program := tea.NewProgram(initModel())
	err := program.Start()
	if err != nil {
		log.Fatalln(err)
	}

}

func (m model) Init() tea.Cmd {
	return tea.Batch(
		memTickFunc(),
		cpuTickFunc(),
	)
}

type memTickMsg time.Time

func memTickFunc() tea.Cmd {
	return tea.Tick(time.Second, func(t time.Time) tea.Msg {
		return memTickMsg(t)
	},
	)
}

type cpuTickMsg time.Time

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch t := msg.(type) {
	case tea.KeyMsg:
		if t.String() == "q" {
			return m, tea.Quit
		}
	case memTickMsg:
		m.mem.UpdateMemory()
		m.cpu.UpdateCpu()
		return m, memTickFunc()
	//case cpuTickMsg:
	//	m.cpu.UpdateCpu()
	//	return m,cpuTickFunc()
	default:
		return m, nil
	}

	return m, nil
}

func cpuTickFunc() tea.Cmd {
	//time.Sleep(1)
	return tea.Tick(time.Second, func(t time.Time) tea.Msg {
		return memTickMsg(t)
	},
	)
}
