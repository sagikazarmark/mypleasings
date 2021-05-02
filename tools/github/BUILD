subinclude("//github:tools")

# https://github.com/github/hub/releases
CONFIG.setdefault("HUB_VERSION", "2.14.2")

hub_binary(
    name = "hub",
    version = CONFIG.HUB_VERSION,
    visibility = ["PUBLIC"],
)

# https://github.com/cli/cli/releases
CONFIG.setdefault("GH_VERSION", "1.9.2")

gh_binary(
    name = "gh",
    version = CONFIG.GH_VERSION,
    visibility = ["PUBLIC"],
)
