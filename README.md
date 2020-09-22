# Pleasings

[![GitHub Workflow Status](https://img.shields.io/github/workflow/status/sagikazarmark/mypleasings/CI?style=flat-square)](https://github.com/sagikazarmark/mypleasings/actions?query=workflow%3ACI)
![Please Version](https://img.shields.io/badge/please%20version-%3E=15.4.0-B879FF.svg?style=flat-square)

**Addons & build rules for [Please](https://please.build).**


**⚠️ This repo is as experimental as it can be! Things may change or disappear without notice. ⚠️**


## Usage

Place the following in a `BUILD` file in your project (eg. in your project root):

```starlark
github_repo(
    name = "pleasings2",
    repo = "sagikazarmark/mypleasings",
    revision = "master",
)
```

Then include include it where necessary.
For example, use the following snippet to generate code from an OpenAPI specification:

```starlark
subinclude("///pleasings2//openapi")

openapi_library(
    name = "openapi",
    spec = "openapi.yaml",
)
```

## FAQ

### Helm end2end tests are failing

If you receive an error like the following:

```
Fail: //test/k8s/charts/hello-world:e2e   0 passed   0 skipped   0 failed   1 errored Took 350ms
Error: TestFailed in e2e
Test failed
exit status 1
Standard output:
Error: Kubernetes cluster unreachable: Get "https://127.0.0.1:57414/version?timeout=32s": dial tcp 127.0.0.1:57414: connect: connection refused
Error: Kubernetes cluster unreachable: Get "https://127.0.0.1:57414/version?timeout=32s": dial tcp 127.0.0.1:57414: connect: connection refused
```

Chances are your [KinD](https://github.com/kubernetes-sigs/kind) config has changed. Try rebuilding the kubeconfig:

```bash
./pleasew build --rebuild //test/k8s:kubeconfig
```

### `sha256sum` rule doesn't work on MacOS

Install the GNU coreutils for the `sha256sum` binary:

```bash
brew install coreutils
```

Or configure a custom tool:

```
[buildconfig]
sha256sum-tool = //path/to/tools:sha256sum
```


## License

The MIT License (MIT). Please see [License File](LICENSE) for more information.
