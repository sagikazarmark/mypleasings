subinclude("//k8s")

helm_package(
    name = "hello-world",
)

helm_lint(
    name = "lint",
    chart = ":hello-world",
)

helm_package(
    name = "hello-world-all",
    chart_yaml = "Chart.yaml",
    values_yaml = "values.yaml",
    readme = "README.md",
    templates = filegroup(
        name = "hello-world-all",
        tag = "templates",
        srcs = glob(["templates/**"], exclude = ["templates/tests", "templates/tests/**"]),
    ),
    tests = filegroup(
        name = "hello-world-all",
        tag = "tests",
        srcs = glob(["templates/tests/**"]),
    ),
    helmignore = ".helmignore",
)

helm_package(
   name = "hello-world-values",
   values_yaml = "//test/k8s/charts/hello-world/values:values.yaml",
)

helm_package(
   name = "hello-world-templates",
   templates = "//test/k8s/charts/hello-world/templates2",
   tests = "//test/k8s/charts/hello-world/templates2/tests",
)

helm_package(
    name = "hello-world-misc",
    srcs = ["config.yaml"],
)
