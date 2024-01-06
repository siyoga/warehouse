ALTER DATABASE users_db SET timezone TO 'Europe/Moscow';
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

ALTER SYSTEM SET wal_level = logical;

CREATE PUBLICATION IF NOT EXISTS "$POSTGRES_PUB_NAME" FOR All TABLES;