genrule(
    name = "templates2",
    srcs = glob(
        ["*"],
        exclude = ["tests"],
    ),
    outs = ["out"],
    cmd = "cp -a $PKG/. $OUT",
    visibility = ["//test/kubernetes/charts/hello-world-tests/..."],
)
