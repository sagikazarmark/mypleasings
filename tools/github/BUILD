subinclude("//github:tools")

CONFIG.setdefault("HUB_VERSION", "2.14.2")

hub_binary(
    name = "hub",
    version = CONFIG.HUB_VERSION,
    visibility = ["PUBLIC"],
)

CONFIG.setdefault("GH_VERSION", "1.0.0")

gh_binary(
    name = "gh",
    version = CONFIG.GH_VERSION,
    visibility = ["PUBLIC"],
)
