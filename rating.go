package getmoe

// Rating defines boorus rating system
type Rating string

// Boorus rating system has safe, questionable and explicit tags
const (
	RatingSafe         Rating = "s"
	RatingQuestionable Rating = "q"
	RatingExplicit     Rating = "e"
)
