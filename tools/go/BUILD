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
