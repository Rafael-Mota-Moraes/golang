INSERT INTO usuarios (nome, nick, email, senha) VALUES
('Usuário 1', 'usuario_1', 'user1@email.com', '$2a$10$PvToyf9CgtxoY7mcoOI7N.DPeKiBI/leWarGatebv5iM6kiYOjENa'),
('Usuário 2', 'usuario_2', 'user2@email.com', '$2a$10$PvToyf9CgtxoY7mcoOI7N.DPeKiBI/leWarGatebv5iM6kiYOjENa'),
('Usuário 3', 'usuario_3', 'user3@email.com', '$2a$10$PvToyf9CgtxoY7mcoOI7N.DPeKiBI/leWarGatebv5iM6kiYOjENa'),
('Usuário 4', 'usuario_4', 'user4@email.com', '$2a$10$PvToyf9CgtxoY7mcoOI7N.DPeKiBI/leWarGatebv5iM6kiYOjENa'),
('Usuário 5', 'usuario_5', 'user5@email.com', '$2a$10$PvToyf9CgtxoY7mcoOI7N.DPeKiBI/leWarGatebv5iM6kiYOjENa'),
('Usuário 6', 'usuario_6', 'user6@email.com', '$2a$10$PvToyf9CgtxoY7mcoOI7N.DPeKiBI/leWarGatebv5iM6kiYOjENa'),
('Usuário 7', 'usuario_7', 'user7@email.com', '$2a$10$PvToyf9CgtxoY7mcoOI7N.DPeKiBI/leWarGatebv5iM6kiYOjENa'),
('Usuário 8', 'usuario_8', 'user8@email.com', '$2a$10$PvToyf9CgtxoY7mcoOI7N.DPeKiBI/leWarGatebv5iM6kiYOjENa'),
('Usuário 9', 'usuario_9', 'user9@email.com', '$2a$10$PvToyf9CgtxoY7mcoOI7N.DPeKiBI/leWarGatebv5iM6kiYOjENa'),
('Usuário 10', 'usuario_10', 'user10@email.com', '$2a$10$PvToyf9CgtxoY7mcoOI7N.DPeKiBI/leWarGatebv5iM6kiYOjENa'),
('Usuário 11', 'usuario_11', 'user11@email.com', '$2a$10$PvToyf9CgtxoY7mcoOI7N.DPeKiBI/leWarGatebv5iM6kiYOjENa'),
('Usuário 12', 'usuario_12', 'user12@email.com', '$2a$10$PvToyf9CgtxoY7mcoOI7N.DPeKiBI/leWarGatebv5iM6kiYOjENa'),
('Usuário 13', 'usuario_13', 'user13@email.com', '$2a$10$PvToyf9CgtxoY7mcoOI7N.DPeKiBI/leWarGatebv5iM6kiYOjENa'),
('Usuário 14', 'usuario_14', 'user14@email.com', '$2a$10$PvToyf9CgtxoY7mcoOI7N.DPeKiBI/leWarGatebv5iM6kiYOjENa'),
('Usuário 15', 'usuario_15', 'user15@email.com', '$2a$10$PvToyf9CgtxoY7mcoOI7N.DPeKiBI/leWarGatebv5iM6kiYOjENa');


insert into seguidores(usuario_id, seguidor_id) 
values
(1, 2),
(3, 2),
(2, 1),
(1, 3);

insert into publicacoes (titulo, conteudo, autor_id)
values
("P1", "publicacao p1", 1),
("P2", "publicacao p2", 2),
("P3", "publicacao p3", 3);