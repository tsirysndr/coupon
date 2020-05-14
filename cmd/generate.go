/*
Copyright © 2020 Tsiry Sandratraina
All rights reserved.

Redistribution and use in source and binary forms, with or without
modification, are permitted provided that the following conditions are met:

1. Redistributions of source code must retain the above copyright notice,
   this list of conditions and the following disclaimer.

2. Redistributions in binary form must reproduce the above copyright notice,
   this list of conditions and the following disclaimer in the documentation
   and/or other materials provided with the distribution.

3. Neither the name of the copyright holder nor the names of its contributors
   may be used to endorse or promote products derived from this software
   without specific prior written permission.

THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS"
AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE
IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE
ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE
LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR
CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF
SUBSTITUTE GOODS OR SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS
INTERRUPTION) HOWEVER CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN
CONTRACT, STRICT LIABILITY, OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE)
ARISING IN ANY WAY OUT OF THE USE OF THIS SOFTWARE, EVEN IF ADVISED OF THE
POSSIBILITY OF SUCH DAMAGE.
*/
package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
	g "github.com/tsirysndr/coupon/generator"
)

// generateCmd represents the generate command
var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generate new voucher codes",
	Run: func(cmd *cobra.Command, args []string) {
		count, errCount := cmd.Flags().GetInt("count")
		length, errLength := cmd.Flags().GetInt("length")
		charset, _ := cmd.Flags().GetString("charset")
		prefix, _ := cmd.Flags().GetString("prefix")
		postfix, _ := cmd.Flags().GetString("postfix")
		pattern, _ := cmd.Flags().GetString("pattern")

		cfg := g.Config{
			Count:   count,
			Length:  length,
			Charset: charset,
			Prefix:  prefix,
			Postfix: postfix,
			Pattern: pattern,
		}

		if cfg.Pattern == g.Repeat("#", 8) && cfg.Length != 8 {
			cfg.Pattern = g.Repeat("#", cfg.Length)
		}

		codes, err := g.Generate(&cfg)
		if err != nil || errCount != nil || errLength != nil {
			log.Fatal(err)
		}
		for _, item := range codes {
			fmt.Println(item)
		}
	},
}

func init() {
	rootCmd.AddCommand(generateCmd)
	generateCmd.Flags().Int("count", 1, "Number of codes generated.")
	generateCmd.Flags().Int("length", 8, "Number of characters in a generated code (excluding prefix and postfix)")
	generateCmd.Flags().String("charset", g.Charset("alphanumeric"), "Characters that can appear in the code.")
	generateCmd.Flags().String("prefix", "", "A text appended before the code.")
	generateCmd.Flags().String("postfix", "", "A text appended after the code.")
	generateCmd.Flags().String("pattern", g.Repeat("#", 8), "A pattern for codes where hashes (#) will be replaced with random characters.")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// generateCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// generateCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
