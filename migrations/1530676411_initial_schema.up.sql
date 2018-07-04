BEGIN;

CREATE TABLE login (
	account_id uuid NOT NULL PRIMARY KEY,
	source_ip text NOT NULL,
	date_created timestamptz NOT NULL
);


CREATE TABLE registration (
	id uuid NOT NULL PRIMARY KEY,
	source_ip text NOT NULL,
	date_registered timestamptz NOT NULL
);


COMMIT;
