# verify that ZSH_CUSTOM is set
if [ -z "$ZSH_CUSTOM" ]; then
  echo "ZSH_CUSTOM is not set. Please set it to your custom plugins directory."
  exit 1
fi

#  create the folder if it doesn't exist
if [ ! -d "$ZSH_CUSTOM/plugins/snipository" ]; then
  echo "creating $ZSH_CUSTOM/plugins/snipository"
  mkdir $ZSH_CUSTOM/plugins/snipository
fi

# copy the files
echo "installing snipository plugin to $ZSH_CUSTOM=/plugins/snipository..."
cp ./snipository $ZSH_CUSTOM/plugins/snipository/snipository
cp ./snipository.plugin.zsh $ZSH_CUSTOM/plugins/snipository/snipository.plugin.zsh
echo "successfully installed snipository plugin to $ZSH_CUSTOM/plugins/snipository"

echo "adding snipository completion command to your .zshrc"
echo "source <(snipository completion zsh)" >> $HOME/.zshrc
echo "NOTE: if you reapetedly run this script, you will have multiple lines of this command in your .zshrc"

echo "adding export of the history file to your .zshrc"
echo "export HISTFILE=\$HISTFILE" >> $HOME/.zshrc

echo "restart your zsh shell to use snipository"