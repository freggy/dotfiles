# append to .bashrc
PATH=$PATH:/home/yannic/.local/bin
eval "$(direnv hook bash)"
PATH=$PATH:/usr/local/kubebuilder/bin
export PATH=$PATH:/home/yannic/.linkerd2/bin
export PATH=$PATH:/usr/bin/maven/bin
export PATH=$PATH:/home/yannic/bin
export GO111MODULE="on"
