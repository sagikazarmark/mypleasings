subinclude("//github")

genrule(
    name = "asset",
    out = "asset",
    cmd = 'echo asset > "$OUT"',
)

github_release(
    name = "publish",
    assets = [":asset"],
)

genrule(
    name = "notes",
    out = "notes",
    cmd = 'echo notes > "$OUT"',
)

filegroup(
    name = "assets",
    srcs = [
        ":asset",
        ":notes",
    ],
)

github_release(
    name = "publish_advanced",
    assets = [
        ":assets",
        ":asset",
    ],
    notes = ":notes",
)
