# FokBomb

`FokBomb` is an amateur implementation of a Windows 
fork bomb that recursively self replicates into the users Startup Directory, then forks that copy of the application. 
This way, after the system shuts down, the computer starts back up with the new instances of the application. Each instance forking and copying itself into the Startup Directory.

Since there is no reliable way to have an app run on start for Linux, `FokBomb` will do something a little different. 
`FokBomb` will still replicate itself and fork that instance, but the location of the new instance is randomized per app launch.


When ran on a machine with a NVMe SSD, it could write upwards of 1GB/s (Each copy of the executable being ~1.7MB). These speeds might be slower when running on a SATA SSD, and even slower on a SATA HDD. These speeds do not factor in drive utilization for other applications or the OS itself, and were taken in a clean environment.

# Releases

There are several binaries

| Name | Target OS | Working | Obfuscated |
| --- | --- | --- | --- |
| [fokbomb](./build/fokbomb) | Linux x64 | No | No |
| [fokbom_debug.exe](./build/fokbomb_debug.exe) | Windows x64 | Yes | No |
| [fokbomb_garbled](./build/fokbomb_garbled) | Linux x64 | TBD | Yes |
| [fokbomb_garbled.exe](./build/fokbomb_garbled.exe) | Windows x64 | Yes | Yes |
| [fokbomb.exe](./build/fokbomb.exe) | Windows x64 | Yes | No |

# Arguments

To enter a forced debug mode even when compiled without, give the argument `::`. This argument can be anywhere, it
just needs to appear at least once as a standalone string. This will cause `FokBomb` to exit just after getting all configurations.
With this enabled, nothing will be copied or spawned.

Example:

```bash
$ ./fokbomb.exe this text will be ignored ::
```
# Project Layout

<img src="./flowchart.drawio.svg" alt="Flowchart diagram">

# Disclaimer

All the binaries of `FokBomb` should be used for authorized penetration testing and/or educational purposes only. Any misuse of this software will not be the responsibility of the author or of any other collaborator. Use it at your own machines and/or with the machine owner's permission.

`FokBomb` is licensed under [MIT](./LICENSE)