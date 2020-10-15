subinclude("//go:tools")

CONFIG.setdefault("GOLANGCI_LINT_VERSION", "1.31.0")

golangci_lint_binary(
    name = "golangci-lint",
    labels = ["go"],
    version = CONFIG.GOLANGCI_LINT_VERSION,
    visibility = ["PUBLIC"],
)

CONFIG.setdefault("GOTESTSUM_VERSION", "0.5.3")

gotestsum_binary(
    name = "gotestsum",
    labels = ["go"],
    version = CONFIG.GOTESTSUM_VERSION,
    visibility = ["PUBLIC"],
)

CONFIG.setdefault("MGA_VERSION", "0.4.1")

mga_binary(
    name = "mga",
    labels = ["go"],
    version = CONFIG.MGA_VERSION,
    visibility = ["PUBLIC"],
)

CONFIG.setdefault("PKGER_VERSION", "0.17.1")

pkger_binary(
    name = "pkger",
    labels = ["go"],
    version = CONFIG.PKGER_VERSION,
    visibility = ["PUBLIC"],
)

CONFIG.setdefault("GO_BINDATA_VERSION", "3.21.0")

go_bindata_binary(
    name = "go-bindata",
    labels = ["go"],
    version = CONFIG.GO_BINDATA_VERSION,
    visibility = ["PUBLIC"],
)
