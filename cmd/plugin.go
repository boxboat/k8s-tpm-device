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
	"os"
	"os/signal"
	"syscall"

	"github.com/boxboat/k8s-tpm-device/pkg/common"
	"github.com/boxboat/k8s-tpm-device/pkg/plugin"
	"github.com/intel/intel-device-plugins-for-kubernetes/pkg/deviceplugin"
	"github.com/spf13/cobra"
)

type PluginArgs struct {
	capacity  int
	namespace string
}

var pluginArgs PluginArgs

func runPlugin() {
	tpmPlugin, err := plugin.NewTpmDevicePlugin(
		plugin.WithDeviceCapacity(pluginArgs.capacity))
	common.ExitIfError(err)
	manager := deviceplugin.NewManager(pluginArgs.namespace, tpmPlugin)

	go func() {
		manager.Run()
	}()

	// listen for shutdown signal
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)
	<-signalChan
}

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "run plugin",
	Long:  `Run TPM kubernetes device plugin`,
	Run: func(cmd *cobra.Command, args []string) {
		runPlugin()
	},
}

func init() {
	rootCmd.AddCommand(runCmd)
	runCmd.Flags().IntVar(
		&pluginArgs.capacity,
		"capacity",
		1,
		"desired number of pods that can access the tpm device")
	runCmd.Flags().StringVar(
		&pluginArgs.namespace,
		"namespace",
		"tpm.boxboat.io",
		"device namespace to reference")
}
