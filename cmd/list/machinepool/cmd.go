/*
Copyright (c) 2020 Red Hat, Inc.

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

package machinepool

import (
	"context"
	"fmt"

	cmv1 "github.com/openshift-online/ocm-sdk-go/clustersmgmt/v1"
	"github.com/spf13/cobra"

	"github.com/openshift/rosa/pkg/machinepool"
	"github.com/openshift/rosa/pkg/ocm"
	"github.com/openshift/rosa/pkg/output"
	"github.com/openshift/rosa/pkg/rosa"
)

const (
	use     = "machinepools"
	short   = "List cluster machine pools"
	long    = "List machine pools configured on a cluster."
	example = `  # List all machine pools on a cluster named "mycluster"
  rosa list machinepools --cluster=mycluster`
)

var (
	aliases = []string{"machinepool", "machine-pools", "machine-pool"}
)

func NewListMachinePoolCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:     use,
		Short:   short,
		Long:    long,
		Aliases: aliases,
		Example: example,
		Args:    cobra.NoArgs,
		Run:     rosa.DefaultRunner(rosa.RuntimeWithOCM(), ListMachinePoolRunner()),
	}

	output.AddFlag(cmd)
	ocm.AddClusterFlag(cmd)
	return cmd
}

func ListMachinePoolRunner() rosa.CommandRunner {
	return func(_ context.Context, runtime *rosa.Runtime, cmd *cobra.Command, _ []string) error {
		clusterKey := runtime.GetClusterKey()

		cluster := runtime.FetchCluster()
		if cluster.State() != cmv1.ClusterStateReady &&
			cluster.State() != cmv1.ClusterStateHibernating {
			return fmt.Errorf("Cluster '%s' is not yet ready", clusterKey)
		}

		service := machinepool.NewMachinePoolService()
		err := service.ListMachinePools(runtime, clusterKey, cluster)
		if err != nil {
			return fmt.Errorf("Failed to list machinepools: %s", err)
		}
		return nil
	}
}
