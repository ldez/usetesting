# UseTesting

Detects when some calls can be replaced by methods from the testing package.

## Examples

### `os.Chdir`

```go
func TestExample(t *testing.T) {
	os.Chdir("")
	// ...
}
```

It can be replaced by:

```go
func TestExample(t *testing.T) {
	t.Chdir("")
    // ...
}
```

### `context.Background`

```go
func TestExample(t *testing.T) {
    ctx := context.Background()
	// ...
}
```

It can be replaced by:

```go
func TestExample(t *testing.T) {
    ctx := t.Context()
    // ...
}
```

### `context.TODO`

```go
func TestExample(t *testing.T) {
    ctx := context.TODO()
	// ...
}
```

It can be replaced by:

```go
func TestExample(t *testing.T) {
    ctx := t.Context()
    // ...
}
```

## TODO

- [ ] add option for each element
- [ ] only for go1.24+, inside linter or inside golangci-lint?

## References

- https://tip.golang.org/doc/go1.24#testingpkgtesting
