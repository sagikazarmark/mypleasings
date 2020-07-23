CONFIG.setdefault("HELM_VERSION", "3.2.4")

download_binary(
    binary_name = "helm",
    url = f"https://get.helm.sh/helm-v{CONFIG.HELM_VERSION}-{CONFIG.HOSTOS}-{CONFIG.HOSTARCH}.tar.gz",
    download_out = f"helm-{CONFIG.HELM_VERSION}.tar.gz",
    visibility = ["PUBLIC"],
)
