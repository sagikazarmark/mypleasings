subinclude("//go")

filegroup(
    name = "lib",
    srcs = ["lib.go"],
    visibility = ["PUBLIC"],
)

go_testx(
    name = "test",
    srcs = [":lib", "lib_test.go"],
)


subinclude("//go:compat")

go_library(
    name = "lib_compat",
    srcs = ["lib.go"],
    visibility = ["PUBLIC"],
)

go_test(
    name = "test_compat",
    srcs = [":lib_compat", "lib_test.go"],
)
