subinclude("//go")

go_build(
    name = "build",
    export_deps = [":bindata"],
    deps = ["//test/go/lib"],
)

go_bindata(
    name = "bindata",
    srcs = [
        "go.mod",
    ],
    prefix = "$PKG",
)

go_binary(
    name = "bindata_test_bin",
    srcs = [
        ":bindata",
        "//test/go/bindata:main",
    ],
)

gentest(
    name = "bindata_test",
    data = [":bindata_test_bin"],
    no_test_output = True,
    test_cmd = '"$DATA"',
)
