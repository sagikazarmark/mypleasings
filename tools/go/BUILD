CONFIG.setdefault("GOLANGCI_LINT_VERSION", "1.28.3")

download_binary(
    binary_name = "golangci-lint",
    url = f"https://github.com/golangci/golangci-lint/releases/download/v{CONFIG.GOLANGCI_LINT_VERSION}/golangci-lint-{CONFIG.GOLANGCI_LINT_VERSION}-{CONFIG.HOSTOS}-{CONFIG.HOSTARCH}.tar.gz",
    download_out = f"golangci-lint-{CONFIG.GOLANGCI_LINT_VERSION}.tar.gz",
    visibility = ["PUBLIC"],
)

CONFIG.setdefault("GOTESTSUM_VERSION", "0.5.1")

download_binary(
    binary_name = "gotestsum",
    url = f"https://github.com/gotestyourself/gotestsum/releases/download/v{CONFIG.GOTESTSUM_VERSION}/gotestsum_{CONFIG.GOTESTSUM_VERSION}_{CONFIG.HOSTOS}_{CONFIG.HOSTARCH}.tar.gz",
    download_out = f"gotestsum-{CONFIG.GOTESTSUM_VERSION}.tar.gz",
    visibility = ["PUBLIC"],
)

CONFIG.setdefault("MGA_VERSION", "0.2.1")

download_binary(
    binary_name = "mga",
    url = f"https://github.com/sagikazarmark/mga/releases/download/v{CONFIG.MGA_VERSION}/mga_{CONFIG.HOSTOS}_{CONFIG.HOSTARCH}.tar.gz",
    download_out = f"mga-{CONFIG.MGA_VERSION}.tar.gz",
    visibility = ["PUBLIC"],
)
