package cmd

import (
	"errors"
	"fmt"
	"github.com/bob2325168/bbs/framework/go-admin/cmd/app"
	"github.com/bob2325168/bbs/framework/go-admin/common/global"
	"github.com/go-admin-team/go-admin-core/sdk/pkg"
	"os"

	"github.com/spf13/cobra"

	"github.com/bob2325168/bbs/framework/go-admin/cmd/api"
	"github.com/bob2325168/bbs/framework/go-admin/cmd/config"
	"github.com/bob2325168/bbs/framework/go-admin/cmd/migrate"
	"github.com/bob2325168/bbs/framework/go-admin/cmd/version"
)

var rootCmd = &cobra.Command{
	Use:          "go-admin",
	Short:        "go-admin",
	SilenceUsage: true,
	Long:         `go-admin`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			tip()
			return errors.New(pkg.Red("requires at least one arg"))
		}
		return nil
	},
	PersistentPreRunE: func(*cobra.Command, []string) error { return nil },
	Run: func(cmd *cobra.Command, args []string) {
		tip()
	},
}

func tip() {
	usageStr := `欢迎使用 ` + pkg.Green(`go-admin `+global.Version) + ` 可以使用 ` + pkg.Red(`-h`) + ` 查看命令`
	usageStr1 := `也可以参考 https://doc.go-admin.dev/guide/ksks.html 里边的【启动】章节`
	fmt.Printf("%s\n", usageStr)
	fmt.Printf("%s\n", usageStr1)
}

func init() {
	rootCmd.AddCommand(api.StartCmd)
	rootCmd.AddCommand(migrate.StartCmd)
	rootCmd.AddCommand(version.StartCmd)
	rootCmd.AddCommand(config.StartCmd)
	rootCmd.AddCommand(app.StartCmd)
}

//Execute : apply commands
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(-1)
	}
}
