subinclude("//kubernetes:tools")

# https://github.com/kubernetes/kubectl/releases
CONFIG.setdefault("KUBECTL_VERSION", "1.21.0")

kubectl_binary(
    name = "kubectl",
    version = CONFIG.KUBECTL_VERSION,
    visibility = ["PUBLIC"],
)

# https://github.com/kubernetes-sigs/kind/releases
CONFIG.setdefault("KIND_VERSION", "0.10.0")

kind_binary(
    name = "kind",
    version = CONFIG.KIND_VERSION,
    visibility = ["PUBLIC"],
)

# https://github.com/helm/helm/releases
CONFIG.setdefault("HELM_VERSION", "3.5.4")

helm_binary(
    name = "helm",
    version = CONFIG.HELM_VERSION,
    visibility = ["PUBLIC"],
)

# https://github.com/norwoodj/helm-docs/releases
CONFIG.setdefault("HELM_DOCS_VERSION", "1.5.0")

helm_docs_binary(
    name = "helm-docs",
    version = CONFIG.HELM_DOCS_VERSION,
    visibility = ["PUBLIC"],
)
