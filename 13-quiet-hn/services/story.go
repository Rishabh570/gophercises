package services

import (
	"fmt"
	"quietHN/helpers"
	"quietHN/hn"
	"quietHN/models"
	"sort"
	"time"
)

func GetStories(client *hn.Client, numStories int, cache *models.Cache) ([]models.Item, error) {
	cache.CacheMutex.Lock()
	defer cache.CacheMutex.Unlock()
	fmt.Println("🔍 GetStories start")

	cachedStories := cache.GetStoriesFromCache()
	if cachedStories != nil {
		fmt.Println("🔍 Returning top stories from cache")
		return cachedStories, nil
	}

	ids, err := client.TopItems()
	if err != nil {
		return nil, err
	}

	var stories []models.Item
	for len(stories) < numStories {
		// calculate number of stories required
		storiesRequired := (numStories - len(stories)) * 5 / 4

		// > THOUGHT PROCESS - In Node.js, execution doesn't go below that line until the async operation is complete for a particular client.
		//     Whereas in go, it executes the whole GetStories fn, including getNStories, and doesn't wait for the async operation to complete since it's real concurrency;
		//     this means multiple clients could be in be in the getNStories function at the same time
		// > We use len(results) as start index to ensure we start fetching from where we left off in the previous iteration
		stories = getNStories(client, storiesRequired, len(stories), ids)
	}

	cache.SetStoriesInCache(stories, time.Now().Add(10*time.Second))

	fmt.Println("🔍 GetStories success")
	return stories, nil
}

func getNStories(client *hn.Client, n int, startIndex int, ids []int) []models.Item {
	var resultChan = make(chan models.Result)
	var results []models.Result

	// Fire n goroutines to fetch stories concurrently
	for i := startIndex; i < startIndex+n; i++ {
		go func(idx, id int) {
			hnItem, err := client.GetItem(id)
			if err != nil {
				resultChan <- models.Result{Index: idx, Err: err}
			}

			resultChan <- models.Result{Index: idx, Item: helpers.ParseHNItem(hnItem)}
		}(i, ids[i])
	}

	// append to results only if result is a story
	for i := 0; i < n; i++ {
		// This is a blocking operation, even though we fired multiple goroutines above
		// execution won't proceed until a value is received from the itemChanel for ALL iterations
		res := <-resultChan
		if res.Err != nil {
			continue
		}

		if helpers.IsStoryLink(res.Item) {
			results = append(results, res)
		}
	}

	// Sort results based on their original index (we want to maintain the order TopItems gave us)
	sort.Slice(results, func(i, j int) bool {
		return results[i].Index < results[j].Index
	})

	var stories []models.Item
	for _, res := range results {
		stories = append(stories, res.Item)
	}

	return stories
}
