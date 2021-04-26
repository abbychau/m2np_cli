package shell

type directory struct {
	name   string
	parent *directory
	child  []*directory
	action action
}

func (d *directory) addChild(child *directory) {
	d.child = append(d.child, child)
	child.parent = d
}

func newDirectory(name string, actions ...map[string]action) *directory {
	return &directory{
		name: name,
		action: func(s *shell, inputs []byte) (string, error) {
			return distribute(s, inputs, actions...)
		},
	}
}

var root = newDirectory("/", rootActions, baseActions)

// var posts = newDirectory("posts", postsActions, baseActions)
var inbox = newDirectory("inbox", postsActions, baseActions)
var outbox = newDirectory("outbox", postsActions, baseActions)
var followers = newDirectory("followers", followersActions, baseActions)
var followings = newDirectory("followings", followersActions, baseActions)

func init() {
	for _, v := range []*directory{inbox, outbox, followers, followings} {
		root.addChild(v)
	}

}
