package metrics

import "testing"

func TestPriceForModelAliasAnthropicFableAndOpus48(t *testing.T) {
	cases := []struct {
		model string
		want  ModelPrice
	}{
		{
			model: "claude-fable-5",
			want:  ModelPrice{Provider: "anthropic", Model: "claude-fable-5", InputPerM: 10, CacheReadPerM: 1, CacheWritePerM: 12.5, OutputPerM: 50},
		},
		{
			model: "anthropic/claude-fable-5",
			want:  ModelPrice{Provider: "anthropic", Model: "claude-fable-5", InputPerM: 10, CacheReadPerM: 1, CacheWritePerM: 12.5, OutputPerM: 50},
		},
		{
			model: "claude-opus-4-8",
			want:  ModelPrice{Provider: "anthropic", Model: "claude-opus-4.8", InputPerM: 5, CacheReadPerM: 0.5, CacheWritePerM: 6.25, OutputPerM: 25},
		},
	}

	for _, tc := range cases {
		got, ok := PriceForModelAlias(tc.model)
		if !ok {
			t.Fatalf("PriceForModelAlias(%q) did not resolve", tc.model)
		}
		if got != tc.want {
			t.Fatalf("PriceForModelAlias(%q) = %+v, want %+v", tc.model, got, tc.want)
		}
	}
}

func TestPriceForModelAliasCurrentOpenAIAndNoBroadFallback(t *testing.T) {
	cases := []struct {
		model string
		want  ModelPrice
	}{
		{
			model: "gpt-5.4-nano",
			want:  ModelPrice{Provider: "openai", Model: "gpt-5.4-nano", InputPerM: 0.20, CacheReadPerM: 0.02, CacheWritePerM: 0.20, OutputPerM: 1.25},
		},
		{
			model: "gpt-5.2-codex",
			want:  ModelPrice{Provider: "openai", Model: "gpt-5.2-codex", InputPerM: 1.75, CacheReadPerM: 0.175, CacheWritePerM: 1.75, OutputPerM: 14},
		},
		{
			model: "openai/gpt-4.1",
			want:  ModelPrice{Provider: "openai", Model: "gpt-4.1", InputPerM: 2, CacheReadPerM: 0.5, CacheWritePerM: 2, OutputPerM: 8},
		},
	}

	for _, tc := range cases {
		got, ok := PriceForModelAlias(tc.model)
		if !ok {
			t.Fatalf("PriceForModelAlias(%q) did not resolve", tc.model)
		}
		if got != tc.want {
			t.Fatalf("PriceForModelAlias(%q) = %+v, want %+v", tc.model, got, tc.want)
		}
	}

	if _, ok := PriceForModelAlias("gpt-5.5-mini"); ok {
		t.Fatal("gpt-5.5-mini must remain unmapped until OpenAI publishes a public price")
	}
}

func TestPriceForModelAliasDeepSeekGeminiKimiAndZhipu(t *testing.T) {
	cases := []struct {
		model string
		want  ModelPrice
	}{
		{
			model: "deepseek/deepseek-v4-pro",
			want:  ModelPrice{Provider: "deepseek", Model: "v4-pro", InputPerM: 0.435, CacheReadPerM: 0.003625, CacheWritePerM: 0.435, OutputPerM: 0.87},
		},
		{
			model: "deepseek-chat",
			want:  ModelPrice{Provider: "deepseek", Model: "v4-flash", InputPerM: 0.14, CacheReadPerM: 0.0028, CacheWritePerM: 0.14, OutputPerM: 0.28},
		},
		{
			model: "gemini-3.1-pro-preview-customtools",
			want:  ModelPrice{Provider: "google", Model: "gemini-3.1-pro", InputPerM: 2, CacheReadPerM: 0.2, CacheWritePerM: 2, OutputPerM: 12},
		},
		{
			model: "gemini-2.5-flash-lite",
			want:  ModelPrice{Provider: "google", Model: "gemini-2.5-flash-lite", InputPerM: 0.10, CacheReadPerM: 0.01, CacheWritePerM: 0.10, OutputPerM: 0.40},
		},
		{
			model: "kimi-k2.7-code-highspeed",
			want:  ModelPrice{Provider: "kimi", Model: "k2.7-code-highspeed", InputPerM: 1.90, CacheReadPerM: 0.38, CacheWritePerM: 1.90, OutputPerM: 8},
		},
		{
			model: "glm-5.2",
			want:  ModelPrice{Provider: "zhipu", Model: "glm-5.2", InputPerM: 1.4, CacheReadPerM: 0.26, CacheWritePerM: 1.4, OutputPerM: 4.4},
		},
	}

	for _, tc := range cases {
		got, ok := PriceForModelAlias(tc.model)
		if !ok {
			t.Fatalf("PriceForModelAlias(%q) did not resolve", tc.model)
		}
		if got != tc.want {
			t.Fatalf("PriceForModelAlias(%q) = %+v, want %+v", tc.model, got, tc.want)
		}
	}

	for _, unmapped := range []string{
		"gemini-3.1-pro-preview-ioa",
		"gemini-3.5-flash-pro",
		"kimi-k2.6-ioa",
		"glm-5.1-ioa",
	} {
		if _, ok := PriceForModelAlias(unmapped); ok {
			t.Fatalf("%q must remain unmapped because its public API price is not verified", unmapped)
		}
	}
}
