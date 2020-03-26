# MySQL Partitioning on Go for Multitenancy


## テーブル名にテナント識別子を入れて分割[WIP]

```sh
make run-t
make login
```

### まとめ

#### GOOD

- クエリが早い
- テーブルで他のテナントのレコードが分離されている

#### BAD

- マイグレーション時に、すべてのテナントのテーブルに変更を加える必要がある
- データベース内にテーブルがあふれかえる

## テーブル内のカラムにテナントの識別子を入れてマルチテナント[WIP]

### まとめ

#### GOOD

- 管理が容易

#### BAD

- クエリが遅い
- 同一のテーブル内に全テナントのレコードが同居する

## テナントごとのデータベース分離[WIP]

#### GOOD

- 管理が容易
- クエリが早い

#### BAD

- マイグレーション時に、すべてのテナントのデータベースに変更を加える必要がある

## パーティショニング

```sh
make run-p
make login
```

Refs: <https://dev.mysql.com/doc/refman/5.6/ja/partitioning-types.html>

- [`LIST COLUMNS パーティショニング`](https://dev.mysql.com/doc/refman/5.6/ja/partitioning-columns-list.html)が良さそう


### テーブル作成

- パーティショニングを有効にするカラムは `PRIMARY KEY` である必要がある
- パーティショニングを有効にする場合、一つ以上のパーティションを設定する必要がある

```sql
CREATE TABLE users (
  id int(11) NOT NULL AUTO_INCREMENT,
  company_id int(11) NOT NULL,
  name varchar(255) COLLATE utf8mb4_bin NOT NULL,
  age int(11) NOT NULL,
  PRIMARY KEY (id, company_id)
)
ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin
PARTITION BY LIST COLUMNS(company_id) (
  PARTITION zero_company VALUES IN(0)
);
```

- ここでは絶対に入らないはずの0を設定した

### パーティションの追加

- `ALTER` で追加する必要がある

```sql
ALTER TABLE users ADD PARTITION (
  PARTITION demo_bank VALUES IN(1)
);
```

- ここで `ADD` しないと 既存のパーティションが失われるから注意

```sql
/*
これをやると既存のパーティションである zero_company が消滅する
*/
ALTER TABLE users PARTITION BY LIST COLUMNS(company_id) (
  PARTITION demo_bank VALUES IN(1)
);
```

### レコードの取得

- これは良い気がする

```sql
SELECT *
FROM users PARTITION (demo_bank);
```

- `WHERE`句は後ろ

```
SELECT *
FROM users PARTITION (demo_bank)
WHERE age = 20;
```

### テーブル定義の確認

- ひとクセ必要

#### `DESCRIBE`

- `DESC users;` ではわからない

```sql
mysql> DESC users;
+------------+--------------+------+-----+---------+----------------+
| Field      | Type         | Null | Key | Default | Extra          |
+------------+--------------+------+-----+---------+----------------+
| id         | int(11)      | NO   | PRI | NULL    | auto_increment |
| company_id | int(11)      | NO   | PRI | NULL    |                |
| name       | varchar(255) | NO   |     | NULL    |                |
| age        | int(11)      | NO   |     | NULL    |                |
+------------+--------------+------+-----+---------+----------------+
```

#### `SHOW CREATE`

- どのパーティションにどのidが入っているか分かる

```sql
SHOW
CREATE TABLE users\G;
*************************** 1. row ***************************
       Table: users
Create Table: CREATE TABLE `users` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `company_id` int(11) NOT NULL,
  `name` varchar(255) COLLATE utf8mb4_bin NOT NULL,
  `age` int(11) NOT NULL,
  PRIMARY KEY (`id`,`company_id`)
) ENGINE=InnoDB AUTO_INCREMENT=501 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin
/*!50500 PARTITION BY LIST  COLUMNS(company_id)
(PARTITION zero_company VALUES IN (0) ENGINE = InnoDB,
 PARTITION demo_bank VALUES IN (1) ENGINE = InnoDB,
 PARTITION demo_shop VALUES IN (2) ENGINE = InnoDB,
 PARTITION demo_restaurant VALUES IN (3) ENGINE = InnoDB,
 PARTITION demo_hotel VALUES IN (4) ENGINE = InnoDB,
 PARTITION demo_system VALUES IN (5) ENGINE = InnoDB) */
```

#### `EXPLAIN PARTITIONS`

- 一つのカラムにいっぱい入っていて見にくい

```sql
EXPLAIN PARTITIONS
SELECT *
FROM users\G;
*************************** 1. row ***************************
           id: 1
  select_type: SIMPLE
        table: users
   partitions: default_company,demo_bank,demo_shop,demo_restaurant,demo_hotel,demo_system
         type: ALL
possible_keys: NULL
          key: NULL
      key_len: NULL
          ref: NULL
         rows: 500
     filtered: 100.00
        Extra: NULL
```

#### `INFORMATION_SCHEMA.PARTITIONS`

- 欲しい情報が全部取れる(名前、ID、パーティション内レコード数等)

```sql
SELECT TABLE_NAME,
       PARTITION_NAME,
       PARTITION_DESCRIPTION,
       TABLE_ROWS
FROM INFORMATION_SCHEMA.PARTITIONS
WHERE TABLE_NAME = 'users'
  AND TABLE_SCHEMA = 'test_db';
+------------+-----------------+-----------------------+------------+
| TABLE_NAME | PARTITION_NAME  | PARTITION_DESCRIPTION | TABLE_ROWS |
+------------+-----------------+-----------------------+------------+
| users      | default_company | 0                     |          0 |
| users      | demo_bank       | 1                     |        100 |
| users      | demo_shop       | 2                     |        100 |
| users      | demo_restaurant | 3                     |        100 |
| users      | demo_hotel      | 4                     |        100 |
| users      | demo_system     | 5                     |        100 |
+------------+-----------------+-----------------------+------------+
```

### まとめ

#### GOOD

- クエリが早い
- 簡単にパーティション内のレコードを取得できる

#### BAD

- パーティションの追加時に`ALTER`しないといけない
- パーティションの追加時にすべてのテーブルに`ALTER`しないといけない
- パーティションの確認が面倒


どうなんだろう...

### TODO

- Go で問題がないか確認する
