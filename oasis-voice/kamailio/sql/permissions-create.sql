CREATE TABLE IF NOT EXISTS kamailio_trusted (
    id SERIAL PRIMARY KEY NOT NULL,
    src_ip VARCHAR(50) NOT NULL,
    proto VARCHAR(4) NOT NULL,
    from_pattern VARCHAR(64) DEFAULT NULL,
    ruri_pattern VARCHAR(64) DEFAULT NULL,
    tag VARCHAR(64),
    priority INTEGER DEFAULT 0 NOT NULL
);

CREATE INDEX trusted_peer_idx ON kamailio_trusted (src_ip);

INSERT INTO version (table_name, table_version) values ('kamailio_trusted','6');

CREATE TABLE IF NOT EXISTS kamailio_address (
    id SERIAL PRIMARY KEY NOT NULL,
    grp INTEGER DEFAULT 1 NOT NULL,
    ip_addr VARCHAR(50) NOT NULL,
    mask INTEGER DEFAULT 32 NOT NULL,
    port SMALLINT DEFAULT 0 NOT NULL,
    tag VARCHAR(64)
);

INSERT INTO version (table_name, table_version) values ('kamailio_address','6');

