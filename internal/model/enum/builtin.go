package enum

//go:generate go-enum --names --values --nocase --marshal --mustparse --sqlnullint

/*
Boolean
ENUM(

	False, // false
	True, // true

)
*/
type Boolean int
