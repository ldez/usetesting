# UseTesting

Detects when some calls can be replaced by methods from the testing package.

## Usages

### Inside golangci-lint

Recommended.

```yml
linters-settings:
  usetesting:
    # Enable/disable `context.Background()` detections.
    # Disabled if Go < 1.24.
    # Default: true
    context-background: false
    
    # Enable/disable `context.TODO()` detections.
    # Disabled if Go < 1.24.
    # Default: true
    context-todo: false
    
    # Enable/disable `os.Chdir()` detections.
    # Disabled if Go < 1.24.
    # Default: true
    os-chdir: false
    
    # Enable/disable `os.MkdirTemp()` detections.
    # Default: true
    os-mkdir-temp: false
    
    # Enable/disable `os.Setenv()` detections.
    # Default: false
    os-setenv: true
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
  -osmkdirtemp
        Enable/disable os.MkdirTemp() detections (default true)
  -ossetenv
        Enable/disable os.Setenv() detections (default true)
...
```

## Examples

### `os.MkdirTemp`

```go
func TestExample(t *testing.T) {
	os.MkdirTemp("", "")
	// ...
}
```

It can be replaced by:

```go
func TestExample(t *testing.T) {
	t.TempDir()
    // ...
}
```

### `os.Setenv`

```go
func TestExample(t *testing.T) {
	os.Setenv("", "")
	// ...
}
```

It can be replaced by:

```go
func TestExample(t *testing.T) {
	t.Setenv("", "")
    // ...
}
```

### `os.Chdir` (Go >= 1.24)

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

### `context.Background` (Go >= 1.24)

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

### `context.TODO` (Go >= 1.24)

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
