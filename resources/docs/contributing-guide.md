# Contribution Guide

- [# Folder Structure](#-folder-structure)
- [# For Security Issues](#-for-security-issues)

---

We welcome all developers to contribute to this project, this documentation will help us achieve to work on the same standards.

{#-folder-structure}

## [#](#-folder-structure) Folder Structure

- `/.build/`
- `/.vscode/`
- `/app/`
- `/cmd/`
- `/databases/`
- `/internal/`
- `/pkg/`
- `/registrar/`
- `/resources/`
- `/storage/`
  - `/framework`
    - ...
  - `/logs`
    - ...
- `/stubs/`
- `/env & /env.local`
  - It is the place we store all configurations, the `.env` holds the default config, while the `env.{APP_ENV` will be loaded after we are able to capture the `APP_ENV`
- `/serve, /build, /preview, & /run`
  - These are shell commands to help you **serve**, **build**, **preview** or **run** a console command.

{#-for-security-issues}

## [#](#-for-security-issues) For Security Issues

If you found any security concerns, please send a direct email to **daison12006013@gmail.com**, the title of the email should have at least a word "Lucid".
