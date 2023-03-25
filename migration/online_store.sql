DROP DATABASE IF EXISTS online_store;

CREATE DATABASE online_store;

\connect online_store;

SET client_encoding = 'UTF8';

CREATE TABLE country(
                        country_id UUID NOT NULL DEFAULT gen_random_uuid(),
                        country_name VARCHAR(40) NOT NULL,
                        PRIMARY KEY (country_id)
);

CREATE TABLE city(
                     city_id UUID NOT NULL DEFAULT gen_random_uuid(),
                     city_name VARCHAR(40) NOT NULL,
                     country_id UUID NOT NULL,
                     PRIMARY KEY (city_id),
                     FOREIGN KEY (country_id)
                         REFERENCES country (country_id)
                         ON DELETE CASCADE
);

CREATE TABLE image(
                      image_path VARCHAR(64) NOT NULL,
                      PRIMARY KEY (image_path)
);

CREATE TABLE "user"(
                       user_id UUID NOT NULL DEFAULT gen_random_uuid(),
                       username VARCHAR(40) NOT NULL,
                       password VARCHAR(64) NOT NULL,
                       city_id UUID DEFAULT NULL,
                       image_path VARCHAR(64) DEFAULT NULL,
                       role VARCHAR(6) NOT NULL,
                       PRIMARY KEY (user_id),
                       CHECK (role IN ('buyer', 'seller')),
                       FOREIGN KEY (city_id)
                           REFERENCES city (city_id)
                           ON DELETE SET NULL,
                       FOREIGN KEY (image_path)
                           REFERENCES image (image_path)
                           ON DELETE SET NULL
);

CREATE TABLE color(
                      color_id UUID NOT NULL DEFAULT gen_random_uuid(),
                      color_name VARCHAR(40) NOT NULL,
                      PRIMARY KEY (color_id)
);

CREATE TABLE category(
                         category_id UUID NOT NULL DEFAULT gen_random_uuid(),
                         category_name VARCHAR(40) NOT NULL,
                         PRIMARY KEY (category_id)
);

CREATE TABLE product(
                        product_id UUID NOT NULL DEFAULT gen_random_uuid(),
                        product_name VARCHAR(40) NOT NULL,
                        user_id UUID NOT NULL,
                        unit_price INTEGER NOT NULL,
                        units_in_stock INTEGER NOT NULL,
                        color_id UUID DEFAULT NULL,
                        category_id UUID DEFAULT NULL,
                        PRIMARY KEY (product_id),
                        FOREIGN KEY (user_id)
                            REFERENCES "user" (user_id)
                            ON DELETE CASCADE,
                        FOREIGN KEY (color_id)
                            REFERENCES color (color_id)
                            ON DELETE SET NULL,
                        FOREIGN KEY (category_id)
                            REFERENCES category (category_id)
                            ON DELETE SET NULL
);

CREATE TABLE product_image(
                              image_path VARCHAR(64) NOT NULL,
                              product_id UUID NOT NULL,
                              FOREIGN KEY (product_id)
                                  REFERENCES product (product_id)
                                  ON DELETE CASCADE,
                              FOREIGN KEY (image_path)
                                  REFERENCES image (image_path)
                                  ON DELETE CASCADE
);

CREATE TABLE offer(
                      offer_id UUID NOT NULL DEFAULT gen_random_uuid(),
                      product_id UUID,
                      offered_unit_price INTEGER NOT NULL,
                      desired_quantity INTEGER NOT NULL,
                      sender_id UUID,
                      recipient_id UUID,
                      offer_type VARCHAR(8) NOT NULL,
                      status VARCHAR(8) NOT NULL,
                      PRIMARY KEY (offer_id),
                      CHECK (offer_type IN ('purchase', 'selling')),
                      CHECK (status IN ('pending', 'approved', 'rejected')),
                      FOREIGN KEY (product_id)
                          REFERENCES product (product_id)
                          ON DELETE SET NULL,
                      FOREIGN KEY (sender_id)
                          REFERENCES "user"
                          ON DELETE SET NULL,
                      FOREIGN KEY (recipient_id)
                          REFERENCES "user"
                          ON DELETE SET NULL
);

INSERT INTO country (country_name)
VALUES
    ('Норвегия'),
    ('Ирландия'),
    ('Швейцария');

INSERT INTO city (city_name, country_id)
VALUES
    ('Акерсхус', (SELECT country_id FROM country WHERE country_name = 'Норвегия')),
    ('Бускеруд', (SELECT country_id FROM country WHERE country_name = 'Норвегия')),
    ('Вестфолл', (SELECT country_id FROM country WHERE country_name = 'Норвегия')),

    ('Дублин', (SELECT country_id FROM country WHERE country_name = 'Ирландия')),
    ('Голуэй', (SELECT country_id FROM country WHERE country_name = 'Ирландия')),
    ('Карлоу', (SELECT country_id FROM country WHERE country_name = 'Ирландия')),

    ('Цюрих', (SELECT country_id FROM country WHERE country_name = 'Швейцария')),
    ('Женева', (SELECT country_id FROM country WHERE country_name = 'Швейцария')),
    ('Базель', (SELECT country_id FROM country WHERE country_name = 'Швейцария'));

INSERT INTO color (color_name)
VALUES
    ('Бежевый'),
    ('Бордовый'),
    ('Оранжевый');

INSERT INTO category (category_name)
VALUES
    ('Транспорт'),
    ('Недвижимость'),
    ('Электроника'),
    ('Одежда');
