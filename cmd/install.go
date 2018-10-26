package cmd

import (
	"github.com/spf13/cobra"
)

func init() {
	var src, dst, force string
	var installCmd = &cobra.Command{
		Use:   "install",
		Short: "download & install a resource from netbrain-hub",
		Long:  `download & install a resource from netbrain-hub`,
		Example: `
# install overall-health-monitor qapp from http://hub.netbrain.com/qapp/overall-health-monitor 
  netbrain install qapp --src overall-health-monitor -d public/monitor -f`,
		Run: func(cmd *cobra.Command, args []string) {

		},
	}
	installCmd.Flags().StringVarP(&src, "src", "s", "", "resource name.")
	installCmd.Flags().StringVarP(&dst, "dst", "d", "", "install location.")
	installCmd.Flags().StringVarP(&force, "force", "f", "TCP", "protocol.")
	RootCmd.AddCommand(installCmd)
}

