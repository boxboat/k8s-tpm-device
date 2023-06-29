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

package plugin

import (
	"fmt"
	"os"
	"path"
	"regexp"
	"time"

	"github.com/boxboat/k8s-tpm-device/pkg/common"
	"github.com/intel/intel-device-plugins-for-kubernetes/pkg/deviceplugin"
	"github.com/pkg/errors"
	"k8s.io/kubelet/pkg/apis/deviceplugin/v1beta1"
)

const (
	deviceDir    = "/dev"
	tpmRegex     = `^tpmrm[0-9]+$`
	deviceType   = "tpmrm"
	scanInterval = 5 * time.Second
)

type TpmDevicePlugin struct {
	deviceHostPath string
	deviceRegex    *regexp.Regexp
	sharedCapacity int
	ticker         *time.Ticker
	scanDone       chan bool
}

func WithDeviceHostPath(path string) Option {
	return optFn(func(opts *options) error {
		opts.deviceHostPath = path
		return nil
	})
}

func WithDeviceCapacity(capacity int) Option {
	return optFn(func(opts *options) error {
		if capacity <= 0 {
			return errors.New("capacity must be greater than 0")
		}
		opts.capacity = capacity
		return nil
	})
}

func NewTpmDevicePlugin(opts ...Option) (*TpmDevicePlugin, error) {
	var o options
	for _, opt := range opts {
		if opt != nil {
			if err := opt.configurePlugin(&o); err != nil {
				return nil, err
			}
		}
	}
	if o.deviceHostPath == "" {
		o.deviceHostPath = deviceDir
	}
	return &TpmDevicePlugin{
		deviceHostPath: o.deviceHostPath,
		deviceRegex:    regexp.MustCompile(tpmRegex),
		sharedCapacity: o.capacity,
		ticker:         time.NewTicker(scanInterval),
		scanDone:       make(chan bool, 1),
	}, nil
}

func (t *TpmDevicePlugin) Scan(notifier deviceplugin.Notifier) error {
	defer t.ticker.Stop()
	previouslyFound := -1
	for {
		select {
		case <-t.scanDone:
			return nil
		case <-t.ticker.C:
			devTree, err := t.scan()
			if err != nil {
				common.Log.Errorf("failed to scan: %v", err)
			}
			found := len(devTree)
			if found != previouslyFound {
				common.Log.Infof("tpm scan devices found: %d", found)
				previouslyFound = found
			}
			notifier.Notify(devTree)
		}
	}
}

func (t *TpmDevicePlugin) scan() (deviceplugin.DeviceTree, error) {
	files, err := os.ReadDir(t.deviceHostPath)
	if err != nil {
		return nil, errors.Wrap(err, "cannot read device directory")
	}

	devTree := deviceplugin.NewDeviceTree()
	for _, f := range files {
		var nodes []v1beta1.DeviceSpec

		if !t.deviceRegex.MatchString(f.Name()) {
			continue
		}

		devPath := path.Join(t.deviceHostPath, f.Name())
		common.Log.Debugf("adding %s to tpmrm %s", devPath, f.Name())
		nodes = append(nodes, v1beta1.DeviceSpec{
			HostPath:      devPath,
			ContainerPath: devPath,
			Permissions:   "rw",
		})

		if len(nodes) > 0 {
			for i := 0; i < t.sharedCapacity; i++ {
				devID := fmt.Sprintf("%s-%d", f.Name(), i)
				common.Log.Debugf("device ID: %s for device: %+v", devID, f)
				devTree.AddDevice(
					deviceType,
					devID,
					deviceplugin.NewDeviceInfo(
						v1beta1.Healthy,
						nodes,
						nil,
						nil,
						nil))
			}
		}
	}

	return devTree, nil
}

type options struct {
	capacity       int
	deviceHostPath string
}

type Option interface {
	configurePlugin(opts *options) error
}

type optFn func(opts *options) error

func (opt optFn) configurePlugin(opts *options) error {
	return opt(opts)
}
