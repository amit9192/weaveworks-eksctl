package utils

import (
	"encoding/json"
	"fmt"

	"github.com/kris-nova/logger"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	api "github.com/weaveworks/eksctl/pkg/apis/eksctl.io/v1alpha5"
	"github.com/weaveworks/eksctl/pkg/cfn/builder"
	"github.com/weaveworks/eksctl/pkg/ctl/cmdutils"
	"github.com/weaveworks/eksctl/pkg/eks"
	"github.com/weaveworks/eksctl/pkg/printers"
	"github.com/weaveworks/eksctl/pkg/spot"
	"github.com/weaveworks/goformation/v4"
)

func updateSpotOceanCluster(cmd *cmdutils.Cmd) {
	cfg := api.NewClusterConfig()
	cmd.ClusterConfig = cfg

	cmd.SetDescription("update-spot-ocean-cluster", "Update Spot Ocean cluster", "")

	cmd.CobraCommand.RunE = func(_ *cobra.Command, args []string) error {
		cmd.NameArg = cmdutils.GetNameArg(args)
		return doUpdateSpotOceanCluster(cmd)
	}

	cmd.FlagSetGroup.InFlagSet("General", func(fs *pflag.FlagSet) {
		cmdutils.AddConfigFileFlag(fs, &cmd.ClusterConfigFile)
		cmdutils.AddApproveFlag(fs, cmd)
		cmdutils.AddTimeoutFlag(fs, &cmd.ProviderConfig.WaitTimeout)
	})

	cmdutils.AddCommonFlagsForAWS(cmd.FlagSetGroup, &cmd.ProviderConfig, false)
}

func doUpdateSpotOceanCluster(cmd *cmdutils.Cmd) error {
	// Load configuration from file.
	// -------------------------------------------------------------------------

	if err := cmdutils.NewUtilsSpotOceanUpdateCluster(cmd).Load(); err != nil {
		return err
	}

	cfg := cmd.ClusterConfig
	meta := cmd.ClusterConfig.Metadata

	cmdutils.LogRegionAndVersionInfo(meta)

	ctl, err := cmd.NewCtl()
	if err != nil {
		return err
	}

	if err := ctl.CheckAuth(); err != nil {
		return err
	}

	if ok, err := ctl.CanOperate(cfg); !ok {
		return err
	}

	if err := checkVersion(ctl, cfg.Metadata); err != nil {
		return err
	}

	if err := ctl.LoadClusterIntoSpec(cfg); err != nil {
		return errors.Wrapf(err, "getting existing configuration for cluster %q", cfg.Metadata.Name)
	}

	nodeGroupService := eks.NewNodeGroupService(cfg, ctl.Provider)
	if err := nodeGroupService.Normalize(cmdutils.ToNodePools(cfg)); err != nil {
		return err
	}

	printer := printers.NewJSONPrinter()
	if err := printer.LogObj(logger.Debug, "cfg.json = \\\n%s\n", cfg); err != nil {
		return err
	}

	// Determine which nodegroup should be used to build the Ocean cluster from.
	// -------------------------------------------------------------------------

	ng, err := spot.DetermineOceanNodeGroup(cfg.NodeGroups)
	if err != nil {
		return err
	}

	stackManager := ctl.NewStackManager(cfg)
	stacks, err := stackManager.DescribeNodeGroupStacks()
	if err != nil {
		return err
	}

	stack := spot.GetStackByNodeGroupName(ng.Name, stacks)
	if stack == nil {
		return fmt.Errorf("ocean: couldn't find stack for nodegroup %q", ng.Name)
	}

	// Build a new stack from the default nodegroup.
	// -------------------------------------------------------------------------

	logger.Info("building nodegroup stack %q", ng.Name)
	clusterName := makeClusterStackName(cfg.Metadata.Name)
	newStack := builder.NewNodeGroupResourceSet(ctl.Provider, cfg, clusterName,
		stack.Tags, ng, false, false)
	if err := newStack.AddAllResources(); err != nil {
		return err
	}

	// Extract the NodeGroup resource from the new resource set.
	// -------------------------------------------------------------------------

	newTemplate := newStack.Template()
	newNodeGroup := newTemplate.Resources["NodeGroup"]
	newNodeGroupBytes, err := json.Marshal(newNodeGroup)
	if err != nil {
		return err
	}

	newNodeGroupCFN := struct{ Properties map[string]interface{} }{}
	if err = json.Unmarshal(newNodeGroupBytes, &newNodeGroupCFN); err != nil {
		return err
	}

	// Extract the NodeGroup resource from the existing resource set.
	// -------------------------------------------------------------------------

	existingTemplateBody, err := stackManager.GetStackTemplate(*stack.StackName)
	if err != nil {
		return err
	}

	existingTemplate, err := goformation.ParseJSON([]byte(existingTemplateBody))
	if err != nil {
		return fmt.Errorf("unexpected error parsing nodegroup template: %w", err)
	}

	existingNodeGroup, err := existingTemplate.GetCustomResourceWithName("NodeGroup")
	if err != nil {
		return fmt.Errorf("unable to find custom resource: %w", err)
	}

	// Override the resource properties of the Ocean cluster.
	// -------------------------------------------------------------------------

	existingNodeGroup.Properties["ocean"] = newNodeGroupCFN.Properties["ocean"]

	// Update the stack.
	// -------------------------------------------------------------------------

	existingTemplateJSON, err := existingTemplate.JSON()
	if err != nil {
		return err
	}

	if !cmd.Plan {
		if err := stackManager.UpdateNodeGroupStack(ng.Name, string(existingTemplateJSON)); err != nil {
			return fmt.Errorf("error updating nodegroup stack: %w", err)
		}
	}

	cmdutils.LogPlanModeWarning(cmd.Plan)
	return nil
}

// makeClusterStackName generates the name of the cluster stack.
func makeClusterStackName(clusterName string) string {
	return "eksctl-" + clusterName + "-cluster"
}

// checkVersion configures the version based on control plane version.
func checkVersion(ctl *eks.ClusterProvider, meta *api.ClusterMeta) error {
	version := ctl.ControlPlaneVersion()
	if version == "" {
		return fmt.Errorf("unable to get control plane version")
	}

	meta.Version = version
	logger.Info("will use version %s for new nodegroup(s) based on "+
		"control plane version", meta.Version)

	return nil
}
