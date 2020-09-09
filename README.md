# Pleasings

[![GitHub Workflow Status](https://img.shields.io/github/workflow/status/sagikazarmark/mypleasings/CI?style=flat-square)](https://github.com/sagikazarmark/mypleasings/actions?query=workflow%3ACI)
![Please Version](https://img.shields.io/badge/please%20version-%3E=15.3.1--beta.1-B879FF.svg?style=flat-square)

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


## License

The MIT License (MIT). Please see [License File](LICENSE) for more information.
