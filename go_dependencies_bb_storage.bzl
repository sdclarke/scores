load("@bazel_gazelle//:deps.bzl", "go_repository")

def bb_storage_go_dependencies():
    go_repository(
        name = "com_github_google_uuid",
        importpath = "github.com/google/uuid",
        sha256 = "7e330758f7c81d9f489493fb7ae0e67d06f50753429758b64f25ad5fb2727e21",
        strip_prefix = "uuid-1.1.0",
        urls = ["https://github.com/google/uuid/archive/v1.1.0.tar.gz"],
    )

    go_repository(
        name = "com_github_prometheus_client_golang",
        importpath = "github.com/prometheus/client_golang",
        sha256 = "5f6ca8740a08622ae0a19c32b1026b8233bfd943a1f4af34963d326ab5fa94e5",
        strip_prefix = "client_golang-0.9.2",
        urls = ["https://github.com/prometheus/client_golang/archive/v0.9.2.tar.gz"],
    )

    go_repository(
        name = "com_github_google_go_jsonnet",
        commit = "0959f85501584da690e34871b31e280242226e1f",
        importpath = "github.com/google/go-jsonnet",
        patches = ["@com_github_buildbarn_bb_storage//:patches/com_github_google_go_jsonnet/astgen.diff"],
    )

    go_repository(
        name = "com_github_gorilla_mux",
        importpath = "github.com/gorilla/mux",
        commit = "49c01487a141b49f8ffe06277f3dca3ee80a55fa",
    )

    go_repository(
        name = "org_golang_google_grpc",
        build_file_proto_mode = "disable",
        commit = "6af3d372ceca1a2c17f8c7789446a3488a91b8c6",
        importpath = "google.golang.org/grpc",
    )

    go_repository(
        name = "com_github_prometheus_common",
        importpath = "github.com/prometheus/common",
        sha256 = "3a02a3c8d881ef0be78fc58d63d14979ce0226f91ab52b2d67a11bc120e054dd",
        strip_prefix = "common-0.2.0",
        urls = ["https://github.com/prometheus/common/archive/v0.2.0.tar.gz"],
    )

    go_repository(
        name = "com_github_prometheus_client_model",
        commit = "5c3871d89910bfb32f5fcab2aa4b9ec68e65a99f",
        importpath = "github.com/prometheus/client_model",
    )

    go_repository(
        name = "com_github_prometheus_procfs",
        commit = "ae68e2d4c00fed4943b5f6698d504a5fe083da8a",
        importpath = "github.com/prometheus/procfs",
    )

    go_repository(
        name = "com_github_prometheus_client_golang",
        importpath = "github.com/prometheus/client_golang",
        sha256 = "5f6ca8740a08622ae0a19c32b1026b8233bfd943a1f4af34963d326ab5fa94e5",
        strip_prefix = "client_golang-0.9.2",
        urls = ["https://github.com/prometheus/client_golang/archive/v0.9.2.tar.gz"],
    )

    go_repository(
        name = "com_github_matttproud_golang_protobuf_extensions",
        commit = "c12348ce28de40eed0136aa2b644d0ee0650e56c",
        importpath = "github.com/matttproud/golang_protobuf_extensions",
    )

    go_repository(
        name = "com_github_beorn7_perks",
        commit = "3a771d992973f24aa725d07868b467d1ddfceafb",
        importpath = "github.com/beorn7/perks",
    )
