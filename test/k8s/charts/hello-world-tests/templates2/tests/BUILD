genrule(
    name = "tests",
    srcs = glob(["*"]),
    cmd = 'cp -a $PKG/. $OUT',
    outs = ["out"],
    visibility = ["//test/k8s/charts/hello-world-tests/..."],
)
