subinclude("//k8s:tools")

CONFIG.setdefault("HELM_VERSION", "3.2.4")

helm_binary(
    name = "helm",
    version = CONFIG.HELM_VERSION,
    visibility = ["PUBLIC"],
)
