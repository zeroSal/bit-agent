[![Release](https://github.com/zeroSal/bit-agent/actions/workflows/go.yml/badge.svg)](https://github.com/zeroSal/bit-agent/actions/workflows/go.yml)
# bit-agent
Unofficial Bitwarden SSH agent.

![bit-agent-icon](https://github.com/zeroSal/bit-agent/assets/38191926/a6baaad0-61be-4305-b55b-d78935edb00e)


The icon is a mashup of two [Freepik](https://www.flaticon.com/free-icons/partnership) icons.

## Why this project?
Bitwarden stands out as one of the premier password managers available today. However, it lacks a critical feature for software developers: the management and injection of SSH keys. This project addresses this significant shortcoming, thereby enabling even advanced users to fully appreciate and be satisfied with this password manager.

## OS compatibility
 - Linux
 - MacOS

## Architecture compatibility
 - x64 (both Intel and AMD)
 - arm64

## Dependecies
 - [Bitwarden CLI](https://bitwarden.com/help/cli/)

## Installation
The binary has a built-in wizard that installs the softwares on the system.

### Download the binary
The release binary name is like `bit-agent-{OS}-{ARCH}`, where `OS` is the kernel name and the `ARCH` is the CPU architecture.
```bash
# Linux on x64
curl https://github.com/zeroSal/bit-agent/releases/latest/download/bit-agent-linux-amd64

# Linux on ARM
curl https://github.com/zeroSal/bit-agent/releases/latest/download/bit-agent-linux-arm64

# MacOS on x64
curl https://github.com/zeroSal/bit-agent/releases/latest/download/bit-agent-darwin-amd64

# MacOS on ARM
curl https://github.com/zeroSal/bit-agent/releases/latest/download/bit-agent-darwin-arm64
```

### Set the execution permissions
The binary needs execution permissions.
```bash
chmod +x bit-agent
```

### Run the install procedure
This procedure needs root privileges.
```bash
sudo ./bit-agent --install
```

## Run it
A symlink to `/usr/local/bin` is created, so the user can run the binary directly from the PATH.
```
bit-agent
```