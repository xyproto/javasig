package javasig

import "testing"

func TestIsJavaFunc(t *testing.T) {
	cases := []struct {
		line     string
		expected bool
	}{
		{"public void run() {", true},
		{"private static <T> List<T> filter(List<T> x);", true},
		{"new Something();", false},
		{"return x;", false},
		{"super(foo);", false},
		{"@Override public int compareTo(Object o) {", true},
		{"main(args);", false},
		{"abstract <T> T make();", true},
		{"void renew();", true},
		{"String toString();", true},
		{"throw new RuntimeException();", false},
		{"@interface Foo {}", false},
		{"public static void main(String[] args) {", true},
		{"  static <K,V> Map<K,V> createMap();", true},
		{"Random rnd = new Random();", false},
	}
	for _, c := range cases {
		actual := Is(c.line)
		if actual != c.expected {
			t.Errorf("Is(%q) = %v; want %v", c.line, actual, c.expected)
		}
	}
}
