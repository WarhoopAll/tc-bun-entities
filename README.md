### TrinityCore Database Models for Go (BUN ORM)

This repository contains Go models for TrinityCore databases (`auth`, `world`, and `characters`) designed for use with [Bun ORM](https://bun.uptrace.dev).  
It simplifies integration with TrinityCore databases, allowing for faster development of custom tools and services in Go.

**Databases:**
- `auth` – Authentication and account data.
- `world` – Game world data, creatures, items, etc.
- `characters` – Player character data.

**Usage:**
```go
import "github.com/WarhoopAll/tc-bun-entities/auth"
import "github.com/WarhoopAll/tc-bun-entities/world"
import "github.com/WarhoopAll/tc-bun-entities/characters"
