filegroup(
    name = "go",
    srcs = [
        "embed.build_defs",
        "go.build_defs",
        "misc.build_defs",
    ],
    visibility = ["PUBLIC"],
)

filegroup(
    name = "tools",
    srcs = ["tools.build_defs"],
    visibility = ["PUBLIC"],
)

filegroup(
    name = "compat",
    srcs = [
        "compat.build_defs",
        ":go",
    ],
    visibility = ["PUBLIC"],
)

export_file(
    name = "golangci.defaults.yaml",
    src = "golangci.defaults.yaml",
    visibility = ["PUBLIC"],
)
