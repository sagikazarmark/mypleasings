filegroup(
    name = "go",
    srcs = ["go.build_defs", "embed.build_defs", "misc.build_defs"],
    visibility = ["PUBLIC"],
)

filegroup(
    name = "tools",
    srcs = ["tools.build_defs"],
    visibility = ["PUBLIC"],
)

filegroup(
    name = "compat",
    srcs = [":go", "compat.build_defs"],
    visibility = ["PUBLIC"],
)

export_file(
    name = "golangci.defaults.yaml",
    src = "golangci.defaults.yaml",
    visibility = ["PUBLIC"],
)
