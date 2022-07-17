load("@bazel_gazelle//:def.bzl", "gazelle", "gazelle_binary")
load("@rules_proto_grpc//go:defs.bzl", "go_proto_compile")

#load("@rules_proto_grpc//go:defs.bzl", "go_grpc_compile")
load("@rules_proto_grpc//grpc-gateway:defs.bzl", "gateway_grpc_compile")
load("@rules_proto_grpc//grpc-gateway:defs.bzl", "gateway_openapiv2_compile")

gazelle_binary(
    name = "gazelle-protobuf",
    languages = [
        "@bazel_gazelle//language/go",
        "@bazel_gazelle//language/proto",
    ],
)

# 自动生成 BUILD.bazel, 这里排查了包含go文件的目录，我们先不用 bazel 构建 go 代码
# gazelle:prefix github.com/0xhuk/bazel-go-grpc-example
# gazelle:exclude internal
# gazelle:exclude cmd
# gazelle:proto package
# gazelle:proto_group go_package
# gazelle:go_generate_proto false
gazelle(
    name = "gazelle",
    gazelle = ":gazelle-protobuf",
)

# 更新依赖
gazelle(
    name = "gazelle-update-repos",
    args = [
        "-from_file=go.mod",
        "-to_macro=deps.bzl%go_dependencies",
        "-prune",
    ],
    command = "update-repos",
)

# 生成proto文件
go_proto_compile(
    name = "pure-proto",
    protos = [
        "//proto/common:common_proto",
    ],
)

# 生成pb文件
#go_grpc_compile(
#    name = "grpc-proto",
#    protos = [
#        "//proto/common:common_proto",
#    ],
#)

# 生成对应的 grpc-gateway go 文件 构建文件路径看命令输出
gateway_grpc_compile(
    name = "grpc-gateway",
    protos = [
        "//proto:user_proto",
    ],
)

gateway_openapiv2_compile(
    name = "grpc-openapi",
    protos = [
        "//proto:user_proto",
    ],
)
