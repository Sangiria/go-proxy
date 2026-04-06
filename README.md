# go-proxy

> ⚠️ Disclaimer: This project is under active development and is currently in an "alpha" state. Expect bugs, incomplete features, and breaking changes.

A desktop GUI application for managing proxy nodes on Linux. It allows you to manually add servers or use subscriptions to start/stop proxying with a single click.
Traffic is routed via [Xray](https://github.com/xtls/xray-core).

## Requirements
- OS: Linux x86_64 (Tested on Ubuntu/Linux Mint/Kubuntu)
- Dependencies: The bundle includes most requirements, but ensure your system has standard libraries for Qt/PySide6 applications.

## Quick Start

### 1. Download the Release
Go to the [Releases page](https://github.com/Sangiria/go-proxy/releases) and download the latest .zip archive.

### 2. Extract and Run
To start the application, simply run `./run.sh` in your terminal.

> 💡Local Proxy Note: This application provides a local SOCKS5/HTTP proxy (default: 127.0.0.1:20808). It does not automatically route your entire system traffic. You must manually configure your browser or other applications to use this local address.
