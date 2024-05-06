# FokBomb

`FokBomb` is an amateur implementation of a Windows 
fork bomb that recursively self replicates into the users Startup Directory, then forks that copy of the application. 
This way, after the system shuts down, the computer starts back up with the new instances of the application. Each instance forking and copying itself into the Startup Directory.

This application was written in [Go](https://go.dev/dl/), so is very fast.
When ran on a machine with a NVMe SSD, it could write upwards of 1GB/s. These speeds might be slower when running on a SATA SSD, and even slower on a SATA HDD.

# Disclaimer

All the binaries of `FokBomb` should be used for authorized penetration testing and/or educational purposes only. Any misuse of this software will not be the responsibility of the author or of any other collaborator. Use it at your own machines and/or with the machine owner's permission.

`FokBomb` is licensed under [MIT](./LICENSE)