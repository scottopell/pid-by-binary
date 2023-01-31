# pid-by-binary

Given a path to an executable binary, print the pids found running that binary.

Only support systems with a `/proc` filesystem mounted

Example:
```
$ ./pid-by-binary -m ./bin/agent/agent
Pid: 506406     MemInfo: {"rss":117919744,"vms":1619230720,"hwm":0,"data":0,"stack":0,"locked":0,"swap":0}
```