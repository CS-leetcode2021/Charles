# BoltDB的使用规范

---

## Open a DB

```go
db, err := bolt.Open("my.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
```


---

## Transactions

-   读写功能
```go
err := db.Update(func(tx *bolt.Tx) error {
	...
	return nil
})

```

- 只读
```go
err := db.View(func(tx *bolt.Tx) error {
	...
	return nil
})

```

- 批量读写

