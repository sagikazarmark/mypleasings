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

CONFIG.setdefault("JQ_VERSION", "1.6")

if CONFIG.HOSTOS == 'linux':
    _jq_host = CONFIG.HOSTOS
elif CONFIG.HOSTOS == 'darwin':
    _jq_host = 'osx'

remote_file(
    name = 'jq',
    url = [
        f"https://github.com/stedolan/jq/releases/download/jq-{CONFIG.JQ_VERSION}/jq-{_jq_host}64",
        f"https://github.com/stedolan/jq/releases/download/jq-{CONFIG.JQ_VERSION}/jq-{_jq_host}-amd64",
    ],
    out = "jq",
    binary = True,
    visibility = ["PUBLIC"],
)
