package getmoe_test

import (
	"fmt"

	"github.com/leonidboykov/getmoe"
)

func ExampleTags() {
	t := getmoe.NewTags("first_tag", "second_tag")
	t.And("and_this_tag")
	t.No("except_this_tag")
	t.WithoutRating(getmoe.RatingQuestionable, getmoe.RatingExplicit)

	fmt.Println(t)
	// Output: first_tag second_tag and_this_tag -except_this_tag -rating:q -rating:e
}
