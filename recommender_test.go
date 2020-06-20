package recommender_test

import (
	"fmt"
	"log"
	"testing"

	"github.com/manson/recommender"
)

func TestLike(t *testing.T) {
	// log.Printf("TestLike")

	r, err := recommender.NewRecommender(true)
	if err != nil {
		log.Fatal(err)
	}
	defer r.Close()

	niko, _ := recommender.NewUser("1", "Niko Kovacevic")
	aubreigh, _ := recommender.NewUser("2", "Aubreigh Brunschwig")

	denver, _ := recommender.NewItem("Denver", "Denver")
	phoenix, _ := recommender.NewItem("Phoenix", "Phoenix")
	pittsburgh, _ := recommender.NewItem("Pittsburgh", "Pittsburgh")
	portland, _ := recommender.NewItem("Portland", "Portland")
	losAngeles, _ := recommender.NewItem("Los_Angeles", "Los Angeles")
	miami, _ := recommender.NewItem("Miami", "Miami")

	// GetLikedItems should return no items at this point
	items, err := r.GetLikedItems(niko)
	if err != nil {
		t.Errorf("Error: %s", err)
	}
	if l := len(items); l != 0 {
		t.Errorf("There should be 0 items. There are %d.", l)
	}

	// GetUsersWhoLike should return no users at this point
	users, err := r.GetUsersWhoLike(portland)
	if err != nil {
		t.Errorf("Error: %s", err)
	}
	if l := len(users); l != 0 {
		t.Errorf("There should be 0 users. There are %d.", l)
	}

	// Add some likes
	r.Like(niko, phoenix)
	r.Like(niko, denver)
	r.Like(niko, pittsburgh)
	r.Like(aubreigh, phoenix)
	r.Like(aubreigh, portland)

	// Get the liked items
	items, err = r.GetLikedItems(niko)
	if err != nil {
		t.Errorf("Error: %s", err)
	}
	if l := len(items); l != 3 {
		t.Errorf("There should be 3 items. There are %d.", l)
	}

	// Add some dislikes
	r.Dislike(niko, phoenix)
	r.Dislike(niko, miami)
	r.Dislike(niko, losAngeles)

	// Add some more likes, with some overlapping and previously disliked
	r.Like(niko, phoenix)
	r.Like(niko, portland)
	r.Like(niko, pittsburgh)
	r.Like(aubreigh, phoenix)

	// There should only be one new item, 4 total
	items, err = r.GetLikedItems(niko)
	if err != nil {
		t.Errorf("Error: %s", err)
	}
	if l := len(items); l != 4 {
		t.Errorf("There should be 4 items. There are %d: %v", l, items)
	}
}

func TestDisLike(t *testing.T) {
	// log.Printf("TestDislike")

	r, err := recommender.NewRecommender(true)
	if err != nil {
		log.Fatal(err)
	}
	defer r.Close()

	niko, _ := recommender.NewUser("1", "Niko Kovacevic")

	phoenix, _ := recommender.NewItem("Phoenix", "Phoenix")
	losAngeles, _ := recommender.NewItem("Los_Angeles", "Los Angeles")
	miami, _ := recommender.NewItem("Miami", "Miami")

	// GetLikedItems should return no items at this point
	items, err := r.GetDislikedItems(niko)
	if err != nil {
		t.Errorf("Error: %s", err)
	}
	if l := len(items); l != 0 {
		t.Errorf("There should be 0 items. There are %d.", l)
	}

	// GetUsersWhoDislike should return no users at this point
	users, err := r.GetUsersWhoDislike(miami)
	if err != nil {
		t.Errorf("Error: %s", err)
	}
	if l := len(users); l != 0 {
		t.Errorf("There should be 0 users. There are %d.", l)
	}

	// Add some dislikes
	r.Dislike(niko, phoenix)
	r.Dislike(niko, miami)

	// Get the disliked items
	items, err = r.GetDislikedItems(niko)
	if err != nil {
		t.Errorf("Error: %s", err)
	}
	if l := len(items); l != 2 {
		t.Errorf("There should be 2 items. There are %d.", l)
	}

	// Like some items
	r.Like(niko, phoenix)

	// Add some more dislikes, with some overlapping and previously liked
	r.Dislike(niko, phoenix)
	r.Dislike(niko, losAngeles)

	// There should only be one new item, 3 total
	items, err = r.GetDislikedItems(niko)
	if err != nil {
		t.Errorf("Error: %s", err)
	}
	if l := len(items); l != 3 {
		t.Errorf("There should be 3 items. There are %d.", l)
	}
}

func TestGetRatings(t *testing.T) {
	// log.Printf("TestGetRatings")

	r, err := recommender.NewRecommender(true)
	if err != nil {
		log.Fatal(err)
	}
	defer r.Close()

	niko, _ := recommender.NewUser("1", "Niko Kovacevic")

	phoenix, _ := recommender.NewItem("Phoenix", "Phoenix")
	losAngeles, _ := recommender.NewItem("Los_Angeles", "Los Angeles")
	miami, _ := recommender.NewItem("Miami", "Miami")
	pittsburgh, _ := recommender.NewItem("Pittsburgh", "Pittsburgh")
	boulder, _ := recommender.NewItem("Boulder", "Boulder")
	seattle, _ := recommender.NewItem("Seattle", "Seattle")

	// GetLikedItems should return no items at this point
	ratings, err := r.GetRatings(niko)
	if err != nil {
		t.Errorf("Error: %s", err)
	}
	if len(ratings) != 0 {
		t.Errorf("There should be 0 items. There are %d.", len(ratings))
	}

	// Add some likes and dislikes
	r.Dislike(niko, phoenix)
	r.Dislike(niko, miami)
	r.Dislike(niko, losAngeles)
	r.Like(niko, pittsburgh)
	r.Like(niko, boulder)
	r.Like(niko, seattle)

	// There should be six ratings
	ratings, err = r.GetRatings(niko)
	if err != nil {
		t.Errorf("Error: %s", err)
	}
	if len(ratings) != 6 {
		t.Errorf("There should be six items. There are %d.", len(ratings))
	}

	/*
		// Print ratings (notice order varies because of concurrency)
		for _, rating := range ratings {
			fmt.Printf("%v\n", rating)
		}
	*/
}

func TestGetUsersWhoRated(t *testing.T) {
	// log.Printf("TestGetUsersWhoRated")

	r, err := recommender.NewRecommender(true)
	if err != nil {
		log.Fatal(err)
	}
	defer r.Close()

	niko, _ := recommender.NewUser("1", "Niko Kovacevic")
	aubreigh, _ := recommender.NewUser("2", "Aubreigh Brunschwig")
	johnny, _ := recommender.NewUser("3", "Johnny Bernard")
	amanda, _ := recommender.NewUser("4", "Amanda Hunt")
	nick, _ := recommender.NewUser("5", "Nick Evers")

	phoenix, _ := recommender.NewItem("Phoenix", "Phoenix")
	pittsburgh, _ := recommender.NewItem("Pittsburgh", "Pittsburgh")

	// GetUsersWhoRated should return no users at this point
	users, err := r.GetUsersWhoRated(pittsburgh)
	if err != nil {
		t.Errorf("Error: %s", err)
	}
	if len(users) != 0 {
		t.Errorf("There should be 0 users. There are %d.", len(users))
	}

	// Add some likes and dislikes
	r.Dislike(niko, phoenix)
	r.Dislike(aubreigh, phoenix)
	r.Like(johnny, phoenix)
	r.Like(amanda, phoenix)
	r.Like(niko, pittsburgh)
	r.Like(nick, pittsburgh)

	// GetUsersWhoRated should return 4 users at this point
	users, err = r.GetUsersWhoRated(phoenix)
	if err != nil {
		t.Errorf("Error: %s", err)
	}
	if len(users) != 4 {
		t.Errorf("There should be 4 users. There are %d:", len(users))
		for _, user := range users {
			fmt.Printf("%v\n", user)
		}
	}

	// GetUsersWhoRated should return 2 users at this point
	users, err = r.GetUsersWhoRated(pittsburgh)
	if err != nil {
		t.Errorf("Error: %s", err)
	}
	if len(users) != 2 {
		t.Errorf("There should be 2 users. There are %d:", len(users))
		for _, user := range users {
			fmt.Printf("%v\n", user)
		}
	}
}

func TestGetRatingNeighbors(t *testing.T) {
	// log.Printf("TestGetRatingNeighbors")

	r, err := recommender.NewRecommender(true)
	if err != nil {
		log.Fatal(err)
	}
	defer r.Close()

	niko, _ := recommender.NewUser("1", "Niko Kovacevic")
	aubreigh, _ := recommender.NewUser("2", "Aubreigh Brunschwig")
	johnny, _ := recommender.NewUser("3", "Johnny Bernard")
	amanda, _ := recommender.NewUser("4", "Amanda Hunt")
	nick, _ := recommender.NewUser("5", "Nick Evers")

	phoenix, _ := recommender.NewItem("Phoenix", "Phoenix")
	pittsburgh, _ := recommender.NewItem("Pittsburgh", "Pittsburgh")

	// GetRatingNeighbors should return no users at this point
	neighbors, err := r.GetRatingNeighbors(niko)
	if err != nil {
		t.Errorf("Error: %s", err)
	}
	if len(neighbors) != 0 {
		t.Errorf("There should be 0 users. There are %d:", len(neighbors))
		for _, neighbor := range neighbors {
			fmt.Printf("%v\n", neighbor)
		}
	}

	// Add some likes and dislikes
	r.Dislike(niko, phoenix)
	r.Dislike(aubreigh, phoenix)
	r.Like(johnny, phoenix)
	r.Like(amanda, phoenix)
	r.Like(niko, pittsburgh)
	r.Dislike(aubreigh, pittsburgh)
	r.Like(nick, pittsburgh)

	// GetRatingNeighbors should return five users at this point
	neighbors, err = r.GetRatingNeighbors(niko)
	if err != nil {
		t.Errorf("Error: %s", err)
	}
	if len(neighbors) != 4 {
		t.Errorf("There should be 4 users. There are %d:", len(neighbors))
		for _, neighbor := range neighbors {
			fmt.Printf("%v\n", neighbor)
		}
	}
}

func TestSimilarity(t *testing.T) {
	//log.Printf("TestSimilarity")

	r, err := recommender.NewRecommender(true)
	if err != nil {
		log.Fatal(err)
	}
	defer r.Close()

	niko, _ := recommender.NewUser("1", "Niko Kovacevic")
	aubreigh, _ := recommender.NewUser("2", "Aubreigh Brunschwig")
	johnny, _ := recommender.NewUser("3", "Johnny Bernard")
	nick, _ := recommender.NewUser("5", "Nick Evers")

	phoenix, _ := recommender.NewItem("Phoenix", "Phoenix")
	pittsburgh, _ := recommender.NewItem("Pittsburgh", "Pittsburgh")
	boulder, _ := recommender.NewItem("Boulder", "Boulder")
	losAngeles, _ := recommender.NewItem("Los_Angeles", "Los Angeles")
	portland, _ := recommender.NewItem("Portland", "Portland")
	seattle, _ := recommender.NewItem("Seattle", "Seattle")

	// GetSimilarity should return nothing at this point
	nikoSims, err := r.GetSimilarity(niko)
	if err != nil {
		t.Errorf("Error: %s", err)
	}
	if len(nikoSims) != 0 {
		t.Errorf("There should be 0 similarities. There are %d:", len(nikoSims))
		for _, similarity := range nikoSims {
			fmt.Printf("%v\n", similarity)
		}
	}

	// Add some likes and dislikes
	r.Dislike(niko, phoenix)
	r.Like(niko, pittsburgh)
	r.Like(niko, boulder)
	r.Dislike(niko, losAngeles)
	r.Like(niko, portland)
	r.Like(niko, seattle)

	r.Dislike(aubreigh, phoenix)
	r.Dislike(aubreigh, pittsburgh)
	r.Like(aubreigh, boulder)
	r.Like(aubreigh, losAngeles)
	r.Like(aubreigh, portland)
	r.Like(aubreigh, seattle)

	r.Like(johnny, phoenix)
	r.Like(johnny, losAngeles)

	r.Like(nick, pittsburgh)
	r.Like(nick, portland)

	// GetSimilarity should return three similarities at this point
	nikoSims, err = r.GetSimilarity(niko)
	if err != nil {
		t.Errorf("Error: %s", err)
	}
	if len(nikoSims) != 3 {
		t.Errorf("There should be 3 similarities. There are %d:", len(nikoSims))
		for _, similarity := range nikoSims {
			fmt.Printf("%v\n", similarity)
		}
	}

	// Get other users's similarities
	aubreighSims, err := r.GetSimilarity(aubreigh)
	if err != nil {
		t.Errorf("Error: %s", err)
	}
	nickSims, err := r.GetSimilarity(nick)
	if err != nil {
		t.Errorf("Error: %s", err)
	}
	johnnySims, err := r.GetSimilarity(johnny)
	if err != nil {
		t.Errorf("Error: %s", err)
	}

	// Test commutativity
	if nikoSims[aubreigh.ID].Index != aubreighSims[niko.ID].Index {
		t.Errorf("Similarity(Niko, Aubreigh) should equaul Similarity(Aubreigh, Niko).")
	}
	if nikoSims[nick.ID].Index != nickSims[niko.ID].Index {
		t.Errorf("Similarity(Niko, Nick) should equaul Similarity(Nick, Niko).")
	}
	if nikoSims[johnny.ID].Index != johnnySims[niko.ID].Index {
		t.Errorf("Similarity(Niko, Johnny) should equaul Similarity(Johnny, Niko).")
	}

	// Test values
	if float32(nikoSims[aubreigh.ID].Index) != float32(2.0/6.0) {
		t.Errorf("Similarity(Niko, Aubreigh) should be %f. Actually %f", float32(2.0/6.0), nikoSims[aubreigh.ID].Index)
	}
	if float32(nikoSims[nick.ID].Index) != float32(1) {
		t.Errorf("Similarity(Niko, Nick) should be %f. Actually %f", float32(1), nikoSims[nick.ID].Index)
	}
	if float32(nikoSims[johnny.ID].Index) != float32(-1) {
		t.Errorf("Similarity(Niko, Johnny) should be %f. Actually %f", float32(-1), nikoSims[johnny.ID].Index)
	}
}

func TestSuggestions(t *testing.T) {
	log.Printf("TestSuggestions")

	r, err := recommender.NewRecommender(true)
	if err != nil {
		log.Fatal(err)
	}
	defer r.Close()

	niko, _ := recommender.NewUser("1", "Niko Kovacevic")
	aubreigh, _ := recommender.NewUser("2", "Aubreigh Brunschwig")
	johnny, _ := recommender.NewUser("3", "Johnny Bernard")
	amanda, _ := recommender.NewUser("4", "Amanda Hunt")
	nick, _ := recommender.NewUser("5", "Nick Evers")
	katie, _ := recommender.NewUser("Katie Yoder", "Katie Yoder")
	matt, _ := recommender.NewUser("Matt Rolland", "Matt Rolland")
	bekah, _ := recommender.NewUser("Bekah Sandoval", "Bekah Sandoval")
	bill, _ := recommender.NewUser("Bill Taggart", "Bill Taggart")
	megan, _ := recommender.NewUser("Megan Murzyn", "Megan Murzyn")

	ashland, _ := recommender.NewItem("Ashland, Oregon", "Ashland, Oregon")
	austin, _ := recommender.NewItem("Austin_Texas", "Austin, Texas")
	boulder, _ := recommender.NewItem("Boulder, Colorado", "Boulder, Colorado")
	denver, _ := recommender.NewItem("Denver, Colorado", "Denver, Colorado")
	flagstaff, _ := recommender.NewItem("Flagstaff, Arizona", "Flagstaff, Arizona")
	houston, _ := recommender.NewItem("Houston, Texas", "Houston, Texas")
	lasVegas, _ := recommender.NewItem("Las Vegas, Nevada", "Las Vegas, Nevada")
	losAngeles, _ := recommender.NewItem("Los Angeles, California", "Los Angeles, California")
	newYork, _ := recommender.NewItem("New York, New York", "New York, New York")
	philadelphia, _ := recommender.NewItem("Philadelphia, Pennsylvania", "Philadelphia, Pennsylvania")
	phoenix, _ := recommender.NewItem("Phoenix, Arizona", "Phoenix, Arizona")
	pittsburgh, _ := recommender.NewItem("Pittsburgh, Pennsylvania", "Pittsburgh, Pennsylvania")
	portlandOR, _ := recommender.NewItem("Portland, Oregon", "Portland, Oregon")
	portlandME, _ := recommender.NewItem("Portland, Maine", "Portland, Maine")
	princeton, _ := recommender.NewItem("Princeton, New Jersey", "Princeton, New Jersey")
	sacramento, _ := recommender.NewItem("Sacramento, California", "Sacramento, California")
	sanFrancisco, _ := recommender.NewItem("San Francisco, California", "San Francisco, California")
	santaFe, _ := recommender.NewItem("Santa Fe, New Mexico", "Santa Fe, New Mexico")
	seattle, _ := recommender.NewItem("Seattle, Washington", "Seattle, Washington")
	tacoma, _ := recommender.NewItem("Tacoma, Washington", "Tacoma, Washington")
	tucson, _ := recommender.NewItem("Tucson, Arizona", "Tucson, Arizona")

	// GetSuggestions should return nothing at this point
	nikoSuggestions, err := r.GetSuggestions(niko)
	if err != nil {
		t.Errorf("Error: %s", err)
	}
	if len(nikoSuggestions) != 0 {
		t.Errorf("There should be 0 suggestions. There are %d:", len(nikoSuggestions))
		for _, suggestion := range nikoSuggestions {
			fmt.Printf("%v\n", suggestion)
		}
	}

	// Add some likes and dislikes
	r.Like(niko, boulder)
	r.Like(niko, pittsburgh)
	r.Like(niko, seattle)
	r.Dislike(niko, lasVegas)
	r.Dislike(niko, losAngeles)
	r.Dislike(niko, phoenix)

	r.Like(aubreigh, ashland)
	r.Like(aubreigh, boulder)
	r.Like(aubreigh, denver)
	r.Like(aubreigh, flagstaff)
	r.Like(aubreigh, losAngeles)
	r.Like(aubreigh, portlandOR)
	r.Like(aubreigh, sanFrancisco)
	r.Like(aubreigh, seattle)
	r.Dislike(aubreigh, lasVegas)
	r.Dislike(aubreigh, phoenix)
	r.Dislike(aubreigh, pittsburgh)
	r.Dislike(aubreigh, tacoma)

	r.Like(johnny, phoenix)
	r.Like(johnny, flagstaff)
	r.Like(johnny, losAngeles)
	r.Like(johnny, sanFrancisco)
	r.Like(johnny, lasVegas)
	r.Like(johnny, portlandME)
	r.Dislike(johnny, sacramento)
	r.Dislike(johnny, santaFe)

	r.Like(amanda, losAngeles)
	r.Like(amanda, flagstaff)
	r.Like(amanda, sanFrancisco)
	r.Like(amanda, portlandME)
	r.Like(amanda, santaFe)
	r.Dislike(amanda, lasVegas)
	r.Dislike(amanda, phoenix)
	r.Dislike(amanda, sacramento)

	r.Like(nick, pittsburgh)
	r.Like(nick, portlandOR)
	r.Like(nick, seattle)
	r.Like(nick, ashland)
	r.Like(nick, austin)
	r.Dislike(nick, houston)
	r.Dislike(nick, philadelphia)

	r.Like(katie, portlandOR)
	r.Like(katie, seattle)
	r.Like(katie, ashland)
	r.Like(katie, austin)
	r.Like(katie, houston)
	r.Dislike(katie, pittsburgh)

	r.Like(matt, flagstaff)
	r.Like(matt, tucson)
	r.Like(matt, denver)
	r.Like(matt, boulder)
	r.Like(matt, portlandOR)
	r.Like(matt, santaFe)
	r.Like(matt, newYork)
	r.Dislike(matt, phoenix)
	r.Dislike(matt, losAngeles)
	r.Dislike(matt, lasVegas)

	r.Like(bekah, flagstaff)
	r.Like(bekah, tucson)
	r.Like(bekah, denver)
	r.Like(bekah, boulder)
	r.Like(bekah, portlandOR)
	r.Like(bekah, tacoma)
	r.Like(bekah, seattle)
	r.Like(bekah, newYork)
	r.Dislike(bekah, phoenix)
	r.Dislike(bekah, losAngeles)
	r.Dislike(bekah, lasVegas)
	r.Dislike(bekah, sacramento)

	r.Like(bill, flagstaff)
	r.Like(bill, denver)
	r.Like(bill, portlandOR)
	r.Like(bill, philadelphia)
	r.Like(bill, princeton)
	r.Like(bill, newYork)
	r.Like(bill, phoenix)
	r.Like(bill, losAngeles)
	r.Like(bill, lasVegas)
	r.Dislike(bill, tucson)
	r.Dislike(bill, santaFe)
	r.Dislike(bill, pittsburgh)
	r.Dislike(bill, seattle)
	r.Dislike(bill, boulder)

	r.Like(megan, flagstaff)
	r.Like(megan, philadelphia)
	r.Like(megan, princeton)
	r.Like(megan, newYork)
	r.Like(megan, phoenix)
	r.Dislike(megan, tucson)

	if err := r.UpdateSuggestions(niko); err != nil {
		t.Errorf("Error: %s", err)
	}

	// GetSuggestions should return nothing at this point
	nikoSuggestions, err = r.GetSuggestions(niko)
	if err != nil {
		t.Errorf("Error: %s", err)
	}
	if len(nikoSuggestions) != 15 {
		t.Errorf("There should be 15 suggestions. There are %d:", len(nikoSuggestions))
		for _, suggestion := range nikoSuggestions {
			fmt.Printf("%v\n", suggestion)
		}
	}
	for _, suggestion := range nikoSuggestions {
		fmt.Printf("%s - %v\n", suggestion.Item.ID, suggestion)
	}

}
