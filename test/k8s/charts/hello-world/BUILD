subinclude("//k8s")

helm_package(
    name = "hello-world",
)

helm_lint(
    name = "lint",
    chart = ":hello-world",
)

helm_e2e_test(
    name = "e2e",
    chart = ":hello-world",
    kubeconfig = "//test/k8s:kubeconfig",
    run_tests = True,
)
