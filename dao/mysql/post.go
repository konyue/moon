package mysql

import "moon/models"

func CreatePost(p *models.Post) (err error) {
	sqlStr := `insert into post(
	post_id,title,content,author_id,community_id)
	value (?,?,?,?,?)
	`
	_, err = db.Exec(sqlStr, p.ID, p.Title, p.Content, p.AuthorId, p.CommunityID)
	return

}
