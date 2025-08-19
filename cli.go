package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func cli_init() *cobra.Command {
	var rootCmd = &cobra.Command{
		Use:   "sfpkg",
		Short: "A lightweight software package manager for Segfault systems",
		Long: `Segfault Package Manager (sfpkg) allows you to install and manage software packages
on your Segfault root server using an Alpine Linux container environment.`,
	}

	var initCmd = &cobra.Command{
		Use:   "init",
		Short: "Initialize the container environment",
		Long: `Initialize the container environment for package management.
This will set up a clean Alpine Linux environment for installing software.
Warning: Running this again will delete all existing installations and changes!`,
		Run: func(cmd *cobra.Command, args []string) {
			if err := c_init(); err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
		},
	}

	var apkCmd = &cobra.Command{
		Use:                "apk [apk-args...]",
		Short:              "Access Alpine package manager (apk)",
		Long:               "Execute apk commands within the container environment. Usage is identical to Alpine Linux's native apk tool.",
		DisableFlagParsing: true,
		Run: func(cmd *cobra.Command, args []string) {
			if err := c_exec(append([]string{"apk"}, args...)); err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
		},
	}

	var runCmd = &cobra.Command{
		Use:                "run [command [args...]]",
		Short:              "Run a command inside the container",
		Long:               "Execute any command within the container environment. The command will have access to your home directory and current working directory.",
		DisableFlagParsing: true,
		Run: func(cmd *cobra.Command, args []string) {
			if err := c_exec(args); err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
		},
	}

	var linkCmd = &cobra.Command{
		Use:   "link [package]",
		Short: "Link a package to make it available system-wide",
		Long: `Link an installed package to your host system, making it accessible as a native command.
This allows you to use containerized applications directly from your terminal.`,
		Args: cobra.ExactArgs(1),
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
