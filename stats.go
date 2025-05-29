package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"time"
)

const (
	defaultQuestionsDir = "questions"
	version             = "1.0.0"
)

// ANSI color codes
const (
	colorReset  = "\033[0m"
	colorRed    = "\033[31m"
	colorGreen  = "\033[32m"
	colorYellow = "\033[33m"
	colorBlue   = "\033[34m"
	colorPurple = "\033[35m"
	colorCyan   = "\033[36m"
	colorWhite  = "\033[37m"
	colorBold   = "\033[1m"
	colorDim    = "\033[2m"
)

// Language display mappings
var languageDisplay = map[string]string{
	"py":    "Python",
	"js":    "JavaScript",
	"ts":    "TypeScript",
	"java":  "Java",
	"cpp":   "C++",
	"c":     "C",
	"go":    "Go",
	"rs":    "Rust",
	"php":   "PHP",
	"rb":    "Ruby",
	"swift": "Swift",
	"kt":    "Kotlin",
	"cs":    "C#",
	"scala": "Scala",
	"r":     "R",
	"sql":   "SQL",
}

// Solution file patterns
var solutionPatterns = []*regexp.Regexp{
	regexp.MustCompile(`^(solution|Solution)\.(\w+)$`),
	regexp.MustCompile(`^(main|Main)\.(\w+)$`),
	regexp.MustCompile(`^(solve|Solve)\.(\w+)$`),
}

type Config struct {
	QuestionsDir string
	Verbose      bool
	NoColor      bool
}

type Stats struct {
	QuestionMap    map[string][]string `json:"questions"`
	LanguageCount  map[string]int      `json:"language_count"`
	TotalCount     int                 `json:"total_questions"`
	MultiLangCount int                 `json:"multi_language_questions"`
	GeneratedAt    string              `json:"generated_at"`
}

type CLI struct {
	config *Config
}

func NewCLI() *CLI {
	return &CLI{
		config: &Config{},
	}
}

func (c *CLI) colorize(text string, color string) string {
	if c.config.NoColor {
		return text
	}
	return color + text + colorReset
}

func (c *CLI) printError(msg string) {
	fmt.Fprintf(os.Stderr, "%s Error: %s\n", c.colorize("❌", colorRed), msg)
}

func (c *CLI) printSuccess(msg string) {
	fmt.Printf("%s %s\n", c.colorize("✅", colorGreen), msg)
}

func (c *CLI) printInfo(msg string) {
	fmt.Printf("%s %s\n", c.colorize("ℹ️", colorBlue), msg)
}

func (c *CLI) getLanguageDisplayName(lang string) string {
	if display, exists := languageDisplay[strings.ToLower(lang)]; exists {
		return display
	}
	return strings.ToUpper(lang)
}

func (c *CLI) scanQuestionsDirectory(questionsDir string) (*Stats, error) {
	questionMap := make(map[string][]string)
	languageCount := make(map[string]int)

	if c.config.Verbose {
		fmt.Printf("Scanning directory: %s\n", questionsDir)
	}

	err := filepath.WalkDir(questionsDir, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if !d.IsDir() {
			return nil
		}

		// Skip the root directory
		if path == questionsDir {
			return nil
		}

		questionName := filepath.Base(path)

		// Read files in this question directory
		files, err := os.ReadDir(path)
		if err != nil {
			if c.config.Verbose {
				fmt.Printf("Warning: Cannot read directory %s: %v\n", path, err)
			}
			return nil
		}

		for _, file := range files {
			if file.IsDir() {
				continue
			}

			fileName := file.Name()
			for _, pattern := range solutionPatterns {
				matches := pattern.FindStringSubmatch(fileName)
				if len(matches) >= 3 {
					lang := strings.ToLower(matches[2])
					questionMap[questionName] = append(questionMap[questionName], lang)
					languageCount[lang]++
					break
				}
			}
		}

		return filepath.SkipDir // Don't recurse into subdirectories
	})

	if err != nil {
		return nil, fmt.Errorf("failed to scan directory: %v", err)
	}

	multiLangCount := 0
	for _, langs := range questionMap {
		if len(langs) > 1 {
			multiLangCount++
		}
	}

	return &Stats{
		QuestionMap:    questionMap,
		LanguageCount:  languageCount,
		TotalCount:     len(questionMap),
		MultiLangCount: multiLangCount,
		GeneratedAt:    time.Now().Format(time.RFC3339),
	}, nil
}

func (c *CLI) printSummary(stats *Stats) {
	fmt.Printf("\n%s\n", c.colorize("📈 SUMMARY", colorBold+colorCyan))
	fmt.Printf("├─ %s %s questions solved\n",
		c.colorize("📊", colorBlue),
		c.colorize(fmt.Sprintf("%d", stats.TotalCount), colorBold+colorGreen))
	fmt.Printf("├─ %s %s programming languages used\n",
		c.colorize("🔤", colorBlue),
		c.colorize(fmt.Sprintf("%d", len(stats.LanguageCount)), colorBold+colorYellow))
	fmt.Printf("└─ %s %s questions solved in multiple languages\n",
		c.colorize("🔁", colorBlue),
		c.colorize(fmt.Sprintf("%d", stats.MultiLangCount), colorBold+colorPurple))
	fmt.Println()
}

func (c *CLI) printLanguageBreakdown(stats *Stats) {
	if len(stats.LanguageCount) == 0 {
		c.printInfo("No languages found.")
		return
	}

	// Sort languages by count (descending)
	type langCount struct {
		lang  string
		count int
	}

	var sorted []langCount
	totalSolutions := 0
	for lang, count := range stats.LanguageCount {
		sorted = append(sorted, langCount{lang, count})
		totalSolutions += count
	}

	sort.Slice(sorted, func(i, j int) bool {
		return sorted[i].count > sorted[j].count
	})

	fmt.Printf("%s\n", c.colorize("🔤 LANGUAGE USAGE BREAKDOWN", colorBold+colorCyan))

	// Calculate column widths
	maxLangWidth := 0
	maxCountWidth := 0
	for _, lc := range sorted {
		displayName := c.getLanguageDisplayName(lc.lang)
		if len(displayName) > maxLangWidth {
			maxLangWidth = len(displayName)
		}
		countStr := strconv.Itoa(lc.count)
		if len(countStr) > maxCountWidth {
			maxCountWidth = len(countStr)
		}
	}

	if maxLangWidth < 8 {
		maxLangWidth = 8
	}
	if maxCountWidth < 5 {
		maxCountWidth = 5
	}

	// Print header
	fmt.Printf("┌─%s─┬─%s─┬─%s─┐\n",
		strings.Repeat("─", maxLangWidth),
		strings.Repeat("─", maxCountWidth+2),
		strings.Repeat("─", 8))
	fmt.Printf("│ %s │ %s │ %s │\n",
		c.colorize(fmt.Sprintf("%-*s", maxLangWidth, "Language"), colorBold),
		c.colorize(fmt.Sprintf("%*s", maxCountWidth+2, "Count"), colorBold),
		c.colorize(fmt.Sprintf("%8s", "Percent"), colorBold))
	fmt.Printf("├─%s─┼─%s─┼─%s─┤\n",
		strings.Repeat("─", maxLangWidth),
		strings.Repeat("─", maxCountWidth+2),
		strings.Repeat("─", 8))

	// Print data rows
	for _, lc := range sorted {
		displayName := c.getLanguageDisplayName(lc.lang)
		percentage := float64(lc.count) / float64(totalSolutions) * 100
		fmt.Printf("│ %s │ %s │ %s │\n",
			c.colorize(fmt.Sprintf("%-*s", maxLangWidth, displayName), colorCyan),
			c.colorize(fmt.Sprintf("%*d", maxCountWidth+2, lc.count), colorGreen),
			c.colorize(fmt.Sprintf("%7.1f%%", percentage), colorYellow))
	}

	fmt.Printf("└─%s─┴─%s─┴─%s─┘\n",
		strings.Repeat("─", maxLangWidth),
		strings.Repeat("─", maxCountWidth+2),
		strings.Repeat("─", 8))
	fmt.Println()
}

func (c *CLI) printMultiLanguageQuestions(stats *Stats) {
	multiLang := make(map[string][]string)
	for question, langs := range stats.QuestionMap {
		if len(langs) > 1 {
			multiLang[question] = langs
		}
	}

	if len(multiLang) == 0 {
		c.printInfo("No questions solved in multiple languages yet.")
		return
	}

	fmt.Printf("%s\n", c.colorize("🔁 MULTI-LANGUAGE SOLUTIONS", colorBold+colorCyan))

	// Sort questions alphabetically
	var questions []string
	for question := range multiLang {
		questions = append(questions, question)
	}
	sort.Strings(questions)

	// Calculate column widths
	maxQuestionWidth := 0
	maxLangWidth := 0
	for _, question := range questions {
		if len(question) > maxQuestionWidth {
			maxQuestionWidth = len(question)
		}

		langs := multiLang[question]
		var displayLangs []string
		for _, lang := range langs {
			displayLangs = append(displayLangs, c.getLanguageDisplayName(lang))
		}
		sort.Strings(displayLangs)
		langStr := strings.Join(displayLangs, ", ")
		if len(langStr) > maxLangWidth {
			maxLangWidth = len(langStr)
		}
	}

	if maxQuestionWidth < 8 {
		maxQuestionWidth = 8
	}
	if maxLangWidth < 9 {
		maxLangWidth = 9
	}

	// Print header
	fmt.Printf("┌─%s─┬─%s─┬─%s─┐\n",
		strings.Repeat("─", maxQuestionWidth),
		strings.Repeat("─", maxLangWidth),
		strings.Repeat("─", 5))
	fmt.Printf("│ %s │ %s │ %s │\n",
		c.colorize(fmt.Sprintf("%-*s", maxQuestionWidth, "Question"), colorBold),
		c.colorize(fmt.Sprintf("%-*s", maxLangWidth, "Languages"), colorBold),
		c.colorize(fmt.Sprintf("%5s", "Count"), colorBold))
	fmt.Printf("├─%s─┼─%s─┼─%s─┤\n",
		strings.Repeat("─", maxQuestionWidth),
		strings.Repeat("─", maxLangWidth),
		strings.Repeat("─", 5))

	// Print data rows
	for _, question := range questions {
		langs := multiLang[question]
		var displayLangs []string
		for _, lang := range langs {
			displayLangs = append(displayLangs, c.getLanguageDisplayName(lang))
		}
		sort.Strings(displayLangs)
		langStr := strings.Join(displayLangs, ", ")

		fmt.Printf("│ %s │ %s │ %s │\n",
			c.colorize(fmt.Sprintf("%-*s", maxQuestionWidth, question), colorCyan),
			c.colorize(fmt.Sprintf("%-*s", maxLangWidth, langStr), colorGreen),
			c.colorize(fmt.Sprintf("%5d", len(langs)), colorYellow))
	}

	fmt.Printf("└─%s─┴─%s─┴─%s─┘\n",
		strings.Repeat("─", maxQuestionWidth),
		strings.Repeat("─", maxLangWidth),
		strings.Repeat("─", 5))
	fmt.Println()
}

func (c *CLI) printAllQuestions(stats *Stats, showFiles bool) {
	if len(stats.QuestionMap) == 0 {
		c.printInfo("No questions found.")
		return
	}

	fmt.Printf("%s\n", c.colorize("📝 ALL QUESTIONS", colorBold+colorCyan))

	// Sort questions alphabetically
	var questions []string
	for question := range stats.QuestionMap {
		questions = append(questions, question)
	}
	sort.Strings(questions)

	// Calculate column widths
	maxQuestionWidth := 0
	maxLangWidth := 0
	for _, question := range questions {
		if len(question) > maxQuestionWidth {
			maxQuestionWidth = len(question)
		}

		langs := stats.QuestionMap[question]
		var displayLangs []string
		for _, lang := range langs {
			displayLangs = append(displayLangs, c.getLanguageDisplayName(lang))
		}
		sort.Strings(displayLangs)
		langStr := strings.Join(displayLangs, ", ")
		if len(langStr) > maxLangWidth {
			maxLangWidth = len(langStr)
		}
	}

	if maxQuestionWidth < 8 {
		maxQuestionWidth = 8
	}
	if maxLangWidth < 9 {
		maxLangWidth = 9
	}

	// Print header
	headerFormat := "┌─%s─┬─%s─┐\n"
	rowFormat := "│ %s │ %s │\n"
	separatorFormat := "├─%s─┼─%s─┤\n"
	footerFormat := "└─%s─┴─%s─┘\n"

	fmt.Printf(headerFormat,
		strings.Repeat("─", maxQuestionWidth),
		strings.Repeat("─", maxLangWidth))
	fmt.Printf(rowFormat,
		c.colorize(fmt.Sprintf("%-*s", maxQuestionWidth, "Question"), colorBold),
		c.colorize(fmt.Sprintf("%-*s", maxLangWidth, "Languages"), colorBold))
	fmt.Printf(separatorFormat,
		strings.Repeat("─", maxQuestionWidth),
		strings.Repeat("─", maxLangWidth))

	// Print data rows
	for _, question := range questions {
		langs := stats.QuestionMap[question]
		var displayLangs []string
		for _, lang := range langs {
			displayLangs = append(displayLangs, c.getLanguageDisplayName(lang))
		}
		sort.Strings(displayLangs)
		langStr := strings.Join(displayLangs, ", ")

		fmt.Printf(rowFormat,
			c.colorize(fmt.Sprintf("%-*s", maxQuestionWidth, question), colorCyan),
			c.colorize(fmt.Sprintf("%-*s", maxLangWidth, langStr), colorGreen))
	}

	fmt.Printf(footerFormat,
		strings.Repeat("─", maxQuestionWidth),
		strings.Repeat("─", maxLangWidth))
	fmt.Println()
}

func (c *CLI) exportStats(stats *Stats, outputFile string) error {
	if outputFile == "" {
		outputFile = fmt.Sprintf("leetcode_stats_%s.json",
			time.Now().Format("20060102_150405"))
	}

	data, err := json.MarshalIndent(stats, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal JSON: %v", err)
	}

	err = os.WriteFile(outputFile, data, 0644)
	if err != nil {
		return fmt.Errorf("failed to write file: %v", err)
	}

	c.printSuccess(fmt.Sprintf("Statistics exported to: %s", c.colorize(outputFile, colorBold+colorGreen)))
	return nil
}

func (c *CLI) printUsage() {
	fmt.Printf(`%s - LeetCode Statistics CLI v%s

%s
  Analyze your LeetCode solutions with style!

%s
  %s [OPTIONS] [COMMAND]

%s
  stats     Show comprehensive statistics (default)
  languages Show language breakdown only  
  multi     Show multi-language solutions only
  count     Show total question count only
  list      List all questions with details
  export    Export statistics to JSON file
  help      Show this help message
  version   Show version information

%s
  -d, --dir string     Questions directory path (default: "%s")
  -v, --verbose        Enable verbose output
  -o, --output string  Output file for export command
  --no-color          Disable colored output
  -h, --help          Show help

%s
  %s                    # Show all stats
  %s --dir /path/to/questions    # Use custom directory  
  %s languages                   # Show only language breakdown
  %s list                        # List all questions
  %s export -o my_stats.json     # Export to specific file

`,
		c.colorize("🚀 LeetCode Stats", colorBold+colorCyan), version,
		c.colorize("DESCRIPTION", colorBold+colorYellow),
		c.colorize("USAGE", colorBold+colorYellow),
		c.colorize("leetcode-stats", colorCyan),
		c.colorize("COMMANDS", colorBold+colorYellow),
		c.colorize("OPTIONS", colorBold+colorYellow), defaultQuestionsDir,
		c.colorize("EXAMPLES", colorBold+colorYellow),
		c.colorize("leetcode-stats", colorCyan),
		c.colorize("leetcode-stats", colorCyan),
		c.colorize("leetcode-stats", colorCyan),
		c.colorize("leetcode-stats", colorCyan),
		c.colorize("leetcode-stats", colorCyan))
}

func (c *CLI) Run() {
	// Custom flag set for better control
	flagSet := flag.NewFlagSet("leetcode-stats", flag.ExitOnError)
	flagSet.Usage = func() {
		c.printUsage()
	}

	// Define flags
	dir := flagSet.String("d", defaultQuestionsDir, "Questions directory path")
	dirLong := flagSet.String("dir", defaultQuestionsDir, "Questions directory path")
	verbose := flagSet.Bool("v", false, "Enable verbose output")
	verboseLong := flagSet.Bool("verbose", false, "Enable verbose output")
	output := flagSet.String("o", "", "Output file for export command")
	outputLong := flagSet.String("output", "", "Output file for export command")
	noColor := flagSet.Bool("no-color", false, "Disable colored output")
	help := flagSet.Bool("h", false, "Show help")
	helpLong := flagSet.Bool("help", false, "Show help")

	// Parse flags
	flagSet.Parse(os.Args[1:])

	// Handle help
	if *help || *helpLong {
		c.printUsage()
		return
	}

	// Set config
	if *dirLong != defaultQuestionsDir {
		c.config.QuestionsDir = *dirLong
	} else {
		c.config.QuestionsDir = *dir
	}
	c.config.Verbose = *verbose || *verboseLong
	c.config.NoColor = *noColor

	// Check if NO_COLOR environment variable is set
	if os.Getenv("NO_COLOR") != "" {
		c.config.NoColor = true
	}

	// Get command
	args := flagSet.Args()
	command := "stats" // default command
	if len(args) > 0 {
		command = args[0]
	}

	// Handle version command
	if command == "version" {
		fmt.Printf("LeetCode Stats CLI v%s\n", version)
		return
	}

	// Check if questions directory exists
	if _, err := os.Stat(c.config.QuestionsDir); os.IsNotExist(err) {
		c.printError(fmt.Sprintf("Questions directory '%s' not found.", c.config.QuestionsDir))
		fmt.Printf("Create the directory or specify a different path with --dir flag.\n")
		os.Exit(1)
	}

	// Scan directory
	stats, err := c.scanQuestionsDirectory(c.config.QuestionsDir)
	if err != nil {
		c.printError(err.Error())
		os.Exit(1)
	}

	if stats.TotalCount == 0 {
		c.printError("No LeetCode solutions found.")
		fmt.Printf("Make sure your solutions are in the '%s' directory.\n", c.config.QuestionsDir)
		os.Exit(1)
	}

	// Handle commands
	switch command {
	case "stats", "":
		c.printSummary(stats)
		c.printLanguageBreakdown(stats)
		c.printMultiLanguageQuestions(stats)

	case "languages", "lang":
		c.printLanguageBreakdown(stats)

	case "multi":
		c.printMultiLanguageQuestions(stats)

	case "count":
		fmt.Printf("\n%s Total LeetCode Questions: %s\n\n",
			c.colorize("📊", colorBlue),
			c.colorize(fmt.Sprintf("%d", stats.TotalCount), colorBold+colorGreen))

	case "list":
		c.printAllQuestions(stats, false)

	case "export":
		outputFile := *outputLong
		if *output != "" {
			outputFile = *output
		}
		if err := c.exportStats(stats, outputFile); err != nil {
			c.printError(err.Error())
			os.Exit(1)
		}

	case "help":
		c.printUsage()

	default:
		c.printError(fmt.Sprintf("Unknown command: %s", command))
		fmt.Println("Run 'leetcode-stats help' for usage information.")
		os.Exit(1)
	}
}

func main() {
	cli := NewCLI()
	cli.Run()
}
