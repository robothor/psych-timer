load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library", "go_embed_data")
load("@bazel_gazelle//:def.bzl", "gazelle")

# gazelle:prefix github.com/robothor/psych-timer
gazelle(name = "gazelle")

go_library(
    name = "go_default_library",
    srcs = [
        "main.go",
        "mindware_emitter.go",
        "psych_timer.go",
        ":static",  # keep
    ],
    importpath = "github.com/robothor/psych-timer",
    visibility = ["//visibility:private"],
    deps = [
        "@com_github_docopt_docopt_go//:go_default_library",
        "@com_github_faiface_beep//:go_default_library",
        "@com_github_faiface_beep//speaker:go_default_library",
        "@com_github_faiface_beep//wav:go_default_library",
        "@com_github_gorilla_websocket//:go_default_library",
        "@com_github_mitchellh_copystructure//:go_default_library",
        "@com_github_skratchdot_open_golang//open:go_default_library",
        "@com_github_spf13_viper//:go_default_library",
    ],
)

go_binary(
    name = "psych-timer",
    data = [
        "beep.wav",
        "beepbeep.wav",
        "interoception_config.yaml",
        "hrv_config.yaml",
    ],
    embed = [
        ":go_default_library",
    ],
    visibility = ["//visibility:public"],
)

go_embed_data(
    name = "static",
    package = "main",
    srcs = glob(["static/**"]),
    string = True,
)
