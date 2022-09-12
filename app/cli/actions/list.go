package actions

import (
	"fmt"
	logger "github.com/glonner/pkg/log"
	"github.com/glonner/pkg/ux"
	"github.com/glonner/services"
	"github.com/pterm/pterm"
)

type ListDataResult struct {
	RepositoryName string
	CloneURL       string
}

var (
	AllowedFormats = []string{"table", "csv", "json"}
)

type ListActionArgs struct {
	Output string
}

func checkOutputFormat(output string) bool {
	for _, format := range AllowedFormats {
		if format == output {
			return true
		}
	}

	return false
}

func ListAction(args GlobalRequiredArgs, specifics ListActionArgs, logger *logger.ILogger) {
	if specifics.Output == "" {
		specifics.Output = "table"
	}

	if !checkOutputFormat(specifics.Output) {
		ux.OutError("Invalid output format. Please use one of the following: table, csv, json",
			"", false)
		return
	}

	ghSvc := services.NewGitHubSvc(logger)
	repositories, err := ghSvc.GetRepositories(args.Organization, args.Token)

	if err != nil {
		// TODO: Handler error, accordingly.
		return
	}

	var data []ListDataResult
	for _, repo := range repositories {
		repoDataSimplified := ghSvc.GetSimplifiedRepositoryData(repo)
		for _, item := range repoDataSimplified {
			data = append(data, ListDataResult{
				RepositoryName: item.Name,
				CloneURL:       item.CloneURL,
			})
		}
	}

	ListShowAsTable(data, args.Organization)
}

func ListShowAsTable(data []ListDataResult, org string) {
	var tableHeaders [][]string
	tableHeaders = append(tableHeaders, []string{"Repository Name", "Clone URL"})

	var tableData [][]string
	for _, item := range data {
		var row [][]string
		row = append(row, []string{item.RepositoryName, item.CloneURL})
		tableData = append(tableData, row...)
	}

	tableData = append(tableHeaders, tableData...)

	_ = ux.ShowHeader(fmt.Sprintf("Summary of [%s] Repositories", org), "info")
	_ = pterm.DefaultTable.WithHasHeader().WithData(tableData).Render()
}
