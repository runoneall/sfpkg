# Segfault Package Manager (sfpkg)

A lightweight software package manager for Segfault.

With this tool, you can install any software you like on your Segfault root server!

First, install sfpkg onto your system like this:

```shell
wget https://github.com/runoneall/sfpkg/releases/download/v2/sfpkg
chmod +x sfpkg
mv sfpkg /sec/usr/bin
```

Then, initialize a container to store your installed packages—don’t worry, sfpkg handles this automatically:

```shell
sfpkg init
```

> 💡 Tip: If you want to start over, just run `sfpkg init` again.  
> Be aware: this will delete all existing installations and changes!

Once that's done, you’re ready to install software. sfpkg uses the latest Alpine Linux image, so anything available on Alpine Linux can be installed on Segfault!

Here’s something even more exciting: sfpkg automatically syncs your desktop environment, home directory, and current working directory into the container—so you can even run GUI applications!

> ⚠️ CAUTION: This is a powerful feature, but use it carefully.  
> Avoid running sfpkg (or any software installed through it) inside the `/sec/root` directory (or its subdirectories) unless you know what you’re doing.

Use the `sfpkg apk` command to install packages. It works exactly like Alpine Linux’s native `apk`!

For example:

```shell
sfpkg apk add uv
```

But wait—it gets better. You can even install `apk` itself onto your Segfault root server!

```shell
sfpkg link apk
```

Now you can use `apk` directly:

```shell
apk add uv
```

After installation, you can typically run software using:

```shell
sfpkg run uv
```

But remember the `sfpkg link` command? You’re not limited to just `apk`. You can link any software installed inside the container to your Segfault root server—making it feel just like a native application!

Like this:

```shell
sfpkg link uv && uv
```
