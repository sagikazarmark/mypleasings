subinclude("//kubernetes")

helm_package(
    name = "hello-world-all",
    chart_yaml = "Chart.yaml",
    helmignore = ".helmignore",
    readme = "README.md",
    templates = filegroup(
        name = "hello-world-all",
        srcs = glob(
            ["templates/**"],
            exclude = [
                "templates/tests",
                "templates/tests/**",
            ],
        ),
        tag = "templates",
    ),
    tests = filegroup(
        name = "hello-world-all",
        srcs = glob(["templates/tests/**"]),
        tag = "tests",
    ),
    values_yaml = "values.yaml",
)

helm_package(
    name = "hello-world-values",
    values_yaml = "//test/kubernetes/charts/hello-world-tests/values:values.yaml",
)

helm_package(
    name = "hello-world-templates",
    templates = "//test/kubernetes/charts/hello-world-tests/templates2",
    tests = "//test/kubernetes/charts/hello-world-tests/templates2/tests",
)

helm_package(
    name = "hello-world-misc",
    srcs = ["config.yaml"],
)
