# sqlb
[![Build Status](https://travis-ci.org/tuanphung/sqlb.svg?branch=master)](https://travis-ci.org/tuanphubg/sqlb) [![Coverage Status](https://coveralls.io/repos/github/tuanphung/sqlb/badge.svg)](https://coveralls.io/github/tuanphung/sqlb)

A lightweight package which provides an convenience way for you to construct SQL statements. We focus on producing SQL statements with arguments and leave it to you to decide how execution is, you can either use the standard `database/sql` or `sqlx`.

## Features
* Support most common SQL statements (SELECT, FROM, WHERE, HAVING, ORDER BY, LIMIT, OFFSET, etc...).
* Statement chaining, e.g: `sb.Select("*").From("users")`.
* Expression combination to build up `WHERE` clause.
* Support Postgresql argument placeholder ($).

## Installation
```go
import "github.com/Masterminds/squirrel"
```

## Usages

#### Most popular statement
```go
import sb "github.com/tuanphung/sqlb"

sql, args, err := sb.Select("id", "name").From("user").Where(sb.Eq{"id", 1}).ToExpr()
// SELECT id, name FROM user WHERE id = 1
```

#### A complex statement
```go
import sb "github.com/tuanphung/sqlb"

sql, args, err := sb.Select("id", "name").From("user").Where(Or{Eq{"foo", "bar"}, Eq{"id", 1}}).Offset(0).Limit(10).ToExpr()
// SELECT id, name FROM user WHERE (foo = 'bar' OR id = 1) LIMIT 10
```

#### Not enough statement? Use Raw
```go
import sb "github.com/tuanphung/sqlb"

sql, args, err := sb.Raw("EXPLAIN").Select("id", "name").From("user").Where(Or{Eq{"foo", "bar"}, Eq{"id", 1}}).Offset(0).Limit(10).ToExpr()
// EXPLAIN SELECT id, name FROM user WHERE (foo = 'bar' OR id = 1) LIMIT 10
```

## FAQs

## License
sqlb is released under the
[MIT License](http://www.opensource.org/licenses/MIT).
