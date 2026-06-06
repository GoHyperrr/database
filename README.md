# database — Hyperrr Pluggable Database Drivers Module

[![Go Reference](https://pkg.go.dev/badge/github.com/GoHyperrr/database.svg)](https://pkg.go.dev/github.com/GoHyperrr/database)

This repository contains the pluggable database driver adapters for the Hyperrr engine, specifically GORM PostgreSQL dialect configuration.

---

## 🛠️ Active Providers

* **PostgreSQL GORM Dialect**: Dynamically registers the GORM PostgreSQL driver and connection provider on initialization.

---

## 🛠️ Usage

This module is imported dynamically by the core Hyperrr engine during startup to support external PostgreSQL storage backends.

To learn more about how to configure database backends, see the [Hyperrr Developer Guide](https://github.com/GoHyperrr/hyperrr/blob/main/developer.md).
