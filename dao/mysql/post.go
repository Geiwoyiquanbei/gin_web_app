package mysql

import (
	"strconv"
	"web_app/models"
)

func CreatePost(p *models.Post) {
	sqlStr := `insert into post (post_id,title ,content,author_id,community_id) values (?,?,?,?,?)`
	db.Exec(sqlStr, p.ID, p.Title, p.Content, p.AuthorID, p.CommunityID)
}
func GetPostDetail(id int64) (data *models.Post, err error) {
	sqlStr := `select post_id ,title ,content ,author_id ,community_id,create_time from post where post_id= ? `
	data = new(models.Post)
	err = db.Get(data, sqlStr, id)
	if err != nil {
		return nil, err
	}
	return data, nil
}
func GetPostList(limit, offset int64) (data []*models.Post, err error) {
	//limit 条数  offset 页数
	sqlStr := `select post_id ,title ,content ,author_id ,community_id,create_time from post  ORDER BY create_time DESC limit ?,?`
	data = make([]*models.Post, 0, 2)
	db.Select(&data, sqlStr, (offset-1)*limit, limit)
	return data, nil
}
func GetPostByOrder(ids []string) (postList []*models.Post, err error) {
	for index, _ := range ids {
		sqlStr := `select post_id ,title ,content ,author_id ,community_id ,create_time
	from post where post_id = ?`
		tmp := models.Post{}
		id, _ := strconv.ParseInt(ids[index], 10, 64)
		err := db.Get(&tmp, sqlStr, id)
		if err != nil {
			return nil, err
		}
		data := models.Post{}
		data = tmp
		postList = append(postList, &data)
	}
	//sqlStr := `select post_id ,title ,content ,author_id ,community_id ,create_time
	//from post where id in (?)
	//order by FIND_IN_SET(post_id,?)
	//`
	return postList, nil

}
