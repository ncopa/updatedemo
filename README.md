# Auto respawn on update demo

Demo on how an executable can respawn after an update

## How to build
```
go build updatedemo.go
```


## How to run
```
./updatedemo
```

## Rebuild a new version
From a different console, rebuild it with a new version:
```
go build -ldflags "-X main.Version=0.2" updatedemo.go
```

The program should replace the running process with the newer version.
It will keep the same PID and a supervisor like systemd will not notice it.

