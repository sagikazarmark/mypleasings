subinclude("//container")

hadolint(
    name = "hadolint",
)

hadolint(
    name = "hadolint_config",
    dockerfile = "Dockerfile.config",
    config = "config.yaml",
)

hadolint(
    name = "hadolint_trusted_registry",
    dockerfile = "Dockerfile.config",
    trusted_registries = ["quay.io", ""],
)

hadolint(
    name = "hadolint_ignore",
    dockerfile = "Dockerfile.bad",
    ignore = ["DL3006", "DL3007"],
)
