subinclude("//go:tools")

CONFIG.setdefault("GOLANGCI_LINT_VERSION", "1.30.0")

golangci_lint_binary(
    name = "golangci-lint",
    version = CONFIG.GOLANGCI_LINT_VERSION,
    visibility = ["PUBLIC"],
)

CONFIG.setdefault("GOTESTSUM_VERSION", "0.5.3")

gotestsum_binary(
    name = "gotestsum",
    version = CONFIG.GOTESTSUM_VERSION,
    visibility = ["PUBLIC"],
)

CONFIG.setdefault("MGA_VERSION", "0.2.1")

mga_binary(
    name = "mga",
    version = CONFIG.MGA_VERSION,
    visibility = ["PUBLIC"],
)

CONFIG.setdefault("PKGER_VERSION", "0.17.0")

pkger_binary(
    name = "pkger",
    version = CONFIG.PKGER_VERSION,
    visibility = ["PUBLIC"],
)
