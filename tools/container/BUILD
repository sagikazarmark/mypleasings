subinclude("//container:tools")

# https://github.com/hadolint/hadolint/releases
CONFIG.setdefault("HADOLINT_VERSION", "2.1.0")

hadolint_binary(
    name = "hadolint",
    version = CONFIG.HADOLINT_VERSION,
    visibility = ["PUBLIC"],
)
