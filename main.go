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

package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"
	"strings"

	docopt "github.com/docopt/docopt-go"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/klog/v2"
	"k8s.io/klog/v2/klogr"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/gianlucam76/pod-log-level/internal/commands"
	"github.com/gianlucam76/pod-log-level/internal/utils"
	"github.com/gianlucam76/pod-log-level/lib"
)

func main() {
	doc := `Usage:
	helper [options] <command> [<args>...]

    log-level      Allows changing the log verbosity.

Options:
	-h --help          Show this screen.

Description:
  The helper command line tool is used to display/set log level.
  See 'helper <command> --help' to read about a specific subcommand.
 
  To reach cluster:
  - KUBECONFIG environment variable pointing at a file
  - In-cluster config if running in cluster
  - $HOME/.kube/config if exists
`
	klog.InitFlags(nil)

	ctx := context.Background()
	scheme, restConfig, clientSet, c := initializeManagementClusterAccess()
	utils.InitalizeManagementClusterAcces(scheme, restConfig, clientSet, c)

	parser := &docopt.Parser{
		HelpHandler:   docopt.PrintHelpOnly,
		OptionsFirst:  true,
		SkipHelpFlags: false,
	}

	logger := klogr.New()
	opts, err := parser.ParseArgs(doc, nil, "")
	if err != nil {
		var userError docopt.UserError
		if errors.As(err, &userError) {
			logger.V(lib.LogInfo).Info(fmt.Sprintf(
				"Invalid option: 'helper %s'. Use flag '--help' to read about a specific subcommand.\n",
				strings.Join(os.Args[1:], " "),
			))
		}
		os.Exit(1)
	}

	if opts["<command>"] != nil {
		command := opts["<command>"].(string)
		args := append([]string{command}, opts["<args>"].([]string)...)
		var err error

		switch command {
		case "log-level":
			err = commands.LogLevel(ctx, args, logger)
		default:
			err = fmt.Errorf("unknown command: %q\n%s", command, doc)
		}

		if err != nil {
			logger.V(lib.LogInfo).Info(fmt.Sprintf("%v\n", err))
		}
	}
}

func initializeManagementClusterAccess() (*runtime.Scheme, *rest.Config, *kubernetes.Clientset, client.Client) {
	scheme, err := utils.GetScheme()
	if err != nil {
		werr := fmt.Errorf("failed to get scheme %w", err)
		log.Fatal(werr)
	}

	restConfig := ctrl.GetConfigOrDie()
	restConfig.QPS = 100
	restConfig.Burst = 100

	cs, err := kubernetes.NewForConfig(restConfig)
	if err != nil {
		werr := fmt.Errorf("error in getting access to K8S: %w", err)
		log.Fatal(werr)
	}

	c, err := client.New(restConfig, client.Options{Scheme: scheme})
	if err != nil {
		werr := fmt.Errorf("failed to connect: %w", err)
		log.Fatal(werr)
	}

	return scheme, restConfig, cs, c
}
