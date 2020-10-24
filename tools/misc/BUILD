subinclude("//misc:tools")

CONFIG.setdefault("LICENSEI_VERSION", "0.3.1")

licensei_binary(
    name = "licensei",
    version = CONFIG.LICENSEI_VERSION,
    visibility = ["PUBLIC"],
)

CONFIG.setdefault("MIGRATE_VERSION", "4.13.0")

migrate_binary(
    name = "migrate",
    version = CONFIG.MIGRATE_VERSION,
    visibility = ["PUBLIC"],
)

CONFIG.setdefault("JQ_VERSION", "1.6")

jq_binary(
    name = "jq",
    version = CONFIG.JQ_VERSION,
    visibility = ["PUBLIC"],
)

CONFIG.setdefault("YQ_VERSION", "3.4.1")

yq_binary(
    name = "yq",
    version = CONFIG.YQ_VERSION,
    visibility = ["PUBLIC"],
)

# This is an exception, because semver is too general
CONFIG.setdefault("SEMVER_TOOL_VERSION", "3.0.0")

semver_script(
    name = "semver",
    version = CONFIG.SEMVER_TOOL_VERSION,
    visibility = ["PUBLIC"],
)

CONFIG.setdefault("FOSSA_VERSION", "1.1.0")

fossa_binary(
    name = "fossa",
    version = CONFIG.FOSSA_VERSION,
    visibility = ["PUBLIC"],
)

releaser_script(
    name = "releaser",
    visibility = ["PUBLIC"],
)
