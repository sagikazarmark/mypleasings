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
