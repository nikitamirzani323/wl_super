package entities

type Model_game struct {
	Game_id             int    `json:"game_id"`
	Game_idcategame     string `json:"game_idcategame"`
	Game_idprovidergame string `json:"game_idprovidergame"`
	Game_name           string `json:"game_name"`
	Game_imgcover       string `json:"game_imgcover"`
	Game_imgthumb       string `json:"game_imgthumb"`
	Game_endpointurl    string `json:"game_endpointurl"`
	Game_status         string `json:"game_status"`
	Game_create         string `json:"game_create"`
	Game_update         string `json:"game_update"`
}

type Controller_game struct {
	Master string `json:"master" validate:"required"`
}
type Controller_gamesave struct {
	Sdata               string `json:"sdata" validate:"required"`
	Page                string `json:"page" validate:"required"`
	Game_id             int    `json:"game_id"`
	Game_idcategame     string `json:"game_idcategame" validate:"required"`
	Game_idprovidergame string `json:"game_idprovidergame" validate:"required"`
	Game_name           string `json:"game_name" validate:"required"`
	Game_imgcover       string `json:"game_imgcover"`
	Game_imgthumb       string `json:"game_imgthumb"`
	Game_endpointurl    string `json:"game_endpointurl"`
	Game_status         string `json:"game_status" validate:"required"`
}
