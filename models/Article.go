package models

import "goModWork/databases"

func CreateArticle(article databases.Article) error {
	return databases.ArticleDAO.CreateArticle(databases.Article{
		Uid : article.Uid,
		LableId: article.LableId,
		Title: article.Title,
		Content: article.Content,
		Pic: article.Pic,
	})
}

func FirstArticle() (databases.Article, error) {
	return databases.ArticleDAO.FirstArtilce()
}

func UpdateArticle(article databases.Article) error {
	return databases.ArticleDAO.UpdateArticle(databases.Article{
		Uid : article.Uid,
		LableId: article.LableId,
		Title: article.Title,
		Content: article.Content,
		Pic: article.Pic,
	})
}

func DeleteArticle() error {
	return databases.ArticleDAO.DeleteArticle()
}

func FindArticleLimit(page,lableId int, keyword string) ([]*databases.Article, error) {
	return databases.ArticleDAO.FindArticleLimit(page, lableId, keyword)
}

func FinArticleCount(keyword string) int {
	return databases.ArticleDAO.FinArticleCount(keyword)
}