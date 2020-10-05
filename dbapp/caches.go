package main

func fillCaches() {
	usersCache.Fill(ctl.getAllUsers())
	articlesCache.Fill(ctl.getArticles())
}
