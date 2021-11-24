package cmd

import (
	"github.com/CheckmarxDev/sast-correlation-engine/engine/internals/input"
	"github.com/CheckmarxDev/sast-correlation-engine/engine/internals/output"
	"github.com/CheckmarxDev/sast-correlation-engine/engine/pkg/engine"
	"github.com/CheckmarxDev/sast-correlation-engine/engine/pkg/query"
	"github.com/spf13/cobra"
)

var (
	queriesPath   string
	inputPath     string
	inputProjects string
	outputPath    string

	cli = &cobra.Command{
		Use:   "scan",
		Short: "Start correlation scan",
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {

			queries := query.LoadQueries(queriesPath)
			rawInput := input.RawRead(inputPath)

			rawProjects := input.RawRead(inputProjects)

			results := engine.Run(rawInput, rawProjects, queries)

			output.Write(results, outputPath)
		},
	}
)

func init() { //nolint

	cli.Flags().StringVarP(&queriesPath, "queries-path", "q", "./queries/security", "path to directory with queries")
	cli.Flags().StringVarP(&inputPath, "input-path", "i", "./input", "path to read results from")
	cli.Flags().StringVarP(&outputPath, "output-path", "o", "./output", "path to write results to")
	cli.Flags().StringVarP(&inputProjects, "input-projects", "p", "./input", "path to read projects from")
}

func Execute() error {
	return cli.Execute()
}
