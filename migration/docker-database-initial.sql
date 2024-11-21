CREATE TABLE mangas (
    id SERIAL PRIMARY KEY,
    nome VARCHAR(255),
    preco NUMERIC(10, 2),
    vnumero INTEGER,
    mespub SMALLINT,
    anopub INTEGER,
    img TEXT
);

INSERT INTO mangas (nome, preco, vnumero, mespub, anopub, img)
VALUES
    ('Acho Que Meu Filho é Gay', 37.9, 1, 4, 2024, 'https://d14d9vp3wdof84.cloudfront.net/image/589816272436/image_gtifsapffd6sf5odn1d3tffl7u/-S897-FWEBP'),
    ('Dandadan', 40.9, 9, 4, 2024, 'https://d14d9vp3wdof84.cloudfront.net/image/589816272436/image_ujjnp8t94h53v67m20teu4hg2a/-S897-FWEBP'),
    ('Insones - Caçando Estrelas Depois Da Aula', 40.9, 12, 4, 2024, 'https://d14d9vp3wdof84.cloudfront.net/image/589816272436/image_ukl28ckf9t7g509ip0bi017d7e/-S897-FWEBP'),
    ('Kemono Jihen - Incidentes Sobrenaturais', 39.9, 19, 4, 2024, 'https://d14d9vp3wdof84.cloudfront.net/image/589816272436/image_fltm8qqm8l41hb57357nmllv5g/-S897-FWEBP'),
    ('Komi Não Consegue Se Comunicar', 40.9, 25, 4, 2024, 'https://d14d9vp3wdof84.cloudfront.net/image/589816272436/image_mou9inc74d44d9ckjfl3gd572h/-S897-FWEBP'),
    ('Wind Breaker', 39.9, 9, 4, 2024, 'https://d14d9vp3wdof84.cloudfront.net/image/589816272436/image_68v3ahddrt7dh8iibqr40uas11/-S897-FWEBP'),
    ('Yomotsuhegui - O Fruto Do Mundo Dos Mortos', 40.9, 3, 4, 2024, 'https://d14d9vp3wdof84.cloudfront.net/image/589816272436/image_g4mrml92096v1e73h51bvqai4r/-S897-FWEBP');
