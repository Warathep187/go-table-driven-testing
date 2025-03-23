---
marp: true
theme: default
class: lead
---

# Table-Driven Tests in Go

---

## Why Table-Driven Tests?

- **Problem**: Writing multiple test cases for a function can lead to repetitive code.
- **Solution**: Table-driven tests allow us to define a list of test cases and iterate over them, reducing redundancy and improving maintainability.

---

## Example: Sum Function

```go
func Sum(a, b int) int {
	return a + b
}
```

---

## Traditional Test Case

```go
func TestSum(t *testing.T) {
	// Test case 1: Basic addition
	got := Sum(1, 2)
	want := 3
	if got != want {
		t.Errorf("Sum(1, 2) = %d; want %d", got, want)
	}
}
```

---

## Table-Driven Test

```go
func TestSum(t *testing.T) {
	tests := []struct {
		a, b, want int
	}{
		{1, 2, 3},
		{-1, 1, 0},
		{0, 0, 0},
	}

	for _, test := range tests {
		got := Sum(test.a, test.b)
		if got != test.want {
			t.Errorf("Sum(%d, %d) = %d; want %d", test.a, test.b, got, test.want)
		}
	}
}
```

---

## Benefits of Table-Driven Tests

- **Readability**: Test cases are more readable and easier to understand.
- **Maintainability**: Adding new test cases is straightforward and doesn't require modifying existing code.
