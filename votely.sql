-- Tabela de usuários (apenas o administrador)
CREATE TABLE IF NOT EXISTS users (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    username TEXT NOT NULL UNIQUE,
    password TEXT NOT NULL
);

-- Tabela de eleições
CREATE TABLE IF NOT EXISTS elections (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    title TEXT NOT NULL,
    description TEXT
);

-- Tabela de candidatos vinculados a uma eleição
CREATE TABLE IF NOT EXISTS candidates (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    election_id INTEGER NOT NULL,
    name TEXT NOT NULL,
    FOREIGN KEY (election_id) REFERENCES elections(id)
);

-- Tabela de blocos para votos
CREATE TABLE IF NOT EXISTS blocks (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    candidate_id INTEGER NOT NULL,
    previous_hash TEXT NOT NULL,
    hash TEXT NOT NULL,
    FOREIGN KEY (candidate_id) REFERENCES candidates(id)
);

