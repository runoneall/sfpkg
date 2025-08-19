package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func cli_init() *cobra.Command {
	var rootCmd = &cobra.Command{
		Use:   "sfpkg",
		Short: "a simple software package manager for segfault",
	}

	var initCmd = &cobra.Command{
		Use:   "init",
		Short: "init udocker and container",
		Run: func(cmd *cobra.Command, args []string) {
			if err := c_init(); err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
		},
	}

	var apkCmd = &cobra.Command{
		Use:                "apk",
		Short:              "package manager",
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
		Short:              "run software",
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
		Short: "link software to system",
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
