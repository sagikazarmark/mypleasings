subinclude("//go")

filegroup(
    name = "lib",
    srcs = ["lib.go"],
    visibility = ["PUBLIC"],
)

go_testx(
    name = "test",
    srcs = [
        "lib_test.go",
        ":lib",
    ],
)

subinclude("//go:compat")

go_library(
    name = "lib_compat",
    srcs = ["lib.go"],
    visibility = ["PUBLIC"],
)

go_test(
    name = "test_compat",
    srcs = [
        "lib_test.go",
        ":lib_compat",
    ],
)

gotestsum_test(
    name = "gotest",
)
