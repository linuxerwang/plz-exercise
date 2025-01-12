package(default_visibility = ["PUBLIC"])

go_toolchain(
    name = "toolchain",
    version = "1.23.4",
)

go_stdlib(name = "std")

go_repo(
    name = "etcd_api",
    licences = ["Apache-2.0"],
    module = "go.etcd.io/etcd/api/v3",
    version = "v3.5.17",
)

go_repo(
    name = "etcd_cli",
    licences = ["Apache-2.0"],
    module = "go.etcd.io/etcd/client/v3",
    version = "v3.5.17",
    deps = [
        ":grpc",
        ":zap",
    ],
)

go_repo(
    name = "grpc",
    module = "google.golang.org/grpc",
    version = "v1.69.2",
)

go_repo(
    name = "zap",
    module = "go.uber.org/zap",
    version = "v1.27.0",
)
