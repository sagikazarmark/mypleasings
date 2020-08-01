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
