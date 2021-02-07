filegroup(
    name = "go",
    srcs = [
        "embed.build_defs",
        "go.build_defs",
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
