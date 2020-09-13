genrule(
    name = "templates2",
    srcs = glob(["*"], exclude = ["tests"]),
    cmd = 'cp -a $PKG/. $OUT',
    outs = ["out"],
    visibility = ["//test/k8s/charts/hello-world-tests/..."],
)
