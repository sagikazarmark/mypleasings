genrule(
    name = "tests",
    srcs = glob(["*"]),
    outs = ["out"],
    cmd = "cp -a $PKG/. $OUT",
    visibility = ["//test/kubernetes/charts/hello-world-tests/..."],
)
