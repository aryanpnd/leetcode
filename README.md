<p align="center">
  <img src="https://leetcode.com/static/images/LeetCode_logo.png" alt="LeetCode" width="120"/>
</p>

# LeetCode Practice Repository

A simple repository for tracking my LeetCode problem-solving journey with automated organization and statistics.

## 👤 Profile Links

- **LeetCode Profile**: [aryanpnd](https://leetcode.com/u/aryanpnd)
- **GitHub Repository**: [aryanpnd/leetcode](https://github.com/aryanpnd/leetcode)

## 🚀 Quick Start

### Clone the Repository
```bash
git clone https://github.com/aryanpnd/leetcode.git
cd leetcode
```

### Setup Git Hooks (Optional)
```bash
# Copy the pre-commit hook for auto-organization
cp .githooks/pre-commit .git/hooks/pre-commit
chmod +x .git/hooks/pre-commit
```

The git hook automatically organizes files like `121.problem-name.go` into folders like `121.problem-name/solution.go` when you commit.

## 📈 Statistics Tool

Track your progress with the built-in statistics tool.

### Usage
```bash
# Show all statistics
./stats.exe

# Language breakdown
./stats.exe languages

# Export to JSON
./stats.exe export
```

### Commands
- `stats` - Show comprehensive statistics (default)
- `languages` - Language breakdown only
- `count` - Total question count
- `list` - List all questions
- `export` - Export to JSON

### Sample Output
```
📈 SUMMARY
├─ 📊 8 questions solved
├─ 🔤 3 programming languages used
└─ 🔁 2 questions solved in multiple languages

🔤 LANGUAGE USAGE BREAKDOWN
┌────────────┬───────┬──────────┐
│ Language   │ Count │  Percent │
├────────────┼───────┼──────────┤
│ Go         │     4 │     50.0% │
│ JavaScript │     3 │     37.5% │
│ C++        │     1 │     12.5% │
└────────────┴───────┴──────────┘
```