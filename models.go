package main

type User struct {
	ID        string `json:"id"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	Name      string `json:"name"`
	ApiKey    string `json:"api_key"`
}

func databaseUserToUser(dbUser UsersDBModel) User {
	return User{
		ID:        dbUser.UUID,
		Name:      dbUser.Name,
		CreatedAt: dbUser.CreatedAt,
		UpdatedAt: dbUser.UpdatedAt,
		ApiKey:    dbUser.ApiKey,
	}
}
