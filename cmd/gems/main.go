package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/spf13/cobra"
)

type Stone struct {
	Name  string
	Value int
}

var categories = [][]Stone{
	{ // Lesser Semiprecious Stones
		{"Agate", 5}, {"Alabaster", 5}, {"Azurite", 5}, {"Hematite", 5},
		{"Lapis lazuli", 5}, {"Malachite", 5}, {"Obsidian", 5}, {"Pearl, irregular freshwater", 5},
		{"Pyrite", 5}, {"Rhodochrosite", 5}, {"Quartz, rock crystal", 5}, {"Shell", 5},
		{"Tiger’s-eye", 5}, {"Turquoise", 5},
	},
	{ // Moderate Semiprecious Stones
		{"Bloodstone", 25}, {"Carnelian", 25}, {"Chrysoprase", 25}, {"Citrine", 25},
		{"Ivory", 25}, {"Jasper", 25}, {"Moonstone", 25}, {"Onyx", 25},
		{"Peridot", 25}, {"Quartz, milky, rose, or smoky", 25}, {"Sard", 25}, {"Sardonyx", 25},
		{"Spinel, red or green", 25}, {"Zircon", 25},
	},
	{ // Greater Semiprecious Stones
		{"Amber", 50}, {"Amethyst", 50}, {"Chrysoberyl", 50}, {"Coral", 50},
		{"Garnet", 50}, {"Jade", 50}, {"Jet", 50}, {"Pearl, saltwater", 50},
		{"Spinel, deep blue", 50}, {"Tourmaline", 50},
	},
	{ // Lesser Precious Stones
		{"Aquamarine", 500}, {"Opal", 500}, {"Pearl, black", 500}, {"Topaz", 500},
	},
	{ // Moderate Precious Stones
		{"Diamond, small", 1000}, {"Emerald", 1000}, {"Ruby, small", 1000}, {"Sapphire", 1000},
	},
	{ // Greater Precious Stones
		{"Diamond, large", 5000}, {"Emerald, brilliant green", 5000}, {"Ruby, large", 5000}, {"Star sapphire", 5000},
	},
}

func highestAffordableCategory(budget int) int {
	// start from the highest category and go to lower categories
	for i := len(categories) - 1; i >= 0; i-- {
		// check if we can afford at least one gem from this category and two from the lowest
		if budget >= categories[i][0].Value+2*categories[0][0].Value {
			return i
		}
	}

	// if no category found, return -1
	return -1
}

func generateStones(target int, numStones int) {
	var stones []Stone
	total := 0

	// Get the highest affordable category.
	categoryIndex := highestAffordableCategory(target)
	if categoryIndex < 0 {
		fmt.Println("Budget too small for any stones")
		return
	}

	for len(stones) < numStones && total < target {
		// While total value is less than the target, keep adding stones, biased towards lesser value.
		if rand.Float32() < 0.6 || len(stones) >= numStones {
			// 60% chance of choosing a gem from a lesser category (0 to len(categories)-2) or
			// we already added the requested number of stones
			categoryIndex = rand.Intn(len(categories) - 1)
		} else {
			// 40% chance of choosing a gem from any category
			categoryIndex = rand.Intn(len(categories))
		}

		// Check if we can afford at least one gem from this category.
		if categories[categoryIndex][0].Value > target-total {
			continue
		}

		gem := categories[categoryIndex][rand.Intn(len(categories[categoryIndex]))]
		count := 1 // Always add only one stone to respect the numStones limit
		value := count * gem.Value

		// Check if adding this gem would exceed the target value.
		if total+value > target {
			break
		}

		stones = append(stones, Stone{fmt.Sprintf("%s (%dx)", gem.Name, count), value})
		total += value
	}

	fmt.Printf("Generated stones with total value %d:\n", total)
	for _, stone := range stones {
		fmt.Printf("%s worth %d\n", stone.Name, stone.Value)
	}
}

var rootCmd = &cobra.Command{
	Use:   "stones",
	Short: "Stones generates a collection of precious stones within a target value",
	Long: `A comprehensive CLI tool built in Go for generating a collection 
of precious stones based on a target value provided by the user.`,
	Run: func(cmd *cobra.Command, args []string) {
		target, _ := cmd.Flags().GetInt("target")
		numStones, _ := cmd.Flags().GetInt("numStones") // Add this line
		if target <= 0 || numStones <= 0 {              // Update this line
			fmt.Println("Error: both target value and number of stones should be positive integers.")
			os.Exit(1)
		}
		rand.Seed(time.Now().UnixNano())
		generateStones(target, numStones) // Update this line
	},
}

func init() {
	rootCmd.PersistentFlags().Int("target", 200, "target value for generateStones")
	rootCmd.PersistentFlags().Int("numStones", 10, "number of stones to generate") // Add this line
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
