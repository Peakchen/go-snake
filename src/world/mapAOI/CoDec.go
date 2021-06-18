package mapAOI

const (

	bytesPerChar = 1 // How many bytes to encode a character of a string
    bytesPerID = 2 // How many bytes to encode numerical id"s (a maximum id of 2^16 = 65536 seems reasonable for a small game, "real" games should use at least 3 bytes)
    booleanBytes = 1 // How many bytes to use to represent booleans (= 8 booleans per byte allocated),
    stampBytes = 4 // How many bytes to encode timestamp (a timestamp takes more room than 4 bytes, but only the last 4 bytes are relevant, since the time spans incoded in the remaining ones are too big to be useful)

)

type TtileSchema struct {
	propertiesBytes int
	numerical struct{
		x int
		y int
	}
}

type TplayerRouteSchema struct {
	propertiesBytes int
	numerical struct {
		orientation int
		delta 		int
	}
	standAlone struct {
		end *TtileSchema
	}
}

type TmonsterRouteSchema struct{
	propertiesBytes int
	numerical struct {
		delta int
	}
	arrays struct {
		path *TtileSchema
	}
}

type Tint16schema struct {
	Primitive 	bool
	Type 		string
	Bytes 		int
}

type TplayerSchema struct{
	propertiesBytes int
	numerical struct {
		id int
		x int
		y int
		weapon int
		armor int
		aoi int
		targetID int
	}
	strings []string
	booleans []string
	standAlone struct {
		route *TplayerRouteSchema
	}
}

type TitemSchema struct {
	propertiesBytes int
	numerical struct {
		id int
		x  int
		y  int
		itemID int
	}

	booleans []string
}

type TmonsterSchema struct {
	propertiesBytes int
	numerical struct {
		id int
		x int
		y int
		targetID int
		lastHitter int
		monster int
	}
	booleans []string
	standAlone struct {
		route *TmonsterRouteSchema
	}
	
}

type TglobalUpdateSchema struct {
	propertiesBytes int // How many bytes to use to indicate the presence/absence of fields in the object; Limits the number of encodable fields to 8*propertiesBytes
	arrays struct {
		newplayers *TplayerSchema
		newitems *TitemSchema
		newmonsters *TmonsterSchema
		disconnected *Tint16schema
	}
	maps struct {
		players *TplayerSchema
		monsters *TmonsterSchema
		items *TitemSchema
	}
}

type ThpSchema struct {
	propertiesBytes int32
	numerical struct {
        hp int
        from int
    }
	booleans []string
}

type TlocalUpdateSchema struct {
	propertiesBytes int
	numerical struct{
		life int
		x int
		y int
	}
	booleans []string
	arrays struct {
		hp *ThpSchema
		killed *Tint16schema
		used *Tint16schema
	}
}

var (

	int16schema = &Tint16schema{
		Primitive: true,
		Type: "int",
		Bytes: 2,
	}

	tileSchema = &TtileSchema{
		propertiesBytes: 1,
		numerical: struct{
			x int
			y int
		}{
			x: 2,
			y: 2,
		},
	}
	
	playerRouteSchema = &TplayerRouteSchema{
		propertiesBytes: 1,

		numerical : struct {
			orientation int
			delta 		int
		}{
			orientation: 1,
			delta: 2,
		},

		standAlone : struct {
			end *TtileSchema
		}{
			end: tileSchema,
		},
	}

	monsterRouteSchema = &TmonsterRouteSchema{
		propertiesBytes: 1,
		numerical : struct {
			delta int
		}{
			delta: 2,
		},
		arrays : struct {
			path *TtileSchema
		}{
			path : tileSchema,
		},
	}

	playerSchema = &TplayerSchema{
		propertiesBytes: 2,
		numerical : struct {
			id int
			x int
			y int
			weapon int
			armor int
			aoi int
			targetID int
		}{
			id: bytesPerID,
			x: 2,
			y: 2,
			weapon: 1,
			armor: 1,
			aoi: 2,
			targetID: bytesPerID,
		},
		strings : []string{"name"},
		booleans: []string{"inFight","alive"},
		standAlone : struct{
			route *TplayerRouteSchema
		}{
			route: playerRouteSchema,
		},
	}
	
	itemSchema = &TitemSchema{
		propertiesBytes: 2,
		numerical : struct {
			id int
			x  int
			y  int
			itemID int
		}{
			id: bytesPerID,
			x: 2,
			y: 2,
			itemID: 1,
		},
		booleans: []string{"visible","respawn","chest","inChest","loot"},
	}

	monsterSchema = &TmonsterSchema{
		propertiesBytes: 2,
		numerical : struct {
			id int
			x int
			y int
			targetID int
			lastHitter int
			monster int
		}{
			id: bytesPerID,
			x: 2,
			y: 2,
			targetID: bytesPerID,
			lastHitter: bytesPerID,
			monster: 1,
		},
		booleans: []string{"inFight","alive"},
		standAlone : struct {
			route *TmonsterRouteSchema
		}{
			route: monsterRouteSchema,
		},
	}

	globalUpdateSchema = &TglobalUpdateSchema{
		propertiesBytes: 1, // How many bytes to use to indicate the presence/absence of fields in the object; Limits the number of encodable fields to 8*propertiesBytes
		arrays : struct {
			newplayers *TplayerSchema
			newitems *TitemSchema
			newmonsters *TmonsterSchema
			disconnected *Tint16schema
		}{
			newplayers: playerSchema,
			newitems: itemSchema,
			newmonsters: monsterSchema,
			disconnected: int16schema,
		},
		maps : struct {
			players *TplayerSchema
			monsters *TmonsterSchema
			items *TitemSchema
		}{
			players: playerSchema,
			monsters: monsterSchema,
			items: itemSchema,
		},
	}

	hpSchema = &ThpSchema{
		propertiesBytes: 1,
		numerical : struct {
			hp int
			from int
		}{
			hp : 1,
			from: bytesPerID,
		},
		booleans : []string{"target"},
	}

	localUpdateSchema = &TlocalUpdateSchema{
		propertiesBytes: 1,
		numerical : struct{
			life int
			x int
			y int
		}{
			life : 1,
			x: 2,
			y: 2,
		},
		booleans : []string{"noPick"},
		arrays : struct {
			hp *ThpSchema
			killed *Tint16schema
			used *Tint16schema
		}{
			hp : hpSchema,
			killed: int16schema,
			used: int16schema,
		},
	}

	finalUpdateSchema = struct{
		propertiesBytes int
		numerical struct {
			latency int
			stamp int
			nbconnected int
		}
		standAlone struct{
			global *TglobalUpdateSchema
			local *TlocalUpdateSchema
		}
	}{
		propertiesBytes: 1,
		numerical : struct {
			latency int
			stamp int
			nbconnected int
		}{
			latency: 2,
			stamp: stampBytes,
			nbconnected: 1,
		},
		standAlone : struct{
			global *TglobalUpdateSchema
			local *TlocalUpdateSchema
		}{
			global : globalUpdateSchema,
			local : localUpdateSchema,
		},
	}

	initializationSchema = struct{
		propertiesBytes int
		numerical struct{
			stamp int
			nbconnected int
			nbAOIhorizontal int
			lastAOIid int
		}
		standAlone struct{
			player *TplayerSchema
		}
	}{
		propertiesBytes: 1,
		numerical : struct{
			stamp int
			nbconnected int
			nbAOIhorizontal int
			lastAOIid int
		}{
			stamp: stampBytes,
			nbconnected: 1,
			nbAOIhorizontal: 1,
			lastAOIid: 2,
		},
		standAlone : struct{
			player *TplayerSchema
		}{
			player: playerSchema,
		},
	}
)