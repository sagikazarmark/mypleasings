subinclude("//go:tools")

# https://github.com/golangci/golangci-lint/releases
CONFIG.setdefault("GOLANGCI_LINT_VERSION", "1.37.1")

golangci_lint_binary(
    name = "golangci-lint",
    labels = ["go"],
    version = CONFIG.GOLANGCI_LINT_VERSION,
    visibility = ["PUBLIC"],
)

# https://github.com/gotestyourself/gotestsum/releases
CONFIG.setdefault("GOTESTSUM_VERSION", "1.6.2")

gotestsum_binary(
    name = "gotestsum",
    labels = ["go"],
    version = CONFIG.GOTESTSUM_VERSION,
    visibility = ["PUBLIC"],
)

# https://github.com/sagikazarmark/mga/releases
CONFIG.setdefault("MGA_VERSION", "0.4.2")

mga_binary(
    name = "mga",
    labels = ["go"],
    version = CONFIG.MGA_VERSION,
    visibility = ["PUBLIC"],
)

# https://github.com/markbates/pkger/releases
CONFIG.setdefault("PKGER_VERSION", "0.17.1")

pkger_binary(
    name = "pkger",
    labels = ["go"],
    version = CONFIG.PKGER_VERSION,
    visibility = ["PUBLIC"],
)

# https://github.com/kevinburke/go-bindata/releases
CONFIG.setdefault("GO_BINDATA_VERSION", "3.22.0")

go_bindata_binary(
    name = "go-bindata",
    labels = ["go"],
    version = CONFIG.GO_BINDATA_VERSION,
    visibility = ["PUBLIC"],
)

# https://github.com/tcncloud/wollemi/releases
CONFIG.setdefault("WOLLEMI_VERSION", "0.2.1")

wollemi_binary(
    name = "wollemi",
    labels = ["go"],
    version = CONFIG.WOLLEMI_VERSION,
    visibility = ["PUBLIC"],
)

wollemi_wrapper(
    name = "wollemi-wrapper",
    binary = ":wollemi",
    labels = ["go"],
    visibility = ["PUBLIC"],
)
