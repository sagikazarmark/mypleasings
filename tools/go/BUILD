subinclude("//go:tools")

CONFIG.setdefault("GOLANGCI_LINT_VERSION", "1.28.3")

golangci_lint_binary(
    name = "golangci-lint",
    version = CONFIG.GOLANGCI_LINT_VERSION,
    visibility = ["PUBLIC"],
)

CONFIG.setdefault("GOTESTSUM_VERSION", "0.5.1")

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
