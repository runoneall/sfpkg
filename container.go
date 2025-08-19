package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func c_run(args []string) error {
	real_args := append([]string{"udocker", "--allow-root"}, args...)

	fmt.Println(">>>", strings.Join(real_args, " "))
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
		return fmt.Errorf("不能完成 udocker 安装")
	}

	c_run([]string{"rmi", "alpine"})
	if err := c_run([]string{"pull", "alpine"}); err != nil {
		return fmt.Errorf("不能完成 alpine 系统下载")
	}

	c_run([]string{"rm", "sfpkg-container"})
	if err := c_run([]string{"create", "--force", "--name=sfpkg-container", "alpine"}); err != nil {
		return fmt.Errorf("不能创建应用容器 sfpkg-container")
	}

	if err := c_exec([]string{"apk", "add", "font-noto-cjk", "--no-cache"}); err != nil {
		return fmt.Errorf("不能安装 font-noto-cjk 字体")
	}

	return nil
}

func c_linkout(name string) error {
	env_SHELL := os.Getenv("SHELL")

	fix_name := name
	if _, err := exec.LookPath(name); err == nil {
		fix_name = fmt.Sprintf("sf%s", name)
		fmt.Println("由于", name, "已经安装了，新软件将被安装为", fix_name)
	}

	return os.WriteFile(
		fmt.Sprintf("/sec/usr/bin/%s", fix_name),
		[]byte(fmt.Sprintf(`#!%s
sfpkg run %s $@`, env_SHELL, name)),
		0755,
	)
}
