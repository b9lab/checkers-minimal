# Minimal checkers module

This repository demonstrates how to create a minimal checkers module to be integrated into the minimal Cosmos chain found [here](https://github.com/cosmosregistry/chain-minimal).

## Versions

Versions used here are:

* Go: 1.21.1
* Cosmos SDK: v0.50.1

## Progressive feature branches

The project is created with a clean list of commits in order to demonstrate the natural progression of the project. In this sense, there is no late commit that fixes an error introduced earlier. If there is a fix for an error introduced earlier, the fix should be squashed with the earlier commit that introduced the error. This may require some conflict resolution.

Having a clean list of commits makes it possible to do meaningful `compare`s.

To make it easier to link to the content at the different stages of the project's progression, a number of branches have been created at commits that have `Add branch the-branch-name` as message. Be careful with the commit message as it depends on it matching the `"Add branch [0-9a-zA-Z\-]*"` regular expression.

The script `push-branches.sh` is used to extract these commits and force push them to the appropriate branch in the repository. After having made changes, you should run this script, and also manually force push to `main`.

Here is the list of branches:

* [`initial-protobuf`](../../tree/initial-protobuf)
* [`prevent-error`](../../tree/prevent-error), [diff](../../compare/initial-protobuf..prevent-error)
* [`first-proto-gen`](../../tree/first-proto-gen), [diff](../../compare/prevent-error..first-proto-gen)
