package main

func get_players(who ...string){

}

// Inserts new player into db.
func insert_new_player(string player){
	player = player.strtolower(player)

	if($this->db->player_exists($player)){
		$this->add_out("Player '$player' already exists.","msg","ERROR")
	} else {
		$this->db->insert_new_player($player)
		$this->add_out("New player $player entered with 0 points.","msg","OK")
	}
	return;
}

