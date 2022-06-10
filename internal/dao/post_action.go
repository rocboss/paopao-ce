package dao

const (
	idxActNop indexActionT = iota + 1
	idxActCreatePost
	idxActUpdatePost
	idxActDeletePost
	idxActStickPost
)

type indexActionT uint8

func (a indexActionT) String() string {
	switch a {
	case idxActNop:
		return "no operator"
	case idxActCreatePost:
		return "create post"
	case idxActUpdatePost:
		return "update post"
	case idxActDeletePost:
		return "delete post"
	case idxActStickPost:
		return "stick post"
	default:
		return "unknow action"
	}
}
