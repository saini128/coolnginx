# CoolNginx

This CLI tool provides an intuitive way to manage Nginx configurations without requiring in-depth knowledge of Nginx syntax. It allows users to make changes, add new server blocks, configure proxies, and more, through interactive prompts and a simplified interface.

## Features

* **Interactive Configuration:** User-friendly CLI prompts guide you through the configuration process.
* **Simplification of Code:** Using SLMs the tool provides breif information of whats happning with the configs.
* **Abstraction Layers:** Simplifies complex Nginx directives and blocks.
* **Configuration Validation:** Ensures valid Nginx syntax using `nginx -t`.
* **Configuration Storage:** Stores simplified configuration data in a local BoltDB database for fast access and retrieval.
* **Backup Functionality:** Automatically creates backups of original configuration files.
* **Add, Modify, and Delete:** Easily add, modify, and delete server blocks, locations, and directives.
* **Preview Changes:** Shows a preview of changes before saving.
* **Cross-Platform:** Built with Go, providing cross-platform compatibility.

## Getting Started

### Prerequisites

* Go (version 1.22 or later)
* Nginx installed on your system
* BoltDB (installed with go get go.etcd.io/bbolt)

### Initialization

* It is checking if nginx is running or not
* It is check if AI agent (for now support is for groq only) is present or not. If not then it picks env `GROQ_API_KEY` from os

### Configuration Storage

The tool uses BoltDB to store simplified versions of your Nginx configurations. The database file `nginx_configs.db` is created to store all the data.
### Configuration Backups

Original Nginx configuration files are backed up before any changes are made. Backup files are stored in the same directory as the original files, with a `.bak` extension.

### Example Workflow

1.  The tool parses the Nginx configuration and looks for any manual changes made.
2.  The tool updates the user understandable language in the db using AI
2.  The tool prompts the user to select an action (e.g., add a website, configure a proxy).
3.  The tool guides the user through the necessary steps with interactive prompts.
4.  The tool takes backups and modifies the configuration data model based on user input.
5.  The tool validates the modified configuration using `nginx -t`.
6.  The tool saves the modified configuration and creates a backup of the original.

### Dependencies

* `go.etcd.io/bbolt`

### Contributing

Contributions are welcome! Please feel free to submit pull requests or open issues for bug reports or feature requests.

### License

This project is licensed under the [Your License] License.