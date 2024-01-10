# Snipository

Snipository is a Go-based project that provides a way to manage your bash snippets without leaving your terminal.

## Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes.
you can also install snipository as a ZSH plugin.

### Prerequisites

- Go
- ZSH Shell

### Building

To build a local version of the project, run the following command:

```makefile
make build-local
```
This will create a binary named `snipository` in your current directory.

### Installing as a ZSH Plugin

To install snipository as a ZSH plugin, first run the following script from zsh:

```shellscript
ZSH_CUSTOM=$ZSH_CUSTOM ./install_as_plugin.sh
```
the `ZSH_CUSTOM` variable is the path to your custom plugins directory, it should be set if you are using oh-my-zsh.
it is required here like this because the script runs on a different shell session, and thus, does not have access to the original`ZSH_CUSTOM` variable.

this will:
1. create a directory named `snipository` in your custom plugins directory.
2. copy the `snipository.plugin.zsh` file to the newly created directory.
3. copy the `snipository` binary to the newly created directory.
4. adding `snipository` completion command to your .zshrc
5. adding `export HISTFILE=$HISTFILE` to your .zshrc (this is required for the `snipository` completion command to work properly
   the HISTFILE is the path for the file that contains the history of your commands, and required as env var that the snipository binary can access)

next, you need to add `snipository` to your plugins list in your `.zshrc` file:
```shellscript
plugins=(... snipository)
```

Then, restart your ZSH shell to use snipository.

