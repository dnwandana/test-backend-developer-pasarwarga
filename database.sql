# create database
CREATE DATABASE test_backend_developer_pasarwarga;

# use database
USE test_backend_developer_pasarwarga;

# create article table
CREATE TABLE articles
(
    id          INT          NOT NULL AUTO_INCREMENT,
    title       VARCHAR(100) NOT NULL,
    slug        VARCHAR(100) NOT NULL,
    category_id INT          NOT NULL,
    content     TEXT         NOT NULL,
    created_at  DATETIME     NOT NULL DEFAULT NOW(),
    updated_at  DATETIME     NULL ON UPDATE NOW(),
    deleted_at  DATETIME     NULL,
    UNIQUE (slug),
    PRIMARY KEY (id)
) ENGINE = InnoDB;

# create category table
CREATE TABLE categories
(
    id            INT         NOT NULL AUTO_INCREMENT,
    category_name VARCHAR(30) NOT NULL,
    category_slug VARCHAR(30) NOT NULL,
    created_at    DATETIME    NOT NULL DEFAULT NOW(),
    updated_at    DATETIME    NULL ON UPDATE NOW(),
    deleted_at    DATETIME    NULL,
    UNIQUE (category_name),
    PRIMARY KEY (id)
) ENGINE = InnoDB