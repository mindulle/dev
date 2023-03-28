/*
Copyright © 2023 mindulle mindullestudio@gmail.com
*/
package cmd

import (
	"fmt"

	"github.com/charmbracelet/glamour"
	"github.com/spf13/cobra"
)

// refactorCmd represents the refactor command
var refactorCmd = &cobra.Command{
	Use:   "refactor",
	Short: "Fuzzy search refactoring techniques.",
	Long: `You can fuzzy search refactoring techniques with brief descriptions
or see a more well-explained article on refactoring.guru.

When to refactor:
Rule of Three
1. When you’re doing something for the first time, just get it done.
2. When you’re doing something similar for the second time, cringe at having to repeat but do the same thing anyway.
3. When you’re doing something for the third time, start refactoring.

How to refactor:
Checklist of refactoring done right way.
- [ ] The code should become cleaner.
- [ ] New functionality shouldn’t be created during refactoring.
- [ ] All existing tests must pass after refactoring.
`,
	Run: func(cmd *cobra.Command, args []string) {
		in := `![img](https://refactoring.guru/images/refactoring/diagrams/Push%20Down%20Field%20-%20Before.png)`

		out, _ := glamour.Render(in, "dark")
		fmt.Print(out)
	},
}

func init() {
	rootCmd.AddCommand(refactorCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// refactorCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// refactorCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
