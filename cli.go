package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func cli_init() *cobra.Command {
	var rootCmd = &cobra.Command{
		Use:   "sfpkg",
		Short: "基于容器化技术的隔离式软件包管理器",
	}

	var initCmd = &cobra.Command{
		Use:   "init",
		Short: "初始化容器",
		Run: func(cmd *cobra.Command, args []string) {
			if err := c_init(); err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
		},
	}

	var apkCmd = &cobra.Command{
		Use:                "apk",
		Short:              "软件管理程序",
		DisableFlagParsing: true,
		Run: func(cmd *cobra.Command, args []string) {
			if err := c_exec(append([]string{"apk"}, args...)); err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
		},
	}

	var runCmd = &cobra.Command{
		Use:                "run",
		Short:              "运行软件",
		DisableFlagParsing: true,
		Run: func(cmd *cobra.Command, args []string) {
			if err := c_exec(args); err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
		},
	}

	var linkCmd = &cobra.Command{
		Use:   "link",
		Short: "添加软件链接",
		Run: func(cmd *cobra.Command, args []string) {
			if err := c_linkout(args[0]); err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
		},
	}

	rootCmd.AddCommand(initCmd, apkCmd, runCmd, linkCmd)
	return rootCmd
}
