subinclude("//github")

genrule(
    name = "asset",
    cmd = 'echo asset > "$OUT"',
    out = "asset",
)

github_release(
    name = "publish",
    assets = [":asset"],
)

genrule(
    name = "notes",
    cmd = 'echo notes > "$OUT"',
    out = "notes",
)

filegroup(
    name = "assets",
    srcs = [":asset", ":notes"],
)

github_release(
    name = "publish_advanced",
    assets = [":assets", ":asset"],
    notes = ":notes",
)
