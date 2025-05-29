package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

// Map of folder names to their LeetCode problem numbers and names
var problemMap = map[string]string{
	"EliminateMaximumNumberofMonsters":        "1921.eliminate-maximum-number-of-monsters",
	"DetermineIfaCellIsReachableAtAGivenTime": "2849.determine-if-a-cell-is-reachable-at-a-given-time",
	"RestoreTheArrayFromAdjacentPairs":        "1743.restore-the-array-from-adjacent-pairs",
	"findTheWinnerOfAnArrayGame":              "1535.find-the-winner-of-an-array-game",
	"SeatReservationManager":                  "1845.seat-reservation-manager",
	"SearchInsertPosition":                    "35.search-insert-position",
	"repeatedSubstringPattern":                "459.repeated-substring-pattern",
	"containerWithMostWater":                  "11.container-with-most-water",
	"medianOfTwoSortedArrays":                 "4.median-of-two-sorted-arrays",
	"transposeMatrix":                         "867.transpose-matrix",
	"palindromeNumber":                        "9.palindrome-number",
	"fibonacciNumber":                         "509.fibonacci-number",
	"countingBits":                            "338.counting-bits",
	"reverseInteger":                          "7.reverse-integer",
	"validParentheses":                        "20.valid-parentheses",
	"spiralMatrix":                            "54.spiral-matrix",
	"rotateImage":                             "48.rotate-image",
	"binarySearch":                            "704.binary-search",
	"twoSum":                                  "1.two-sum",
	"counter":                                 "2620.counter",
	"sleep":                                   "2621.sleep",
}

func main() {
	// Get the current working directory
	currentDir, err := os.Getwd()
	if err != nil {
		log.Fatal("Error getting current directory:", err)
	}

	fmt.Printf("Working in directory: %s\n", currentDir)
	fmt.Println("Starting folder renaming process...")

	// Read all directories in the current path
	entries, err := os.ReadDir(currentDir)
	if err != nil {
		log.Fatal("Error reading directory:", err)
	}

	renamedCount := 0
	skippedCount := 0

	for _, entry := range entries {
		if entry.IsDir() {
			folderName := entry.Name()

			// Check if this folder is in our problem map
			if newName, exists := problemMap[folderName]; exists {
				oldPath := filepath.Join(currentDir, folderName)
				newPath := filepath.Join(currentDir, newName)

				// Check if target directory already exists
				if _, err := os.Stat(newPath); err == nil {
					fmt.Printf("⚠️  Skipping '%s' -> '%s' (target already exists)\n", folderName, newName)
					skippedCount++
					continue
				}

				// Rename the folder
				err := os.Rename(oldPath, newPath)
				if err != nil {
					fmt.Printf("❌ Error renaming '%s' to '%s': %v\n", folderName, newName, err)
				} else {
					fmt.Printf("✅ Renamed: '%s' -> '%s'\n", folderName, newName)
					renamedCount++
				}
			} else {
				// Try to auto-generate name if not in map
				if autoName := generateName(folderName); autoName != "" {
					fmt.Printf("🔍 Unknown folder '%s' - suggested name: '%s'\n", folderName, autoName)
					fmt.Printf("   Add to problemMap if you want to rename this folder\n")
				}
			}
		}
	}

	fmt.Printf("\n📊 Summary:\n")
	fmt.Printf("   ✅ Successfully renamed: %d folders\n", renamedCount)
	fmt.Printf("   ⚠️  Skipped: %d folders\n", skippedCount)
	fmt.Printf("   🔍 Unknown folders found - check output above\n")
}

// generateName attempts to convert a folder name to kebab-case
func generateName(folderName string) string {
	// Skip if it already looks like a proper format
	if matched, _ := regexp.MatchString(`^\d+\.`, folderName); matched {
		return ""
	}

	// Convert PascalCase/camelCase to kebab-case
	re := regexp.MustCompile(`([a-z])([A-Z])`)
	kebab := re.ReplaceAllString(folderName, `${1}-${2}`)
	kebab = strings.ToLower(kebab)

	// Add placeholder number
	return fmt.Sprintf("xxxx.%s", kebab)
}
