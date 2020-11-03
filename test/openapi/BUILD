subinclude("//openapi")

remote_file(
    name = "spec",
    hashes = ["3ecf7a17ebc5ce3bb9acf5fbcefd4cfc24d25797"],
    labels = ["manual"],
    licences = ["Apache-2.0"],
    url = "https://raw.githubusercontent.com/banzaicloud/pipeline/2140cd586b3cdccc4811176a363affa0096f64c4/apis/pipeline/pipeline.yaml",
)

openapi_library(
    name = "openapi",
    additional_properties = {
        "packageName": "pipeline",
        "withGoCodegenComment": "true",
        "sourceFolder": "pipeline",
    },
    generator = "go-server",
    labels = ["docker"],
    spec = ":spec",
    strip = ["{Dockerfile,go.*,README.md,main.go,go/api*.go,go/logger.go,go/routers.go}"],
)

openapi_validate(
    name = "validate",
    labels = ["docker"],
    spec = ":spec",
)

remote_file(
    name = "openapi_3_1",
    hashes = ["28ea3039742c17a45bdfbf4557a7bdd44717eee82f822a5d8ed28b364538dc50"],
    labels = ["manual"],
    licences = ["Apache-2.0"],
    url = "https://raw.githubusercontent.com/OAI/OpenAPI-Specification/f8449d1a893cc6a811c6f3d87e88b05761dc5397/examples/v3.1/webhook-example.json",
)

spectral_lint(
    name = "lint",
    spec = ":openapi_3_1",
)
