load("@bazel_gazelle//:def.bzl", "DEFAULT_LANGUAGES", "gazelle", "gazelle_binary")

# gazelle:prefix github.com/arianics/bazel-buf-demo
gazelle(
    name = "gazelle",
    gazelle = ":gazelle-buf",
)

gazelle(
    name = "gazelle-update-go-repos",
    args = [
        "-from_file=go.mod",
        "-to_macro=deps.bzl%go_dependencies",
        "-prune",
    ],
    command = "update-repos",
    gazelle = ":gazelle-buf",
)

gazelle(
    name = "gazelle-update-buf-repos",
    args = [
        # This can also be `buf.yaml` and `buf.lock`.
        "--from_file=buf.work.yaml",
        # This is optional but recommended, if absent gazelle
        # will add the rules directly to WORKSPACE
        "-to_macro=deps.bzl%buf_dependencies",
        # Deletes outdated repo rules
        "-prune",
    ],
    command = "update-repos",
    gazelle = ":gazelle-buf",
)

gazelle_binary(
    name = "gazelle-buf",
    languages = DEFAULT_LANGUAGES + [
        # Loads the Buf extension
        "@rules_buf//gazelle/buf:buf",
        # NOTE: This needs to be loaded after the proto language
    ],
)
