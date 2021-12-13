# kbt
Tool for analyzing [Konvoy 1.X diagnostics bundle](https://support.d2iq.com/hc/en-us/articles/4409215513236-Create-a-Konvoy-1-x-Support-Bundle)


## Features

* Prints information about pods in the CSV format 

## Installation

### macOS

1. Download and unpack the binary:

```
$ curl -O -L https://github.com/adyatlov/kbt/releases/latest/download/kbt_darwin_amd64.tar.gz && tar -zxvf kbt_darwin_amd64.tar.gz
```

2. Move the `kbt` binary to one of the directories in the `PATH`.

### Linux

1. Download and unpack the binary:

```
$ curl -O -L https://github.com/adyatlov/kbt/releases/latest/download/kbt_linux_amd64.tar.gz && tar -zxvf kbt_linux_amd64.tar.gz
```

2. Move the `kbt` binary to one of the directories in the `PATH`.

### Windows

1. Download [the command](https://github.com/adyatlov/kbt/releases/latest/download/kbt_windows_amd64.tar.gz)
2. Extract it from the archive and move the `kbt` binary to one of the folders in the `PATH`.

### From sources

1. Install [Go compiler](https://golang.org/dl/).
2. Run the following command in your terminal:

```bash
$ go get github.com/adyatlov/kbt
```

## Usage

Launch the following command to see the list of commands:

```
$ kbt help
```

## How to release

1. Install [GoReleaser](https://goreleaser.com/install/).
2. Create [Github personal access token](https://help.github.com/en/articles/creating-a-personal-access-token-for-the-command-line)
    with the `repo` scope and export it as an environment variable called `GITHUB_TOKEN`:

  	```bash
  	$ export GITHUB_TOKEN=<your personal GitHub access token>
  	```

    Please find more information about this step [here](https://goreleaser.com/environment/).
3. Create a Git tag which adheres to [semantic versioning](https://semver.org/) and
    push it to GitHub:

    ```bash
    $ git tag -a v0.1.0 -m "Release v0.1.0"
    $ git push origin v0.1.0
    ```

    If you made a mistake on this step, you can delete the tag remotely and locally:

    ```bash
    $ git push origin :refs/tags/v0.1.0
    $ git tag --delete v0.1.0
    ```

4. Test that the build works with the following command:

    ```bash
    $ goreleaser release --skip-publish --rm-dist
    ```

5. If everything is fine publish the build with the following command:

    ```bash
	$ goreleaser release --rm-dist
    ```

