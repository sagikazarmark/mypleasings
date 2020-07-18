package(default_visibility = ["PUBLIC"])

CONFIG.setdefault("GOLANGCI_LINT_VERSION", "1.28.3")

genrule(
    name = "golangci-lint",
    srcs = [remote_file(
        name = "golangci-lint",
        _tag = "download",
        url = f"https://github.com/golangci/golangci-lint/releases/download/v{CONFIG.GOLANGCI_LINT_VERSION}/golangci-lint-{CONFIG.GOLANGCI_LINT_VERSION}-{CONFIG.HOSTOS}-{CONFIG.HOSTARCH}.tar.gz",
        out = f"golangci-lint-{CONFIG.GOLANGCI_LINT_VERSION}.tar.gz",
    )],
    outs = ["golangci-lint"],
    binary = True,
    cmd = '"$TOOL" x "$SRCS" golangci-lint -o tmp && mv $(find tmp -name "golangci-lint") "$OUT"',
    tools = [CONFIG.JARCAT_TOOL],
)

CONFIG.setdefault("GOTESTSUM_VERSION", "0.5.1")

genrule(
    name = "gotestsum",
    srcs = [remote_file(
        name = "gotestsum",
        _tag = "download",
        url = f"https://github.com/gotestyourself/gotestsum/releases/download/v{CONFIG.GOTESTSUM_VERSION}/gotestsum_{CONFIG.GOTESTSUM_VERSION}_{CONFIG.HOSTOS}_{CONFIG.HOSTARCH}.tar.gz",
        out = f"gotestsum-{CONFIG.GOTESTSUM_VERSION}.tar.gz",
    )],
    outs = ["gotestsum"],
    binary = True,
    cmd = '"$TOOL" x "$SRCS" gotestsum -o tmp && mv $(find tmp -name "gotestsum") "$OUT"',
    tools = [CONFIG.JARCAT_TOOL],
)

CONFIG.setdefault("MGA_VERSION", "0.2.1")

genrule(
    name = "mga",
    srcs = [remote_file(
        name = "mga",
        _tag = "download",
        url = f"https://github.com/sagikazarmark/mga/releases/download/v{CONFIG.MGA_VERSION}/mga_{CONFIG.HOSTOS}_{CONFIG.HOSTARCH}.tar.gz",
        out = f"mga-{CONFIG.MGA_VERSION}.tar.gz",
    )],
    outs = ["mga"],
    binary = True,
    cmd = '"$TOOL" x "$SRCS" mga -o tmp && mv $(find tmp -name "mga") "$OUT"',
    tools = [CONFIG.JARCAT_TOOL],
)
