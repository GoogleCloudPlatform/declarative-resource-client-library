# Builds

This code is meant to be imported using the standard `go build` rules, but it
also supports being built by bazel.  These two build systems are kept in sync
by a tool called "gazelle".  Gazelle will need to be run by hand from time to
time.

## When you add a new dependency, run
```
$ bazel run :gazelle -- update-repos -from_file=go.mod
```

## When you add a new file, run
```
$ bazel run :gazelle
```

Then commit your changes and push them up.  This should only modify WORKSPACE,
BUILD, and BUILD.bazel files.
