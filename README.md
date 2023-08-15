# Dotfiles

This repository contains my configuration for my workstation and tooling to keep it up to date.
All files are mapped from `rootfs` to `/` on the local filesystem.  

When setting up the workstation the first time do the following:
```
./bootstrap
```
this will install all prerequisites needed.

## TODO 

**Sync config**

```
df sync
```

**Upgrade packages and config**
```
df upg
```

**Only update configs**
```
df up --cfg
df up -c
```

**Update packages only**
```
df upgrade --pkg
df upgrade -p
```




