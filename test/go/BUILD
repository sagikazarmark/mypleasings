subinclude("//go")

go_build(
    name = "build",
    deps = ["//test/go/lib"],
    exported_deps = [":bindata"],
)

golangci_lint_config(
    name = "golangci.yml",
)

go_bindata(
    name = "bindata",
    srcs = ["go.mod", ":golangci.yml"],
    prefix = '$PKG',
)

go_binary(
    name = "bindata_test_bin",
    srcs = ["//test/go/bindata:main", ":bindata"],
)

gentest(
    name = "bindata_test",
    test_cmd = '"$DATA"',
    data = [":bindata_test_bin"],
    no_test_output = True,
)
