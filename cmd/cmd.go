package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

const (
	sourceMySQL = "mysql"
)

var (
	cmdParam struct {
		Source string
		Output string
		Dsn    string
		Table  []string
	}

	genCmd = &cobra.Command{
		Use:   "dbtogo",
		Short: "数据库表 -> go结构体",
		Run: func(cmd *cobra.Command, args []string) {
			switch cmdParam.Source {
			case sourceMySQL:
				Run(new(MysqlGen))
			}
		},
	}
)

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the genCmd.
func Execute() {
	err := genCmd.Execute()
	if err != nil {
		os.Exit(0)
	}
}

func init() {
	genCmd.CompletionOptions.DisableDefaultCmd = true

	genCmd.Flags().StringVarP(&cmdParam.Source, "source", "s", sourceMySQL, "数据源(支持:mysql), 默认mysql")
	genCmd.Flags().StringVarP(&cmdParam.Output, "output", "o", "", "文件输出目录, 未指定则只打印到终端")
	genCmd.Flags().StringVarP(&cmdParam.Dsn, "dsn", "d", "", "dsn链接")
	genCmd.Flags().StringSliceVarP(&cmdParam.Table, "table", "t", []string{}, "表名, 多个逗号隔开")

	genCmd.MarkFlagRequired("dsn")
	genCmd.MarkFlagRequired("table")
}
