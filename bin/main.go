package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/charmbracelet/lipgloss"
	"github.com/guptarohit/asciigraph"
	"github.com/joho/godotenv"
	"github.com/pioz/dexcommer"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
}

func main() {
	username := os.Getenv("DEXCOMMER_USERNAME")
	password := os.Getenv("DEXCOMMER_PASSWORD")
	applicationId := os.Getenv("DEXCOMMER_APPLICATION_ID")
	if applicationId == "" {
		applicationId = "d89443d2-327c-4a6f-89e5-496bbb0317db"
	}

	var count, minutes int
	var help bool

	flag.IntVar(&minutes, "m", 1440, "Retrieve last c glucose values in last m minutes")
	flag.IntVar(&count, "c", 5, "Retrieve last c glucose values")
	flag.BoolVar(&help, "h", false, "Print this help")
	flag.Parse()
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Username, password and applicationId are set via ENV variables or .env file:\nDEXCOMMER_USERNAME\nDEXCOMMER_PASSWORD\nDEXCOMMER_APPLICATION_ID\n\nUsage of %s:\n", os.Args[0])

		flag.PrintDefaults()
	}
	if help {
		flag.Usage()
		os.Exit(0)
	}

	session := dexcommer.NewSession(username, password, applicationId)
	glucoseValues := session.ReadLastestGlucoseValues(minutes, count)

	data := make([]float64, count)
	for i := 0; i < len(glucoseValues); i++ {
		data[i] = float64(glucoseValues[len(glucoseValues)-i-1].Value)
	}

	graph := asciigraph.Plot(data, asciigraph.Width(80), asciigraph.Height(10))
	fmt.Println(graph)

	for i := len(glucoseValues) - 1; i >= 0; i-- {
		var style = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("#FAFAFA")).
			Background(lipgloss.Color(color(glucoseValues[i].Value))).
			PaddingLeft(1).
			PaddingRight(1)

		v := style.Render(fmt.Sprintf("%d", glucoseValues[i].Value))
		fmt.Printf("%s %s %s\n", glucoseValues[i].Date.Format("15:04"), v, arrow(glucoseValues[i].Trend))
	}
}

func arrow(value string) string {
	switch value {
	case "Flat":
		return "➡️"
	case "FortyFiveDown":
		return "↘️"
	case "FortyFiveUp":
		return "↗️"
	case "NintyDown":
		return "⬇️"
	case "NintyUp":
		return "⬆️"
	default:
		return "❓"
	}
}

func color(value int) string {
	if value < 70 {
		return "#f5335a"
	}
	if value > 130 {
		return "#5574f2"
	}
	return "#39bf47"
}
