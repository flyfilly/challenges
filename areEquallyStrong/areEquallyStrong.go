package main

func main() {
	areEquallyStrong(10, 15, 15, 10)
}

func areEquallyStrong(yourLeft int, yourRight int, friendsLeft int, friendsRight int) bool {
	//return yourLeft+yourRight == friendsLeft+friendsRight
	return (yourLeft == friendsLeft || yourRight == friendsRight || yourLeft == friendsRight || yourRight == friendsLeft) && (yourLeft+yourRight) == (friendsLeft+friendsRight)
}
