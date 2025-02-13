package delete_user_restaurant

type Request struct {
	ID     int64 `json:"id"`
	UserID int64 `json:"user_id"`
}
