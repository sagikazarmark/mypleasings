genrule(
    name = "values.yaml",
    cmd = "echo 'key: value' > $OUT",
    outs = ["values.yaml"],
    visibility = ["//test/k8s/charts/hello-world-tests/..."],
)
