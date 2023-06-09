/*
Copyright 2023

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package loglevel

import (
	"context"
	"fmt"
	"os"
	"strings"

	docopt "github.com/docopt/docopt-go"
	"github.com/olekukonko/tablewriter"
)

func showLogSetting(ctx context.Context) error {
	componentConfiguration, err := collectLogLevelConfiguration(ctx)
	if err != nil {
		return err
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"COMPONENT NAMESPACE", "COMPONENT IDENTIFIER", "VERBOSITY"})
	genRow := func(namespace, identifier, verbosity string) []string {
		return []string{
			namespace,
			identifier,
			verbosity,
		}
	}

	for _, c := range componentConfiguration {
		table.Append(genRow(c.component.Namespace, c.component.Identifier, string(c.logSeverity)))
	}

	table.Render()
	return nil
}

// Show displays information about log verbosity (if set)
func Show(ctx context.Context, args []string) error {
	doc := `Usage:
  helper log-level show
Options:
  -h --help             Show this screen.
     
Description:
  The log-level show command shows information about current log verbosity.
`
	parsedArgs, err := docopt.ParseArgs(doc, nil, "1.0")
	if err != nil {
		return fmt.Errorf(
			"invalid option: 'helper %s'. Use flag '--help' to read about a specific subcommand",
			strings.Join(args, " "),
		)
	}
	if len(parsedArgs) == 0 {
		return nil
	}

	return showLogSetting(ctx)
}
