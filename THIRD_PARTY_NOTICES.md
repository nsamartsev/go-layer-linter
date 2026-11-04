# Third-party notices

The command uses the following third-party Go modules. The corresponding license texts are included in [`licenses/`](licenses/).

| Module | Version | License | License text |
| --- | --- | --- | --- |
| `github.com/spf13/cobra` | v1.9.1 | Apache-2.0 | [`Apache-2.0.txt`](licenses/Apache-2.0.txt) |
| `github.com/inconshreveable/mousetrap` | v1.1.0 | Apache-2.0 | [`Apache-2.0-mousetrap.txt`](licenses/Apache-2.0-mousetrap.txt) |
| `github.com/spf13/pflag` | v1.0.6 | BSD-3-Clause | [`BSD-3-Clause-pflag.txt`](licenses/BSD-3-Clause-pflag.txt) |
| `gopkg.in/yaml.v2` | v2.4.0 | Apache-2.0; libyaml-derived files retain MIT notices | [`Apache-2.0-yaml.v2.txt`](licenses/Apache-2.0-yaml.v2.txt), [`yaml.v2-NOTICE.txt`](licenses/yaml.v2-NOTICE.txt), [`MIT-libyaml.txt`](licenses/MIT-libyaml.txt) |

These are the runtime modules used by the command (including the Windows-only `mousetrap` dependency). The Go module graph also contains documentation and test dependencies that are not linked into the released command.
