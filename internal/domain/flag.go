package domain


type Flag struct{
	id	string `json:"id"`
	key string `json:"key"`
	name string `json:"name"`
	description string `json:"description"`
	created_at string `json:"created_at"`
	updated_at string `json:"updated_at"`
	archieved_at string `json:"archieved_at"`
}
