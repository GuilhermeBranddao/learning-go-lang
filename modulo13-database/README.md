# 🗄️ Módulo 13 - Banco de Dados

## 🎯 Objetivos
Trabalhar com bancos de dados em Go usando `database/sql`.

## 🔑 Comparação com Python

| Aspecto | Python | Go |
|---------|--------|-----|
| **ORM** | SQLAlchemy | GORM, sqlx |
| **Driver** | psycopg2, mysql | lib/pq, go-sql-driver |
| **Raw SQL** | cursor.execute() | db.Query(), db.Exec() |

---

## 📝 Básico (SQLite)

```go
import (
    "database/sql"
    _ "github.com/mattn/go-sqlite3"
)

db, err := sql.Open("sqlite3", "database.db")
defer db.Close()

// Query
rows, err := db.Query("SELECT * FROM users")
defer rows.Close()

for rows.Next() {
    var id int
    var nome string
    rows.Scan(&id, &nome)
}

// Insert
result, err := db.Exec("INSERT INTO users (nome) VALUES (?)", "João")
```

---

## 📋 Tópicos

1. **database/sql**
2. **CRUD operations**
3. **Prepared statements**
4. **Transactions**
5. **GORM (ORM)**

---

**Anterior:** [Módulo 12 - HTTP](../modulo12-http/)  
**Próximo:** [Módulo 14 - Context](../modulo14-context/)
