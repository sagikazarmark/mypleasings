subinclude("//misc:tools")

CONFIG.setdefault("LICENSEI_VERSION", "0.2.0")

licensei_binary(
    name = "licensei",
    version = CONFIG.LICENSEI_VERSION,
    visibility = ["PUBLIC"],
)

CONFIG.setdefault("MIGRATE_VERSION", "4.11.0")

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

CONFIG.setdefault("HUB_VERSION", "2.14.2")

hub_binary(
    name = "hub",
    version = CONFIG.HUB_VERSION,
    visibility = ["PUBLIC"],
)
