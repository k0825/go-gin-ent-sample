# ローカル実行
```
go run .
```

# ダミーデータ作成

## マイグレーション
```
atlas migrate apply \
  --dir "file://ent/migrate/migrations" \
  --url "postgres://postgres:password@db:5432/todolist?sslmode=disable"
```

## postgresでダミーデータ作成
```
psql -U postgres -d todolist
```

```
-- Insert sample data into "todos" table
INSERT INTO "todos" ("id", "title", "description", "image", "starts_at", "ends_at", "created_at", "updated_at")
VALUES ('e6630c4c-aa2d-4a6a-b1f9-70f7298bf5c3', 'Sample Todo 1', 'This is a sample todo.', 'https://example.com/image1.jpg', '2022-01-01 00:00:00', '2022-01-02 00:00:00', '2022-01-01 00:00:00', '2022-01-01 00:00:00'),
       ('2f8bfee1-f186-4be6-a7ae-e54c2349e101', 'Sample Todo 2', 'This is another sample todo.', 'https://example.com/image2.jpg', '2022-01-03 00:00:00', '2022-01-04 00:00:00', '2022-01-01 00:00:00', '2022-01-01 00:00:00');

-- Insert sample data into "tags" table
INSERT INTO "tags" ("id", "keyword", "todo_id")
VALUES ('c608fbf8-d235-4b44-b692-314b1f0dee4c', 'sample', 'e6630c4c-aa2d-4a6a-b1f9-70f7298bf5c3'),
       ('7160eb23-37f7-4680-9659-aceda5a6c998', 'test', 'e6630c4c-aa2d-4a6a-b1f9-70f7298bf5c3'),
       ('55a5f43e-684b-49a5-800b-27e554cedbc7', 'example', '2f8bfee1-f186-4be6-a7ae-e54c2349e101');
```