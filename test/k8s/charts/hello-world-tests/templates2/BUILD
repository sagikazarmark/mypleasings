genrule(
    name = "templates2",
    srcs = glob(
        ["*"],
        exclude = ["tests"],
    ),
    outs = ["out"],
    cmd = "cp -a $PKG/. $OUT",
    visibility = ["//test/k8s/charts/hello-world-tests/..."],
)
