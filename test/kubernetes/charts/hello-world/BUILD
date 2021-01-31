subinclude("//kubernetes")

helm_docs(
    name = "README.md",
)

helm_docs(
    name = "README2.md",
    out = "README2.md",
    templates = ["README1.gotmpl.md", "README2.gotmpl.md"],
)

helm_chart(
    name = "chart",
    srcs = [":README.md"],
)

helm_lint(
    name = "lint",
    chart = ":chart",
)

helm_package(
    name = "hello-world",
    chart = ":chart",
)

helm_e2e_test(
    name = "e2e",
    chart = ":hello-world",
    kubeconfig = "//test/kubernetes:kubeconfig",
    labels = ["kind"],
    run_tests = True,
)
