package adm

import (
	"github.com/redhat-openshift-ecosystem/provider-certification-tool/pkg/cmd/adm/baseline"
	"github.com/redhat-openshift-ecosystem/provider-certification-tool/pkg/cmd/adm/generate"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var admCmd = &cobra.Command{
	Use:   "adm",
	Short: "Administrative commands.",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			if err := cmd.Help(); err != nil {
				log.Errorf("error loading help(): %v", err)
			}
		}
	},
}

func init() {
	admCmd.AddCommand(parseMetricsCmd)
	admCmd.AddCommand(parseEtcdLogsCmd)
	admCmd.AddCommand(setupNodeCmd)
	admCmd.AddCommand(cleanerCmd)
	admCmd.AddCommand(baseline.NewCmdBaseline())
	admCmd.AddCommand(generate.NewCommandGenerate())
}

func NewCmdAdm() *cobra.Command {
	return admCmd
}
