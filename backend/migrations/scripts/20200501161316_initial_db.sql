-- +migrate Up
CREATE TABLE IF NOT EXISTS types(
  id INT NOT NULL PRIMARY KEY AUTO_INCREMENT,
  name VARCHAR(100) NOT NULL
);

CREATE TABLE IF NOT EXISTS regions(
  id INT NOT NULL PRIMARY KEY AUTO_INCREMENT,
  name VARCHAR(100) NOT NULL
);

CREATE TABLE IF NOT EXISTS pokemons(
  id INT NOT NULL PRIMARY KEY AUTO_INCREMENT,
  region_id INT NOT NULL,
  evolution INT NULL,
  number INT NOT NULL,
  name VARCHAR(100) NOT NULL,
  image VARCHAR(200) NOT NULL,
  created_at DATETIME NOT NULL DEFAULT NOW(),
  updated_at DATETIME NULL,
  deleted_at DATETIME NULL,

  CONSTRAINT fk_region_id FOREIGN KEY(region_id) REFERENCES regions(id),
  CONSTRAINT fk_evolution FOREIGN KEY(evolution) REFERENCES pokemons(id)
);

CREATE TABLE IF NOT EXISTS pokemon_type(
  pokemon_id INT NOT NULL,
  type_id INT NOT NULL,
  PRIMARY KEY(pokemon_id, type_id),

  CONSTRAINT fk_type_id FOREIGN KEY(type_id) REFERENCES types(id),
  CONSTRAINT fk_pokemon_id FOREIGN KEY(pokemon_id) REFERENCES pokemons(id)
);

-- +migrate Down
DROP TABLE IF EXISTS pokemon_type;
DROP TABLE IF EXISTS pokemons;
DROP TABLE IF EXISTS types;
DROP TABLE IF EXISTS regions;
