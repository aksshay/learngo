# Gorm

### Installation

```bash
go get -u gorm.io/gorm
go get -u gorm.io/driver/sqlite
```

```sql
create table users (
    ID          bigserial PRIMARY KEY,
    user_name   text,
    first_name  text,
    last_name   text,
    email       text,
    password    text,
    created_at  timestamp,
    deleted_at  timestamp,
    updated_at  timestamp
);
```

### Created service

Added services.go file that is responsible for creating and returning a GORM DB
object.

### models

Updated users.go to handle the crud operations. Emdedded gorm model struct is
used for primary key ID.

### controllers

Updated the controllers file, that create a struct from models and it is used
for calling all the crud related operations.
