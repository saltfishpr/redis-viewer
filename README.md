# Redis Viewer

A tool to view Redis data in terminal.

![user interface](images/ui.png)

## Install

Go version >=1.16: `go install github.com/SaltFishPr/redis-viewer@latest`

or

Download from [Release](https://github.com/SaltFishPr/redis-viewer/releases).

## Usage:

| KeyBoard | Description            |
| :------: | :--------------------- |
|  ctrl+c  | exit redis viewer      |
|    ↑     | previous key           |
|    ↓     | next key               |
|    ←     | previous page          |
|    →     | next page              |
|    r     | reload data from redis |
|    s     | enter search keywords  |
|  enter   | apply scan with text   |

|    Mouse    | Description        |
| :---------: | :----------------- |
|  scroll up  | detail scroll up   |
| scroll down | detail scroll down |

config file directory:

-   Windows: `%USERPROFILE%/redis-viewer.yml`
-   Linux: `~/.config/redis-viewer/redis-viewer.yml`

example config file:

```yaml
addrs:
    - 127.0.0.1:6380
    - 127.0.0.1:6381
    - 127.0.0.1:6382
    - 127.0.0.1:6383
    - 127.0.0.1:6384
    - 127.0.0.1:6385

db:
username:
password:

master_name:
```

## Support:

-   client, sentinel and cluster mode.
-   `string, hash, list, set, zset` key types.

## Note:

In Windows, you should change system encoding to `UTF-8` before run this program.

## TODOs:

-   [x] Add load animation.
-   [ ] Friendly value detail.
-   [ ] Add log view.

Build with [bubbletea](https://github.com/charmbracelet/bubbletea).
