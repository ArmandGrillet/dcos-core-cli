# CHANGELOG

## 1.14-patch.x

* Breaking changes

  * `dcos task ls` without any argument to get the list of all tasks files is not supported anymore.

* Features

  * Add `dcos task download` to download task sandbox files
  * Add a `--user` option to `dcos task exec`
  * Add an `--all` option to `dcos node log`
  * Add job task ID(s) when printing history with `dcos job history --json`

## 1.13-patch.0

* Features

  * Add color support to `dcos node log`
  * Add a public IP field to `dcos node list`
  * Add `--user` flag to `dcos service log`
  * Add journalctl format options to `dcos node log`: `json-pretty`, `json`, `cat`, `short`