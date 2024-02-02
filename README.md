Backend for te-eme project.

Primera idea de base de datos relacional:

CREATE TABLE Usuarios (
    ID INT PRIMARY KEY,
    Alias VARCHAR(100) NOT NULL
);

CREATE TABLE Juegos (
    ID INT PRIMARY KEY,
    Nombre VARCHAR(100) NOT NULL
);

CREATE TABLE Torneos (
    ID INT PRIMARY KEY,
    ID_Juego INT,
    FOREIGN KEY (ID_Juego) REFERENCES Juegos(ID)
);

CREATE TABLE Partidas (
    ID INT PRIMARY KEY,
    ID_Torneo INT,
    FOREIGN KEY (ID_Torneo) REFERENCES Torneos(ID)
);

CREATE TABLE Ganadores (
    ID_Partida INT,
    ID_Usuario INT,
    FOREIGN KEY (ID_Partida) REFERENCES Partidas(ID),
    FOREIGN KEY (ID_Usuario) REFERENCES Usuarios(ID)
);

CREATE TABLE Perdedores (
    ID_Partida INT,
    ID_Usuario INT,
    FOREIGN KEY (ID_Partida) REFERENCES Partidas(ID),
    FOREIGN KEY (ID_Usuario) REFERENCES Usuarios(ID)
);
