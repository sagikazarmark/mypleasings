subinclude("//misc:tools")

# https://github.com/goph/licensei/releases
CONFIG.setdefault("LICENSEI_VERSION", "0.3.1")

licensei_binary(
    name = "licensei",
    version = CONFIG.LICENSEI_VERSION,
    visibility = ["PUBLIC"],
)

# https://github.com/golang-migrate/migrate/releases
CONFIG.setdefault("MIGRATE_VERSION", "4.14.1")

migrate_binary(
    name = "migrate",
    version = CONFIG.MIGRATE_VERSION,
    visibility = ["PUBLIC"],
)

# https://github.com/stedolan/jq/releases/
CONFIG.setdefault("JQ_VERSION", "1.6")

jq_binary(
    name = "jq",
    version = CONFIG.JQ_VERSION,
    visibility = ["PUBLIC"],
)

# https://github.com/mikefarah/yq/releases
CONFIG.setdefault("YQ_VERSION", "4.7.1")

yq_binary(
    name = "yq",
    version = CONFIG.YQ_VERSION,
    visibility = ["PUBLIC"],
)

# "TOOL" in config var name is an exception, because semver is too general
# https://github.com/fsaintjacques/semver-tool/releases
CONFIG.setdefault("SEMVER_TOOL_VERSION", "3.2.0")

semver_script(
    name = "semver",
    version = CONFIG.SEMVER_TOOL_VERSION,
    visibility = ["PUBLIC"],
)

# https://github.com/fossas/fossa-cli/releases
CONFIG.setdefault("FOSSA_VERSION", "1.1.7")

fossa_binary(
    name = "fossa",
    version = CONFIG.FOSSA_VERSION,
    visibility = ["PUBLIC"],
)

releaser_script(
    name = "releaser",
    visibility = ["PUBLIC"],
)
