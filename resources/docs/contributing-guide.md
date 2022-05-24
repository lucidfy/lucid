# Contribution Guide

- [# Standards](#-standards)
  - [# Encapsulation](#-encapsulation)
- [# For Security Issues](#-for-security-issues)

---

We welcome all developers to contribute to this project, this documentation will help us achieve to work on the same standards.

{#-standards}

## [#](#-standards) Standards

{#-encapsulation}

### [#](#-encapsulation) Encapsulation

Scoped variables shall be declared as `snake_case`, while functions shall be in `lowerCamelCase`, for example:

```go
var first_name string = "John"
last_name := "Doe"
func getMethod() {}
func (ms *MyStruct) getFiles() {}
```

Accessible variables & functions shall be in `UpperCamelCase`, for example:

```go
func GetKey() {}
func (ms *MyStruct) District() {}
```

{#-for-security-issues}

## [#](#-for-security-issues) For Security Issues

If you found any security concerns, please send a direct email to **daison12006013@gmail.com**, the title of the email should have at least a word "Lucid".
