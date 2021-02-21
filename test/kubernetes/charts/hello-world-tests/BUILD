subinclude("//kubernetes")

filegroup(
    name = "values.yaml",
    srcs = ["values.yaml"],
)

helm_chart(
    name = "chart-all",
    srcs = [
        "README.md",
        "config.yaml",
    ],
    chart_yaml = "Chart.yaml",
    helmignore = ".helmignore",
    templates = filegroup(
        name = "chart-all",
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
        name = "chart-all",
        srcs = glob(["templates/tests/**"]),
        tag = "tests",
    ),
    values_file = ":values.yaml",
)

helm_package(
    name = "hello-world-tests",
    chart = ":chart-all",
)
