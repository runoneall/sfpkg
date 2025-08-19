package main

import (
	"fmt"
	"os"
	"os/exec"
)

func c_run(args []string) error {
	real_args := append([]string{"UDOCKER_LOGLEVEL=0", "/usr/bin/udocker", "--allow-root"}, args...)
	cmd := exec.Command(real_args[0], real_args[1:]...)

	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}

func c_exec(args []string) error {
	env_PWD := os.Getenv("PWD")
	env_LOGNAME := os.Getenv("LOGNAME")

	return c_run(append([]string{
		"run",
		fmt.Sprintf("--workdir=%s", env_PWD),
		fmt.Sprintf("--volume=%s:%s", env_PWD, env_PWD),
		"--volume=/tmp:/tmp",
		"--hostenv",
		"--hostauth",
		"--bindhome",
		fmt.Sprintf("--user=%s", env_LOGNAME),
		"sfpkg-container",
	}, args...))
}

func c_init() error {
	if err := c_run([]string{"install", "--force", "--purge"}); err != nil {
		return fmt.Errorf("can not complete udocker installation")
	}

	c_run([]string{"rmi", "alpine"})
	if err := c_run([]string{"pull", "alpine"}); err != nil {
		return fmt.Errorf("can not pull alpine image")
	}

	c_run([]string{"rm", "sfpkg-container"})
	if err := c_run([]string{"create", "--force", "--name=sfpkg-container", "alpine"}); err != nil {
		return fmt.Errorf("can not create sfpkg-container")
	}

	if err := c_exec([]string{"apk", "add", "font-noto-cjk", "--no-cache"}); err != nil {
		return fmt.Errorf("can not install font-noto-cjk")
	}

	return nil
}

func c_linkout(name string) error {
	env_SHELL := os.Getenv("SHELL")

	fix_name := name
	if path, err := exec.LookPath(name); err == nil {
		fix_name = fmt.Sprintf("sf%s", name)
		fmt.Printf("Because %s is already installed in %s\n", name, path)
		fmt.Printf("We will create a startup script to it as %s\n", fix_name)
	}

	fmt.Printf("The script will installed at /sec/usr/bin/%s\n", fix_name)
	fmt.Printf("You can use \"rm /sec/usr/bin/%s\" to uninstall it\n", fix_name)
	return os.WriteFile(
		fmt.Sprintf("/sec/usr/bin/%s", fix_name),
		[]byte(fmt.Sprintf(`#!%s
sfpkg run %s $@`, env_SHELL, name)),
		0755,
	)
}
