package service

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
)

func GetTopMenuItems(logFile string) ([]string, error) {

	topMenuItems := make([]string, 0)
	file, err := os.Open(logFile)
	if err != nil {
		return topMenuItems, err
	}
	defer file.Close()

	encountered := make(map[string]bool)
	menuItemCounts := make(map[string]int)

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, ",")
		if len(parts) != 2 {
			log.Printf("Invalid log entry: %s\n", line)
			continue
		}
		eaterId := strings.TrimSpace(parts[0])
		foodMenuId := strings.TrimSpace(parts[1])
		if encountered[eaterId+foodMenuId] {
			return topMenuItems, fmt.Errorf("Food menu id repeated for the given eater")
		} else {
			encountered[eaterId+foodMenuId] = true
		}
		menuItemCounts[foodMenuId]++
	}
	// Create a slice of keys
	keys := make([]string, 0, len(menuItemCounts))
	for key := range menuItemCounts {
		keys = append(keys, key)
	}

	// Sort the slice based on values and keys
	sort.Slice(keys, func(i, j int) bool {
		if menuItemCounts[keys[i]] != menuItemCounts[keys[j]] {
			return menuItemCounts[keys[i]] > menuItemCounts[keys[j]]
		}
		return keys[i] < keys[j]
	})

	topMenuItems = append(topMenuItems, keys[0], keys[1], keys[2])

	return topMenuItems, nil
}
