subinclude("//k8s")

helm_package(
    name = "hello-world",
)

helm_lint(
    name = "lint",
)

helm_package2(
    name = "hello-world2",
)

helm_package2(
    name = "hello-world3",
    templates = "//test/k8s/charts/hello-world/templates2"
)

helm_lint2(
    name = "lint2",
    chart = ":hello-world3",
)

helm_e2e_test("e2e", ":hello-world3")
