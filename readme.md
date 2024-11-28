# UseTesting

Detects when some calls can be replaced by methods from the testing package.

## Usages

### Inside golangci-lint

Recommended.

```yml
linters-settings:
  usetesting:
    # Enable/disable `context.Background()` detections.
    # Default: true
    contextbackground: false
    
    # Enable/disable `context.TODO()` detections.
    # Default: true
    contexttodo: false
    
    # Enable/disable `os.Chdir()` detections.
    # Default: true
    oschdir: false
```

### As a CLI

```
usetesting: Reports uses of functions with replacement inside the testing package.

Usage: usetesting [-flag] [package]

Flags:
  -contextbackground
        Enable/disable context.Background() detections (default true)
  -contexttodo
        Enable/disable context.TODO() detections (default true)
  -oschdir
        Enable/disable os.Chdir() detections (default true)

...
```

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

## References

- https://tip.golang.org/doc/go1.24#testingpkgtesting
