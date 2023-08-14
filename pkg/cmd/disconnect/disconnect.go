/*
Copyright © 2021 The LitmusChaos Authors

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
package disconnect

import (
	"github.com/spf13/cobra"
)

// disconnectCmd represents the disconnect command
var DisconnectCmd = &cobra.Command{
	Use: "disconnect",
	Short: `Disconnect resources for LitmusChaos Execution plane.
		Examples:
		#disconnect a Chaos Infra
		litmusctl disconnect chaos-infra c520650e-7cb6-474c-b0f0-4df07b2b025b --project-id=c520650e-7cb6-474c-b0f0-4df07b2b025b

		Note: The default location of the config file is $HOME/.litmusconfig, and can be overridden by a --config flag
	`,
}
