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
	"strings"

	docopt "github.com/docopt/docopt-go"

	v1alpha1 "github.com/gianlucam76/pod-log-level/api/v1alpha1"
)

func updateLogSetting(ctx context.Context, logSeverity v1alpha1.LogLevel,
	component v1alpha1.Component) error {

	cc, err := collectLogLevelConfiguration(ctx)
	if err != nil {
		return nil
	}

	found := false
	spec := make([]v1alpha1.ComponentConfiguration, len(cc))

	for i, c := range cc {
		if c.component.Namespace == component.Namespace &&
			c.component.Identifier == component.Identifier {

			spec[i] = v1alpha1.ComponentConfiguration{
				Component: c.component,
				LogLevel:  logSeverity,
			}
			found = true
			break
		} else {
			spec[i] = v1alpha1.ComponentConfiguration{
				Component: c.component,
				LogLevel:  c.logSeverity,
			}
		}
	}

	if !found {
		spec = append(spec,
			v1alpha1.ComponentConfiguration{
				Component: component,
				LogLevel:  logSeverity,
			},
		)
	}

	return updateLogLevelConfiguration(ctx, spec)
}

// Set displays/changes log verbosity for a given component
func Set(ctx context.Context, args []string) error {
	doc := `Usage:
  helper log-level set --namespace=<namespace> --identifier=<identifier> (--info|--debug|--verbose)
Options:
  -h --help                    Show this screen.
     --namespace=<namespace>   Namespace of the component for which log severity is being set.
     --identifier=<identifier> Identifier of the component for which log severity is being set.
     --info                    Set log severity to info.
     --debug                   Set log severity to debug.
     --verbose                 Set log severity to verbose.
	 
Description:
  The log-level set command set log severity for the specified component.
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

	namespace := ""
	if passedNamespace := parsedArgs["--namespace"]; passedNamespace != nil {
		namespace = passedNamespace.(string)
	}

	identifier := ""
	if passedIdentifier := parsedArgs["--identifier"]; passedIdentifier != nil {
		identifier = passedIdentifier.(string)
	}

	info := parsedArgs["--info"].(bool)
	debug := parsedArgs["--debug"].(bool)
	verbose := parsedArgs["--verbose"].(bool)

	var logSeverity v1alpha1.LogLevel
	if info {
		logSeverity = v1alpha1.LogLevelInfo
	} else if debug {
		logSeverity = v1alpha1.LogLevelDebug
	} else if verbose {
		logSeverity = v1alpha1.LogLevelVerbose
	}

	return updateLogSetting(ctx, logSeverity, v1alpha1.Component{Namespace: namespace, Identifier: identifier})
}
