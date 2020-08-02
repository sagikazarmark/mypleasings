subinclude("//go")

go_mod()

go_build(
    name = "build",
    deps = ["//test/go/lib"],
)

go_build(
    name = "build_not_in_tree",
    in_tree = False,
    deps = ["//test/go/lib"],
)

subinclude("//go:compat")

go_binary(
    name = "build_compat",
    srcs = ["main.go"],
    deps = ["//test/go/lib:lib_compat"],
)
