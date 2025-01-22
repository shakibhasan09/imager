-- +goose Up
-- +goose StatementBegin

-- projects begin
CREATE TABLE IF NOT EXISTS projects (
  uuid text NOT NULL PRIMARY KEY,
  name text NOT NULL,
  created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TRIGGER IF NOT EXISTS update_projects_updated_at
BEFORE UPDATE ON projects
FOR EACH ROW
BEGIN
  UPDATE projects
  SET updated_at = CURRENT_TIMESTAMP
  WHERE uuid = OLD.uuid;
END;
-- projects end

-- variants begin
CREATE TABLE IF NOT EXISTS variants (
  uuid text NOT NULL PRIMARY KEY,
  project_uuid text NOT NULL,
  name text NOT NULL,
  height integer NOT NULL,
  width integer NOT NULL,
  fit text NOT NULL,
  metadata boolean NOT NULL,
  created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX index_variants ON variants (project_uuid);

CREATE TRIGGER IF NOT EXISTS update_variants_updated_at
BEFORE UPDATE ON variants
FOR EACH ROW
BEGIN
  UPDATE variants
  SET updated_at = CURRENT_TIMESTAMP
  WHERE uuid = OLD.uuid;
END;
-- variants end

-- images begin
CREATE TABLE IF NOT EXISTS images (
  uuid text NOT NULL PRIMARY KEY,
  project_uuid text NOT NULL,
  name text NOT NULL,
  metadata text,
  created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX index_images ON images (project_uuid);

CREATE TRIGGER IF NOT EXISTS update_images_updated_at
BEFORE UPDATE ON images
FOR EACH ROW
BEGIN
  UPDATE images
  SET updated_at = CURRENT_TIMESTAMP
  WHERE uuid = OLD.uuid;
END;
-- images end

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS projects;
DROP TRIGGER IF EXISTS update_projects_updated_at;

DROP TABLE IF EXISTS variants;
DROP TRIGGER IF EXISTS update_variants_updated_at;

DROP TABLE IF EXISTS images;
DROP TRIGGER IF EXISTS update_images_updated_at;
-- +goose StatementEnd
