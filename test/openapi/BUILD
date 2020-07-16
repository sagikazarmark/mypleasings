subinclude("//openapi")

remote_file(
    name = "spec",
    url = "https://raw.githubusercontent.com/banzaicloud/pipeline/2140cd586b3cdccc4811176a363affa0096f64c4/apis/pipeline/pipeline.yaml",
    hashes = ["3ecf7a17ebc5ce3bb9acf5fbcefd4cfc24d25797"],
    licences = ["Apache-2.0"],
)

openapi_library(
    name = "openapi",
    spec = ":spec",
    generator = "go-server",
    additional_properties = {
        "packageName": "pipeline",
        "withGoCodegenComment": "true",
        "sourceFolder": "pipeline",
    },
    strip = ["{Dockerfile,go.*,README.md,main.go,go/api*.go,go/logger.go,go/routers.go}"],
)

openapi_validate(
    name = "validate",
    spec = ":spec",
)
