// Copyright 2020 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package env

import (
	"runtime"
	"strings"

	"github.com/GoogleCloudPlatform/opentelemetry-operations-collector/internal/version"
)

func Create() error {
	userAgent, err := getUserAgent()
	if err != nil {
		return err
	}

	os.Setenv("USERAGENT", userAgent)
	return nil
}

func getUserAgent() (string, error) {
	userAgent := fmt.Sprintf(
		"Google Cloud Metrics Agent/%v (TargetPlatform=%v; Framework=OpenTelemetry Collector)",
		version.Version,
		strings.Title(runtime.GOOS),
	)

	return userAgent, nil
}
