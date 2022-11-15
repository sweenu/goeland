package filters

import (
	"testing"
	"time"

	"github.com/slurdge/goeland/internal/goeland"
	"github.com/stretchr/testify/assert"
)

func createTestSource() goeland.Source {

	return goeland.Source{
		Name:  "Test name",
		Title: "Test title",
		URL:   "http://test.com",
		Entries: []goeland.Entry{
			{
				UID:         "1",
				Title:       "First entry",
				Content:     "<h1>test 1</h1>",
				URL:         "http://test.com/blog1",
				Date:        time.Now(),
				ImageURL:    "http://test.com/blog1_img",
			},
			{
				UID:     "2",
				Title:   "Second entry",
				Content: "<h1>test 2</h1>",
				URL:     "http://test.com/blog2",
				Date:    time.Now().AddDate(0, 0, -2), // two days ago
			},
			{
				UID:     "3",
				Title:   "Third entry",
				Content: "<h1>test 3</h1>",
				URL:     "http://test.com/blog3",
				Date:    time.Date(2000, time.January, 1, 0, 0, 0, 0, time.UTC),
			},
		},
	}
}

func TestFilters(t *testing.T) {
	t.Run("all", func(t *testing.T) {
		source := createTestSource()
		entries := source.Entries
		filterAll(&source, nil)
		assert.Equal(t, entries, source.Entries)
	})
	t.Run("none", func(t *testing.T) {
		source := createTestSource()
		filterNone(&source, nil)
		assert.Nil(t, source.Entries)
	})
	t.Run("first", func(t *testing.T) {
		source := createTestSource()
		entries := source.Entries
		filterFirst(&source, &filterParams{})
		assert.Equal(t, entries[:1], source.Entries)
	})
	t.Run("first(2)", func(t *testing.T) {
		source := createTestSource()
		entries := source.Entries
		filterFirst(&source, &filterParams{args: []string{"2"}})
		assert.Equal(t, entries[:2], source.Entries)
	})
	t.Run("last", func(t *testing.T) {
		source := createTestSource()
		entries := source.Entries
		filterLast(&source, &filterParams{})
		assert.Equal(t, entries[len(entries)-1:], source.Entries)
	})
	t.Run("last(2)", func(t *testing.T) {
		source := createTestSource()
		entries := source.Entries
		filterLast(&source, &filterParams{args: []string{"2"}})
		assert.Equal(t, entries[len(entries)-2:], source.Entries)
	})
	t.Run("random(2)", func(t *testing.T) {
		source := createTestSource()
		filterRandom(&source, &filterParams{args: []string{"2"}})
		assert.Len(t, source.Entries, 2)
	})
	t.Run("reverse", func(t *testing.T) {
		source := createTestSource()
		entries := source.Entries
		// Reverse entries
		for i, j := 0, len(entries)-1; i < j; i, j = i+1, j-1 {
			entries[i], entries[j] = entries[j], entries[i]
		}

		filterReverse(&source, nil)
		assert.Equal(t, entries, source.Entries)
	})
	t.Run("today", func(t *testing.T) {
		source := createTestSource()
		entries := source.Entries
		filterToday(&source, nil)
		assert.Equal(t, entries[:1], source.Entries)
	})
	t.Run("lasthours", func(t *testing.T) {
		source := createTestSource()
		entries := source.Entries
		filterLastHour(&source, &filterParams{})
		assert.Equal(t, entries[:1], source.Entries)
	})
	t.Run("lasthours(49)", func(t *testing.T) {
		source := createTestSource()
		entries := source.Entries
		filterLastHour(&source, &filterParams{args: []string{"49"}})
		assert.Equal(t, entries[:2], source.Entries)
	})
	// t.Run("digest", func(t *testing.T) {
	// 	source := createTestSource()
	// 	entries := source.Entries
	// 	filterDigest(&source, &filterParams{})
	// 	assert.Equal(t, entries[:2], source.Entries)
	// })
	// t.Run("digest(4)", func(t *testing.T) {
	// 	source := createTestSource()
	// 	entries := source.Entries
	// 	filterDigest(&source, &filterParams{args: []string{"4"}})
	// 	assert.Equal(t, entries[:2], source.Entries)
	// })
	// t.Run("combine", func(t *testing.T) {
	// 	source := createTestSource()
	// 	entries := source.Entries
	// 	filterCombine(&source, nil)
	// 	assert.Equal(t, entries[:2], source.Entries)
	// })
	// t.Run("links", func(t *testing.T) {
	// 	source := createTestSource()
	// 	entries := source.Entries
	// 	filterRelativeLinks(&source, nil)
	// 	assert.Equal(t, entries[:2], source.Entries)
	// })
	// t.Run("embedimage", func(t *testing.T) {
	// 	source := createTestSource()
	// 	entries := source.Entries
	// 	filterEmbedImage(&source, &filterParams{args: []string{"49"}})
	// 	assert.Equal(t, entries[:2], source.Entries)
	// })
	// t.Run("replace", func(t *testing.T) {
	// 	source := createTestSource()
	// 	entries := source.Entries
	// 	filterReplace(&source, &filterParams{args: []string{"testreplace"}})
	// 	assert.Equal(t, entries[:2], source.Entries)
	// })
	t.Run("includelink", func(t *testing.T) {
		source := createTestSource()
		filterIncludeLink(&source, nil)
		for entry := range source.Entries {
			assert.True(t, entry.IncludeLink)
		}
	})
	// t.Run("sanitize", func(t *testing.T) {
	// 	source := createTestSource()
	// 	entries := source.Entries
	// 	filterSanitize(&source, nil)
	// 	assert.Equal(t, entries[:2], source.Entries)
	// })
	t.Run("toc", func(t *testing.T) {
		source := createTestSource()
		entries := source.Entries
		filterToc(&source, &filterParams{})
		assert.Equal(t, entries[:2], source.Entries)
	})
	t.Run("toc(title)", func(t *testing.T) {
		source := createTestSource()
		entries := source.Entries
		filterToc(&source, &filterParams{args: []string{"title"}})
		assert.Equal(t, entries[:2], source.Entries)
	})
	t.Run("toc without entry", func(t *testing.T) {
		source := createTestSource()
		entries := source.Entries
		filterToc(&source, &filterParams{args: []string{"title"}})
		assert.Equal(t, entries[:2], source.Entries)
	})
	t.Run("limitwords", func(t *testing.T) {
		source := createTestSource()
		entries := source.Entries
		filterLimitWords(&source, &filterParams{})
		assert.Equal(t, entries[:2], source.Entries)
	})
	t.Run("limitwords(5)", func(t *testing.T) {
		source := createTestSource()
		entries := source.Entries
		filterLimitWords(&source, &filterParams{args: []string{"5"}})
		assert.Equal(t, entries[:2], source.Entries)
	})
}
