/*
Copyright (c) 2024 Red Hat, Inc.

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

package detach

import (
	"github.com/spf13/cobra"

	policy "github.com/openshift/rosa/cmd/detach/policy"
)

func NewRosaDetachCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:    "detach",
		Short:  "Detach AWS resource",
		Long:   "Detach AWS resource",
		Hidden: true,
		Args:   cobra.NoArgs,
	}
	cmd.AddCommand(policy.NewDetachPolicyCommand())
	return cmd
}
