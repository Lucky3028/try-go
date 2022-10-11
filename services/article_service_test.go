package services_test

import "testing"

func BenchmarkGetArticle(b *testing.B) {
	articleId := 1

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if _, err := articleService.GetArticle(articleId); err != nil {
			b.Error(err)
			break
		}

	}
}
