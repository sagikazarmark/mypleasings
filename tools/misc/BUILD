CONFIG.setdefault("LICENSEI_VERSION", "0.2.0")

genrule(
    name = "licensei",
    srcs = [remote_file(
        name = "licensei",
        _tag = "download",
        url = f"https://github.com/goph/licensei/releases/download/v{CONFIG.LICENSEI_VERSION}/licensei_{CONFIG.HOSTOS}_{CONFIG.HOSTARCH}.tar.gz",
        out = f"licensei-{CONFIG.LICENSEI_VERSION}.tar.gz",
    )],
    outs = ["licensei"],
    binary = True,
    cmd = '"$TOOL" x "$SRCS" licensei -o tmp && mv $(find tmp -name "licensei") "$OUT"',
    tools = [CONFIG.JARCAT_TOOL],
    visibility = ["PUBLIC"],
)

CONFIG.setdefault("MIGRATE_VERSION", "4.11.0")

genrule(
    name = "migrate",
    srcs = [remote_file(
        name = "migrate",
        _tag = "download",
        url = f"https://github.com/golang-migrate/migrate/releases/download/v{CONFIG.MIGRATE_VERSION}/migrate.{CONFIG.HOSTOS}-{CONFIG.HOSTARCH}.tar.gz",
        out = f"migrate-{CONFIG.MIGRATE_VERSION}.tar.gz",
    )],
    outs = ["migrate"],
    binary = True,
    cmd = f'"$TOOL" x "$SRCS" migrate -o tmp && mv $(find tmp -name "migrate.{CONFIG.HOSTOS}-{CONFIG.HOSTARCH}") "$OUT"',
    tools = [CONFIG.JARCAT_TOOL],
    visibility = ["PUBLIC"],
)
