/*
Copyright 2024. Open Data Hub Authors

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

package models

import (
	"github.com/opendatahub-io/ai-edge/cli/pkg/commands/common"
	"github.com/opendatahub-io/ai-edge/cli/pkg/commands/flags"
	"github.com/spf13/cobra"
)

var addCmd = common.NewCmd(
	"add <model-name> <model-description> <model-version>",
	"Add a model and version to the model registry",
	`Add a model to the model registry

This command allows you to add a model and version to the model registry along with the build parameters from the 
provided parameters file.`,
	cobra.ExactArgs(3),
	[]flags.Flag{flags.FlagModelRegistryURL.SetParentFlag(), flags.FlagParams},
	common.SubCommandAdd,
	NewTeaModel,
)
