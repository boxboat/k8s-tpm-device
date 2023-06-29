// Copyright © 2023 BoxBoat, an IBM Company
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//	http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"fmt"
	"os"

	"github.com/boxboat/k8s-tpm-device/pkg/common"
	log "github.com/sirupsen/logrus"

	"github.com/spf13/cobra"
)

var (
	CfgFile string
	debug   bool
)

// rootCmdPersistentPreRunE configures logging
func rootCmdPersistentPreRunE(cmd *cobra.Command, args []string) error {
	common.Log.SetOutput(os.Stdout)
	common.Log.SetFormatter(&log.TextFormatter{
		FullTimestamp: true,
	})
	if debug {
		common.Log.SetLevel(log.DebugLevel)
	} else {
		common.Log.SetLevel(log.InfoLevel)
	}
	return nil
}

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:               "k8s-tpm-device",
	Short:             "K8s TPM Device Plugin",
	Long:              `K8s Trusted Platform Module Device Plugin`,
	PersistentPreRunE: rootCmdPersistentPreRunE,
}

// Execute adds all child commands to the root command and sets flags appropriately.
func Execute(version string) {
	rootCmd.Version = version
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().BoolVarP(&debug, "debug", "", false, "debug output")
}
