subinclude("//go")

go_build(
    name = "build",
    deps = ["//test/go/lib"],
)

subinclude("//go:compat")

go_binary(
    name = "build_compat",
    srcs = ["main.go"],
    deps = ["//test/go/lib:lib_compat"],
)

golangci_lint_config(
    name = "golangci.yml",
)
