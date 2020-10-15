subinclude("//go")

go_build(
    name = "build",
    export_deps = [":bindata"],
    deps = ["//test/go/lib"],
)

golangci_lint_config(
    name = "golangci.yml",
)

go_bindata(
    name = "bindata",
    srcs = [
        "go.mod",
        ":golangci.yml",
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
