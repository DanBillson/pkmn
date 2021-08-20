package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/mtslzr/pokeapi-go"
	"github.com/spf13/cobra"
)

// entryCmd represents the entry command
var entryCmd = &cobra.Command{
	Use:   "entry",
	Short: "get a pokedex entry by name or #",
	Long: `get a pokedex entry by name or #`,
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		p, err := pokeapi.Pokemon(args[0])
		if err != nil {
			fmt.Printf("Entry for pokemon %v could not be found\n", args[0])
			os.Exit(1)
		}
		fmt.Printf("Pokedex Entry: #%03d\n", p.ID)
		fmt.Printf("Name: %s\n", strings.Title(p.Name))
		if len(p.Types) == 1 {
			fmt.Printf("Type: %s\n", strings.Title(p.Types[0].Type.Name))
		} else {
			fmt.Printf("Types: %s/%s\n", strings.Title(p.Types[0].Type.Name), strings.Title(p.Types[1].Type.Name))
		}
		fmt.Println("Stats:")
		for _, v := range p.Stats {
			fmt.Printf("  %s: %d\n", formatKey(v.Stat.Name), v.BaseStat)
		}
	},
}

func init() {
	rootCmd.AddCommand(entryCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// entryCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// entryCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func formatKey(s string) string {
	return strings.Title(strings.Join(strings.Split(s, "-"), " "))
}