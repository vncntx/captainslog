package main

// Please install zap before running these benchmarks

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"testing"
	"time"

	"github.com/vincentfiestada/captainslog/v2"
	"github.com/vincentfiestada/captainslog/v2/format"
)

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func main() {

	results := []string{}

	results = append(results, runBenchmark("stdlib", benchmarkStdlib))
	results = append(results, runBenchmark("captainslog", benchmarkCaptainsLog))
	results = append(results, runBenchmark("captainslog (json)", benchmarkCaptainsLogJSON))
	results = append(results, runBenchmark("captainslog (minimal)", benchmarkCaptainsLogMinimal))

	fmt.Printf("\n%21s\n", "Benchmark Results")
	for _, res := range results {
		fmt.Println(res)
	}
}

func runBenchmark(name string, test func(*testing.B)) string {
	result := testing.Benchmark(test)
	return fmt.Sprintf("%25s: %10d ns/op, %4d allocs/op, %5d bytes/op", name, result.NsPerOp(), result.AllocsPerOp(), result.AllocedBytesPerOp())
}

func benchmarkStdlib(b *testing.B) {
	log := log.New(os.Stdout, "", 0)

	b.RunParallel(func(i *testing.PB) {
		for i.Next() {
			log.Printf("%s %s, a=%#v, b=%#v, c=%#v",
				time.Now().Format(captainslog.ISO8601),
				randomMessage(20),
				rand.Int(),
				rand.Int(),
				rand.Int(),
			)
		}
	})
}

func benchmarkCaptainsLog(b *testing.B) {
	log := captainslog.NewLogger()
	log.Stdout = os.Stdout

	b.RunParallel(func(i *testing.PB) {
		for i.Next() {
			log.Fields(
				log.I("a", rand.Int()),
				log.I("b", rand.Int()),
				log.I("c", rand.Int()),
			).Info(randomMessage(20))
		}
	})
}

func benchmarkCaptainsLogJSON(b *testing.B) {
	log := captainslog.NewLogger()
	log.Stdout = os.Stdout
	log.Format = format.JSON

	b.RunParallel(func(i *testing.PB) {
		for i.Next() {
			log.Fields(
				log.I("a", rand.Int()),
				log.I("b", rand.Int()),
				log.I("c", rand.Int()),
			).Info(randomMessage(20))
		}
	})
}

func benchmarkCaptainsLogMinimal(b *testing.B) {
	log := captainslog.NewLogger()
	log.Stdout = os.Stdout
	log.Format = format.Minimal

	b.RunParallel(func(i *testing.PB) {
		for i.Next() {
			log.Fields(
				log.I("a", rand.Int()),
				log.I("b", rand.Int()),
				log.I("c", rand.Int()),
			).Info(randomMessage(20))
		}
	})
}

func randomMessage(length int) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}
	return string(b)
}
