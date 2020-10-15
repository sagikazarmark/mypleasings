genrule(
    name = "tests",
    srcs = glob(["*"]),
    outs = ["out"],
    cmd = "cp -a $PKG/. $OUT",
    visibility = ["//test/k8s/charts/hello-world-tests/..."],
)
