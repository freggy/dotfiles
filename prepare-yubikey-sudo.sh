#!/bin/bash
ykman otp chalresp 2 -g -t  
mkdir -m0755 -p ~/.yubico
sudo ykpamcfg -2
